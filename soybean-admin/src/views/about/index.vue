<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { getBuildTime, getVersionTag } from '@/utils/version';
import { $t } from '@/locales';

const defaultCheckUrl = 'https://raw.githubusercontent.com/liyan-lucky/rustdesk-api-server-pro/main/VERSION';
const storageKey = 'rustdesk-api-update-check-url';
const checkUrl = ref(localStorage.getItem(storageKey) || defaultCheckUrl);
const runningVersion = ref(getVersionTag().replace(/^v/, ''));
const latestVersion = ref('');
const compatVersion = ref('');
const checking = ref(false);
const errorMessage = ref('');
const hasUpdate = computed(() => latestVersion.value && compareVersions(latestVersion.value, runningVersion.value) > 0);

function normalizeVersion(value: string) {
  const match = value.trim().match(/v?(\d+\.\d+\.\d+)/);
  return match?.[1] || '';
}

function compareVersions(a: string, b: string) {
  const left = a.split('.').map(Number);
  const right = b.split('.').map(Number);
  for (let i = 0; i < 3; i += 1) {
    if ((left[i] || 0) !== (right[i] || 0)) return (left[i] || 0) - (right[i] || 0);
  }
  return 0;
}

async function loadRuntimeVersion() {
  try {
    const response = await fetch('/api/version', { cache: 'no-store' });
    const payload = await response.json();
    runningVersion.value = payload?.compat_target?.server?.version || normalizeVersion(payload?.version || '') || runningVersion.value;
    compatVersion.value = payload?.compat_target?.client?.version || '';
  } catch { /* build-injected version remains available */ }
}

async function checkUpdate() {
  checking.value = true;
  errorMessage.value = '';
  latestVersion.value = '';
  try {
    const url = new URL(checkUrl.value, window.location.origin);
    if (!['http:', 'https:'].includes(url.protocol)) throw new Error($t('page.about.invalidUrl'));
    const response = await fetch(url.toString(), { cache: 'no-store', headers: { Accept: 'application/json, text/plain' } });
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    const text = await response.text();
    let candidate = text;
    try {
      const json = JSON.parse(text);
      candidate = json.version || json.latest_version || json.tag_name || json.server?.version || text;
    } catch { /* plain VERSION file */ }
    latestVersion.value = normalizeVersion(candidate);
    if (!latestVersion.value) throw new Error($t('page.about.invalidResponse'));
    localStorage.setItem(storageKey, url.toString());
  } catch (error) {
    errorMessage.value = error instanceof Error ? error.message : String(error);
  } finally { checking.value = false; }
}

function restoreDefault() {
  checkUrl.value = defaultCheckUrl;
  localStorage.setItem(storageKey, defaultCheckUrl);
}

onMounted(() => { loadRuntimeVersion(); checkUpdate(); });
</script>

<template>
  <NSpace vertical size="large">
    <NCard :title="$t('route.about')" :bordered="false">
      <NDescriptions bordered label-placement="left" :column="1">
        <NDescriptionsItem :label="$t('page.about.runningVersion')"><NTag type="success">{{ runningVersion }}</NTag></NDescriptionsItem>
        <NDescriptionsItem :label="$t('page.about.buildTime')">{{ getBuildTime() || '-' }}</NDescriptionsItem>
        <NDescriptionsItem :label="$t('page.about.compatVersion')">{{ compatVersion || '-' }}</NDescriptionsItem>
        <NDescriptionsItem :label="$t('page.about.latestVersion')">{{ latestVersion || '-' }} <NTag v-if="latestVersion" class="ml-8px" :type="hasUpdate ? 'warning' : 'success'">{{ hasUpdate ? $t('page.about.updateAvailable') : $t('page.about.upToDate') }}</NTag></NDescriptionsItem>
      </NDescriptions>
    </NCard>
    <NCard :title="$t('page.about.updateCheck')" :bordered="false">
      <NAlert type="info" class="mb-16px">{{ $t('page.about.urlTip') }}</NAlert>
      <NSpace vertical>
        <NInput v-model:value="checkUrl" :placeholder="$t('page.about.urlPlaceholder')" />
        <NSpace><NButton type="primary" :loading="checking" @click="checkUpdate">{{ $t('page.about.checkNow') }}</NButton><NButton @click="restoreDefault">{{ $t('page.about.restoreDefault') }}</NButton></NSpace>
        <NAlert v-if="errorMessage" type="error">{{ $t('page.about.checkFailed') }}: {{ errorMessage }}</NAlert>
      </NSpace>
    </NCard>
  </NSpace>
</template>
