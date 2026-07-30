<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { useAuthStore } from '@/store/modules/auth';
import { fetchCaptcha, fetchUserLogin } from '@/service/api/auth';
import { localStg } from '@/utils/storage';

defineOptions({
  name: 'UserLogin'
});

const authStore = useAuthStore();
const router = useRouter();
const { formRef, validate } = useNaiveForm();

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
  const { data, error } = await fetchUserLogin(model);
  if (!error && data?.token) {
    localStg.set('token', data.token);
    localStg.set('userType', 'user');
    authStore.token = data.token;
    router.push('/my-devices');
  } else if (error?.response?.data?.message === 'CaptchaError') {
    handleCaptcha();
  }
}

async function handleCaptcha() {
  const c = await fetchCaptcha();
  captcha.id = c.data?.id || '';
  captcha.img = c.data?.img || '';
  model.captchaId = captcha.id || '';
}

onMounted(() => {
  handleCaptcha();
});
</script>

<template>
  <NForm ref="formRef" :model="model" :rules="rules" size="large" :show-label="false">
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
      <div class="pl-8px">
        <img width="152" height="40" class="cursor-pointer" :src="captcha.img" @click="handleCaptcha" />
      </div>
    </NFormItem>
    <NSpace vertical :size="16">
      <NButton
        attr-type="submit"
        type="primary"
        size="large"
        round
        block
        @click="handleSubmit"
      >
        {{ $t('common.confirm') }}
      </NButton>
      <NButton text type="primary" @click="router.push('/login/pwd-login')">
        {{ $t('page.login.userLogin.switchToAdmin') }}
      </NButton>
    </NSpace>
  </NForm>
</template>

<style scoped></style>
