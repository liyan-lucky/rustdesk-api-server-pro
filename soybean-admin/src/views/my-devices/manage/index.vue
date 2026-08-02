<script setup lang="tsx">
import { onMounted, reactive, ref } from 'vue';
import { NButton, NPopconfirm, NSpace, NTag } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbAllList, fetchAbSharedAdd, fetchAbSharedDelete, fetchAbSharedUpdate } from '@/service/api/address-book';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<Api.AddressBook.AddressBook[]>([]);
const modalVisible = ref(false);
const editing = ref(false);
const form = reactive({ guid: '', name: '', note: '', rule: 1, max_peer: 0 });

const columns = [
  { key: 'owner', title: $t('dataMap.ab.owner'), align: 'center' as const },
  { key: 'name', title: $t('dataMap.ab.name'), align: 'center' as const },
  { key: 'note', title: $t('dataMap.ab.note'), align: 'center' as const },
  { key: 'guid', title: $t('dataMap.ab.guid'), align: 'center' as const },
  { key: 'rule', title: $t('dataMap.ab.rule'), align: 'center' as const },
  { key: 'max_peer', title: $t('dataMap.ab.max_peer'), align: 'center' as const },
  {
    key: 'shared', title: $t('dataMap.ab.shared'), align: 'center' as const,
    render: (row: Api.AddressBook.AddressBook) => <NTag type={row.shared ? 'success' : 'default'}>{row.shared ? $t('common.yesOrNo.yes') : $t('dataMap.ab.personal')}</NTag>
  },
  {
    key: 'actions', title: $t('common.action'), align: 'center' as const,
    render: (row: Api.AddressBook.AddressBook) => row.shared ? (
      <NSpace justify="center">
        <NButton size="small" quaternary type="primary" onClick={() => openEdit(row)}>{$t('common.edit')}</NButton>
        <NPopconfirm onPositiveClick={() => remove(row)}>
          {{ default: () => $t('common.confirmDelete'), trigger: () => <NButton size="small" quaternary type="error">{$t('common.delete')}</NButton> }}
        </NPopconfirm>
      </NSpace>
    ) : <NTag size="small">{$t('dataMap.ab.personalReadOnly')}</NTag>
  }
];

async function loadData() {
  loading.value = true;
  try {
    const { data: res, error } = await fetchAbAllList();
    if (!error && res) data.value = res;
  } finally { loading.value = false; }
}

function openAdd() {
  editing.value = false;
  Object.assign(form, { guid: '', name: '', note: '', rule: 1, max_peer: 0 });
  modalVisible.value = true;
}

function openEdit(row: Api.AddressBook.AddressBook) {
  editing.value = true;
  Object.assign(form, row);
  modalVisible.value = true;
}

async function submit() {
  if (!form.name.trim()) return window.$message?.warning($t('dataMap.ab.nameRequired'));
  const result = editing.value ? await fetchAbSharedUpdate(form) : await fetchAbSharedAdd(form);
  if (!result.error) {
    window.$message?.success(editing.value ? $t('common.updateSuccess') : $t('common.addSuccess'));
    modalVisible.value = false;
    loadData();
  }
}

async function remove(row: Api.AddressBook.AddressBook) {
  const { error } = await fetchAbSharedDelete([row.guid]);
  if (!error) { window.$message?.success($t('common.deleteSuccess')); loadData(); }
}

onMounted(loadData);
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.my-devices_manage')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra><NButton type="primary" size="small" @click="openAdd">{{ $t('common.add') }}</NButton></template>
      <NDataTable :columns="columns" :data="data" size="small" :flex-height="!appStore.isMobile" :scroll-x="1200" :loading="loading" :row-key="(row: any) => row.id" class="sm:h-full" />
    </NCard>
    <NModal v-model:show="modalVisible" preset="card" :title="editing ? $t('common.edit') : $t('common.add')" class="max-w-560px">
      <NForm label-placement="left" label-width="120">
        <NFormItem :label="$t('dataMap.ab.name')"><NInput v-model:value="form.name" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.note')"><NInput v-model:value="form.note" type="textarea" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.rule')"><NSelect v-model:value="form.rule" :options="[{ label: $t('dataMap.ab.read'), value: 1 }, { label: $t('dataMap.ab.readWrite'), value: 2 }, { label: $t('dataMap.ab.fullControl'), value: 3 }]" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.max_peer')"><NInputNumber v-model:value="form.max_peer" :min="0" /></NFormItem>
      </NForm>
      <template #footer><NSpace justify="end"><NButton @click="modalVisible = false">{{ $t('common.cancel') }}</NButton><NButton type="primary" @click="submit">{{ $t('common.confirm') }}</NButton></NSpace></template>
    </NModal>
  </div>
</template>
