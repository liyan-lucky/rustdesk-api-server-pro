<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { $t } from '@/locales';
import { fetchStat } from '@/service/api/home';

defineOptions({
  name: 'ChangeLogs'
});

interface LogItem {
  id: number;
  content: string;
  version: string;
  time: string;
}

const compatVersion = ref('latest');

const logs = computed<LogItem[]>(() => [
  {
    id: 1,
    content: `${$t('page.home.changeLogs')}: RustDesk API compatibility update`,
    version: compatVersion.value,
    time: new Date().toISOString().slice(0, 10)
  }
]);

onMounted(async () => {
  const stat = (await fetchStat()).data;
  if (stat?.compatVersion) {
    compatVersion.value = stat.compatVersion;
  }
});
</script>

<template>
  <NCard :title="$t('page.home.changeLogs')" :bordered="false" size="small" segmented class="card-wrapper">
    <NList>
      <NListItem v-for="item in logs" :key="item.id">
        <template #prefix>
          <SoybeanAvatar class="size-48px!" />
        </template>
        <NThing :title="item.content" :title-extra="item.version" :description="item.time" />
      </NListItem>
    </NList>
  </NCard>
</template>

<style scoped></style>
