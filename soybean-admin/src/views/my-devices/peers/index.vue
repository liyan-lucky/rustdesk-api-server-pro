<script setup lang="tsx">
import { onMounted, reactive, ref } from 'vue';
import { NTag, NButton, NSpace, NPopconfirm } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbPeers, fetchAbPeerAdd, fetchAbPeerUpdate, fetchAbPeerDelete, fetchAbAllList, fetchAbPersonal } from '@/service/api/address-book';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<Api.AddressBook.Peer[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);
const currentAbGuid = ref('');
const abOptions = ref<{ label: string; value: string }[]>([]);
const modalVisible = ref(false);
const editing = ref(false);
const form = reactive({ id: '', username: '', hostname: '', platform: '', alias: '', hash: '', password: '', note: '', tagText: '' });

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
        <NButton type="primary" size="small" quaternary onClick={() => openEdit(row)}>{$t('common.edit')}</NButton>
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
    if (currentAbGuid.value) params.ab = currentAbGuid.value;
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
  const { error } = await fetchAbPeerDelete(currentAbGuid.value, [row.id]);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    loadData();
  }
}

function handlePageChange(page: number) { currentPage.value = page; loadData(); }
function handlePageSizeChange(size: number) { pageSize.value = size; currentPage.value = 1; loadData(); }

function openAdd() {
  editing.value = false;
  Object.assign(form, { id: '', username: '', hostname: '', platform: '', alias: '', hash: '', password: '', note: '', tagText: '' });
  modalVisible.value = true;
}

function openEdit(row: Api.AddressBook.Peer) {
  editing.value = true;
  Object.assign(form, row, { password: '', tagText: (row.tags || []).join(', ') });
  modalVisible.value = true;
}

async function submit() {
  if (!form.id.trim()) return window.$message?.warning($t('dataMap.ab.deviceIdRequired'));
  const payload = { ...form, tags: form.tagText.split(',').map(item => item.trim()).filter(Boolean) };
  const result = editing.value ? await fetchAbPeerUpdate(currentAbGuid.value, payload) : await fetchAbPeerAdd(currentAbGuid.value, payload);
  if (!result.error) {
    window.$message?.success(editing.value ? $t('common.updateSuccess') : $t('common.addSuccess'));
    modalVisible.value = false;
    loadData();
  }
}

function handleAbChange() { currentPage.value = 1; loadData(); }

onMounted(async () => {
  const { data: books } = await fetchAbAllList();
  if (books?.length) {
    abOptions.value = books.map(book => ({ label: `${book.name} (${book.owner})`, value: book.guid }));
    currentAbGuid.value = books[0].guid;
  } else {
    const { data: personal } = await fetchAbPersonal();
    if (personal) { currentAbGuid.value = personal.guid; abOptions.value = [{ label: $t('dataMap.ab.personal'), value: personal.guid }]; }
  }
  loadData();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.my-devices_peers')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra><NSpace><NSelect v-model:value="currentAbGuid" :options="abOptions" size="small" style="width: 260px" @update:value="handleAbChange" /><NButton type="primary" size="small" :disabled="!currentAbGuid" @click="openAdd">{{ $t('common.add') }}</NButton></NSpace></template>
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
    <NModal v-model:show="modalVisible" preset="card" :title="editing ? $t('common.edit') : $t('common.add')" class="max-w-620px">
      <NForm label-placement="left" label-width="110">
        <NFormItem :label="$t('dataMap.ab.rustdesk_id')"><NInput v-model:value="form.id" :disabled="editing" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.alias')"><NInput v-model:value="form.alias" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.username')"><NInput v-model:value="form.username" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.hostname')"><NInput v-model:value="form.hostname" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.platform')"><NInput v-model:value="form.platform" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.tags')"><NInput v-model:value="form.tagText" :placeholder="$t('dataMap.ab.tagsHint')" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.note')"><NInput v-model:value="form.note" type="textarea" /></NFormItem>
      </NForm>
      <template #footer><NSpace justify="end"><NButton @click="modalVisible = false">{{ $t('common.cancel') }}</NButton><NButton type="primary" @click="submit">{{ $t('common.confirm') }}</NButton></NSpace></template>
    </NModal>
  </div>
</template>

<style scoped></style>
