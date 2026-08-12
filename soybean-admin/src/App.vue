<script setup lang="ts">
import { computed } from 'vue';
import { NConfigProvider, darkTheme } from 'naive-ui';
import type { WatermarkProps } from 'naive-ui';
import { useAppStore } from './store/modules/app';
import { useThemeStore } from './store/modules/theme';
import { naiveDateLocales, naiveLocales } from './locales/naive';
import { useWebBackground } from './hooks/common/web-background';

defineOptions({
  name: 'App'
});

const appStore = useAppStore();
const themeStore = useThemeStore();
const { globalBackgroundEnabled, backgroundStyle } = useWebBackground();

const naiveDarkTheme = computed(() => (themeStore.darkMode ? darkTheme : undefined));

const naiveLocale = computed(() => {
  return naiveLocales[appStore.locale];
});

const naiveDateLocale = computed(() => {
  return naiveDateLocales[appStore.locale];
});

const watermarkProps = computed<WatermarkProps>(() => {
  return {
    content: themeStore.watermark?.text || 'SoybeanAdmin',
    cross: true,
    fullscreen: true,
    fontSize: 16,
    lineHeight: 16,
    width: 384,
    height: 384,
    xOffset: 12,
    yOffset: 60,
    rotate: -15,
    zIndex: 9999
  };
});
</script>

<template>
  <NConfigProvider
    :theme="naiveDarkTheme"
    :theme-overrides="themeStore.naiveTheme"
    :locale="naiveLocale"
    :date-locale="naiveDateLocale"
    class="h-full"
  >
    <div class="relative h-full overflow-hidden" :class="{ 'web-background-enabled': globalBackgroundEnabled }">
      <div v-if="globalBackgroundEnabled" class="pointer-events-none absolute inset-0" :style="backgroundStyle"></div>
      <AppProvider>
        <RouterView class="relative z-1 bg-layout" />
        <NWatermark v-if="themeStore.watermark?.visible" v-bind="watermarkProps" />
      </AppProvider>
    </div>
  </NConfigProvider>
</template>

<style>
.web-background-enabled .bg-layout {
  background-color: rgb(var(--base-color) / 72%) !important;
  backdrop-filter: blur(5px);
}

/* 玻璃透明效果 - 通过 --glass-blur 和 --glass-opacity 动态控制 */
/* 覆盖布局容器 */
html.glass-effect .bg-container {
  background-color: rgb(var(--container-bg-color) / var(--glass-opacity)) !important;
  backdrop-filter: blur(var(--glass-blur));
  box-shadow: 0 8px 24px rgb(0 0 0 / 12%);
}

html.glass-effect .bg-layout {
  background-color: rgb(var(--layout-bg-color) / var(--glass-opacity-weak)) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 覆盖 Naive UI 组件 CSS 变量，使所有卡片/弹窗/抽屉背景透明 */
html.glass-effect .n-card {
  --n-color: rgb(var(--container-bg-color) / var(--glass-opacity)) !important;
  --n-color-modal: rgb(var(--container-bg-color) / var(--glass-opacity-strong)) !important;
  --n-color-popover: rgb(var(--container-bg-color) / var(--glass-opacity-strong)) !important;
  --n-color-embedded: transparent !important;
  backdrop-filter: blur(var(--glass-blur));
  box-shadow: 0 8px 24px rgb(0 0 0 / 12%);
}

html.glass-effect .n-card.n-card--bordered {
  --n-border-color: rgb(var(--primary-color) / 15%) !important;
}

/* 卡片操作区/底部也透明 */
html.glass-effect .n-card .n-card__action,
html.glass-effect .n-card .n-card__footer {
  background-color: transparent !important;
}

/* 抽屉 */
html.glass-effect .n-drawer-content {
  --n-color: rgb(var(--container-bg-color) / var(--glass-opacity-strong)) !important;
  background-color: rgb(var(--container-bg-color) / var(--glass-opacity-strong)) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 弹窗 */
html.glass-effect .n-modal .n-card {
  --n-color: rgb(var(--container-bg-color) / var(--glass-opacity-strong)) !important;
  backdrop-filter: blur(var(--glass-blur));
  box-shadow: 0 20px 55px rgb(0 0 0 / 28%);
}

/* 标签页、表格透明 */
html.glass-effect .n-tabs .n-tabs-nav {
  background-color: transparent !important;
}

html.glass-effect .n-table .n-table-wrapper,
html.glass-effect .n-data-table .n-data-table-wrapper {
  background-color: transparent !important;
}

/* 列表/菜单/其他容器透明 */
html.glass-effect .n-list,
html.glass-effect .n-menu,
html.glass-effect .n-layout-sider,
html.glass-effect .n-layout-header {
  background-color: rgb(var(--container-bg-color) / var(--glass-opacity)) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 弹出层透明 */
html.glass-effect .n-popover,
html.glass-effect .n-popselect-menu {
  --n-color: rgb(var(--container-bg-color) / var(--glass-opacity-strong)) !important;
  background-color: rgb(var(--container-bg-color) / var(--glass-opacity-strong)) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 输入框、选择器等表单容器保持轻微透明 */
html.glass-effect .n-input,
html.glass-effect .n-base-selection {
  --n-color: rgb(var(--container-bg-color) / var(--glass-opacity-weak)) !important;
}
</style>
