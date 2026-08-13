<script lang="ts" setup>
import { computed } from 'vue';
import { $t } from '@/locales';
import { useRouterPush } from '@/hooks/common/router';
import { useAuthStore } from '@/store/modules/auth';
import { useRouteStore } from '@/store/modules/route';

defineOptions({ name: 'ExceptionBase' });

type ExceptionType = '403' | '404' | '500';

interface Props {
  /**
   * Exception type
   *
   * - 403: no permission
   * - 404: not found
   * - 500: service error
   */
  type: ExceptionType;
}

const props = defineProps<Props>();

const { routerPushByKey } = useRouterPush();
const authStore = useAuthStore();
const routeStore = useRouteStore();

const iconMap: Record<ExceptionType, string> = {
  '403': 'no-permission',
  '404': 'not-found',
  '500': 'service-error'
};

const icon = computed(() => iconMap[props.type]);

/**
 * 返回首页
 *
 * 当用户已登录但 auth route 尚未初始化（例如直接通过收藏访问 /#/403）时，
 * 需要先初始化 userInfo 和 auth route，确保 root 路由的 redirect 已根据角色
 * 更新（普通用户更新为 /user/profile），避免跳转到 /home 后因无权限再次回到 403 形成死循环。
 */
async function handleBackHome() {
  if (!authStore.isLogin) {
    routerPushByKey('login', { params: { module: 'pwd-login' } });
    return;
  }
  if (!routeStore.isInitAuthRoute) {
    await authStore.initUserInfo();
    if (authStore.isLogin) {
      await routeStore.initAuthRoute();
    }
  }
  await routerPushByKey('root');
}
</script>

<template>
  <div class="size-full min-h-520px flex-col-center gap-24px overflow-hidden">
    <div class="flex text-400px text-primary">
      <SvgIcon :local-icon="icon" />
    </div>
    <NButton type="primary" @click="handleBackHome">{{ $t('common.backToHome') }}</NButton>
  </div>
</template>

<style scoped></style>
