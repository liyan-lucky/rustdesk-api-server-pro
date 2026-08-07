<script setup lang="tsx">
import { computed, onMounted, ref } from 'vue';
import { NButton, NPopconfirm, NSpace, NTag } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';
import { fetchContainerLogList, fetchContainerLogClear } from '@/service/api/audit';

const appStore = useAppStore();
const authStore = useAuthStore();
const isAdmin = computed(() => authStore.userInfo.roles.includes('R_SUPER'));
const loading = ref(false);
const clearing = ref(false);
const data = ref<any[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);

type NTagType = 'default' | 'primary' | 'info' | 'success' | 'warning' | 'error';

const levelColors: Record<string, NTagType> = {
  INFO: 'success',
  WARN: 'warning',
  ERROR: 'error'
};

const sourceColors: Record<string, NTagType> = {
  system: 'info',
  admin: 'warning',
  client: 'default',
  portal: 'success',
  http: 'default'
};

const searchLevel = ref<string | null>(null);
const searchSource = ref<string | null>(null);
const searchPath = ref('');
const searchStatusCode = ref('');

const levelOptions = [
  { label: 'INFO', value: 'INFO' },
  { label: 'WARN', value: 'WARN' },
  { label: 'ERROR', value: 'ERROR' }
];

const sourceOptions = [
  { label: 'system', value: 'system' },
  { label: 'admin', value: 'admin' },
  { label: 'client', value: 'client' },
  { label: 'portal', value: 'portal' },
  { label: 'http', value: 'http' }
];

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const, width: 60 },
  {
    key: 'timestamp',
    title: $t('dataMap.containerLog.timestamp'),
    align: 'center' as const,
    width: 160
  },
  {
    key: 'level',
    title: $t('dataMap.containerLog.level'),
    align: 'center' as const,
    width: 80,
    render: (row: any) => (
      <NTag bordered={false} type={levelColors[row.level] || 'default'} size="small">
        {row.level}
      </NTag>
    )
  },
  {
    key: 'source',
    title: $t('dataMap.containerLog.source'),
    align: 'center' as const,
    width: 80,
    render: (row: any) => (
      <NTag bordered={false} type={sourceColors[row.source] || 'default'} size="small">
        {row.source}
      </NTag>
    )
  },
  {
    key: 'message',
    title: $t('dataMap.containerLog.message'),
    ellipsis: { tooltip: true }
  },
  {
    key: 'method',
    title: $t('dataMap.containerLog.method'),
    align: 'center' as const,
    width: 70
  },
  {
    key: 'path',
    title: $t('dataMap.containerLog.path'),
    ellipsis: { tooltip: true }
  },
  {
    key: 'status_code',
    title: $t('dataMap.containerLog.status_code'),
    align: 'center' as const,
    width: 70,
    render: (row: any) => {
      const code = row.status_code || 0;
      let type: NTagType = 'default';
      if (code >= 500) type = 'error';
      else if (code >= 400) type = 'warning';
      else if (code >= 200 && code < 300) type = 'success';
      return <NTag bordered={false} type={type} size="small">{code || '-'}</NTag>;
    }
  },
  {
    key: 'duration_ms',
    title: $t('dataMap.containerLog.duration_ms'),
    align: 'center' as const,
    width: 80,
    render: (row: any) => {
      const ms = row.duration_ms || 0;
      let type: NTagType = 'default';
      if (ms > 1000) type = 'error';
      else if (ms > 500) type = 'warning';
      return <NTag bordered={false} type={type} size="small">{ms}ms</NTag>;
    }
  },
  {
    key: 'user_name',
    title: $t('dataMap.containerLog.user_name'),
    align: 'center' as const,
    width: 90
  },
  {
    key: 'client_ip',
    title: $t('dataMap.containerLog.client_ip'),
    align: 'center' as const,
    width: 120,
    ellipsis: { tooltip: true }
  }
];

async function loadData() {
  loading.value = true;
  try {
    const params: any = {
      current: currentPage.value,
      size: pageSize.value
    };
    if (searchLevel.value) params.level = searchLevel.value;
    if (searchSource.value) params.source = searchSource.value;
    if (searchPath.value) params.path = searchPath.value;
    if (searchStatusCode.value) params.status_code = searchStatusCode.value;

    const { data: res, error } = await fetchContainerLogList(params);
    if (!error && res) {
      data.value = res.records || [];
      total.value = res.total || 0;
    }
  } finally {
    loading.value = false;
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

async function handleClearAll() {
  try {
    clearing.value = true;
    const { error } = await fetchContainerLogClear();
    if (!error) {
      window.$message?.success($t('common.deleteSuccess'));
      loadData();
    }
  } finally {
    clearing.value = false;
  }
}

function handleSearch() {
  currentPage.value = 1;
  loadData();
}

function handleReset() {
  searchLevel.value = null;
  searchSource.value = null;
  searchPath.value = '';
  searchStatusCode.value = '';
  currentPage.value = 1;
  loadData();
}

onMounted(() => {
  loadData();
});
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :bordered="false" size="small" class="card-wrapper">
      <NCollapse>
        <NCollapseItem :title="$t('common.search')" name="system-log-search">
          <NForm label-placement="left" :label-width="80">
            <NGrid responsive="screen" item-responsive>
              <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.containerLog.level')" path="level">
                <NSelect v-model:value="searchLevel" :options="levelOptions" clearable />
              </NFormItemGi>
              <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.containerLog.source')" path="source">
                <NSelect v-model:value="searchSource" :options="sourceOptions" clearable />
              </NFormItemGi>
              <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.containerLog.path')" path="path">
                <NInput v-model:value="searchPath" />
              </NFormItemGi>
              <NFormItemGi span="24 s:12 m:6" :label="$t('dataMap.containerLog.status_code')" path="status_code">
                <NInput v-model:value="searchStatusCode" />
              </NFormItemGi>
              <NFormItemGi span="24 m:12">
                <NSpace class="w-full" justify="end">
                  <NButton @click="handleReset">
                    <template #icon>
                      <icon-ic-round-refresh class="text-icon" />
                    </template>
                    {{ $t('common.reset') }}
                  </NButton>
                  <NButton type="primary" ghost @click="handleSearch">
                    <template #icon>
                      <icon-ic-round-search class="text-icon" />
                    </template>
                    {{ $t('common.search') }}
                  </NButton>
                </NSpace>
              </NFormItemGi>
            </NGrid>
          </NForm>
        </NCollapseItem>
      </NCollapse>
    </NCard>

    <NCard :title="$t('route.audit_system-logs')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <NSpace size="small" justify="end">
          <NPopconfirm v-if="isAdmin" @positive-click="handleClearAll">
            {{ $t('common.confirmClear') }}
            <template #trigger>
              <NButton type="error" size="small" :loading="clearing">{{ $t('common.clear') }}</NButton>
            </template>
          </NPopconfirm>
          <NButton size="small" @click="loadData">
            <template #icon>
              <icon-mdi-refresh class="text-icon" :class="{ 'animate-spin': loading }" />
            </template>
            {{ $t('common.refresh') }}
          </NButton>
        </NSpace>
      </template>
      <NDataTable
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="1400"
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
  </div>
</template>

<style scoped></style>
