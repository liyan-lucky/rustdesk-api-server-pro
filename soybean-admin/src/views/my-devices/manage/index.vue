<script setup lang="tsx">
import { computed, onMounted, reactive, ref } from 'vue';
import { NButton, NPopconfirm, NSpace, NTag, NInput, NSelect } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbAllList, fetchAbSharedAdd, fetchAbSharedDelete, fetchAbSharedUpdate } from '@/service/api/address-book';
import { fetchUserList } from '@/service/api/user_management';
import { useAuthStore } from '@/store/modules/auth';
import { downloadCsv, parseCsv } from '@/utils/csv';

const appStore = useAppStore();
const authStore = useAuthStore();
const isAdmin = computed(() => authStore.userInfo.roles.includes('R_SUPER'));
const loading = ref(false);
const data = ref<Api.AddressBook.AddressBook[]>([]);
const modalVisible = ref(false);
const editing = ref(false);
const form = reactive({ guid: '', user_id: 0, name: '', note: '', rule: 1, max_peer: 0 });
const userOptions = ref<{ label: string; value: number }[]>([]);
const importInput = ref<HTMLInputElement>();
const filters = reactive({ owner: '', name: '', note: '', guid: '' });
const filteredData = computed(() => data.value.filter(row => Object.entries(filters).every(([key, value]) => !value || String((row as any)[key] || '').toLowerCase().includes(value.toLowerCase()))));
const filterTitle = (label: string, key: keyof typeof filters) => () => <div class="min-w-130px"><div>{label}</div><NInput value={filters[key]} size="tiny" clearable placeholder={$t('common.keywordSearch')} onUpdateValue={value => { filters[key] = value; }} /></div>;

const columns = [
  { key: 'owner', title: filterTitle($t('dataMap.ab.owner'), 'owner'), align: 'center' as const },
  { key: 'name', title: filterTitle($t('dataMap.ab.name'), 'name'), align: 'center' as const },
  { key: 'note', title: filterTitle($t('dataMap.ab.note'), 'note'), align: 'center' as const },
  { key: 'guid', title: filterTitle($t('dataMap.ab.guid'), 'guid'), align: 'center' as const },
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
        {(!row.created_by_admin || isAdmin.value) && <NPopconfirm onPositiveClick={() => remove(row)}>{{ default: () => $t('common.confirmDelete'), trigger: () => <NButton size="small" quaternary type="error">{$t('common.delete')}</NButton> }}</NPopconfirm>}
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
  Object.assign(form, { guid: '', user_id: userOptions.value[0]?.value || 0, name: '', note: '', rule: 1, max_peer: 0 });
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

function exportData() { downloadCsv('address-books.csv', filteredData.value as any[], ['user_id', 'owner', 'guid', 'name', 'note', 'rule', 'max_peer', 'created_by_admin']); }
async function importData(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0]; if (!file) return;
  const rows = await parseCsv(file); let success = 0;
  for (const row of rows) { if (!row.name) continue; const result = await fetchAbSharedAdd({ user_id: Number(row.user_id) || 0, name: row.name, note: row.note, rule: Number(row.rule) || 1, max_peer: Number(row.max_peer) || 0 } as any); if (!result.error) success += 1; }
  window.$message?.success(`${success}/${rows.length}`); (event.target as HTMLInputElement).value = ''; loadData();
}

onMounted(async () => {
  if (isAdmin.value) { const { data: users } = await fetchUserList({ current: 1, size: 1000 }); userOptions.value = (users?.records || []).filter(user => user.id !== undefined).map(user => ({ label: user.username, value: user.id! })); }
  await loadData();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.my-devices_manage')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra><NSpace><input ref="importInput" type="file" accept=".csv,text/csv" class="hidden" @change="importData" /><NButton size="small" @click="importInput?.click()">{{ $t('common.import') }}</NButton><NButton size="small" @click="exportData">{{ $t('common.export') }}</NButton><NButton type="primary" size="small" @click="openAdd">{{ $t('common.add') }}</NButton></NSpace></template>
      <NDataTable :columns="columns" :data="filteredData" size="small" :flex-height="!appStore.isMobile" :scroll-x="1200" :loading="loading" :row-key="(row: any) => row.id" class="sm:h-full" />
    </NCard>
    <NModal v-model:show="modalVisible" preset="card" :title="editing ? $t('common.edit') : $t('common.add')" class="max-w-560px">
      <NForm label-placement="left" label-width="120">
        <NFormItem v-if="isAdmin && !editing" :label="$t('dataMap.ab.owner')"><NSelect v-model:value="form.user_id" :options="userOptions" filterable /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.name')"><NInput v-model:value="form.name" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.note')"><NInput v-model:value="form.note" type="textarea" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.rule')"><NSelect v-model:value="form.rule" :options="[{ label: $t('dataMap.ab.read'), value: 1 }, { label: $t('dataMap.ab.readWrite'), value: 2 }, { label: $t('dataMap.ab.fullControl'), value: 3 }]" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.max_peer')"><NInputNumber v-model:value="form.max_peer" :min="0" /></NFormItem>
      </NForm>
      <template #footer><NSpace justify="end"><NButton @click="modalVisible = false">{{ $t('common.cancel') }}</NButton><NButton type="primary" @click="submit">{{ $t('common.confirm') }}</NButton></NSpace></template>
    </NModal>
  </div>
</template>
