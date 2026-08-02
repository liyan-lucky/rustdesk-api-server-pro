<script setup lang="tsx">
import { onMounted, reactive, ref } from 'vue';
import { NButton, NSpace, NPopconfirm, NColorPicker, NSelect } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { fetchAbTags, fetchAbTagAdd, fetchAbTagUpdate, fetchAbTagRename, fetchAbTagDelete, fetchAbPersonal, fetchAbAllList } from '@/service/api/address-book';

const appStore = useAppStore();
const loading = ref(false);
const data = ref<any[]>([]);
const currentAbGuid = ref('');
const abOptions = ref<{ label: string; value: string }[]>([]);
const modalVisible = ref(false);
const editing = ref(false);
const form = reactive({ old: '', name: '', color: '#4098fc' });

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const },
  { key: 'ab_id', title: $t('dataMap.ab.ab_id'), align: 'center' as const },
  { key: 'name', title: $t('dataMap.ab.tagName'), align: 'center' as const },
  {
    key: 'color',
    title: $t('dataMap.ab.tagColor'),
    align: 'center' as const,
    render: (row: any) => row.color ? <NColorPicker value={colorToHex(row.color)} size="small" disabled modes={['hex']} /> : '-'
  },
  {
    key: 'actions',
    title: $t('common.action'),
    align: 'center' as const,
    render: (row: any) => (
      <NSpace size="small" justify="center">
        <NButton type="primary" size="small" quaternary onClick={() => openEdit(row)}>{$t('common.edit')}</NButton>
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
      abOptions.value = [{ label: `${$t('dataMap.ab.personal')} (${personalRes.guid.slice(0, 8)}...)`, value: personalRes.guid }];
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
  const { error } = await fetchAbTagDelete(currentAbGuid.value, [row.name]);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    loadData();
  }
}

function colorToHex(value: number) { return `#${(value & 0xffffff).toString(16).padStart(6, '0')}`; }
function colorToInt(value: string) { return (0xff000000 | Number.parseInt(value.slice(1), 16)) >>> 0; }

function openAdd() {
  editing.value = false;
  Object.assign(form, { old: '', name: '', color: '#4098fc' });
  modalVisible.value = true;
}

function openEdit(row: any) {
  editing.value = true;
  Object.assign(form, { old: row.name, name: row.name, color: colorToHex(row.color) });
  modalVisible.value = true;
}

async function submit() {
  if (!form.name.trim()) return window.$message?.warning($t('dataMap.ab.nameRequired'));
  let error: unknown;
  if (!editing.value) {
    ({ error } = await fetchAbTagAdd(currentAbGuid.value, { name: form.name, color: colorToInt(form.color) }));
  } else {
    if (form.old !== form.name) ({ error } = await fetchAbTagRename(currentAbGuid.value, { old: form.old, new: form.name }));
    if (!error) ({ error } = await fetchAbTagUpdate(currentAbGuid.value, { name: form.name, color: colorToInt(form.color) }));
  }
  if (!error) {
    window.$message?.success(editing.value ? $t('common.updateSuccess') : $t('common.addSuccess'));
    modalVisible.value = false;
    loadData();
  }
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
        <NSpace><NSelect v-model:value="currentAbGuid" :options="abOptions" size="small" style="width: 260px" @update:value="handleAbChange" /><NButton type="primary" size="small" :disabled="!currentAbGuid" @click="openAdd">{{ $t('common.add') }}</NButton></NSpace>
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
    <NModal v-model:show="modalVisible" preset="card" :title="editing ? $t('common.edit') : $t('common.add')" class="max-w-520px">
      <NForm label-placement="left" label-width="100">
        <NFormItem :label="$t('dataMap.ab.tagName')"><NInput v-model:value="form.name" /></NFormItem>
        <NFormItem :label="$t('dataMap.ab.tagColor')"><NColorPicker v-model:value="form.color" :show-alpha="false" :modes="['hex']" /></NFormItem>
      </NForm>
      <template #footer><NSpace justify="end"><NButton @click="modalVisible = false">{{ $t('common.cancel') }}</NButton><NButton type="primary" @click="submit">{{ $t('common.confirm') }}</NButton></NSpace></template>
    </NModal>
  </div>
</template>

<style scoped></style>
