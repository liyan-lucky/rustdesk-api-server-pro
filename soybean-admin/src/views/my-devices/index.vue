<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NTag } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { localStg } from '@/utils/storage';
import { fetchMyDevices } from '@/service/api/user-portal';

const appStore = useAppStore();
const loading = ref(false);
const devices = ref<Api.Devices.Device[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);

const columns = [
  { key: 'rustdesk_id', title: $t('dataMap.device.rustdesk_id'), align: 'center' as const },
  { key: 'hostname', title: $t('dataMap.device.hostname'), align: 'center' as const },
  { key: 'username', title: $t('dataMap.device.username'), align: 'center' as const },
  { key: 'os', title: $t('dataMap.device.os'), align: 'center' as const },
  { key: 'version', title: $t('dataMap.device.version'), align: 'center' as const },
  {
    key: 'is_online',
    title: $t('page.myDevices.status'),
    align: 'center' as const,
    render: (row: Api.Devices.Device) =>
      row.is_online ? (
        <NTag type="success" size="small">{$t('page.myDevices.online')}</NTag>
      ) : (
        <NTag type="default" size="small">{$t('page.myDevices.offline')}</NTag>
      )
  },
  { key: 'conns', title: $t('page.myDevices.conns'), align: 'center' as const },
  { key: 'updated_at', title: $t('page.myDevices.lastSync'), align: 'center' as const }
];

async function loadData() {
  loading.value = true;
  try {
    const { data, error } = await fetchMyDevices({
      current: currentPage.value,
      size: pageSize.value
    });
    if (!error && data) {
      devices.value = data.records || [];
      total.value = data.total || 0;
    }
  } finally {
    loading.value = false;
  }
}

function handlePageChange(page: number) {
  currentPage.value = page;
  loadData();
}

function handlePageSizeChange(size: number) {
  pageSize.value = size;
  currentPage.value = 1;
  loadData();
}

onMounted(() => {
  const token = localStg.get('token');
  if (!token) {
    window.location.href = '/#/login/user-login';
    return;
  }
  loadData();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-'16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.my-devices')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <NDataTable
        :columns="columns"
        :data="devices"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="962"
        :loading="loading"
        remote
        :row-key="(row: any) => row.rustdesk_id"
        :pagination="{
          page: currentPage,
          pageSize: pageSize,
          itemCount: total,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
          onChange: handlePageChange,
          onUpdatePageSize: handlePageSizeChange
        }"
        class="sm:h-full"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
