<script setup lang="ts">
import { computed } from 'vue';
import { $t } from '@/locales';
import { useAuthStore } from '@/store/modules/auth';

defineOptions({
  name: 'TableHeader'
});

const authStore = useAuthStore();
const isAdmin = computed(() => authStore.userInfo.roles.includes('R_SUPER'));

interface Props {
  itemAlign?: NaiveUI.Align;
  loading?: boolean;
}

defineProps<Props>();

interface Emits {
  (e: 'refresh'): void;
  (e: 'clear'): void;
}

const emit = defineEmits<Emits>();

const columns = defineModel<NaiveUI.TableColumnCheck[]>('columns', {
  default: () => []
});

function refresh() {
  emit('refresh');
}

function handleClear() {
  emit('clear');
}
</script>

<template>
  <NSpace :align="itemAlign" wrap justify="end" class="lt-sm:w-200px">
    <slot name="prefix"></slot>
    <NButton v-if="isAdmin" size="small" type="error" @click="handleClear">
      <template #icon>
        <icon-ic-round-delete class="text-icon" />
      </template>
      {{ $t('common.delete') }}
    </NButton>
    <NButton size="small" @click="refresh">
      <template #icon>
        <icon-mdi-refresh class="text-icon" :class="{ 'animate-spin': loading }" />
      </template>
      {{ $t('common.refresh') }}
    </NButton>
    <TableColumnSetting v-model:columns="columns" />
    <slot name="suffix"></slot>
  </NSpace>
</template>

<style scoped></style>
