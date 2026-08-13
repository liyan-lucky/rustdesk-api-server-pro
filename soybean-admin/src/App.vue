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

/* 玻璃透明效果 - 通过 --glass-bg / --glass-bg-strong / --glass-bg-weak / --glass-blur 动态控制 */
/* 这些变量由 theme store 在 JS 中计算并注入到 :root，避免 rgb(var()/var()) 嵌套兼容性问题 */

/* 覆盖布局容器 */
html.glass-effect .bg-container {
  background-color: var(--glass-bg) !important;
  backdrop-filter: blur(var(--glass-blur));
  box-shadow: 0 8px 24px rgb(0 0 0 / 12%);
}

html.glass-effect .bg-layout {
  background-color: var(--glass-layout-bg) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 覆盖 Naive UI 组件 CSS 变量，使所有卡片/弹窗/抽屉背景透明 */
html.glass-effect .n-card {
  --n-color: var(--glass-bg) !important;
  --n-color-modal: var(--glass-bg-strong) !important;
  --n-color-popover: var(--glass-bg-strong) !important;
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

/* 抽屉 - 覆盖 .n-drawer 和 .n-drawer-content 的 --n-color */
html.glass-effect .n-drawer {
  --n-color: var(--glass-bg-strong) !important;
  background-color: var(--glass-bg-strong) !important;
  backdrop-filter: blur(var(--glass-blur));
}

html.glass-effect .n-drawer-content {
  --n-color: var(--glass-bg-strong) !important;
  background-color: var(--glass-bg-strong) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 弹窗 */
html.glass-effect .n-modal .n-card {
  --n-color: var(--glass-bg-strong) !important;
  backdrop-filter: blur(var(--glass-blur));
  box-shadow: 0 20px 55px rgb(0 0 0 / 28%);
}

/* 标签页透明 */
html.glass-effect .n-tabs .n-tabs-nav {
  background-color: transparent !important;
}

/* NDataTable 数据表 - 覆盖所有 td/th color 变量 */
html.glass-effect .n-data-table {
  --n-td-color: var(--glass-bg) !important;
  --n-td-color-modal: var(--glass-bg-strong) !important;
  --n-td-color-popover: var(--glass-bg-strong) !important;
  --n-td-color-hover: var(--glass-bg-weak) !important;
  --n-td-color-hover-modal: var(--glass-bg-weak) !important;
  --n-td-color-hover-popover: var(--glass-bg-weak) !important;
  --n-td-color-sorting: var(--glass-bg-weak) !important;
  --n-td-color-sorting-modal: var(--glass-bg-weak) !important;
  --n-td-color-sorting-popover: var(--glass-bg-weak) !important;
  --n-td-color-striped: var(--glass-bg-weak) !important;
  --n-td-color-striped-modal: var(--glass-bg-weak) !important;
  --n-td-color-striped-popover: var(--glass-bg-weak) !important;
  --n-th-color: var(--glass-bg) !important;
  --n-th-color-modal: var(--glass-bg-strong) !important;
  --n-th-color-popover: var(--glass-bg-strong) !important;
  --n-th-color-hover: var(--glass-bg-weak) !important;
  --n-th-color-hover-modal: var(--glass-bg-weak) !important;
  --n-th-color-hover-popover: var(--glass-bg-weak) !important;
  --n-th-color-sorting: var(--glass-bg-weak) !important;
  --n-th-color-sorting-modal: var(--glass-bg-weak) !important;
  --n-th-color-sorting-popover: var(--glass-bg-weak) !important;
  --n-merged-td-color: var(--glass-bg) !important;
  --n-merged-td-color-hover: var(--glass-bg-weak) !important;
  --n-merged-td-color-sorting: var(--glass-bg-weak) !important;
  --n-merged-td-color-striped: var(--glass-bg-weak) !important;
  --n-merged-th-color: var(--glass-bg) !important;
  --n-merged-th-color-hover: var(--glass-bg-weak) !important;
  --n-merged-th-color-sorting: var(--glass-bg-weak) !important;
}

html.glass-effect .n-data-table .n-data-table-wrapper {
  background-color: transparent !important;
}

/* NDescriptions 描述列表 - 覆盖 td/th color 变量 */
html.glass-effect .n-descriptions {
  --n-td-color: var(--glass-bg) !important;
  --n-td-color-modal: var(--glass-bg-strong) !important;
  --n-td-color-popover: var(--glass-bg-strong) !important;
  --n-th-color: var(--glass-bg) !important;
  --n-th-color-modal: var(--glass-bg-strong) !important;
  --n-th-color-popover: var(--glass-bg-strong) !important;
  --n-merged-td-color: var(--glass-bg) !important;
  --n-merged-th-color: var(--glass-bg) !important;
  --n-merged-border-color: rgb(var(--primary-color) / 15%) !important;
}

/* 普通表格 */
html.glass-effect .n-table .n-table-wrapper {
  background-color: transparent !important;
}

/* 列表/菜单/其他容器透明 */
html.glass-effect .n-list,
html.glass-effect .n-menu,
html.glass-effect .n-layout-sider,
html.glass-effect .n-layout-header {
  background-color: var(--glass-bg) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 弹出层透明 */
html.glass-effect .n-popover,
html.glass-effect .n-popselect-menu {
  --n-color: var(--glass-bg-strong) !important;
  background-color: var(--glass-bg-strong) !important;
  backdrop-filter: blur(var(--glass-blur));
}

/* 输入框、选择器等表单容器保持轻微透明 */
html.glass-effect .n-input,
html.glass-effect .n-base-selection {
  --n-color: var(--glass-bg-weak) !important;
}
</style>
