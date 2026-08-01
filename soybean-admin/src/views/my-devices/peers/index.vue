<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NTag, NButton, NSpace, NPopconfirm } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbPeers, fetchAbPeerDelete, fetchAbPersonal } from '@/service/api/address-book';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<Api.AddressBook.Peer[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);
const personalGuid = ref('');

const columns = [
  { key: 'id', title: $t('dataMap.ab.rustdesk_id'), align: 'center' as const },
  { key: 'username', title: $t('dataMap.ab.username'), align: 'center' as const },
  { key: 'hostname', title: $t('dataMap.ab.hostname'), align: 'center' as const },
  {
    key: 'tags',
    title: $t('dataMap.ab.tags'),
    align: 'center' as const,
    render: (row: Api.AddressBook.Peer) => {
      if (!row.tags || row.tags.length === 0) return '-';
      return row.tags.map((tag: string) => <NTag size="small" class="mr-4px">{tag}</NTag>);
    }
  },
  { key: 'alias', title: $t('dataMap.ab.alias'), align: 'center' as const },
  { key: 'hash', title: $t('dataMap.ab.hash'), align: 'center' as const },
  {
    key: 'actions',
    title: $t('common.action'),
    align: 'center' as const,
    render: (row: Api.AddressBook.Peer) => (
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
    const params: Record<string, any> = { current: currentPage.value, size: pageSize.value };
    if (personalGuid.value) params.ab = personalGuid.value;
    const { data: res, error } = await fetchAbPeers(params);
    if (!error && res) {
      data.value = res.records || [];
      total.value = res.total || 0;
    }
  } finally {
    loading.value = false;
  }
}

async function handleDelete(row: Api.AddressBook.Peer) {
  const guid = personalGuid.value || 'personal';
  await fetchAbPeerDelete(guid, [row.id]);
  loadData();
}

function handlePageChange(page: number) { currentPage.value = page; loadData(); }
function handlePageSizeChange(size: number) { pageSize.value = size; currentPage.value = 1; loadData(); }

onMounted(async () => {
  const { data: res } = await fetchAbPersonal();
  if (res) personalGuid.value = res.guid;
  loadData();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.my-devices_peers')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
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
