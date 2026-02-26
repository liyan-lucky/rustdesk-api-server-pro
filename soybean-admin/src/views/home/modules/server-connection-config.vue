<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import { fetchServerConfig, fetchServerConnectivity } from '@/service/api/home';

defineOptions({
  name: 'ServerConnectionConfig'
});

type ConfigKey = 'idServer' | 'relayServer' | 'apiServer' | 'key';

interface ConfigFieldMeta {
  key: ConfigKey;
  labelKey: string;
  placeholderKey: string;
}

interface ConfigItem {
  key: ConfigKey;
  label: string;
  value: string;
  placeholder?: string;
  source: 'env' | 'inferred' | 'empty';
}

type ConfigLoadSource = 'remote' | 'memory-cache' | 'session-cache' | '';
type ConfigValueSource = 'env' | 'inferred' | 'empty';
type ConnectivityStatus = 'idle' | Api.Home.ServerConnectivityItem['status'];
type ConnectivityCheckSource = 'remote' | 'cache';

const SERVER_CONFIG_CACHE_TTL_MS = 30 * 1000;
const SERVER_CONFIG_CACHE_KEY = 'home.server-config-cache.v1';
const CONNECTIVITY_CACHE_TTL_MS = 10 * 1000;
let serverConfigCache: Api.Home.ServerConfig | null = null;
let serverConfigCacheAt = 0;
let serverConfigInFlight: Promise<Api.Home.ServerConfig | null> | null = null;
let connectivityCache: Api.Home.ServerConnectivity | null = null;
let connectivityCacheAt = 0;

const CONFIG_FIELDS: ConfigFieldMeta[] = [
  {
    key: 'idServer',
    labelKey: 'page.home.serverConfig.idServer',
    placeholderKey: 'page.home.serverConfig.idServerPlaceholder'
  },
  {
    key: 'relayServer',
    labelKey: 'page.home.serverConfig.relayServer',
    placeholderKey: 'page.home.serverConfig.relayServerPlaceholder'
  },
  {
    key: 'apiServer',
    labelKey: 'page.home.serverConfig.apiServer',
    placeholderKey: 'page.home.serverConfig.apiServerPlaceholder'
  },
  {
    key: 'key',
    labelKey: 'page.home.serverConfig.key',
    placeholderKey: 'page.home.serverConfig.keyPlaceholder'
  }
];
const REQUIRED_KEYS = new Set<ConfigKey>(['key']);

const { t } = useI18n();

const config = ref<Api.Home.ServerConfig>({
  idServer: '',
  relayServer: '',
  apiServer: '',
  key: ''
});
const loading = ref(false);
const checkingConnectivity = ref(false);
const loadError = ref('');
const showKey = ref(false);
const lastLoadedAt = ref(0);
const lastLoadSource = ref<ConfigLoadSource>('');
const lastConnectivityCheckedAt = ref(0);
const lastConnectivityCheckSource = ref<ConnectivityCheckSource | ''>('');
let latestLoadRequestId = 0;
const connectivity = ref<Record<ConfigKey, { status: ConnectivityStatus; message: string; target: string; durationMs?: number }>>({
  idServer: { status: 'idle', message: '', target: '' },
  relayServer: { status: 'idle', message: '', target: '' },
  apiServer: { status: 'idle', message: '', target: '' },
  key: { status: 'idle', message: '', target: '' }
});

function resetConnectivityState() {
  connectivity.value = {
    idServer: { status: 'idle', message: '', target: '' },
    relayServer: { status: 'idle', message: '', target: '' },
    apiServer: { status: 'idle', message: '', target: '' },
    key: { status: 'idle', message: '', target: '' }
  };
  connectivityCache = null;
  connectivityCacheAt = 0;
  lastConnectivityCheckedAt.value = 0;
  lastConnectivityCheckSource.value = '';
}

function normalizeServerConfig(data: Api.Home.ServerConfig): Api.Home.ServerConfig {
  const sources = data.sources || {};
  return {
    idServer: (data.idServer || '').trim(),
    relayServer: (data.relayServer || '').trim(),
    apiServer: (data.apiServer || '').trim(),
    key: (data.key || '').trim(),
    sources: {
      idServer: sources.idServer || 'empty',
      relayServer: sources.relayServer || 'empty',
      apiServer: sources.apiServer || 'empty',
      key: sources.key || 'empty'
    }
  };
}

function readSessionCache() {
  if (typeof window === 'undefined') return;
  try {
    const raw = window.sessionStorage.getItem(SERVER_CONFIG_CACHE_KEY);
    if (!raw) return;
    const parsed = JSON.parse(raw) as { at?: number; data?: Api.Home.ServerConfig };
    if (!parsed?.at || !parsed?.data) return;
    if (Date.now() - parsed.at >= SERVER_CONFIG_CACHE_TTL_MS) return;
    serverConfigCache = normalizeServerConfig(parsed.data);
    serverConfigCacheAt = parsed.at;
  } catch {
    window.sessionStorage.removeItem(SERVER_CONFIG_CACHE_KEY);
  }
}

function writeSessionCache(data: Api.Home.ServerConfig) {
  if (typeof window === 'undefined') return;
  try {
    window.sessionStorage.setItem(SERVER_CONFIG_CACHE_KEY, JSON.stringify({ at: serverConfigCacheAt, data }));
  } catch {
    // ignore quota/private mode errors
  }
}

function clearServerConfigCache() {
  serverConfigCache = null;
  serverConfigCacheAt = 0;
  if (typeof window === 'undefined') return;
  try {
    window.sessionStorage.removeItem(SERVER_CONFIG_CACHE_KEY);
  } catch {
    // ignore
  }
}

async function requestServerConfig() {
  if (!serverConfigInFlight) {
    serverConfigInFlight = (async () => {
      const res = await fetchServerConfig();
      return res.data ? normalizeServerConfig(res.data) : null;
    })().finally(() => {
      serverConfigInFlight = null;
    });
  }
  return serverConfigInFlight;
}

const items = computed<ConfigItem[]>(() =>
  CONFIG_FIELDS.map(field => ({
    key: field.key,
    label: t(field.labelKey),
    value: config.value[field.key],
    placeholder: t(field.placeholderKey),
    source: (config.value.sources?.[field.key] || 'empty') as ConfigValueSource
  }))
);

const maskedKeyValue = computed(() => {
  const value = config.value.key;
  if (!value) return '';
  return '*'.repeat(Math.min(Math.max(value.length, 8), 64));
});

const missingKeys = computed(() =>
  items.value.filter(item => REQUIRED_KEYS.has(item.key) && !item.value.trim()).map(item => item.label)
);

const hasMissingRequired = computed(() => Boolean(missingKeys.value.length));
const canCopyAll = computed(() => !loading.value && !loadError.value && !hasMissingRequired.value);
const connectivityStats = computed(() => {
  const list = Object.values(connectivity.value);
  return {
    ok: list.filter(item => item.status === 'ok').length,
    error: list.filter(item => item.status === 'error').length,
    skip: list.filter(item => item.status === 'skip').length,
    checked: list.some(item => item.status !== 'idle')
  };
});

readSessionCache();
if (serverConfigCache) {
  config.value = serverConfigCache;
  lastLoadedAt.value = serverConfigCacheAt;
  lastLoadSource.value = 'session-cache';
}

const sourceLabel = computed(() => {
  if (!lastLoadSource.value) return '';
  return t(`page.home.serverConfig.sourceType.${lastLoadSource.value}`);
});

const lastUpdatedText = computed(() => {
  if (!lastLoadedAt.value) return '';
  try {
    return new Date(lastLoadedAt.value).toLocaleString();
  } catch {
    return '';
  }
});

const connectivityCheckSourceLabel = computed(() => {
  if (!lastConnectivityCheckSource.value) return '';
  return t(`page.home.serverConfig.connectivity.checkSourceType.${lastConnectivityCheckSource.value}`);
});

const lastConnectivityCheckedText = computed(() => {
  if (!lastConnectivityCheckedAt.value) return '';
  try {
    return new Date(lastConnectivityCheckedAt.value).toLocaleString();
  } catch {
    return '';
  }
});

function getDisplayValue(item: ConfigItem) {
  if (item.key === 'key' && !showKey.value) {
    return maskedKeyValue.value;
  }
  return item.value;
}

function getFieldSourceLabel(source: ConfigValueSource) {
  return t(`page.home.serverConfig.sourceType.${source}`);
}

function getFieldSourceHint(source: ConfigValueSource) {
  return t(`page.home.serverConfig.sourceHint.${source}`);
}

function getConnectivityLabel(status: ConnectivityStatus) {
  return t(`page.home.serverConfig.connectivity.status.${status}`);
}

function getConnectivityClass(status: ConnectivityStatus) {
  if (status === 'ok') return 'is-ok';
  if (status === 'error') return 'is-error';
  if (status === 'skip') return 'is-skip';
  return 'is-idle';
}

function getConnectivityHint(item: ConfigItem) {
  const state = connectivity.value[item.key];
  if (!state || state.status === 'idle') {
    return t('page.home.serverConfig.connectivity.notChecked');
  }

  const parts = [state.message];
  if (state.target) {
    parts.push(`${t('page.home.serverConfig.connectivity.target')}: ${state.target}`);
  }
  if (typeof state.durationMs === 'number') {
    parts.push(`${t('page.home.serverConfig.connectivity.duration')}: ${state.durationMs}ms`);
  }
  return parts.join(' | ');
}

function fallbackCopyText(text: string) {
  const textarea = document.createElement('textarea');
  textarea.value = text;
  textarea.style.position = 'fixed';
  textarea.style.opacity = '0';
  document.body.appendChild(textarea);
  textarea.focus();
  textarea.select();
  document.execCommand('copy');
  document.body.removeChild(textarea);
}

async function writeClipboardText(text: string) {
  if (navigator.clipboard?.writeText) {
    await navigator.clipboard.writeText(text);
    return;
  }
  fallbackCopyText(text);
}

async function copyValue(value: string, label: string) {
  if (!value) {
    window.$message?.warning(t('page.home.serverConfig.copyEmpty', { label }));
    return;
  }

  try {
    await writeClipboardText(value);
    window.$message?.success(t('page.home.serverConfig.copySuccess', { label }));
  } catch {
    try {
      fallbackCopyText(value);
      window.$message?.success(t('page.home.serverConfig.copySuccess', { label }));
    } catch {
      window.$message?.error(t('page.home.serverConfig.copyFailed', { label }));
    }
  }
}

function buildClientConfigText() {
  return items.value.map(item => `${item.label}: ${item.value || '-'}`).join('\n');
}

async function copyAllConfig() {
  if (hasMissingRequired.value) {
    window.$message?.warning(t('page.home.serverConfig.missingTip', { fields: missingKeys.value.join(' / ') }));
    return;
  }
  await copyValue(buildClientConfigText(), t('page.home.serverConfig.copyAll'));
}

async function clearCacheAndReload() {
  clearServerConfigCache();
  resetConnectivityState();
  lastLoadedAt.value = 0;
  lastLoadSource.value = '';
  window.$message?.success(t('page.home.serverConfig.cacheCleared'));
  await loadServerConfig(true);
}

async function checkConnectivity() {
  if (checkingConnectivity.value) return;
  if (connectivityCache && Date.now() - connectivityCacheAt < CONNECTIVITY_CACHE_TTL_MS) {
    connectivity.value = connectivityCache;
    lastConnectivityCheckedAt.value = connectivityCacheAt;
    lastConnectivityCheckSource.value = 'cache';
    window.$message?.success(t('page.home.serverConfig.connectivity.checkedCached'));
    return;
  }

  checkingConnectivity.value = true;
  try {
    const res = await fetchServerConnectivity();
    if (!res.data) {
      window.$message?.warning(t('page.home.serverConfig.connectivity.checkFailed'));
      return;
    }

    connectivity.value = {
      idServer: res.data.idServer,
      relayServer: res.data.relayServer,
      apiServer: res.data.apiServer,
      key: res.data.key
    };
    connectivityCache = res.data;
    connectivityCacheAt = Date.now();
    lastConnectivityCheckedAt.value = connectivityCacheAt;
    lastConnectivityCheckSource.value = 'remote';
    window.$message?.success(t('page.home.serverConfig.connectivity.checked'));
  } catch {
    window.$message?.error(t('page.home.serverConfig.connectivity.checkFailed'));
  } finally {
    checkingConnectivity.value = false;
  }
}

async function loadServerConfig(force = false) {
  if (loading.value) return;

  if (!force && serverConfigCache && Date.now() - serverConfigCacheAt < SERVER_CONFIG_CACHE_TTL_MS) {
    config.value = serverConfigCache;
    loadError.value = '';
    lastLoadedAt.value = serverConfigCacheAt;
    lastLoadSource.value = 'memory-cache';
    return;
  }

  const requestId = ++latestLoadRequestId;
  loading.value = true;
  loadError.value = '';
  try {
    const data = await requestServerConfig();
    if (requestId !== latestLoadRequestId) return;

    if (data) {
      config.value = data;
      serverConfigCache = data;
      serverConfigCacheAt = Date.now();
      lastLoadedAt.value = serverConfigCacheAt;
      lastLoadSource.value = 'remote';
      writeSessionCache(data);
      return;
    }
    loadError.value = t('page.home.serverConfig.fetchFailed');
  } catch {
    if (requestId !== latestLoadRequestId) return;
    loadError.value = t('page.home.serverConfig.fetchFailed');
  } finally {
    if (requestId === latestLoadRequestId) {
      loading.value = false;
    }
  }
}

onMounted(loadServerConfig);

watch(
  () => [config.value.idServer, config.value.relayServer, config.value.apiServer, config.value.key].join('|'),
  (next, prev) => {
    if (prev !== undefined && next !== prev) {
      resetConnectivityState();
    }
  }
);
</script>

<template>
  <NCard :title="$t('page.home.serverConfig.title')" :bordered="false" size="small" class="card-wrapper">
    <template #header-extra>
      <NSpace :size="8">
        <NButton size="small" :loading="loading" @click="loadServerConfig(true)">
          {{ $t('page.home.serverConfig.refresh') }}
        </NButton>
        <NButton size="small" :disabled="loading" @click="clearCacheAndReload">
          {{ $t('page.home.serverConfig.clearCacheReload') }}
        </NButton>
        <NButton size="small" :loading="checkingConnectivity" :disabled="loading" @click="checkConnectivity">
          {{ $t('page.home.serverConfig.connectivity.check') }}
        </NButton>
        <NButton size="small" type="primary" secondary :disabled="!canCopyAll" @click="copyAllConfig">
          {{ $t('page.home.serverConfig.copyAll') }}
        </NButton>
      </NSpace>
    </template>
    <NAlert type="info" :show-icon="false" class="mb-12px">
      {{ $t('page.home.serverConfig.tip') }}
    </NAlert>
    <div v-if="sourceLabel || lastUpdatedText" class="config-meta mb-12px">
      <span v-if="sourceLabel">
        {{ $t('page.home.serverConfig.source') }}: {{ sourceLabel }}
      </span>
      <span v-if="lastUpdatedText">
        {{ $t('page.home.serverConfig.lastUpdated') }}: {{ lastUpdatedText }}
      </span>
    </div>
    <div v-if="connectivityStats.checked" class="config-meta mb-12px">
      <span v-if="connectivityCheckSourceLabel">
        {{ $t('page.home.serverConfig.connectivity.source') }}: {{ connectivityCheckSourceLabel }}
      </span>
      <span v-if="lastConnectivityCheckedText">
        {{ $t('page.home.serverConfig.connectivity.lastChecked') }}: {{ lastConnectivityCheckedText }}
      </span>
      <span class="is-ok">{{ $t('page.home.serverConfig.connectivity.status.ok') }}: {{ connectivityStats.ok }}</span>
      <span class="is-error">{{ $t('page.home.serverConfig.connectivity.status.error') }}: {{ connectivityStats.error }}</span>
      <span class="is-skip">{{ $t('page.home.serverConfig.connectivity.status.skip') }}: {{ connectivityStats.skip }}</span>
    </div>
    <NAlert v-if="loadError" type="warning" :show-icon="false" class="mb-12px">
      {{ loadError }}
    </NAlert>
    <NAlert v-else-if="!loading && hasMissingRequired" type="warning" :show-icon="false" class="mb-12px">
      {{ $t('page.home.serverConfig.missingTip', { fields: missingKeys.join(' / ') }) }}
    </NAlert>
    <NSpace vertical :size="10">
      <div v-for="item in items" :key="item.key" class="config-row">
        <div class="config-label">{{ item.label }}</div>
        <div class="config-badges">
          <NTooltip trigger="hover">
            <template #trigger>
              <div class="config-label-source">{{ getFieldSourceLabel(item.source) }}</div>
            </template>
            {{ getFieldSourceHint(item.source) }}
          </NTooltip>
          <NTooltip trigger="hover">
            <template #trigger>
              <div class="config-label-source" :class="getConnectivityClass(connectivity[item.key].status)">
                {{ getConnectivityLabel(connectivity[item.key].status) }}
              </div>
            </template>
            {{ getConnectivityHint(item) }}
          </NTooltip>
        </div>
        <NInput
          :value="getDisplayValue(item)"
          readonly
          :placeholder="item.placeholder"
          :loading="loading"
          class="config-input"
        />
        <div class="config-actions">
          <NButton
            v-if="item.key === 'key'"
            size="small"
            :disabled="loading || !item.value"
            @click="showKey = !showKey"
          >
            {{ $t(showKey ? 'page.home.serverConfig.hide' : 'page.home.serverConfig.show') }}
          </NButton>
          <NButton size="small" :disabled="loading" @click="copyValue(item.value, item.label)">
            {{ $t('page.home.serverConfig.copy') }}
          </NButton>
        </div>
      </div>
    </NSpace>
  </NCard>
</template>

<style scoped>
.config-row {
  display: grid;
  grid-template-columns: 92px 84px minmax(0, 1fr) auto;
  gap: 8px;
  align-items: center;
}

.config-label {
  font-size: 12px;
  color: var(--n-text-color-2);
}

.config-input {
  min-width: 0;
}

.config-label-source {
  font-size: 12px;
  color: var(--n-text-color-3);
  white-space: nowrap;
}

.config-badges {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.config-label-source.is-ok {
  color: var(--n-success-color);
}

.config-label-source.is-error {
  color: var(--n-error-color);
}

.config-label-source.is-skip {
  color: var(--n-warning-color);
}

.config-actions {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}

.config-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  font-size: 12px;
  color: var(--n-text-color-3);
}

@media (max-width: 640px) {
  .config-row {
    grid-template-columns: 1fr;
  }

  .config-actions {
    justify-content: flex-start;
  }

  .config-meta {
    gap: 8px;
    flex-direction: column;
  }
}
</style>
