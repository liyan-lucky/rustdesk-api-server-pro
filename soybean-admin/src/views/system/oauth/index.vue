<script setup lang="tsx">
import { computed, onMounted, reactive, ref } from 'vue';
import { NButton, NPopconfirm, NSpace, NTag } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { request } from '@/service/request';

const appStore = useAppStore();
const loading = ref(false);
const saving = ref(false);
const testing = ref('');
const showModal = ref(false);
const data = ref<any[]>([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = ref(10);
const providerConfigs = ref<any[]>([]);
const form = reactive({ originalName: '', type: 'github', name: 'github', displayName: 'GitHub', enabled: true, clientId: '', clientSecret: '', redirectUrl: '', scopesText: 'read:user user:email', accountRole: 'admin', bindByEmail: true, autoCreateAdmin: false, autoCreateUser: false, allowedDomainsText: '' });
const defaultCallback = computed(() => `${window.location.origin}/admin/auth/oauth/${form.name || 'github'}/callback`);

function values(value: string) { return value.split(/[\s,;]+/).map(item => item.trim()).filter(Boolean); }
function resetForm(row?: any) {
  Object.assign(form, row ? { originalName: row.name, type: row.type, name: row.name, displayName: row.displayName, enabled: row.enabled, clientId: row.clientId, clientSecret: '', redirectUrl: row.redirectUrl || '', scopesText: (row.scopes || []).join(' '), accountRole: row.accountRole || 'admin', bindByEmail: row.bindByEmail, autoCreateAdmin: row.autoCreateAdmin, autoCreateUser: row.autoCreateUser, allowedDomainsText: (row.allowedEmailDomains || []).join(' ') } : { originalName: '', type: 'github', name: 'github', displayName: 'GitHub', enabled: true, clientId: '', clientSecret: '', redirectUrl: `${window.location.origin}/admin/auth/oauth/github/callback`, scopesText: 'read:user user:email', accountRole: 'admin', bindByEmail: true, autoCreateAdmin: false, autoCreateUser: false, allowedDomainsText: '' });
  showModal.value = true;
}
async function copyText(value: string) {
  try { await navigator.clipboard.writeText(value); } catch {
    const node = document.createElement('textarea'); node.value = value; document.body.appendChild(node); node.select(); document.execCommand('copy'); node.remove();
  }
  window.$message?.success($t('page.oauth.copied'));
}
async function loadProviderConfigs() {
  const { data: res, error } = await request({ url: '/oauth/providers/config' });
  if (!error) providerConfigs.value = Array.isArray(res) ? res : [];
}
async function saveProvider() {
  saving.value = true;
  try {
    const payload = { ...form, scopes: values(form.scopesText), allowedEmailDomains: values(form.allowedDomainsText) };
    const { error } = await request({ url: '/oauth/provider', method: 'post', data: payload });
    if (!error) { showModal.value = false; window.$message?.success($t('common.updateSuccess')); await loadProviderConfigs(); }
  } finally { saving.value = false; }
}
async function deleteProvider(name: string) {
  const { error } = await request({ url: `/oauth/provider/${encodeURIComponent(name)}`, method: 'delete' });
  if (!error) { window.$message?.success($t('common.deleteSuccess')); await loadProviderConfigs(); }
}
async function testProvider(name: string) {
  testing.value = name;
  try {
    const { data: res, error } = await request({ url: `/oauth/provider/${encodeURIComponent(name)}/test`, method: 'post', params: { baseUrl: window.location.origin } });
    if (!error && res?.authorizationUrl) window.$message?.success($t('page.oauth.testSuccess'));
  } finally { testing.value = ''; }
}

const providerColumns = [
  { key: 'displayName', title: $t('page.oauth.displayName') },
  { key: 'name', title: $t('page.oauth.providerName') },
  { key: 'clientId', title: $t('page.oauth.clientId'), ellipsis: { tooltip: true } },
  { key: 'accountRole', title: $t('page.oauth.accountRole') },
  { key: 'enabled', title: $t('dataMap.token.status'), render: (row: any) => <NTag type={row.enabled ? 'success' : 'default'}>{row.enabled ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}</NTag> },
  { key: 'actions', title: $t('common.action'), render: (row: any) => <NSpace><NButton size="small" onClick={() => resetForm(row)}>{$t('common.edit')}</NButton><NButton size="small" loading={testing.value === row.name} onClick={() => testProvider(row.name)}>{$t('page.oauth.testConfig')}</NButton><NButton size="small" onClick={() => copyText(row.redirectUrl || `${window.location.origin}/admin/auth/oauth/${row.name}/callback`)}>{$t('page.oauth.copyCallback')}</NButton><NPopconfirm onPositiveClick={() => deleteProvider(row.name)}>{{ default: () => $t('common.confirmDelete'), trigger: () => <NButton type="error" size="small" quaternary>{$t('common.delete')}</NButton> }}</NPopconfirm></NSpace> }
];
const columns = [
  { key: 'id', title: 'ID' }, { key: 'user_id', title: $t('dataMap.ab.user_id') }, { key: 'provider', title: $t('dataMap.oauth.provider') }, { key: 'subject', title: $t('dataMap.oauth.subject') }, { key: 'email', title: $t('dataMap.oauth.email') }, { key: 'name', title: $t('dataMap.oauth.name') },
  { key: 'is_admin', title: $t('dataMap.token.is_admin'), render: (row: any) => <NTag type={row.is_admin ? 'warning' : 'default'}>{row.is_admin ? $t('dataMap.token.is_admin') : $t('dataMap.user.statusLabel.normal')}</NTag> },
  { key: 'status', title: $t('dataMap.token.status'), render: (row: any) => <NTag type={row.status === 1 ? 'success' : 'error'}>{row.status === 1 ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}</NTag> },
  { key: 'last_login_at', title: $t('dataMap.oauth.last_login_at') }, { key: 'created_at', title: $t('dataMap.audit.created_at') },
  { key: 'actions', title: $t('common.action'), render: (row: any) => <NPopconfirm onPositiveClick={() => handleDelete(row)}>{{ default: () => $t('common.confirmDelete'), trigger: () => <NButton type="error" size="small" quaternary>{$t('common.delete')}</NButton> }}</NPopconfirm> }
];
async function loadData() { loading.value = true; try { const { data: res, error } = await request({ url: '/oauth/accounts', params: { current: currentPage.value, size: pageSize.value } }); if (!error && res) { data.value = res.records || []; total.value = res.total || 0; } } finally { loading.value = false; } }
async function handleDelete(row: any) { const { error } = await request({ url: `/oauth/account/${row.id}`, method: 'delete' }); if (!error) { window.$message?.success($t('common.deleteSuccess')); loadData(); } }
function handlePageChange(page: number) { currentPage.value = page; loadData(); }
function handlePageSizeChange(size: number) { pageSize.value = size; currentPage.value = 1; loadData(); }
onMounted(() => { loadData(); loadProviderConfigs(); });
</script>

<template>
  <div class="flex-col-stretch gap-16px">
    <NCard :title="$t('page.oauth.configTitle')" :bordered="false" size="small">
      <template #header-extra><NButton type="primary" @click="resetForm()">{{ $t('page.oauth.addProvider') }}</NButton></template>
      <NAlert type="info" class="mb-12px">{{ $t('page.oauth.githubOnlyTip') }}</NAlert>
      <NDataTable :columns="providerColumns" :data="providerConfigs" :row-key="(row: any) => row.name" :scroll-x="1000" />
    </NCard>
    <NCard :title="$t('page.oauth.bindingsTitle')" :bordered="false" size="small">
      <NDataTable :columns="columns" :data="data" size="small" :flex-height="!appStore.isMobile" :scroll-x="1200" :loading="loading" remote :row-key="(row: any) => row.id" :pagination="{ page: currentPage, pageSize, itemCount: total, showSizePicker: true, pageSizes: [10, 20, 50], onChange: handlePageChange, onUpdatePageSize: handlePageSizeChange }" />
    </NCard>
    <NModal v-model:show="showModal" preset="card" :title="form.originalName ? $t('page.oauth.editProvider') : $t('page.oauth.addProvider')" class="w-720px max-w-95vw">
      <NForm label-placement="left" label-width="150">
        <NFormItem :label="$t('dataMap.oauth.provider')"><NSelect v-model:value="form.type" :options="[{ label: 'GitHub', value: 'github' }]" /></NFormItem>
        <NFormItem :label="$t('page.oauth.providerName')"><NInput v-model:value="form.name" /></NFormItem>
        <NFormItem :label="$t('page.oauth.displayName')"><NInput v-model:value="form.displayName" /></NFormItem>
        <NFormItem :label="$t('page.oauth.clientId')"><NInput v-model:value="form.clientId" /></NFormItem>
        <NFormItem :label="$t('page.oauth.clientSecret')"><NInput v-model:value="form.clientSecret" type="password" show-password-on="click" :placeholder="$t('page.oauth.secretPlaceholder')" /></NFormItem>
        <NFormItem :label="$t('page.oauth.redirectUrl')"><NInput v-model:value="form.redirectUrl" /><NButton class="ml-8px" @click="form.redirectUrl = defaultCallback">{{ $t('page.oauth.useDefault') }}</NButton></NFormItem>
        <NFormItem :label="$t('page.oauth.scopes')"><NInput v-model:value="form.scopesText" /></NFormItem>
        <NFormItem :label="$t('page.oauth.accountRole')"><NSelect v-model:value="form.accountRole" :options="[{ label: $t('page.oauth.adminRole'), value: 'admin' }, { label: $t('page.oauth.userRole'), value: 'user' }]" /></NFormItem>
        <NFormItem :label="$t('page.oauth.allowedDomains')"><NInput v-model:value="form.allowedDomainsText" :placeholder="$t('page.oauth.listPlaceholder')" /></NFormItem>
        <NFormItem :label="$t('page.oauth.bindByEmail')"><NSwitch v-model:value="form.bindByEmail" /></NFormItem>
        <NFormItem :label="$t('page.oauth.autoCreateAdmin')"><NSwitch v-model:value="form.autoCreateAdmin" /></NFormItem>
        <NFormItem :label="$t('page.oauth.autoCreateUser')"><NSwitch v-model:value="form.autoCreateUser" /></NFormItem>
        <NFormItem :label="$t('dataMap.token.status')"><NSwitch v-model:value="form.enabled" /></NFormItem>
      </NForm>
      <template #footer><NSpace justify="end"><NButton @click="showModal = false">{{ $t('common.cancel') }}</NButton><NButton type="primary" :loading="saving" @click="saveProvider">{{ $t('common.save') }}</NButton></NSpace></template>
    </NModal>
  </div>
</template>
