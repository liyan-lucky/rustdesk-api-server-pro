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
const providers = ref<any[]>([]);

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const },
  { key: 'user_id', title: $t('dataMap.ab.user_id'), align: 'center' as const },
  { key: 'provider', title: $t('dataMap.oauth.provider'), align: 'center' as const },
  { key: 'subject', title: $t('dataMap.oauth.subject'), align: 'center' as const },
  { key: 'email', title: $t('dataMap.oauth.email'), align: 'center' as const },
  { key: 'name', title: $t('dataMap.oauth.name'), align: 'center' as const },
  {
    key: 'is_admin',
    title: $t('dataMap.token.is_admin'),
    align: 'center' as const,
    render: (row: any) => <NTag type={row.is_admin ? 'warning' : 'default'} size="small">{row.is_admin ? $t('dataMap.token.is_admin') : $t('dataMap.user.statusLabel.normal')}</NTag>
  },
  {
    key: 'status',
    title: $t('dataMap.token.status'),
    align: 'center' as const,
    render: (row: any) => <NTag type={row.status === 1 ? 'success' : 'error'} size="small">{row.status === 1 ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}</NTag>
  },
  { key: 'last_login_at', title: $t('dataMap.oauth.last_login_at'), align: 'center' as const },
  { key: 'created_at', title: $t('dataMap.audit.created_at'), align: 'center' as const },
  {
    key: 'actions',
    title: $t('common.action'),
    align: 'center' as const,
    render: (row: any) => (
      <NSpace size="small" justify="center">
        <NPopconfirm onPositiveClick={() => handleDelete(row)}>
          {{ default: () => $t('common.confirmDelete'), trigger: () => <NButton type="error" size="small" quaternary>{$t('common.delete')}</NButton> }}
        </NPopconfirm>
      </NSpace>
    )
  }
];

async function loadData() {
  loading.value = true;
  try {
    const { data: res, error } = await request({ url: '/oauth/accounts', params: { current: currentPage.value, size: pageSize.value } });
    if (!error && res) {
      data.value = res.records || [];
      total.value = res.total || 0;
    }
  } finally {
    loading.value = false;
  }
}

async function loadProviders() {
  const { data: res, error } = await request({ url: '/oauth/providers' });
  if (!error && res) {
    providers.value = Array.isArray(res) ? res : [];
  }
}

async function handleDelete(row: any) {
  await request({ url: `/oauth/account/${row.id}`, method: 'delete' });
  loadData();
}

function handlePageChange(page: number) { currentPage.value = page; loadData(); }
function handlePageSizeChange(size: number) { pageSize.value = size; currentPage.value = 1; loadData(); }

onMounted(() => { loadData(); loadProviders(); });
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.system_oauth')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <NSpace size="small">
          <NTag v-for="p in providers" :key="p.name" :type="p.enabled ? 'success' : 'default'" size="small">
            {{ p.displayName || p.name }} ({{ p.type }})
          </NTag>
        </NSpace>
      </template>
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
