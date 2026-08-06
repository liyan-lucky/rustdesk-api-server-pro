<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';

const appStore = useAppStore();

const errorCodes = ref<ErrorCodeEntry[]>([]);
const errorCodesLoading = ref(false);
const errorCodesError = ref('');
const searchQuery = ref('');
const selectedModule = ref('');

interface ErrorCodeEntry {
  code: string;
  message: string;
  module: string;
  description: string;
  solution: string;
  descriptionEn: string;
  solutionEn: string;
}

const modules = computed(() => {
  const set = new Set(errorCodes.value.map(e => e.module));
  return Array.from(set).sort();
});

const filteredErrorCodes = computed(() => {
  let list = errorCodes.value;
  if (selectedModule.value) {
    list = list.filter(e => e.module === selectedModule.value);
  }
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase();
    list = list.filter(
      e =>
        e.code.toLowerCase().includes(q) ||
        e.message.toLowerCase().includes(q) ||
        e.description.toLowerCase().includes(q) ||
        e.solution.toLowerCase().includes(q)
    );
  }
  return list;
});

async function loadErrorCodes() {
  errorCodesLoading.value = true;
  errorCodesError.value = '';
  try {
    const response = await fetch('/api/errcode', { cache: 'no-store' });
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    const payload = await response.json();
    errorCodes.value = payload?.errorCodes || [];
  } catch (error) {
    errorCodesError.value = error instanceof Error ? error.message : String(error);
  } finally {
    errorCodesLoading.value = false;
  }
}

onMounted(() => {
  loadErrorCodes();
});
</script>

<template>
  <NSpace vertical size="large" class="mt-16px">
    <NCard :bordered="false">
      <NAlert type="info" class="mb-16px">{{ $t('page.about.errcodeTip') }}</NAlert>
      <NSpace align="center" wrap class="mb-16px" :class="appStore.isMobile ? 'flex-col items-stretch' : ''">
        <NInput
          v-model:value="searchQuery"
          :placeholder="$t('page.about.searchPlaceholder')"
          clearable
          :class="appStore.isMobile ? 'w-full' : 'w-300px'"
        />
        <NSelect
          v-model:value="selectedModule"
          :options="modules.map(m => ({ label: m, value: m }))"
          :placeholder="$t('page.about.moduleFilter')"
          clearable
          :class="appStore.isMobile ? 'w-full' : 'w-160px'"
        />
      </NSpace>
      <NSpin :show="errorCodesLoading">
        <NAlert v-if="errorCodesError" type="error" class="mb-16px">{{ errorCodesError }}</NAlert>
        <NEmpty v-else-if="filteredErrorCodes.length === 0" :description="$t('common.noData')" />
        <NCollapse v-else>
          <NCollapseItem
            v-for="entry in filteredErrorCodes"
            :key="entry.code"
            :title="`${entry.code}  ${$t('api.' + entry.message as App.I18n.I18nKey)}`"
            :name="entry.code"
          >
            <NDescriptions bordered label-placement="left" :column="1" size="small">
              <NDescriptionsItem :label="$t('page.about.errCode')">
                <NTag type="warning" size="small">{{ entry.code }}</NTag>
              </NDescriptionsItem>
              <NDescriptionsItem :label="$t('page.about.errMessage')">
                {{ $t(`api.${entry.message}` as App.I18n.I18nKey) }}
              </NDescriptionsItem>
              <NDescriptionsItem :label="$t('page.about.errModule')">
                <NTag size="small">{{ entry.module }}</NTag>
              </NDescriptionsItem>
              <NDescriptionsItem :label="$t('page.about.errDescription')">
                {{ appStore.locale === 'zh-CN' ? entry.description : (entry.descriptionEn || entry.description) }}
              </NDescriptionsItem>
              <NDescriptionsItem :label="$t('page.about.errSolution')">
                <NAlert type="success" :bordered="false">{{ appStore.locale === 'zh-CN' ? entry.solution : (entry.solutionEn || entry.solution) }}</NAlert>
              </NDescriptionsItem>
            </NDescriptions>
          </NCollapseItem>
        </NCollapse>
      </NSpin>
    </NCard>
  </NSpace>
</template>
