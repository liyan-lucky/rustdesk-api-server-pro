<script setup lang="tsx">
import { NTag } from 'naive-ui';
import { fetchErrorLogClear, fetchErrorLogList } from '@/service/api/error-log';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useTable } from '@/hooks/common/table';
import TableHeader from './components/table-header.vue';
import ErrorLogSearch from './components/search.vue';

const appStore = useAppStore();

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
  apiFn: fetchErrorLogList,
  showTotal: true,
  apiParams: {
    current: 1,
    size: 10,
    code: null,
    module: null,
    created_at: null
  },
  columns: () => [
    {
      key: 'id',
      title: 'ID',
      align: 'center'
    },
    {
      key: 'code',
      title: $t('dataMap.errorLog.code'),
      align: 'center',
      render: (row: any) => {
        return (
          <NTag bordered={false} type="error">
            {row.code || '-'}
          </NTag>
        );
      }
    },
    {
      key: 'message',
      title: $t('dataMap.errorLog.message'),
      align: 'center',
      ellipsis: { tooltip: true }
    },
    {
      key: 'module',
      title: $t('dataMap.errorLog.module'),
      align: 'center'
    },
    {
      key: 'path',
      title: $t('dataMap.errorLog.path'),
      align: 'center',
      ellipsis: { tooltip: true }
    },
    {
      key: 'method',
      title: $t('dataMap.errorLog.method'),
      align: 'center'
    },
    {
      key: 'user_name',
      title: $t('dataMap.errorLog.user_name'),
      align: 'center'
    },
    {
      key: 'client_ip',
      title: $t('dataMap.errorLog.client_ip'),
      align: 'center'
    },
    {
      key: 'created_at',
      title: $t('dataMap.errorLog.created_at'),
      align: 'center'
    }
  ]
});

async function handleClear() {
  window.$dialog?.warning({
    title: $t('common.tip'),
    content: $t('common.confirmDelete'),
    positiveText: $t('common.confirm'),
    negativeText: $t('common.cancel'),
    onPositiveClick: async () => {
      await fetchErrorLogClear();
      getData();
    }
  });
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <ErrorLogSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getDataByPage" />

    <NCard :title="$t('route.audit_error-logs')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeader v-model:columns="columnChecks" :loading="loading" @refresh="getData" @clear="handleClear" />
      </template>
      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1200"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
