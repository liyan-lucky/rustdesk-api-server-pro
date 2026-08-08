<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NButton, NSpace } from 'naive-ui';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import {
  fetchWebauthnCredentials,
  fetchDeleteWebauthnCredential,
  fetchRenameWebauthnCredential,
  fetchWebauthnRegisterBegin,
  fetchWebauthnRegisterFinish
} from '@/service/api/auth';
import { isWebAuthnSupported, prepareCreationOptions, serializeCreation } from '@/utils/webauthn';

const appStore = useAppStore();
const credentials = ref<Api.Auth.WebauthnCredential[]>([]);
const loading = ref(false);
const registerLoading = ref(false);
const renameModalVisible = ref(false);
const renameId = ref(0);
const renameName = ref('');
const registerModalVisible = ref(false);
const registerName = ref('');

const columns = [
  { key: 'name', title: $t('page.login.passkey.credentialName') },
  { key: 'createdAt', title: $t('page.login.passkey.createdAt') },
  {
    key: 'lastUsedAt',
    title: $t('page.login.passkey.lastUsedAt'),
    render: (row: Api.Auth.WebauthnCredential) => row.lastUsedAt || '-'
  },
  {
    key: 'actions',
    title: $t('common.action'),
    width: 160,
    render: (row: Api.Auth.WebauthnCredential) => (
      <NSpace size={8}>
        <NButton size="small" tertiary onClick={() => openRename(row)}>
          {$t('page.login.passkey.rename')}
        </NButton>
        <NButton size="small" tertiary type="error" onClick={() => handleDelete(row.id)}>
          {$t('common.delete')}
        </NButton>
      </NSpace>
    )
  }
];

async function loadCredentials() {
  loading.value = true;
  try {
    const { data } = await fetchWebauthnCredentials();
    credentials.value = data || [];
  } catch {
    credentials.value = [];
  } finally {
    loading.value = false;
  }
}

async function handleRegister() {
  if (registerLoading.value) return;
  if (!isWebAuthnSupported()) {
    window.$message?.error($t('page.login.passkey.notSupported'));
    return;
  }
  registerLoading.value = true;
  try {
    const { data: beginData, error: beginError } = await fetchWebauthnRegisterBegin();
    if (beginError || !beginData?.publicKey) {
      window.$message?.error($t('page.login.passkey.registerFailed'));
      return;
    }

    const publicKey = prepareCreationOptions(beginData.publicKey);
    const credential = (await navigator.credentials.create({ publicKey })) as PublicKeyCredential | null;
    if (!credential) return;

    const serialized = serializeCreation(credential);
    const name = registerName.value || new Date().toLocaleString();
    const { error: finishError } = await fetchWebauthnRegisterFinish(serialized, name);
    if (!finishError) {
      window.$message?.success($t('page.login.passkey.registerSuccess'));
      registerModalVisible.value = false;
      registerName.value = '';
      await loadCredentials();
    } else {
      window.$message?.error($t('page.login.passkey.registerFailed'));
    }
  } catch {
    window.$message?.error($t('page.login.passkey.registerFailed'));
  } finally {
    registerLoading.value = false;
  }
}

async function handleDelete(id: number) {
  window.$dialog?.warning({
    title: $t('page.login.passkey.deleteConfirm'),
    positiveText: $t('common.confirm'),
    negativeText: $t('common.cancel'),
    onPositiveClick: async () => {
      const { error } = await fetchDeleteWebauthnCredential(id);
      if (!error) {
        window.$message?.success($t('page.login.passkey.deleteSuccess'));
        await loadCredentials();
      }
    }
  });
}

function openRename(cred: Api.Auth.WebauthnCredential) {
  renameId.value = cred.id;
  renameName.value = cred.name;
  renameModalVisible.value = true;
}

async function confirmRename() {
  const { error } = await fetchRenameWebauthnCredential(renameId.value, renameName.value);
  if (!error) {
    window.$message?.success($t('page.login.passkey.renameSuccess'));
    renameModalVisible.value = false;
    await loadCredentials();
  } else {
    window.$message?.error($t('page.login.passkey.renameFailed'));
  }
}

function openRegister() {
  registerName.value = '';
  registerModalVisible.value = true;
}

onMounted(() => {
  loadCredentials();
});
</script>

<template>
  <NCard :title="$t('page.login.passkey.title')" :bordered="false" :class="appStore.isMobile ? 'w-full' : 'max-w-720px'" class="mt-16px">
    <NSpace vertical size="medium">
      <div class="flex justify-end">
        <NButton type="primary" :loading="registerLoading" @click="openRegister">
          <template #icon><SvgIcon icon="mdi:key-plus" /></template>
          {{ $t('page.login.passkey.register') }}
        </NButton>
      </div>

      <NEmpty v-if="!loading && credentials.length === 0" :description="$t('page.login.passkey.noCredentials')" />

      <NDataTable v-else :columns="columns" :data="credentials" :loading="loading" :bordered="false" responsive />
    </NSpace>
  </NCard>

  <NModal v-model:show="renameModalVisible" preset="dialog" :title="$t('page.login.passkey.rename')">
    <NInput v-model:value="renameName" :placeholder="$t('page.login.passkey.credentialName')" />
    <template #action>
      <NSpace justify="end">
        <NButton @click="renameModalVisible = false">{{ $t('common.cancel') }}</NButton>
        <NButton type="primary" @click="confirmRename">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </template>
  </NModal>

  <NModal v-model:show="registerModalVisible" preset="dialog" :title="$t('page.login.passkey.register')">
    <NSpace vertical size="medium">
      <p>{{ $t('page.login.passkey.enterName') }}</p>
      <NInput v-model:value="registerName" :placeholder="$t('page.login.passkey.namePlaceholder')" />
    </NSpace>
    <template #action>
      <NSpace justify="end">
        <NButton @click="registerModalVisible = false">{{ $t('common.cancel') }}</NButton>
        <NButton type="primary" :loading="registerLoading" @click="handleRegister">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </template>
  </NModal>
</template>
