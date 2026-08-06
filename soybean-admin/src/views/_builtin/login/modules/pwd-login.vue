<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';
import { fetchCaptcha, fetchOAuthLoginUrl, fetchOAuthProviders } from '@/service/api/auth';

defineOptions({
  name: 'PwdLogin'
});

const authStore = useAuthStore();
const route = useRoute();
const { formRef, validate } = useNaiveForm();
const oauthProviders = ref<Api.Auth.OAuthProvider[]>([]);
const activeProvider = ref('');

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
    code: [
      {
        required: true,
        message: $t('page.login.common.codePlaceholder')
      }
    ],
    captchaId: [
      {
        required: true,
        message: $t('page.login.common.codePlaceholder')
      }
    ]
  };
});

async function handleSubmit() {
  await validate();
  const err = await authStore.login(model);
  if (err?.response?.data.message === 'CaptchaError') {
    handleCaptcha();
  }
}

async function handleCaptcha() {
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
  activeProvider.value = provider.name;
  try {
    const redirect =
      typeof route.query.redirect === 'string' && route.query.redirect !== '/' ? route.query.redirect : '/#/login';
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
  handleCaptcha();
  loadOAuthProviders();
});
</script>

<template>
  <NForm ref="formRef" :model="model" :rules="rules" size="medium" :show-label="false">
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
    <NFormItem path="code">
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
        :loading="authStore.loginLoading"
        @click="handleSubmit"
      >
        {{ $t('common.confirm') }}
      </NButton>
      <NDivider v-if="oauthProviders.length > 0" class="!mt-0 !mb-0">{{ $t('page.login.common.thirdPartyLogin') }}</NDivider>
      <div class="grid grid-cols-2 gap-8px">
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
