<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NButton, NSpace, NPopconfirm, NColorPicker } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbTags, fetchAbTagDelete } from '@/service/api/address-book';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<any[]>([]);
const currentAb = ref('personal');

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const },
  { key: 'ab_id', title: $t('dataMap.ab.ab_id'), align: 'center' as const },
  { key: 'name', title: $t('dataMap.ab.tagName'), align: 'center' as const },
  {
    key: 'color',
    title: $t('dataMap.ab.tagColor'),
    align: 'center' as const,
    render: (row: any) => row.color ? <NColorPicker value={`#${row.color.toString(16).padStart(6, '0')}`} size="small" disabled modes={['hex']} /> : '-'
  },
  { key: 'created_at', title: $t('dataMap.audit.created_at'), align: 'center' as const },
  { key: 'updated_at', title: $t('dataMap.ab.updated_at'), align: 'center' as const },
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
    const { data: res, error } = await fetchAbTags(currentAb.value);
    if (!error && res) {
      data.value = res.data || res || [];
    }
  } finally {
    loading.value = false;
  }
}

async function handleDelete(row: any) {
  await fetchAbTagDelete(currentAb.value, [row.name]);
  loadData();
}

onMounted(() => { loadData(); });
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.address_book_tags')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="900"
        :loading="loading"
        :row-key="(row: any) => row.id"
        class="sm:h-full"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
