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

/** 返回主页：根据登录状态智能跳转，避免 403 死循环 */
function handleBackHome() {
  if (!authStore.isLogin) {
    // 未登录跳转登录页
    routerPushByKey('login', { params: { module: 'pwd-login' } });
    return;
  }
  // 已登录跳转 routeHome（会根据用户角色自动适配 home 或 user_profile）
  routerPushByKey(routeStore.routeHome);
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
