import { computed, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import { defineStore } from 'pinia';
import { useLoading } from '@sa/hooks';
import { SetupStoreId } from '@/enum';
import { useRouterPush } from '@/hooks/common/router';
import { fetchGetUserInfo, fetchLogin } from '@/service/api';
import {
  fetchOAuthTicketToken,
  fetchOidcTicketToken,
  fetchWebauthnLoginBegin,
  fetchWebauthnLoginFinish
} from '@/service/api/auth';
import { localStg } from '@/utils/storage';
import { appendVersion, getVersionTag } from '@/utils/version';
import { isWebAuthnSupported, prepareAssertionOptions, serializeAssertion } from '@/utils/webauthn';
import { $t } from '@/locales';
import { useRouteStore } from '../route';
import { useTabStore } from '../tab';
import { clearAuthStorage, getToken } from './shared';

export const useAuthStore = defineStore(SetupStoreId.Auth, () => {
  const route = useRoute();
  const routeStore = useRouteStore();
  const tabStore = useTabStore();
  const { toLogin, redirectFromLogin } = useRouterPush(false);
  const { loading: loginLoading, startLoading, endLoading } = useLoading();

  const token = ref(getToken());

  const userInfo: Api.Auth.UserInfo = reactive({
    userId: '',
    userName: '',
    roles: [],
    buttons: []
  });

  /** is super role in static route */
  const isStaticSuper = computed(() => {
    const { VITE_AUTH_ROUTE_MODE, VITE_STATIC_SUPER_ROLE } = import.meta.env;

    return VITE_AUTH_ROUTE_MODE === 'static' && userInfo.roles.includes(VITE_STATIC_SUPER_ROLE);
  });

  /** Is login */
  const isLogin = computed(() => Boolean(token.value));

  /** Reset auth store */
  async function resetStore() {
    clearAuthStorage();

    token.value = '';
    Object.assign(userInfo, { userId: '', userName: '', roles: [], buttons: [] });

    if (!route.meta.constant) {
      await toLogin();
    }

    tabStore.cacheTabs();
    routeStore.resetStore();
  }

  /**
   * Login
   *
   * @param userName User name
   * @param password Password
   * @param [redirect=true] Whether to redirect after login. Default is `true`
   */
  async function login(model: Api.Form.LoginForm, redirect = true) {
    startLoading();

    const { data: loginToken, error } = await fetchLogin(model);

    if (!error) {
      const pass = await applyTokenAndBootstrap(loginToken);

      if (pass) {
        await routeStore.initAuthRoute();

        if (redirect) {
          await redirectFromLogin();
        }

        if (routeStore.isInitAuthRoute) {
          window.$notification?.success({
            title: `${$t('page.login.common.loginSuccess')} (${getVersionTag()})`,
            content: appendVersion($t('page.login.common.welcomeBack', { userName: userInfo.userName })),
            duration: 4500
          });
        }
      } else {
        resetStore();
      }
    } else {
      resetStore();
    }

    endLoading();
    return error;
  }

  async function applyTokenAndBootstrap(loginToken: Api.Auth.LoginToken) {
    localStg.set('token', loginToken.token);
    if (loginToken.isAdmin !== undefined) {
      localStg.set('isAdmin', loginToken.isAdmin);
    }

    const pass = await getUserInfo();

    if (pass) {
      token.value = loginToken.token;

      return true;
    }

    clearAuthStorage();
    return false;
  }

  async function loginByOidcTicket(ticket: string, redirect = true) {
    startLoading();
    const { data, error } = await fetchOidcTicketToken(ticket);
    if (!error && data?.token) {
      const pass = await applyTokenAndBootstrap(data);
      if (pass) {
        await routeStore.initAuthRoute();
        if (redirect) {
          await redirectFromLogin();
        }
      } else {
        resetStore();
      }
    }
    endLoading();
    return error;
  }

  async function loginByOAuthTicket(ticket: string, redirect = true) {
    startLoading();
    const { data, error } = await fetchOAuthTicketToken(ticket);
    if (!error && data?.token) {
      const pass = await applyTokenAndBootstrap(data);
      if (pass) {
        await routeStore.initAuthRoute();
        if (redirect) {
          await redirectFromLogin();
        }
      } else {
        resetStore();
      }
    }
    endLoading();
    return error;
  }

  /**
   * 通过 Passkey (WebAuthn) 登录
   *
   * @param username 用户名（用于查找已绑定的 Passkey 凭据）
   * @param [redirect=true] 登录后是否跳转
   */
  async function loginByPasskey(username: string, redirect = true) {
    if (!isWebAuthnSupported()) {
      window.$message?.error($t('page.login.passkey.notSupported'));
      return true;
    }

    startLoading();

    try {
      const { data: beginData, error: beginError } = await fetchWebauthnLoginBegin(username);
      if (beginError || !beginData?.publicKey) {
        window.$message?.error($t('page.login.passkey.beginFailed'));
        endLoading();
        return beginError || true;
      }

      const publicKey = prepareAssertionOptions(beginData.publicKey);
      const credential = (await navigator.credentials.get({ publicKey })) as PublicKeyCredential | null;
      if (!credential) {
        endLoading();
        return true;
      }

      const serialized = serializeAssertion(credential);
      const { data: loginToken, error: finishError } = await fetchWebauthnLoginFinish(serialized);
      if (!finishError && loginToken?.token) {
        const pass = await applyTokenAndBootstrap(loginToken);
        if (pass) {
          await routeStore.initAuthRoute();
          if (redirect) {
            await redirectFromLogin();
          }
          if (routeStore.isInitAuthRoute) {
            window.$notification?.success({
              title: `${$t('page.login.common.loginSuccess')} (${getVersionTag()})`,
              content: appendVersion($t('page.login.common.welcomeBack', { userName: userInfo.userName })),
              duration: 4500
            });
          }
        } else {
          resetStore();
        }
      } else {
        resetStore();
      }

      endLoading();
      return finishError;
    } catch (e) {
      window.$message?.error($t('page.login.passkey.cancelled'));
      endLoading();
      return true;
    }
  }

  async function getUserInfo() {
    const { data: info, error } = await fetchGetUserInfo();

    if (!error) {
      // update store
      Object.assign(userInfo, info);

      return true;
    }

    return false;
  }

  async function initUserInfo() {
    const hasToken = getToken();

    if (hasToken) {
      const pass = await getUserInfo();

      if (!pass) {
        await resetStore();
      }
    }
  }

  return {
    token,
    userInfo,
    isStaticSuper,
    isLogin,
    loginLoading,
    resetStore,
    login,
    loginByPasskey,
    loginByOidcTicket,
    loginByOAuthTicket,
    initUserInfo
  };
});
