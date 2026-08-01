<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NButton, NSpace, NPopconfirm, NColorPicker, NSelect } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbTags, fetchAbTagDelete, fetchAbPersonal, fetchAbAllList } from '@/service/api/address-book';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<any[]>([]);
const currentAbGuid = ref('');
const abOptions = ref<{ label: string; value: string }[]>([]);

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

async function loadAbList() {
  const { data: res, error } = await fetchAbAllList();
  if (!error && res) {
    const list = Array.isArray(res) ? res : [];
    abOptions.value = list.map((ab: any) => ({ label: `${ab.name} (${ab.guid.slice(0, 8)}...)`, value: ab.guid }));
    if (!currentAbGuid.value && list.length > 0) {
      currentAbGuid.value = list[0].guid;
      loadData();
    }
  }
  if (abOptions.value.length === 0) {
    const { data: personalRes, error: personalErr } = await fetchAbPersonal();
    if (!personalErr && personalRes) {
      currentAbGuid.value = personalRes.guid;
      abOptions.value = [{ label: `My address book (${personalRes.guid.slice(0, 8)}...)`, value: personalRes.guid }];
      loadData();
    }
  }
}

async function loadData() {
  if (!currentAbGuid.value) return;
  loading.value = true;
  try {
    const { data: res, error } = await fetchAbTags(currentAbGuid.value);
    if (!error && res) {
      data.value = Array.isArray(res) ? res : [];
    }
  } finally {
    loading.value = false;
  }
}

async function handleDelete(row: any) {
  await fetchAbTagDelete(currentAbGuid.value, [row.name]);
  loadData();
}

function handleAbChange(guid: string) {
  currentAbGuid.value = guid;
  loadData();
}

onMounted(() => { loadAbList(); });
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.my-devices_tags')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <NSelect
          v-model:value="currentAbGuid"
          :options="abOptions"
          size="small"
          style="width: 260px"
          @update:value="handleAbChange"
        />
      </template>
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
