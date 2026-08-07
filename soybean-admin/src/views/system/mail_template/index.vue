<script setup lang="tsx">
import { computed, reactive, ref } from 'vue';
import { NButton, NDatePicker, NInput, NSelect, NSpace, NTag } from 'naive-ui';
import { addMailTemplate, fetchMailTemplateList } from '@/service/api/system';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { MailTemplateOptions } from '@/constants/business';
import { downloadCsv, parseCsv } from '@/utils/csv';
import MailTemplateEdit from './components/edit.vue';
import MailTemplateSearch from './components/search.vue';
import TableHeader from './components/table-header.vue';
const appStore = useAppStore();
const authStore = useAuthStore();
const isAdmin = computed(() => authStore.userInfo.roles.includes('R_SUPER'));
const importInput = ref<HTMLInputElement>();
const headerFilters = reactive<{
  id: string;
  name: string;
  type: number | null;
  subject: string;
  created_at: [number, number] | null;
}>({ id: '', name: '', type: null, subject: '', created_at: null });

const tagTypes: any = {
  1: 'success',
  2: 'warning',
  3: 'info'
};

const {
  columns,
  columnChecks,
  data,
  getData,
  getDataByPage,
  loading,
  mobilePagination,
  searchParams,
  resetSearchParams
} = useTable({
  apiFn: fetchMailTemplateList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    id: null,
    // if you want to use the searchParams in Form, you need to define the following properties, and the value is null
    // the value can not be undefined, otherwise the property in Form will not be reactive
    name: null,
    type: null,
    subject: null,
    created_at: null
  },
  columns: () => {
    const filterTitle = (label: string, key: 'id' | 'name' | 'subject') => () => (
      <div class="min-w-130px">
        <div>{label}</div>
        <NInput
          value={headerFilters[key]}
          size="tiny"
          clearable
          placeholder={$t('common.keywordSearch')}
          onUpdateValue={value => {
            headerFilters[key] = value;
            applyHeaderFilters();
          }}
        />
      </div>
    );
    return [
      {
        type: 'selection',
        align: 'center',
        disabled: row => {
          return row.id === 1;
        }
      },
      {
        key: 'id',
        title: filterTitle('ID', 'id'),
        align: 'center'
      },
      {
        key: 'name',
        title: filterTitle($t('dataMap.mailTemplate.name'), 'name'),
        align: 'center'
      },
      {
        key: 'type',
        title: () => (
          <div class="min-w-130px">
            <div>{$t('dataMap.mailTemplate.type')}</div>
            <NSelect
              value={headerFilters.type}
              size="tiny"
              clearable
              options={MailTemplateOptions.map(option => ({ ...option, label: $t(option.label as App.I18n.I18nKey) }))}
              onUpdateValue={value => {
                headerFilters.type = value;
                applyHeaderFilters();
              }}
            />
          </div>
        ),
        align: 'center',
        render: row => {
          let label = '';
          for (const option of MailTemplateOptions) {
            if (option.value === row.type) {
              label = option.label;
            }
          }
          return (
            <NTag bordered={false} type={tagTypes[row.type]}>
              {$t(label as App.I18n.I18nKey)}
            </NTag>
          );
        }
      },
      {
        key: 'subject',
        title: filterTitle($t('dataMap.mailTemplate.subject'), 'subject'),
        align: 'center'
      },
      {
        key: 'created_at',
        title: () => (
          <div class="min-w-250px">
            <div>{$t('dataMap.session.created_at')}</div>
            <NDatePicker
              value={headerFilters.created_at}
              type="daterange"
              size="small"
              clearable
              onUpdateValue={value => {
                headerFilters.created_at = value as [number, number] | null;
                applyHeaderFilters();
              }}
            />
          </div>
        ),
        align: 'center'
      },
      {
        key: 'operate',
        title: $t('common.action'),
        align: 'center',
        render: row => {
          if (!isAdmin.value) return null;
          return (
            <NSpace justify={'center'}>
              <NButton size={'small'} type={'success'} onClick={() => handleEditTable(row)}>
                {$t('common.edit')}
              </NButton>
            </NSpace>
          );
        }
      }
    ];
  }
});

const {
  drawerVisible,
  operateType,
  handleAdd,
  handleEdit,
  editingData,
  checkedRowKeys
  // closeDrawer
} = useTableOperate(data, getData);

function handleEditTable(row: Api.System.MailTemplate) {
  handleEdit(row.id);
}

function applyHeaderFilters() {
  Object.assign(searchParams, headerFilters, {
    created_at: headerFilters.created_at?.map(value => new Date(value).toISOString()) || null
  });
  getDataByPage();
}

async function exportData() {
  const { data: result, error } = await fetchMailTemplateList({ ...searchParams, current: 1, size: 10000 });
  if (!error && result)
    downloadCsv('mail-templates.csv', result.records as any[], ['name', 'type', 'subject', 'contents']);
}

async function importData(event: Event) {
  const file = (event.target as HTMLInputElement).files?.[0];
  if (!file) return;
  const rows = await parseCsv(file);
  let success = 0;
  for (const row of rows) {
    if (!row.name || !row.subject) continue;
    const result = await addMailTemplate({
      name: row.name,
      type: Number(row.type) || 3,
      subject: row.subject,
      contents: row.contents || ''
    } as Api.System.MailTemplate);
    if (!result.error) success += 1;
  }
  window.$message?.success(`${success}/${rows.length}`);
  (event.target as HTMLInputElement).value = '';
  getDataByPage();
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <MailTemplateSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />
    <NCard
      :title="$t('route.system_mail_template')"
      :bordered="false"
      size="small"
      class="sm:flex-1-hidden card-wrapper"
    >
      <template #header-extra>
        <input ref="importInput" type="file" accept=".csv,text/csv" class="hidden" @change="importData" />
        <TableHeader
          v-model:columns="columnChecks"
          :disabled-kill="checkedRowKeys.length === 0"
          :loading="loading"
          @add="handleAdd"
          @import="importInput?.click()"
          @export="exportData"
          @refresh="getData"
        />
      </template>
      <NDataTable
        v-model:checked-row-keys="checkedRowKeys"
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="962"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
      <MailTemplateEdit
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getDataByPage"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
