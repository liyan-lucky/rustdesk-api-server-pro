<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NTag } from 'naive-ui';
import { fetchMySecurityEvents } from '@/service/api/user-portal';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
const appStore = useAppStore();
const loading = ref(false);
const data = ref<any[]>([]);
const total = ref(0);
const page = ref(1);
const size = ref(10);
const columns = [
  { key: 'event', title: $t('dataMap.loginLog.event') },
  { key: 'ip', title: $t('dataMap.audit.ip') },
  {
    key: 'success',
    title: $t('dataMap.loginLog.success'),
    render: (r: any) => (
      <NTag type={r.success ? 'success' : 'error'}>
        {r.success ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}
      </NTag>
    )
  },
  { key: 'reason', title: $t('dataMap.loginLog.reason') },
  { key: 'created_at', title: $t('dataMap.audit.created_at') }
];
async function load() {
  try {
    loading.value = true;
    const { data: r } = await fetchMySecurityEvents({ current: page.value, size: size.value });
    if (r) {
      data.value = r.records;
      total.value = r.total;
    }
  } finally {
    loading.value = false;
  }
}
onMounted(load);
</script>

<template>
  <NCard :title="$t('route.workspace_security')" :bordered="false">
    <NDataTable
      :columns="columns"
      :data="data"
      :loading="loading"
      :flex-height="!appStore.isMobile"
      :scroll-x="900"
      :pagination="{page,pageSize:size,itemCount:total,onChange:(v:number)=>{page=v;load()},onUpdatePageSize:(v:number)=>{size=v;page=1;load()},showSizePicker:true,pageSlot:appStore.isMobile?3:9}"
    />
  </NCard>
</template>
