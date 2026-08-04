<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useAuthStore } from '@/store/modules/auth';
import { fetchUserPortalInfo } from '@/service/api/user-portal';
import { $t } from '@/locales';
const auth = useAuthStore();
const isAdmin = computed(() => auth.userInfo.roles.includes('R_SUPER'));
const profile = ref<any>({});
onMounted(async () => {
  const { data } = await fetchUserPortalInfo();
  if (data) profile.value = data;
});
</script>

<template>
  <NCard :title="$t('route.workspace_profile')" :bordered="false" class="max-w-720px">
    <NDescriptions label-placement="left" bordered :column="1">
      <NDescriptionsItem :label="$t('dataMap.user.username')">{{ profile.username || '-' }}</NDescriptionsItem>
      <NDescriptionsItem :label="$t('dataMap.user.name')">{{ auth.userInfo.userName }}</NDescriptionsItem>
      <NDescriptionsItem :label="$t('dataMap.user.email')">{{ profile.email || '-' }}</NDescriptionsItem>
      <NDescriptionsItem :label="$t('dataMap.user.licensed_devices')">
        {{ profile.licensedDevices ?? 0 }}
      </NDescriptionsItem>
      <NDescriptionsItem :label="$t('page.workspace.accountRole')">
        <NTag :type="isAdmin ? 'warning' : 'info'">
          {{ isAdmin ? $t('page.workspace.adminRole') : $t('page.workspace.userRole') }}
        </NTag>
      </NDescriptionsItem>
      <NDescriptionsItem :label="$t('page.workspace.permissionScope')">
        {{ $t('page.workspace.userScope') }}
      </NDescriptionsItem>
    </NDescriptions>
  </NCard>
</template>
