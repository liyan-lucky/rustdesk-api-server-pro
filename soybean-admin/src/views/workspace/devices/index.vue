<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NTag } from 'naive-ui';
import { fetchMyDevices } from '@/service/api/user-portal';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
const appStore = useAppStore();
const loading = ref(false);
const data = ref<any[]>([]);
const total = ref(0);
const page = ref(1);
const size = ref(10);
const columns = [
  { key: 'rustdesk_id', title: $t('dataMap.device.rustdesk_id') },
  { key: 'hostname', title: $t('dataMap.device.hostname') },
  { key: 'username', title: $t('dataMap.device.username') },
  { key: 'version', title: $t('dataMap.device.version') },
  { key: 'os', title: $t('dataMap.device.os') },
  {
    key: 'is_online',
    title: $t('page.myDevices.status'),
    render: (r: any) => (
      <NTag type={r.is_online ? 'success' : 'default'}>
        {r.is_online ? $t('page.myDevices.online') : $t('page.myDevices.offline')}
      </NTag>
    )
  },
  { key: 'updated_at', title: $t('dataMap.ab.updated_at') }
];
async function load() {
  try {
    loading.value = true;
    const { data: r } = await fetchMyDevices({ current: page.value, size: size.value });
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
  <NCard :title="$t('route.workspace_devices')" :bordered="false">
    <NDataTable
      :columns="columns"
      :data="data"
      :loading="loading"
      :flex-height="!appStore.isMobile"
      :scroll-x="1000"
      :pagination="{page,pageSize:size,itemCount:total,onChange:(v:number)=>{page=v;load()},onUpdatePageSize:(v:number)=>{size=v;page=1;load()},showSizePicker:true,pageSlot:appStore.isMobile?3:9}"
    />
  </NCard>
</template>
