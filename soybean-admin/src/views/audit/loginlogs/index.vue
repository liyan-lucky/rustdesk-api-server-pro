<script setup lang="tsx">
import { computed, onMounted, ref } from 'vue';
import { NTag, NSpace, NInput, NButton, NSelect } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';
import { request } from '@/service/request';

const appStore = useAppStore();
const authStore = useAuthStore();
const isAdmin = computed(() => authStore.userInfo.roles.includes('R_SUPER'));
const loading = ref(false);
const data = ref<any[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);
const searchUsername = ref('');
const searchEvent = ref(null as string | null);

const eventOptions = [
  { label: $t('dataMap.loginLog.allEvents'), value: '' },
  { label: 'api_login', value: 'api_login' },
  { label: 'admin_login', value: 'admin_login' },
  { label: 'web_user_login', value: 'web_user_login' },
  { label: 'api_token_invalid', value: 'api_token_invalid' },
  { label: 'admin_token_invalid', value: 'admin_token_invalid' },
  { label: 'web_user_token_invalid', value: 'web_user_token_invalid' },
  { label: 'api_logout', value: 'api_logout' }
];

const columns = [
  { key: 'id', title: 'ID', align: 'center' as const },
  { key: 'username', title: $t('dataMap.user.username'), align: 'center' as const },
  { key: 'event', title: $t('dataMap.loginLog.event'), align: 'center' as const },
  { key: 'ip', title: $t('dataMap.audit.ip'), align: 'center' as const },
  { key: 'user_agent', title: $t('dataMap.loginLog.userAgent'), align: 'center' as const, ellipsis: { tooltip: true } },
  {
    key: 'success',
    title: $t('dataMap.loginLog.success'),
    align: 'center' as const,
    render: (row: any) => <NTag type={row.success ? 'success' : 'error'} size="small">{row.success ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}</NTag>
  },
  { key: 'reason', title: $t('dataMap.loginLog.reason'), align: 'center' as const },
  { key: 'created_at', title: $t('dataMap.audit.created_at'), align: 'center' as const }
];

async function loadData() {
  loading.value = true;
  try {
    const params: Record<string, any> = { current: currentPage.value, size: pageSize.value };
    if (isAdmin.value && searchUsername.value) params.username = searchUsername.value;
    if (searchEvent.value) params.event = searchEvent.value;
    const { data: res, error } = await request({ url: '/security-audit/list', params });
    if (!error && res) {
      data.value = res.records || [];
      total.value = res.total || 0;
    }
  } finally {
    loading.value = false;
  }
}

function handleSearch() { currentPage.value = 1; loadData(); }
function handlePageChange(page: number) { currentPage.value = page; loadData(); }
function handlePageSizeChange(size: number) { pageSize.value = size; currentPage.value = 1; loadData(); }

onMounted(() => { loadData(); });
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :bordered="false" size="small">
      <NSpace align="center">
        <NInput v-if="isAdmin" v-model:value="searchUsername" :placeholder="$t('dataMap.user.username')" clearable style="width: 200px" @keyup.enter="handleSearch" />
        <NSelect v-model:value="searchEvent" :options="eventOptions" :placeholder="$t('dataMap.loginLog.event')" clearable style="width: 200px" />
        <NButton type="primary" @click="handleSearch">{{ $t('common.search') }}</NButton>
      </NSpace>
    </NCard>

    <NCard :title="$t('route.audit_loginlogs')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
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
  </div>
</template>

<style scoped></style>
