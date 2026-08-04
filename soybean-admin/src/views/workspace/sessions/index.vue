<script setup lang="tsx">
import { onMounted, ref } from 'vue';
import { NButton, NPopconfirm, NTag } from 'naive-ui';
import { fetchMySessions, revokeMySessions } from '@/service/api/user-portal';
import { $t } from '@/locales';
const loading = ref(false);
const data = ref<any[]>([]);
const total = ref(0);
const page = ref(1);
const size = ref(10);
async function revoke(row: any) {
  const { error } = await revokeMySessions([row.id]);
  if (!error) {
    window.$message?.success($t('common.updateSuccess'));
    load();
  }
}
const columns = [
  { key: 'device_name', title: $t('dataMap.token.device_name') },
  { key: 'rustdesk_id', title: $t('dataMap.device.rustdesk_id') },
  { key: 'device_os', title: $t('dataMap.token.device_os') },
  { key: 'expired', title: $t('dataMap.session.expired') },
  { key: 'created_at', title: $t('dataMap.session.created_at') },
  {
    key: 'current',
    title: $t('page.workspace.currentSession'),
    render: (r: any) => (
      <NTag type={r.current ? 'success' : 'default'}>
        {r.current ? $t('common.yesOrNo.yes') : $t('common.yesOrNo.no')}
      </NTag>
    )
  },
  {
    key: 'action',
    title: $t('common.action'),
    render: (r: any) =>
      r.current ? (
        '-'
      ) : (
        <NPopconfirm onPositiveClick={() => revoke(r)}>
          {{
            default: () => $t('page.workspace.revokeConfirm'),
            trigger: () => (
              <NButton size="small" type="error">
                {$t('page.workspace.revoke')}
              </NButton>
            )
          }}
        </NPopconfirm>
      )
  }
];
async function load() {
  loading.value = true;
  const { data: r } = await fetchMySessions({ current: page.value, size: size.value });
  if (r) {
    data.value = r.records;
    total.value = r.total;
  }
  loading.value = false;
}
onMounted(load);
</script>

<template>
  <NCard :title="$t('route.workspace_sessions')" :bordered="false">
    <NDataTable
      :columns="columns"
      :data="data"
      :loading="loading"
      :scroll-x="1100"
      :pagination="{page,pageSize:size,itemCount:total,onChange:(v:number)=>{page=v;load()},onUpdatePageSize:(v:number)=>{size=v;page=1;load()},showSizePicker:true}"
    />
  </NCard>
</template>
