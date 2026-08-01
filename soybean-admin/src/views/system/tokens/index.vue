<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NButton, NSpace, NPopconfirm, NTag } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { request } from '@/service/request';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<any[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const },
  { key: 'username', title: $t('dataMap.user.username'), align: 'center' as const },
  { key: 'rustdesk_id', title: 'RustDesk ID', align: 'center' as const },
  { key: 'device_os', title: $t('dataMap.token.device_os'), align: 'center' as const },
  { key: 'device_name', title: $t('dataMap.token.device_name'), align: 'center' as const },
  { key: 'token_hash', title: $t('dataMap.token.token_hash'), align: 'center' as const, ellipsis: { tooltip: true } },
  {
    key: 'is_admin',
    title: $t('dataMap.token.is_admin'),
    align: 'center' as const,
    render: (row: any) => <NTag type={row.is_admin ? 'warning' : 'default'} size="small">{row.is_admin ? 'Admin' : 'User'}</NTag>
  },
  {
    key: 'status',
    title: $t('dataMap.token.status'),
    align: 'center' as const,
    render: (row: any) => <NTag type={row.status === 1 ? 'success' : 'error'} size="small">{row.status === 1 ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}</NTag>
  },
  { key: 'expired', title: $t('dataMap.session.expired'), align: 'center' as const },
  { key: 'created_at', title: $t('dataMap.audit.created_at'), align: 'center' as const },
  {
    key: 'actions',
    title: $t('common.action'),
    align: 'center' as const,
    render: (row: any) => (
      <NSpace size="small" justify="center">
        <NPopconfirm onPositiveClick={() => handleKill(row)}>
          {{ default: () => $t('common.confirmDelete'), trigger: () => <NButton type="error" size="small" quaternary>{$t('page.user.sessions.kill')}</NButton> }}
        </NPopconfirm>
      </NSpace>
    )
  }
];

async function loadData() {
  loading.value = true;
  try {
    const { data: res, error } = await request({ url: '/tokens/list', params: { current: currentPage.value, size: pageSize.value } });
    if (!error && res) {
      data.value = res.records || [];
      total.value = res.total || 0;
    }
  } finally {
    loading.value = false;
  }
}

async function handleKill(row: any) {
  await request({ url: '/tokens/kill', method: 'post', data: { ids: [row.id] } });
  loadData();
}

function handlePageChange(page: number) { currentPage.value = page; loadData(); }
function handlePageSizeChange(size: number) { pageSize.value = size; currentPage.value = 1; loadData(); }

onMounted(() => { loadData(); });
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.system_tokens')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1200"
        :loading="loading"
        remote
        :row-key="(row: any) => row.id"
        :pagination="{ page: currentPage, pageSize: pageSize, itemCount: total, showSizePicker: true, pageSizes: [10, 20, 50], onChange: handlePageChange, onUpdatePageSize: handlePageSizeChange }"
        class="sm:h-full"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
