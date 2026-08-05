<script setup lang="tsx">
import { computed, onMounted, reactive, ref } from 'vue';
import { NButton, NInput, NPopconfirm, NSelect, NSpace, NTag } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';
import { fetchUserList } from '@/service/api/user_management';
import {
  fetchAbAllList,
  fetchAbPeerAdd,
  fetchAbPeerDelete,
  fetchAbPeerUpdate,
  fetchAbPeers,
  fetchAbPersonal
} from '@/service/api/address-book';
import { downloadCsv, parseCsv } from '@/utils/csv';
import { localizeAddressBookName, normalizeAddressBookFilter } from '@/utils/address-book';

const appStore = useAppStore();
const authStore = useAuthStore();
const isAdmin = computed(() => authStore.userInfo.roles.includes('R_SUPER'));
const loading = ref(false);
const data = ref<Api.AddressBook.Peer[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);
const abOptions = ref<{ label: string; value: string; userId: number }[]>([]);
const userOptions = ref<{ label: string; value: number }[]>([]);
const importInput = ref<HTMLInputElement>();
const modalVisible = ref(false);
const editing = ref(false);
const form = reactive({
  user_id: 0,
  ab_guid: '',
  id: '',
  username: '',
  hostname: '',
  platform: '',
  alias: '',
  hash: '',
  password: '',
  note: '',
  tagText: ''
});
const selectableAbOptions = computed(() =>
  isAdmin.value ? abOptions.value.filter(option => option.userId === form.user_id) : abOptions.value
);
const filters = reactive({
  ab_name: '',
  id: '',
  username: '',
  hostname: '',
  platform: '',
  alias: '',
  tags: '',
  hash: '',
  note: ''
});
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
        currentPage.value = 1;
        loadData();
      }}
    />
  </div>
);

const columns = [
  {
    key: 'ab_name',
    title: filterTitle($t('dataMap.ab.name'), 'ab_name'),
    align: 'center' as const,
    render: (row: Api.AddressBook.Peer) => localizeAddressBookName(row.ab_name)
  },
  { key: 'id', title: filterTitle($t('dataMap.ab.rustdesk_id'), 'id'), align: 'center' as const },
  { key: 'username', title: filterTitle($t('dataMap.ab.username'), 'username'), align: 'center' as const },
  { key: 'hostname', title: filterTitle($t('dataMap.ab.hostname'), 'hostname'), align: 'center' as const },
  {
    key: 'tags',
    title: filterTitle($t('dataMap.ab.tags'), 'tags'),
    align: 'center' as const,
    render: (row: Api.AddressBook.Peer) => {
      if (!row.tags || row.tags.length === 0) return '-';
      return row.tags.map((tag: string) => (
        <NTag size="small" class="mr-4px">
          {tag}
        </NTag>
      ));
    }
  },
  { key: 'platform', title: filterTitle($t('dataMap.ab.platform'), 'platform'), align: 'center' as const },
  { key: 'alias', title: filterTitle($t('dataMap.ab.alias'), 'alias'), align: 'center' as const },
  { key: 'hash', title: filterTitle($t('dataMap.ab.hash'), 'hash'), align: 'center' as const },
  { key: 'note', title: filterTitle($t('dataMap.ab.note'), 'note'), align: 'center' as const },
  {
    key: 'actions',
    title: $t('common.action'),
    align: 'center' as const,
    render: (row: Api.AddressBook.Peer) => (
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

async function loadData() {
  loading.value = true;
  try {
    const params: Record<string, any> = {
      current: currentPage.value,
      size: pageSize.value,
      ...filters,
      ab_name: normalizeAddressBookFilter(filters.ab_name)
    };
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
  const { error } = await fetchAbPeerDelete(row.ab_guid, [row.id]);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    loadData();
  }
}

function handlePageChange(page: number) {
  currentPage.value = page;
  loadData();
}
function handlePageSizeChange(size: number) {
  pageSize.value = size;
  currentPage.value = 1;
  loadData();
}

function openAdd() {
  editing.value = false;
  Object.assign(form, {
    user_id: 0,
    ab_guid: isAdmin.value ? '' : abOptions.value[0]?.value || '',
    id: '',
    username: '',
    hostname: '',
    platform: '',
    alias: '',
    hash: '',
    password: '',
    note: '',
    tagText: ''
  });
  modalVisible.value = true;
}

function openEdit(row: Api.AddressBook.Peer) {
  editing.value = true;
  const book = abOptions.value.find(option => option.value === row.ab_guid);
  Object.assign(form, row, { user_id: book?.userId || 0, password: '', tagText: (row.tags || []).join(', ') });
  modalVisible.value = true;
}

async function submit() {
  if (!form.ab_guid) return window.$message?.warning($t('dataMap.ab.nameRequired'));
  if (!form.id.trim()) return window.$message?.warning($t('dataMap.ab.deviceIdRequired'));
  const payload = {
    ...form,
    tags: form.tagText
      .split(',')
      .map(item => item.trim())
      .filter(Boolean)
  };
  const result = editing.value
    ? await fetchAbPeerUpdate(form.ab_guid, payload)
    : await fetchAbPeerAdd(form.ab_guid, payload);
  if (!result.error) {
    window.$message?.success(editing.value ? $t('common.updateSuccess') : $t('common.addSuccess'));
    modalVisible.value = false;
    loadData();
  }
}

async function exportData() {
  const { data: result, error } = await fetchAbPeers({ current: 1, size: Math.max(total.value, 1000), ...filters });
  if (!error && result)
    downloadCsv('contacts.csv', result.records as any[], [
      'ab_guid',
      'ab_name',
      'owner',
      'id',
      'username',
      'hostname',
      'platform',
      'alias',
      'tags',
      'note'
    ]);
}
async function importData(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0];
  if (!file) return;
  const rows = await parseCsv(file);
  let success = 0;
  for (const row of rows) {
    const guid = row.ab_guid || abOptions.value.find(option => option.label.startsWith(row.ab_name))?.value || '';
    if (!guid || !row.id) continue;
    const result = await fetchAbPeerAdd(guid, { ...row, tags: (row.tags || '').split('|').filter(Boolean) });
    if (!result.error) success += 1;
  }
  window.$message?.success(`${success}/${rows.length}`);
  (event.target as HTMLInputElement).value = '';
  loadData();
}

onMounted(async () => {
  const { data: books } = await fetchAbAllList();
  if (books?.length) {
    abOptions.value = books.map(book => ({
      label: `${localizeAddressBookName(book.name)} (${book.owner})`,
      value: book.guid,
      userId: book.user_id || 0
    }));
  } else {
    const { data: personal } = await fetchAbPersonal();
    if (personal) {
      abOptions.value = [{ label: $t('dataMap.ab.personal'), value: personal.guid, userId: 0 }];
    }
  }
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
    <NCard :title="$t('route.my-devices_peers')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
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
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1200"
        :loading="loading"
        remote
        :row-key="(row: any) => row.id"
        :pagination="{
          page: currentPage,
          pageSize: pageSize,
          itemCount: total,
          showSizePicker: true,
          pageSizes: [10, 20, 50],
          onChange: handlePageChange,
          onUpdatePageSize: handlePageSizeChange
        }"
        class="sm:h-full"
      />
    </NCard>
    <NModal
      v-model:show="modalVisible"
      preset="card"
      :title="editing ? $t('common.edit') : $t('common.add')"
      class="max-w-620px lt-sm:max-w-full lt-sm:w-full"
    >
      <NForm :label-placement="appStore.isMobile ? 'top' : 'left'" :label-width="appStore.isMobile ? undefined : 110">
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
        <NFormItem :label="$t('dataMap.ab.rustdesk_id')">
          <NInput v-model:value="form.id" :disabled="editing" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.ab.alias')"><NInput v-model:value="form.alias" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.username')"><NInput v-model:value="form.username" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.hostname')"><NInput v-model:value="form.hostname" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.platform')"><NInput v-model:value="form.platform" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.tags')">
          <NInput v-model:value="form.tagText" :placeholder="$t('dataMap.ab.tagsHint')" />
        </NFormItem>
        <NFormItem :label="$t('dataMap.ab.note')"><NInput v-model:value="form.note" type="textarea" /></NFormItem>
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
