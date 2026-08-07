<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { fetchUserOverview } from '@/service/api/user-portal';
import { $t } from '@/locales';

const loading = ref(false);
const stats = ref({ devices: 0, sessions: 0, address_books: 0, security_events: 0, licensed_devices: 0 });
onMounted(async () => {
  try {
    loading.value = true;
    const { data } = await fetchUserOverview();
    if (data) stats.value = data;
  } finally {
    loading.value = false;
  }
});
</script>

<template>
  <NSpace vertical size="large">
    <NAlert type="info" :title="$t('page.workspace.scopeTitle')">{{ $t('page.workspace.scopeTip') }}</NAlert>
    <NGrid :cols="1" responsive="screen" :x-gap="16" :y-gap="16" item-responsive>
      <NGi
        v-for="item in [
          { key: 'devices', label: $t('page.workspace.myDevices') },
          { key: 'sessions', label: $t('page.workspace.activeSessions') },
          { key: 'address_books', label: $t('page.workspace.addressBooks') },
          { key: 'security_events', label: $t('page.workspace.securityEvents') }
        ]"
        :key="item.key"
        span="1 s:2 m:1"
      >
        <NCard :loading="loading"><NStatistic :label="item.label" :value="(stats as any)[item.key]" /></NCard>
      </NGi>
    </NGrid>
  </NSpace>
</template>
