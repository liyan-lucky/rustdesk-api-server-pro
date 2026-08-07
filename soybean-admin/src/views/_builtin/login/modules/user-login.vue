<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { $t } from '@/locales';
import { useNaiveForm } from '@/hooks/common/form';
import { fetchCaptcha, fetchUserLogin } from '@/service/api/auth';
import { localStg } from '@/utils/storage';
import { useAuthStore } from '@/store/modules/auth';
import { useRouteStore } from '@/store/modules/route';

defineOptions({
  name: 'UserLogin'
});

const router = useRouter();
const authStore = useAuthStore();
const routeStore = useRouteStore();
const { formRef, validate } = useNaiveForm();
const loading = ref(false);

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
  loading.value = true;
  try {
    const { data, error } = await fetchUserLogin(model);
    if (!error && data?.token) {
      localStg.set('token', data.token);
      localStg.set('userType', 'user');
      if (data.isAdmin !== undefined) {
        localStg.set('isAdmin', data.isAdmin);
      }
      authStore.token = data.token;
      await authStore.initUserInfo();
      await routeStore.initAuthRoute();
      if (data.isAdmin) {
        await router.push({ name: 'home' });
      } else {
        await router.push({ name: 'my-devices_peers' });
      }
    } else if (error?.response?.data?.message === 'CaptchaError') {
      handleCaptcha();
    }
  } finally {
    loading.value = false;
  }
}

async function handleCaptcha() {
  const c = await fetchCaptcha();
  captcha.id = c.data?.id || '';
  captcha.img = c.data?.img || '';
  model.captchaId = captcha.id || '';
}

function switchToAdmin() {
  router.push({ name: 'login', params: { module: 'pwd-login' } });
}

onMounted(() => {
  handleCaptcha();
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
    <NSpace vertical :size="12">
      <NButton attr-type="submit" type="primary" size="medium" round block :loading="loading" @click="handleSubmit">
        {{ $t('common.confirm') }}
      </NButton>
      <NButton text type="primary" @click="switchToAdmin">
        {{ $t('page.login.userLogin.switchToAdmin') }}
      </NButton>
    </NSpace>
  </NForm>
</template>

<style scoped></style>
