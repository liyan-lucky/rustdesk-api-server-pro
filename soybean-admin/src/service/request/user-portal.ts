import { createFlatRequest, BACKEND_ERROR_CODE } from '@sa/axios';
import { $t } from '@/locales';
import { localStg } from '@/utils/storage';
import { showErrorMsg } from './shared';
import type { RequestInstanceState } from './type';

const isHttpProxy = import.meta.env.DEV && import.meta.env.VITE_HTTP_PROXY === 'Y';
const userPortalBaseURL = isHttpProxy ? '/proxy-user-portal' : '/user-portal';

export const userPortalRequest = createFlatRequest<App.Service.Response, RequestInstanceState>(
  {
    baseURL: userPortalBaseURL,
    headers: {}
  },
  {
    async onRequest(config) {
      const { headers } = config;
      const token = localStg.get('token');
      const Authorization = token ? `${token}` : null;
      Object.assign(headers, { Authorization });
      return config;
    },
    isBackendSuccess(response) {
      return String(response.data.code) === import.meta.env.VITE_SERVICE_SUCCESS_CODE;
    },
    async onBackendFail(_response) {
      return null;
    },
    transformBackendResponse(response) {
      return { data: response.data.data, message: response.data.message };
    },
    onError(error) {
      if (error.response?.status === 401) {
        localStg.remove('token');
        localStg.remove('userType');
        window.location.href = '/#/login/user-login';
        return;
      }
      let message = error.message;
      if (error.code === BACKEND_ERROR_CODE) {
        message = error.response?.data?.message || message;
      }
      showErrorMsg(userPortalRequest.state, $t(`api.${message}` as App.I18n.I18nKey));
    }
  }
);
