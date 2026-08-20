<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import type { Component } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';
import { useThemeStore } from '@/store/modules/theme';
import { loginModuleRecord } from '@/constants/app';
import { useWebBackground } from '@/hooks/common/web-background';
import PwdLogin from './modules/pwd-login.vue';
import UserLogin from './modules/user-login.vue';

interface Props {
  /** The login module */
  module?: UnionKey.LoginModule;
}

const props = defineProps<Props>();

const appStore = useAppStore();
const authStore = useAuthStore();
const themeStore = useThemeStore();
const route = useRoute();
const router = useRouter();
const { backgroundStyle } = useWebBackground();
const ticketProcessing = ref(false);

interface LoginModule {
  label: string;
  component: Component;
}

const moduleMap: Record<UnionKey.LoginModule, LoginModule> = {
  'pwd-login': { label: loginModuleRecord['pwd-login'], component: PwdLogin },
  'user-login': { label: loginModuleRecord['user-login'], component: UserLogin }
};

const activeModule = computed(() => moduleMap[props.module || 'pwd-login']);

async function consumeOAuthTicket(ticket: string) {
  if (ticketProcessing.value) return false;
  ticketProcessing.value = true;
  try {
    return await authStore.loginByOAuthTicket(ticket, false);
  } finally {
    ticketProcessing.value = false;
  }
}

function safeLoginRedirect(value: unknown) {
  if (typeof value !== 'string' || !value.startsWith('/') || value.startsWith('//')) return '/';
  if (value.startsWith('/#/')) return value.slice(2) || '/';
  return value;
}

function finishTicketLogin(target: string) {
  const nextUrl = new URL(window.location.href);
  nextUrl.search = '';
  nextUrl.hash = target === '/' ? '#/' : `#${target}`;
  window.location.replace(nextUrl.toString());
}

onMounted(async () => {
  let shouldReplaceQuery = false;
  const rawSearchQuery: Record<string, string> = {};
  new URLSearchParams(window.location.search).forEach((value, key) => {
    rawSearchQuery[key] = value;
  });
  const mergedQuery = { ...rawSearchQuery, ...route.query };
  const q = { ...mergedQuery };

  const oauthError = mergedQuery.oauth_error;
  if (typeof oauthError === 'string' && oauthError) {
    const messageMap: Record<string, string> = {
      'ERR-2208': $t('page.login.common.oauthAccountNotBound'),
      'ERR-2209': $t('page.login.common.oauthProviderUnreachable'),
      'ERR-2210': $t('page.login.common.oauthStateExpired'),
      'ERR-2211': $t('page.login.common.oauthAuthFailed'),
      'ERR-2212': $t('page.login.common.oauthAuthFailed'),
      oauth_account_not_bound: $t('page.login.common.oauthAccountNotBound'),
      oauth_provider_unreachable: $t('page.login.common.oauthProviderUnreachable'),
      oauth_state_expired: $t('page.login.common.oauthStateExpired'),
      oauth_auth_failed: $t('page.login.common.oauthAuthFailed'),
      auth_failed: $t('page.login.common.oauthAuthFailed')
    };
    window.$message?.error(messageMap[oauthError] || $t('page.login.common.oauthAuthFailed'));
    delete q.oauth_error;
    shouldReplaceQuery = true;
  }

  const oauthTicket = mergedQuery.oauth_ticket;
  if (typeof oauthTicket === 'string' && oauthTicket) {
    const target = safeLoginRedirect(mergedQuery.redirect);
    const consumed = await consumeOAuthTicket(oauthTicket);
    delete q.oauth_ticket;
    delete q.oauth_provider;
    delete q.oauth_error;
    if (consumed) {
      finishTicketLogin(target);
    } else {
      window.$message?.error($t('api.RequestError'));
      await router.replace({ path: route.path, query: q, hash: route.hash });
    }
    return;
  }

  const ticket = mergedQuery.oidc_ticket;
  if (typeof ticket === 'string' && ticket) {
    const target = safeLoginRedirect(mergedQuery.redirect);
    const consumed = await authStore.loginByOidcTicket(ticket, false);
    delete q.oidc_ticket;
    delete q.oidc_error;
    if (consumed) {
      finishTicketLogin(target);
    } else {
      window.$message?.error($t('api.RequestError'));
      await router.replace({ path: route.path, query: q, hash: route.hash });
    }
    return;
  }

  if (shouldReplaceQuery) {
    await router.replace({ path: route.path, query: q, hash: route.hash });
  }
});
</script>

<template>
  <div class="relative size-full flex-center overflow-hidden p-16px" :style="backgroundStyle">
    <NCard
      :bordered="false"
      class="login-card relative z-4 w-full max-w-368px rd-12px lt-sm:max-h-[calc(100vh-32px)] lt-sm:overflow-y-auto"
    >
      <div class="w-full">
        <header class="flex-y-center justify-between">
          <SystemLogo class="text-48px text-primary lt-sm:text-36px" />
          <h3 class="text-22px text-primary font-500 lt-sm:text-18px">{{ $t('system.title') }}</h3>
          <div class="i-flex-col">
            <ThemeSchemaSwitch
              :theme-schema="themeStore.themeScheme"
              :show-tooltip="false"
              class="text-20px lt-sm:text-18px"
              @switch="themeStore.toggleThemeScheme"
            />
            <LangSwitch
              :lang="appStore.locale"
              :lang-options="appStore.localeOptions"
              :show-tooltip="false"
              @change-lang="appStore.changeLocale"
            />
          </div>
        </header>
        <main class="pt-16px">
          <h3 class="text-16px text-primary font-medium">{{ $t(activeModule.label) }}</h3>
          <div class="pt-16px">
            <Transition :name="themeStore.page.animateMode" mode="out-in" appear>
              <component :is="activeModule.component" />
            </Transition>
          </div>
        </main>
      </div>
    </NCard>
  </div>
</template>

<style scoped>
.login-card {
  background-color: rgb(var(--card-color) / 88%);
  box-shadow: 0 20px 55px rgb(0 0 0 / 28%);
  backdrop-filter: blur(14px);
}
</style>
