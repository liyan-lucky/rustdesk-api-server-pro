<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { NTag } from 'naive-ui';
import { $t } from '@/locales';
import { localStg } from '@/utils/storage';
import { fetchMyDevices, fetchUserPortalInfo } from '@/service/api/user-portal';

const router = useRouter();
const loading = ref(false);
const devices = ref<Api.Devices.Device[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);
const userInfo = ref<{ userId: string; userName: string; email?: string } | null>(null);

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

async function loadUserInfo() {
  const { data, error } = await fetchUserPortalInfo();
  if (!error && data) {
    userInfo.value = { userId: String(data.userId), userName: data.userName };
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

function handleLogout() {
  localStg.remove('token');
  localStg.remove('userType');
  router.push('/login/user-login');
}

onMounted(() => {
  const token = localStg.get('token');
  if (!token) {
    router.push('/login/user-login');
    return;
  }
  loadUserInfo();
  loadData();
});
</script>

<template>
  <div class="min-h-screen flex flex-col p-24px">
    <div class="flex justify-between items-center mb-24px">
      <div>
        <h2 class="text-24px font-600 text-primary">{{ $t('page.myDevices.title') }}</h2>
        <p v-if="userInfo" class="text-14px mt-4px text-gray-500">
          {{ $t('page.myDevices.welcome', { userName: userInfo.userName || userInfo.userId }) }}
        </p>
      </div>
      <NButton type="error" ghost @click="handleLogout">
        {{ $t('page.myDevices.logout') }}
      </NButton>
    </div>

    <NCard :bordered="false" size="small" class="flex-1-hidden card-wrapper">
      <NDataTable
        :columns="columns"
        :data="devices"
        size="small"
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
      />
    </NCard>
  </div>
</template>

<style scoped></style>
