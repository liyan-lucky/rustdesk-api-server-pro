import type { RouteMeta } from 'vue-router';
import ElegantVueRouter from '@elegant-router/vue/vite';
import type { RouteKey } from '@elegant-router/types';

export function setupElegantRouter() {
  return ElegantVueRouter({
    layouts: {
      base: 'src/layouts/base-layout/index.vue',
      blank: 'src/layouts/blank-layout/index.vue'
    },
    routePathTransformer(routeName, routePath) {
      const key = routeName as RouteKey;

      if (key === 'login') {
        const modules: UnionKey.LoginModule[] = ['pwd-login'];

        const moduleReg = modules.join('|');

        return `/login/:module(${moduleReg})?`;
      }

      return routePath;
    },
    onRouteMetaGen(routeName) {
      const key = routeName as RouteKey;

      const constantRoutes: RouteKey[] = ['login', '403', '404', '500'];

      const meta: Partial<RouteMeta> = {
        title: key,
        i18nKey: `route.${key}` as App.I18n.I18nKey
      };

      if (constantRoutes.includes(key)) {
        meta.constant = true;
      }

      if (key === 'workspace' || key.startsWith('workspace_')) {
        meta.roles = ['R_USER'];
        meta.hideInMenu = true;
      }

      const adminOnlyRoutes: RouteKey[] = [
        'home', 'user_list', 'user_sessions', 'audit_baselogs', 'audit_filetransferlogs',
        'system_mail', 'system_mail_logs', 'system_mail_template', 'system_oauth', 'system_tokens'
      ];
      if (adminOnlyRoutes.includes(key)) {
        meta.roles = ['R_SUPER'];
      }

      if (key === 'user_profile') meta.roles = ['R_USER'];

      const routeMeta: Partial<Record<RouteKey, Partial<RouteMeta>>> = {
        about: { icon: 'mdi:information-outline', order: 99 },
        workspace: { icon: 'mdi:account-circle', order: 2 },
        workspace_overview: { icon: 'mdi:view-dashboard-outline', order: 1 },
        workspace_devices: { icon: 'mdi:monitor-multiple', order: 2 },
        workspace_sessions: { icon: 'mdi:devices', order: 3 },
        workspace_security: { icon: 'mdi:shield-account', order: 4 },
        workspace_profile: { icon: 'mdi:card-account-details-outline', order: 5 },
        user_profile: { icon: 'mdi:card-account-details-outline', order: 1 },
        'my-devices_peers': { order: 1 },
        'my-devices_tags': { order: 2 },
        'my-devices_manage': { order: 3 }
      };

      Object.assign(meta, routeMeta[key]);

      return meta;
    }
  });
}
