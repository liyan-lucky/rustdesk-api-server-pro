<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { fetchServerConfig } from '@/service/api/home';

defineOptions({
  name: 'ServerConnectionConfig'
});

type ConfigKey = 'idServer' | 'relayServer' | 'apiServer' | 'key';

interface ConfigItem {
  key: ConfigKey;
  label: string;
  value: string;
  placeholder?: string;
}

const { t } = useI18n();

const config = ref<Api.Home.ServerConfig>({
  idServer: '',
  relayServer: '',
  apiServer: '',
  key: ''
});
const loading = ref(false);
const loadError = ref('');
const showKey = ref(false);

const items = computed<ConfigItem[]>(() => [
  {
    key: 'idServer',
    label: t('page.home.serverConfig.idServer'),
    value: config.value.idServer,
    placeholder: t('page.home.serverConfig.idServerPlaceholder')
  },
  {
    key: 'relayServer',
    label: t('page.home.serverConfig.relayServer'),
    value: config.value.relayServer,
    placeholder: t('page.home.serverConfig.relayServerPlaceholder')
  },
  {
    key: 'apiServer',
    label: t('page.home.serverConfig.apiServer'),
    value: config.value.apiServer,
    placeholder: t('page.home.serverConfig.apiServerPlaceholder')
  },
  {
    key: 'key',
    label: t('page.home.serverConfig.key'),
    value: config.value.key,
    placeholder: t('page.home.serverConfig.keyPlaceholder')
  }
]);

const maskedKeyValue = computed(() => {
  const value = config.value.key;
  if (!value) return '';
  return '*'.repeat(Math.min(Math.max(value.length, 8), 64));
});

function getDisplayValue(item: ConfigItem) {
  if (item.key === 'key' && !showKey.value) {
    return maskedKeyValue.value;
  }
  return item.value;
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

async function copyValue(value: string, label: string) {
  if (!value) {
    window.$message?.warning(t('page.home.serverConfig.copyEmpty', { label }));
    return;
  }

  try {
    if (navigator.clipboard?.writeText) {
      await navigator.clipboard.writeText(value);
    } else {
      fallbackCopyText(value);
    }
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

onMounted(async () => {
  loading.value = true;
  loadError.value = '';
  try {
    const res = await fetchServerConfig();
    if (res.data) {
      config.value = res.data;
      return;
    }
    loadError.value = t('page.home.serverConfig.fetchFailed');
  } catch {
    loadError.value = t('page.home.serverConfig.fetchFailed');
  } finally {
    loading.value = false;
  }
});
</script>

<template>
  <NCard :title="$t('page.home.serverConfig.title')" :bordered="false" size="small" class="card-wrapper">
    <NAlert type="info" :show-icon="false" class="mb-12px">
      {{ $t('page.home.serverConfig.tip') }}
    </NAlert>
    <NAlert v-if="loadError" type="warning" :show-icon="false" class="mb-12px">
      {{ loadError }}
    </NAlert>
    <NSpace vertical :size="10">
      <div v-for="item in items" :key="item.key" class="config-row">
        <div class="config-label">{{ item.label }}</div>
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
  grid-template-columns: 92px minmax(0, 1fr) auto;
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

.config-actions {
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}

@media (max-width: 640px) {
  .config-row {
    grid-template-columns: 1fr;
  }

  .config-actions {
    justify-content: flex-start;
  }
}
</style>
