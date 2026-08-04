<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { $t } from '@/locales';
import { fetchStat } from '@/service/api/home';
import { getBuildTime, getVersionTag } from '@/utils/version';

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
const serverVersion = ref('');
const buildTime = ref('');
const appVersion = getVersionTag();
const frontendBuildTime = getBuildTime();

const logs = computed<LogItem[]>(() => [
  {
    id: 1,
    content: 'RustDesk API 兼容能力已更新，支持客户端 1.4.9 主流程',
    version: `${compatVersion.value} / ${serverVersion.value || appVersion}`,
    time: buildTime.value || frontendBuildTime || new Date().toISOString().slice(0, 10)
  }
]);

onMounted(async () => {
  const stat = (await fetchStat()).data;
  if (stat?.compatVersion) {
    compatVersion.value = stat.compatVersion;
  }
  if (stat?.serverVersion) {
    serverVersion.value = stat.serverVersion;
  }
  if (stat?.buildTime) {
    buildTime.value = stat.buildTime;
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
