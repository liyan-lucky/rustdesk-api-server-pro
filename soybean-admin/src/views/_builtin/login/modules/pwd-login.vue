<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';
import {
  fetchCaptcha,
  fetchConfirmClientWebauth,
  fetchConfirmOAuthBinding,
  fetchLoginConfig,
  fetchOAuthLoginUrl,
  fetchOAuthProviders
} from '@/service/api/auth';

defineOptions({
  name: 'PwdLogin'
});

const authStore = useAuthStore();
const route = useRoute();
const { formRef, validate } = useNaiveForm();
const oauthProviders = ref<Api.Auth.OAuthProvider[]>([]);
const activeProvider = ref('');
const clientWebauthCompleted = ref(false);
const clientWebauthLoading = ref(false);
const oauthBindingLoading = ref(false);
const captchaEnabled = ref(true);
const clientPollToken = computed(() => {
  if (route.name !== 'client-webauth') return '';
  return typeof route.query.poll_token === 'string' ? route.query.poll_token : '';
});
const oauthBindingTicket = computed(() =>
  typeof route.query.oauth_bind_ticket === 'string' ? route.query.oauth_bind_ticket : ''
);
const oauthAllowCreate = computed(() => route.query.oauth_allow_create === '1');

const model: Api.Form.LoginForm = reactive({
  username: '',
  password: '',
  code: '',
  captchaId: ''
});

const captcha: Api.Auth.Captcha = reactive({
  id: '',
  img: ''
});

const rules = computed<Record<keyof Api.Form.LoginForm, App.Global.FormRule[]>>(() => {
  return {
    username: [
      {
        required: true,
        message: $t('page.user.list.inputUsername')
      }
    ],
    password: [
      {
        required: true,
        message: $t('page.user.list.inputPassword')
      }
    ],
    code: captchaEnabled.value
      ? [{ required: true, message: $t('page.login.common.codePlaceholder') }]
      : [],
    captchaId: captchaEnabled.value
      ? [{ required: true, message: $t('page.login.common.codePlaceholder') }]
      : []
  };
});

async function handleSubmit() {
  await validate();
  if (oauthBindingTicket.value) {
    if (oauthBindingLoading.value) return;
    oauthBindingLoading.value = true;
    try {
      const { data, error } = await fetchConfirmOAuthBinding(oauthBindingTicket.value, model);
      if (!error && (await finishOAuthBinding(data))) return;
      if (error?.response?.data?.message?.includes('CaptchaError')) handleCaptcha();
    } finally {
      oauthBindingLoading.value = false;
    }
    return;
  }
  if (clientPollToken.value) {
    if (clientWebauthLoading.value) return;
    clientWebauthLoading.value = true;
    try {
      const { data, error } = await fetchConfirmClientWebauth(clientPollToken.value, model);
      if (!error && data?.ok) {
        clientWebauthCompleted.value = true;
        window.setTimeout(closeClientWebauthPage, 1200);
        return;
      }
      if (error?.response?.data?.message?.includes('CaptchaError')) handleCaptcha();
    } finally {
      clientWebauthLoading.value = false;
    }
    return;
  }
  const err = await authStore.login(model);
  if (err?.response?.data?.message?.includes('CaptchaError')) {
    handleCaptcha();
  }
}

async function finishOAuthBinding(data?: { ticket: string; client: boolean; redirect: string }) {
  if (data?.client) {
    clientWebauthCompleted.value = true;
    window.setTimeout(closeClientWebauthPage, 1200);
    return true;
  }
  if (data?.ticket) {
    const redirect = typeof data.redirect === 'string' && data.redirect.startsWith('/') ? data.redirect : '/';
    const consumed = await authStore.loginByOAuthTicket(data.ticket, false);
    if (!consumed) {
      window.$message?.error($t('api.RequestError'));
      return false;
    }
    const nextUrl = new URL(window.location.href);
    nextUrl.search = '';
    nextUrl.hash = redirect === '/' ? '#/' : redirect.startsWith('/#/') ? redirect.slice(1) : `#${redirect}`;
    window.location.replace(nextUrl.toString());
    return true;
  }
  return false;
}

async function handleCreateAndLogin() {
  if (!oauthBindingTicket.value || !oauthAllowCreate.value || oauthBindingLoading.value) return;
  if (captchaEnabled.value && !model.code.trim()) {
    window.$message?.warning($t('page.login.common.codePlaceholder'));
    return;
  }
  oauthBindingLoading.value = true;
  try {
    const { data, error } = await fetchConfirmOAuthBinding(oauthBindingTicket.value, {
      ...model,
      username: '',
      password: ''
    });
    if (!error && (await finishOAuthBinding(data))) return;
    if (error?.response?.data?.message?.includes('CaptchaError')) handleCaptcha();
  } finally {
    oauthBindingLoading.value = false;
  }
}

function launchRustDesk() {
  window.location.href = 'rustdesk://config/';
}

function closeClientWebauthPage() {
  window.close();
}

async function handleCaptcha() {
  if (!captchaEnabled.value) return;
  const c = await fetchCaptcha();
  captcha.id = c.data?.id || '';
  captcha.img = c.data?.img || '';
  model.captchaId = captcha.id || '';
}

async function loadOAuthProviders() {
  try {
    const { data } = await fetchOAuthProviders();
    oauthProviders.value = data || [];
  } catch {
    oauthProviders.value = [];
  }
}

async function handleOAuthLogin(provider: Api.Auth.OAuthProvider) {
  if (activeProvider.value) return;
  activeProvider.value = provider.name;
  try {
    const redirect =
      typeof route.query.redirect === 'string' && route.query.redirect.startsWith('/') && !route.query.redirect.startsWith('//') && route.query.redirect !== '/'
        ? route.query.redirect
        : '/';
    const { data, error } = await fetchOAuthLoginUrl(provider.name, redirect);
    if (!error && data?.enabled && data.url) {
      window.location.href = data.url;
      return;
    }
    window.$message?.error($t('page.login.common.providerUnavailable', { provider: provider.displayName }));
  } finally {
    activeProvider.value = '';
  }
}

function providerIcon(type: string) {
  const icons: Record<string, string> = {
    github: 'mdi:github',
    qq: 'ri:qq-fill',
    google: 'logos:google-icon',
    apple: 'ic:baseline-apple',
    oidc: 'mdi:shield-account'
  };
  return icons[type] || 'mdi:login-variant';
}

onMounted(() => {
  fetchLoginConfig().then(({ data }) => {
    captchaEnabled.value = data?.captchaEnabled !== false;
    if (captchaEnabled.value) handleCaptcha();
  });
  if (!clientPollToken.value && !oauthBindingTicket.value) loadOAuthProviders();
});
</script>

<template>
  <NResult v-if="clientWebauthCompleted" status="success" :title="$t('page.login.common.loginSuccess')">
    <template #footer>
      <NSpace vertical :size="12">
        <p class="text-center text-14px text-gray-500">认证信息已发送，RustDesk 客户端会自动完成登录。</p>
        <p class="text-center text-12px text-gray-400">页面将尝试自动关闭；若浏览器阻止关闭，可直接关闭此标签页。</p>
        <NButton type="primary" round block @click="launchRustDesk">
          <template #icon><SvgIcon icon="mdi:remote-desktop" /></template>
          返回 RustDesk
        </NButton>
        <NButton round block @click="closeClientWebauthPage">关闭页面</NButton>
      </NSpace>
    </template>
  </NResult>
  <NForm v-else ref="formRef" :model="model" :rules="rules" size="medium" :show-label="false">
    <NAlert v-if="oauthBindingTicket" type="info" :show-icon="false" class="mb-16px">
      这是该第三方身份首次登录。请输入要绑定的本地账户及其密码；验证成功后，后续可直接使用第三方登录。
    </NAlert>
    <NFormItem path="username">
      <NInput v-model:value="model.username" :placeholder="$t('page.login.common.userNamePlaceholder')" />
    </NFormItem>
    <NFormItem path="password">
      <NInput
        v-model:value="model.password"
        type="password"
        show-password-on="click"
        :placeholder="$t('page.login.common.passwordPlaceholder')"
      />
    </NFormItem>
    <NFormItem v-if="captchaEnabled" path="code">
      <NInput v-model:value="model.code" :clearable="true" :placeholder="$t('page.login.common.codePlaceholder')" />
      <div class="flex-shrink-0 pl-8px">
        <img
          width="120"
          height="40"
          class="cursor-pointer lt-sm:h-28px lt-sm:w-80px"
          :src="captcha.img"
          @click="handleCaptcha"
        />
      </div>
    </NFormItem>
    <NSpace vertical :size="16">
      <div class="flex-y-center justify-between">
        <NCheckbox>{{ $t('page.login.pwdLogin.rememberMe') }}</NCheckbox>
      </div>
      <NButton
        attr-type="submit"
        type="primary"
        size="medium"
        round
        block
        :loading="oauthBindingTicket ? oauthBindingLoading : clientPollToken ? clientWebauthLoading : authStore.loginLoading"
        @click="handleSubmit"
      >
        {{ oauthBindingTicket ? '验证并绑定' : $t('common.confirm') }}
      </NButton>
      <NButton
        v-if="oauthBindingTicket && oauthAllowCreate"
        attr-type="button"
        secondary
        round
        block
        :loading="oauthBindingLoading"
        @click="handleCreateAndLogin"
      >
        创建普通用户并登录
      </NButton>
      <NDivider v-if="!oauthBindingTicket && oauthProviders.length > 0" class="!mt-0 !mb-0">
        {{ $t('page.login.common.thirdPartyLogin') }}
      </NDivider>
      <div v-if="!oauthBindingTicket" class="grid grid-cols-2 gap-8px lt-sm:grid-cols-1">
        <NButton
          v-for="provider in oauthProviders"
          :key="provider.name"
          tertiary
          size="small"
          :loading="activeProvider === provider.name"
          @click="handleOAuthLogin(provider)"
        >
          <template #icon><SvgIcon :icon="providerIcon(provider.type)" /></template>
          {{ $t('page.login.common.continueWith', { provider: provider.displayName }) }}
        </NButton>
      </div>
    </NSpace>
  </NForm>
</template>

<style scoped></style>
