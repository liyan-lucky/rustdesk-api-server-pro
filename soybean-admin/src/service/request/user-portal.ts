import { BACKEND_ERROR_CODE, createFlatRequest } from '@sa/axios';
import { $t } from '@/locales';
import { localStg } from '@/utils/storage';
import { showErrorMsg } from './shared';
import type { RequestInstanceState } from './type';

function parseBackendMessage(raw: string): { code: string; message: string; i18nKey: string } {
  if (raw.startsWith('ERR-')) {
    const colonIndex = raw.indexOf(':');
    if (colonIndex !== -1) {
      const extractedCode = raw.substring(0, colonIndex).trim();
      const extractedMessage = raw.substring(colonIndex + 1).trim();
      return { code: extractedCode, message: extractedMessage, i18nKey: `api.${extractedMessage}` };
    }
  }
  return { code: '', message: raw, i18nKey: `api.${raw}` };
}

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
        localStg.remove('isAdmin');
        window.location.href = '/#/login/pwd-login';
        return;
      }
      let message = error.message;
      if (error.code === BACKEND_ERROR_CODE) {
        message = error.response?.data?.message || message;
      }
      const parsed = parseBackendMessage(message);
      const displayMsg = parsed.code
        ? `[${parsed.code}] ${$t(parsed.i18nKey as App.I18n.I18nKey)}`
        : $t(parsed.i18nKey as App.I18n.I18nKey);
      showErrorMsg(userPortalRequest.state, displayMsg);
    }
  }
);
