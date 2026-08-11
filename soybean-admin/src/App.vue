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

/* 玻璃透明效果 - 开启时所有容器获得毛玻璃透明背景 */
html.glass-effect .bg-container {
  background-color: rgb(var(--container-bg-color) / 88%) !important;
  backdrop-filter: blur(14px);
  box-shadow: 0 8px 24px rgb(0 0 0 / 12%);
}

html.glass-effect .bg-layout {
  background-color: rgb(var(--layout-bg-color) / 80%) !important;
  backdrop-filter: blur(10px);
}

html.glass-effect .n-card {
  background-color: rgb(var(--container-bg-color) / 88%) !important;
  backdrop-filter: blur(14px);
  box-shadow: 0 8px 24px rgb(0 0 0 / 12%);
}

html.glass-effect .n-card.n-card--bordered {
  border-color: rgb(var(--primary-color) / 15%);
}

html.glass-effect .n-drawer-content {
  background-color: rgb(var(--container-bg-color) / 92%) !important;
  backdrop-filter: blur(16px);
}

html.glass-effect .n-modal .n-card {
  background-color: rgb(var(--container-bg-color) / 92%) !important;
  backdrop-filter: blur(16px);
  box-shadow: 0 20px 55px rgb(0 0 0 / 28%);
}

html.glass-effect .n-tabs .n-tabs-nav {
  background-color: transparent !important;
}

html.glass-effect .n-table .n-table-wrapper {
  background-color: transparent !important;
}
</style>
