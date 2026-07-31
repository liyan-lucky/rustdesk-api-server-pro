<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NTag } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbList } from '@/service/api/address-book';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<Api.AddressBook.AddressBook[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const },
  { key: 'owner', title: $t('dataMap.ab.owner'), align: 'center' as const },
  { key: 'name', title: $t('dataMap.ab.name'), align: 'center' as const },
  { key: 'user_id', title: $t('dataMap.ab.user_id'), align: 'center' as const },
  { key: 'guid', title: $t('dataMap.ab.guid'), align: 'center' as const },
  { key: 'rule', title: $t('dataMap.ab.rule'), align: 'center' as const },
  { key: 'max_peer', title: $t('dataMap.ab.max_peer'), align: 'center' as const },
  {
    key: 'shared',
    title: $t('dataMap.ab.shared'),
    align: 'center' as const,
    render: (row: Api.AddressBook.AddressBook) => (
      <NTag type={row.shared ? 'success' : 'default'} size="small">
        {row.shared ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}
      </NTag>
    )
  },
  { key: 'created_at', title: $t('dataMap.audit.created_at'), align: 'center' as const }
];

async function loadData() {
  loading.value = true;
  try {
    const { data: res, error } = await fetchAbList({ current: currentPage.value, size: pageSize.value });
    if (!error && res) {
      data.value = res.records || [];
      total.value = res.total || 0;
    }
  } finally {
    loading.value = false;
  }
}

function handlePageChange(page: number) { currentPage.value = page; loadData(); }
function handlePageSizeChange(size: number) { pageSize.value = size; currentPage.value = 1; loadData(); }

onMounted(() => { loadData(); });
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.address_book_manage')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
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
