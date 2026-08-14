import { computed, effectScope, nextTick, onScopeDispose, ref, toRefs, watch } from 'vue';
import type { Ref } from 'vue';
import { defineStore } from 'pinia';
import { usePreferredColorScheme } from '@vueuse/core';
import { getPaletteColorByNumber } from '@sa/color';
import { SetupStoreId } from '@/enum';
import { localStg } from '@/utils/storage';
import {
  addThemeVarsToGlobal,
  createThemeToken,
  getNaiveTheme,
  initThemeSettings,
  toggleAuxiliaryColorModes,
  toggleCssDarkMode
} from './shared';

/** Theme store */
export const useThemeStore = defineStore(SetupStoreId.Theme, () => {
  const scope = effectScope();
  const osTheme = usePreferredColorScheme();

  /** Theme settings */
  const settings: Ref<App.Theme.ThemeSetting> = ref(initThemeSettings());

  /** Dark mode */
  const darkMode = computed(() => {
    if (settings.value.themeScheme === 'auto') {
      return osTheme.value === 'dark';
    }
    return settings.value.themeScheme === 'dark';
  });

  /** grayscale mode */
  const grayscaleMode = computed(() => settings.value.grayscale);

  /** colourWeakness mode */
  const colourWeaknessMode = computed(() => settings.value.colourWeakness);

  /** glass effect mode */
  const glassEffectMode = computed(() => settings.value.glassEffect);

  /** glass effect blur strength */
  const glassBlur = computed(() => settings.value.glassBlur);

  /** glass effect opacity */
  const glassOpacity = computed(() => settings.value.glassOpacity);

  /** Theme colors */
  const themeColors = computed(() => {
    const { themeColor, otherColor, isInfoFollowPrimary } = settings.value;
    const colors: App.Theme.ThemeColor = {
      primary: themeColor,
      ...otherColor,
      info: isInfoFollowPrimary ? themeColor : otherColor.info
    };
    return colors;
  });

  /** Naive theme */
  const naiveTheme = computed(() => getNaiveTheme(themeColors.value, settings.value.recommendColor));

  /**
   * Settings json
   *
   * It is for copy settings
   */
  const settingsJson = computed(() => JSON.stringify(settings.value));

  /** Reset store */
  function resetStore() {
    const themeStore = useThemeStore();

    themeStore.$reset();
  }

  /**
   * Set theme scheme
   *
   * @param themeScheme
   */
  function setThemeScheme(themeScheme: UnionKey.ThemeScheme) {
    settings.value.themeScheme = themeScheme;
  }

  /**
   * Set grayscale value
   *
   * @param isGrayscale
   */
  function setGrayscale(isGrayscale: boolean) {
    settings.value.grayscale = isGrayscale;
  }

  /**
   * Set colourWeakness value
   *
   * @param isColourWeakness
   */
  function setColourWeakness(isColourWeakness: boolean) {
    settings.value.colourWeakness = isColourWeakness;
  }

  /**
   * Set glass effect value
   *
   * @param isGlassEffect
   */
  function setGlassEffect(isGlassEffect: boolean) {
    settings.value.glassEffect = isGlassEffect;
  }

  /**
   * Set glass blur strength
   *
   * @param blur blur strength in px (0-30)
   */
  function setGlassBlur(blur: number) {
    settings.value.glassBlur = blur;
  }

  /**
   * Set glass opacity
   *
   * @param opacity opacity in percent (0-100)
   */
  function setGlassOpacity(opacity: number) {
    settings.value.glassOpacity = opacity;
  }

  /** Toggle theme scheme */
  function toggleThemeScheme() {
    const themeSchemes: UnionKey.ThemeScheme[] = ['light', 'dark', 'auto'];

    const index = themeSchemes.findIndex(item => item === settings.value.themeScheme);

    const nextIndex = index === themeSchemes.length - 1 ? 0 : index + 1;

    const nextThemeScheme = themeSchemes[nextIndex];

    setThemeScheme(nextThemeScheme);
  }

  /**
   * Update theme colors
   *
   * @param key Theme color key
   * @param color Theme color
   */
  function updateThemeColors(key: App.Theme.ThemeColorKey, color: string) {
    let colorValue = color;

    if (settings.value.recommendColor) {
      // get a color palette by provided color and color name, and use the suitable color

      colorValue = getPaletteColorByNumber(color, 500, true);
    }

    if (key === 'primary') {
      settings.value.themeColor = colorValue;
    } else {
      settings.value.otherColor[key] = colorValue;
    }
  }

  /**
   * Set theme layout
   *
   * @param mode Theme layout mode
   */
  function setThemeLayout(mode: UnionKey.ThemeLayoutMode) {
    settings.value.layout.mode = mode;
  }

  /** Setup theme vars to global */
  function setupThemeVarsToGlobal() {
    const { themeTokens, darkThemeTokens } = createThemeToken(
      themeColors.value,
      settings.value.tokens,
      settings.value.recommendColor
    );
    addThemeVarsToGlobal(themeTokens, darkThemeTokens);
  }
  /**
   * Set layout reverse horizontal mix
   *
   * @param reverse Reverse horizontal mix
   */
  function setLayoutReverseHorizontalMix(reverse: boolean) {
    settings.value.layout.reverseHorizontalMix = reverse;
  }

  /** Cache theme settings */
  function cacheThemeSettings() {
    localStg.set('themeSettings', settings.value);
  }

  // 主题配置变化时即时写入 localStorage，确保关闭浏览器后配置不丢失
  watch(
    settings,
    () => {
      cacheThemeSettings();
    },
    { deep: true }
  );

  // watch store
  scope.run(() => {
    // watch dark mode
    watch(
      darkMode,
      val => {
        toggleCssDarkMode(val);
      },
      { immediate: true }
    );

    watch(
      [grayscaleMode, colourWeaknessMode],
      val => {
        toggleAuxiliaryColorModes(val[0], val[1]);
      },
      { immediate: true }
    );

    // watch glass effect mode, toggle css class on html element and inject css vars
    watch(
      [glassEffectMode, glassBlur, glassOpacity],
      ([enabled, blur, opacity]) => {
        const htmlClass = 'glass-effect';
        const root = document.documentElement;
        // 读取主题色 CSS 变量，生成完整的 rgba 颜色直接注入
        // 透明度语义：数值越高越透明，alpha = (100 - opacity) / 100
        // alphaStrong 用于弹窗/抽屉等需要更不透明的容器（比主背景不透明度高 8%）
        // alphaWeak 用于布局背景等可以更透明的容器（比主背景透明度高 12%）
        const containerBg = getComputedStyle(root).getPropertyValue('--container-bg-color').trim();
        const layoutBg = getComputedStyle(root).getPropertyValue('--layout-bg-color').trim();
        const alpha = Math.max(0, Math.min(1, (100 - opacity) / 100));
        const alphaStrong = Math.max(0, Math.min(1, (100 - opacity + 8) / 100));
        const alphaWeak = Math.max(0, Math.min(1, (100 - opacity - 12) / 100));
        if (containerBg) {
          root.style.setProperty('--glass-bg', `rgb(${containerBg} / ${alpha})`);
          root.style.setProperty('--glass-bg-strong', `rgb(${containerBg} / ${alphaStrong})`);
          root.style.setProperty('--glass-bg-weak', `rgb(${containerBg} / ${alphaWeak})`);
        }
        if (layoutBg) {
          root.style.setProperty('--glass-layout-bg', `rgb(${layoutBg} / ${alphaWeak})`);
        }
        root.style.setProperty('--glass-blur', `${blur}px`);
        if (enabled) {
          root.classList.add(htmlClass);
        } else {
          root.classList.remove(htmlClass);
        }
      },
      { immediate: true }
    );

    // themeColors change, update css vars and storage theme color
    watch(
      themeColors,
      val => {
        setupThemeVarsToGlobal();
        localStg.set('themeColor', val.primary);
        // 主题色变化后需要重新计算 glass 变量（因为依赖 container-bg-color）
        nextTick(() => {
          const root = document.documentElement;
          const containerBg = getComputedStyle(root).getPropertyValue('--container-bg-color').trim();
          const layoutBg = getComputedStyle(root).getPropertyValue('--layout-bg-color').trim();
          const opacity = settings.value.glassOpacity;
          const alpha = Math.max(0, Math.min(1, (100 - opacity) / 100));
          const alphaStrong = Math.max(0, Math.min(1, (100 - opacity + 8) / 100));
          const alphaWeak = Math.max(0, Math.min(1, (100 - opacity - 12) / 100));
          if (containerBg) {
            root.style.setProperty('--glass-bg', `rgb(${containerBg} / ${alpha})`);
            root.style.setProperty('--glass-bg-strong', `rgb(${containerBg} / ${alphaStrong})`);
            root.style.setProperty('--glass-bg-weak', `rgb(${containerBg} / ${alphaWeak})`);
          }
          if (layoutBg) {
            root.style.setProperty('--glass-layout-bg', `rgb(${layoutBg} / ${alphaWeak})`);
          }
        });
      },
      { immediate: true }
    );
  });

  /** On scope dispose */
  onScopeDispose(() => {
    scope.stop();
  });

  return {
    ...toRefs(settings.value),
    darkMode,
    themeColors,
    naiveTheme,
    settingsJson,
    setGrayscale,
    setColourWeakness,
    setGlassEffect,
    setGlassBlur,
    setGlassOpacity,
    resetStore,
    setThemeScheme,
    toggleThemeScheme,
    updateThemeColors,
    setThemeLayout,
    setLayoutReverseHorizontalMix
  };
});
