<script setup lang="tsx">
import { computed, onMounted, reactive, ref } from 'vue';
import { NButton, NColorPicker, NInput, NPopconfirm, NSelect, NSpace } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';
import { fetchUserList } from '@/service/api/user_management';
import {
  fetchAbAllList,
  fetchAbAllTags,
  fetchAbPersonal,
  fetchAbTagAdd,
  fetchAbTagDelete,
  fetchAbTagRename,
  fetchAbTagUpdate
} from '@/service/api/address-book';
import { downloadCsv, parseCsv } from '@/utils/csv';
import { localizeAddressBookName } from '@/utils/address-book';

const appStore = useAppStore();
const authStore = useAuthStore();
const isAdmin = computed(() => authStore.userInfo.roles.includes('R_SUPER'));
const loading = ref(false);
const data = ref<any[]>([]);
const abOptions = ref<{ label: string; value: string; userId: number }[]>([]);
const userOptions = ref<{ label: string; value: number }[]>([]);
const importInput = ref<HTMLInputElement>();
const modalVisible = ref(false);
const editing = ref(false);
const form = reactive({ user_id: 0, ab_guid: '', old: '', name: '', color: '#4098fc' });
const selectableAbOptions = computed(() =>
  isAdmin.value ? abOptions.value.filter(option => option.userId === form.user_id) : abOptions.value
);
const filters = reactive({ ab_name: '', owner: '', name: '' });
const filteredData = computed(() =>
  data.value.filter(row =>
    Object.entries(filters).every(
      ([key, value]) =>
        !value ||
        String(key === 'ab_name' ? localizeAddressBookName(row[key]) : row[key] || '')
          .toLowerCase()
          .includes(value.toLowerCase())
    )
  )
);
const filterTitle = (label: string, key: keyof typeof filters) => () => (
  <div class="min-w-130px">
    <div>{label}</div>
    <NInput
      value={filters[key]}
      size="tiny"
      clearable
      placeholder={$t('common.keywordSearch')}
      onUpdateValue={value => {
        filters[key] = value;
      }}
    />
  </div>
);

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const },
  {
    key: 'ab_name',
    title: filterTitle($t('dataMap.ab.name'), 'ab_name'),
    align: 'center' as const,
    render: (row: any) => localizeAddressBookName(row.ab_name)
  },
  { key: 'owner', title: filterTitle($t('dataMap.ab.owner'), 'owner'), align: 'center' as const },
  { key: 'name', title: filterTitle($t('dataMap.ab.tagName'), 'name'), align: 'center' as const },
  {
    key: 'color',
    title: $t('dataMap.ab.tagColor'),
    align: 'center' as const,
    render: (row: any) =>
      row.color ? <NColorPicker value={colorToHex(row.color)} size="small" disabled modes={['hex']} /> : '-'
  },
  {
    key: 'actions',
    title: $t('common.action'),
    align: 'center' as const,
    render: (row: any) => (
      <NSpace size="small" justify="center">
        <NButton type="primary" size="small" quaternary onClick={() => openEdit(row)}>
          {$t('common.edit')}
        </NButton>
        <NPopconfirm onPositiveClick={() => handleDelete(row)}>
          {{
            default: () => $t('common.confirmDelete'),
            trigger: () => (
              <NButton type="error" size="small" quaternary>
                {$t('common.delete')}
              </NButton>
            )
          }}
        </NPopconfirm>
      </NSpace>
    )
  }
];

async function loadAbList() {
  const { data: res, error } = await fetchAbAllList();
  if (!error && res) {
    const list = Array.isArray(res) ? res : [];
    abOptions.value = list.map((ab: any) => ({
      label: `${localizeAddressBookName(ab.name)} (${ab.owner})`,
      value: ab.guid,
      userId: ab.user_id || 0
    }));
  }
  if (abOptions.value.length === 0) {
    const { data: personalRes, error: personalErr } = await fetchAbPersonal();
    if (!personalErr && personalRes) {
      abOptions.value = [
        {
          label: `${$t('dataMap.ab.personal')} (${personalRes.guid.slice(0, 8)}...)`,
          value: personalRes.guid,
          userId: 0
        }
      ];
    }
  }
}

async function loadData() {
  loading.value = true;
  try {
    const { data: res, error } = await fetchAbAllTags();
    if (!error && res) {
      data.value = Array.isArray(res) ? res : [];
    }
  } finally {
    loading.value = false;
  }
}

async function handleDelete(row: any) {
  const { error } = await fetchAbTagDelete(row.ab_guid, [row.name]);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    loadData();
  }
}

function colorToHex(value: number) {
  return `#${(value & 0xffffff).toString(16).padStart(6, '0')}`;
}
function colorToInt(value: string) {
  return (0xff000000 | Number.parseInt(value.slice(1), 16)) >>> 0;
}

function openAdd() {
  editing.value = false;
  Object.assign(form, {
    user_id: 0,
    ab_guid: isAdmin.value ? '' : abOptions.value[0]?.value || '',
    old: '',
    name: '',
    color: '#4098fc'
  });
  modalVisible.value = true;
}

function openEdit(row: any) {
  editing.value = true;
  const book = abOptions.value.find(option => option.value === row.ab_guid);
  Object.assign(form, {
    user_id: book?.userId || 0,
    ab_guid: row.ab_guid,
    old: row.name,
    name: row.name,
    color: colorToHex(row.color)
  });
  modalVisible.value = true;
}

async function submit() {
  if (!form.ab_guid) return window.$message?.warning($t('dataMap.ab.nameRequired'));
  if (!form.name.trim()) return window.$message?.warning($t('dataMap.ab.nameRequired'));
  let error: unknown;
  if (!editing.value) {
    ({ error } = await fetchAbTagAdd(form.ab_guid, { name: form.name, color: colorToInt(form.color) }));
  } else {
    if (form.old !== form.name) ({ error } = await fetchAbTagRename(form.ab_guid, { old: form.old, new: form.name }));
    if (!error) ({ error } = await fetchAbTagUpdate(form.ab_guid, { name: form.name, color: colorToInt(form.color) }));
  }
  if (!error) {
    window.$message?.success(editing.value ? $t('common.updateSuccess') : $t('common.addSuccess'));
    modalVisible.value = false;
    loadData();
  }
}

function exportData() {
  downloadCsv('address-book-tags.csv', filteredData.value, ['ab_guid', 'ab_name', 'owner', 'name', 'color']);
}
async function importData(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0];
  if (!file) return;
  const rows = await parseCsv(file);
  let success = 0;
  for (const row of rows) {
    const guid = row.ab_guid || '';
    if (!guid || !row.name) continue;
    const result = await fetchAbTagAdd(guid, { name: row.name, color: Number(row.color) || colorToInt('#4098fc') });
    if (!result.error) success += 1;
  }
  window.$message?.success(`${success}/${rows.length}`);
  (event.target as HTMLInputElement).value = '';
  loadData();
}

onMounted(async () => {
  await loadAbList();
  if (isAdmin.value) {
    const { data: users } = await fetchUserList({ current: 1, size: 1000 });
    userOptions.value = (users?.records || [])
      .filter(user => user.id !== undefined)
      .map(user => ({ label: user.username, value: user.id! }));
  }
  loadData();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('route.my-devices_tags')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <NSpace>
          <input ref="importInput" type="file" accept=".csv,text/csv" class="hidden" @change="importData" />
          <NButton size="small" @click="importInput?.click()">{{ $t('common.import') }}</NButton>
          <NButton size="small" @click="exportData">{{ $t('common.export') }}</NButton>
          <NButton type="primary" size="small" :disabled="!abOptions.length" @click="openAdd">
            {{ $t('common.add') }}
          </NButton>
        </NSpace>
      </template>
      <NDataTable
        :columns="columns"
        :data="filteredData"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="900"
        :loading="loading"
        :row-key="(row: any) => row.id"
        class="sm:h-full"
      />
    </NCard>
    <NModal
      v-model:show="modalVisible"
      preset="card"
      :title="editing ? $t('common.edit') : $t('common.add')"
      class="max-w-520px"
    >
      <NForm label-placement="left" label-width="100">
        <NFormItem v-if="isAdmin && !editing" :label="$t('dataMap.ab.owner')">
          <NSelect v-model:value="form.user_id" :options="userOptions" filterable @update:value="form.ab_guid = ''" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.ab.name')">
          <NSelect
            v-model:value="form.ab_guid"
            :options="selectableAbOptions"
            :disabled="editing || (isAdmin && !form.user_id)"
          />
        </NFormItem>
        <NFormItem :label="$t('dataMap.ab.tagName')"><NInput v-model:value="form.name" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.tagColor')">
          <NColorPicker v-model:value="form.color" :show-alpha="false" :modes="['hex']" />
        </NFormItem>
      </NForm>
      <template #footer>
        <NSpace justify="end">
          <NButton @click="modalVisible = false">{{ $t('common.cancel') }}</NButton>
          <NButton type="primary" @click="submit">{{ $t('common.confirm') }}</NButton>
        </NSpace>
      </template>
    </NModal>
  </div>
</template>

<style scoped></style>
