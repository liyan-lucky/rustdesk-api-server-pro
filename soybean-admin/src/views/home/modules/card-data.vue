<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { createReusableTemplate } from '@vueuse/core';
import { useRouterPush } from '@/hooks/common/router';
import { $t } from '@/locales';
import { fetchStat } from '@/service/api/home';
import { fetchAuditLogList } from '@/service/api/audit';
import { fetchDevicesList } from '@/service/api/devices';
import { fetchUserList } from '@/service/api/user_management';

defineOptions({
  name: 'CardData'
});

interface CardData {
  key: 'userCount' | 'deviceCount' | 'onlineCount' | 'visitCount';
  title: string;
  value: number;
  unit: string;
  color: {
    start: string;
    end: string;
  };
  icon: string;
}

interface DetailItem {
  title: string;
  desc: string;
  meta?: string;
}

interface GradientBgProps {
  gradientColor: string;
}

const stat = ref<Api.Home.Stat>({
  userCount: 0,
  deviceCount: 0,
  onlineCount: 0,
  visitsCount: 0
});

const cardData = computed<CardData[]>(() => [
  {
    key: 'userCount',
    title: $t('page.home.userCount'),
    value: stat.value.userCount,
    unit: '',
    color: { start: '#ec4786', end: '#b955a4' },
    icon: 'gravity-ui:person'
  },
  {
    key: 'deviceCount',
    title: $t('page.home.deviceCount'),
    value: stat.value.deviceCount,
    unit: '',
    color: { start: '#865ec0', end: '#5144b4' },
    icon: 'fluent:desktop-32-regular'
  },
  {
    key: 'onlineCount',
    title: $t('page.home.onlineCount'),
    value: stat.value.onlineCount,
    unit: '',
    color: { start: '#56cdf3', end: '#719de3' },
    icon: 'fluent:desktop-checkmark-20-regular'
  },
  {
    key: 'visitCount',
    title: $t('page.home.visitsCount'),
    value: stat.value.visitsCount,
    unit: '',
    color: { start: '#fcbc25', end: '#f68057' },
    icon: 'ant-design:bar-chart-outlined'
  }
]);

const [DefineGradientBg, GradientBg] = createReusableTemplate<GradientBgProps>();
const { routerPushByKey } = useRouterPush();

const drawerVisible = ref(false);
const activeCardKey = ref<CardData['key'] | null>(null);
const detailLoading = ref(false);
const detailError = ref('');
const detailItems = ref<DetailItem[]>([]);

const activeCard = computed(() => {
  if (!activeCardKey.value) return null;
  return cardData.value.find(item => item.key === activeCardKey.value) || null;
});

const viewDetailHint = computed(() => $t('page.home.cardDetail.viewHint'));

const activeCardDesc = computed(() => {
  const key = activeCard.value?.key;
  if (!key) return '';
  const i18nMap: Record<CardData['key'], App.I18n.I18nKey> = {
    userCount: 'page.home.cardDetail.desc.userCount',
    deviceCount: 'page.home.cardDetail.desc.deviceCount',
    onlineCount: 'page.home.cardDetail.desc.onlineCount',
    visitCount: 'page.home.cardDetail.desc.visitCount'
  };
  return $t(i18nMap[key]);
});

const activeCardListTitle = computed(() => {
  const key = activeCard.value?.key;
  if (!key) return '';
  if (key === 'userCount') return $t('page.home.cardDetail.recentUsers');
  if (key === 'visitCount') return $t('page.home.cardDetail.recentVisits');
  return $t('page.home.cardDetail.recentDevices');
});

const drawerTitle = computed(() => activeCard.value?.title || '');

function getGradientColor(color: CardData['color']) {
  return `linear-gradient(to bottom right, ${color.start}, ${color.end})`;
}

function getRouteByCardKey(key: CardData['key']) {
  switch (key) {
    case 'userCount':
      return 'user_list';
    case 'deviceCount':
    case 'onlineCount':
      return 'devices';
    case 'visitCount':
      return 'audit_baselogs';
    default:
      return null;
  }
}

function handleCardClick(item: CardData) {
  activeCardKey.value = item.key;
  drawerVisible.value = true;
  void loadCardDetails(item.key);
}

function goToDetailPage() {
  if (!activeCard.value) return;
  const routeKey = getRouteByCardKey(activeCard.value.key);
  if (!routeKey) return;
  routerPushByKey(routeKey);
  drawerVisible.value = false;
}

async function loadCardDetails(key: CardData['key']) {
  detailLoading.value = true;
  detailError.value = '';
  detailItems.value = [];

  try {
    if (key === 'userCount') {
      const { data, error } = await fetchUserList({ current: 1, size: 6 });
      if (error || !data) throw new Error('user list fetch failed');
      detailItems.value = (data.records || []).map(item => ({
        title: item.username || '-',
        desc: item.email || item.name || '-',
        meta: item.created_at || ''
      }));
    } else if (key === 'visitCount') {
      const { data, error } = await fetchAuditLogList({ current: 1, size: 6 });
      if (error || !data) throw new Error('audit list fetch failed');
      detailItems.value = (data.records || []).map(item => ({
        title: item.username || '-',
        desc: `${item.rustdesk_id || '-'} / ${item.ip || '-'}`,
        meta: item.created_at || ''
      }));
    } else {
      const { data, error } = await fetchDevicesList({ current: 1, size: 6 });
      if (error || !data) throw new Error('device list fetch failed');
      detailItems.value = (data.records || []).map(item => ({
        title: item.hostname || '-',
        desc: `${item.rustdesk_id || '-'} / ${item.username || '-'}`,
        meta: item.created_at || ''
      }));
    }

    if (detailItems.value.length === 0) {
      detailError.value = $t('common.noData');
    }
  } catch {
    detailError.value = $t('common.error');
  } finally {
    detailLoading.value = false;
  }
}

onMounted(async () => {
  const s = (await fetchStat()).data;
  if (s !== null) {
    stat.value = s;
  }
});
</script>

<template>
  <NCard :bordered="false" size="small" class="card-wrapper">
    <DefineGradientBg v-slot="{ $slots, gradientColor }">
      <div class="stat-card rd-8px px-16px pb-10px pt-8px text-white" :style="{ backgroundImage: gradientColor }">
        <component :is="$slots.default" />
      </div>
    </DefineGradientBg>

    <NGrid cols="2 m:2 l:4" responsive="screen" :x-gap="16" :y-gap="16" class="stat-grid">
      <NGi v-for="item in cardData" :key="item.key">
        <GradientBg
          :gradient-color="getGradientColor(item.color)"
          class="flex-1"
          role="button"
          tabindex="0"
          @click="handleCardClick(item)"
          @keydown.enter.prevent="handleCardClick(item)"
          @keydown.space.prevent="handleCardClick(item)"
        >
          <h3 class="stat-title">{{ item.title }}</h3>
          <p class="stat-hint">{{ viewDetailHint }}</p>
          <div class="flex justify-between pt-12px">
            <SvgIcon :icon="item.icon" class="stat-icon" />
            <CountTo :prefix="item.unit" :start-value="1" :end-value="item.value" class="stat-value text-white dark:text-dark" />
          </div>
        </GradientBg>
      </NGi>
    </NGrid>

    <NDrawer v-model:show="drawerVisible" :width="420" placement="right" :trap-focus="false">
      <NDrawerContent v-if="activeCard" :title="drawerTitle" closable>
        <div class="drawer-body">
          <div class="drawer-value">
            <span class="drawer-value-num">{{ activeCard.value }}</span>
            <span v-if="activeCard.unit" class="drawer-value-unit">{{ activeCard.unit }}</span>
          </div>
          <p v-if="activeCardDesc" class="drawer-desc">{{ activeCardDesc }}</p>
          <div class="drawer-list-wrap">
            <div class="drawer-list-title">{{ activeCardListTitle }}</div>
            <NSpin :show="detailLoading">
              <NList v-if="!detailError && detailItems.length > 0" bordered class="drawer-list">
                <NListItem v-for="(it, idx) in detailItems" :key="`${idx}-${it.title}`">
                  <NThing :title="it.title" :description="it.desc">
                    <template #header-extra>
                      <span class="drawer-item-meta">{{ it.meta }}</span>
                    </template>
                  </NThing>
                </NListItem>
              </NList>
              <div v-else class="drawer-empty">{{ detailError }}</div>
            </NSpin>
          </div>
          <div class="drawer-actions">
            <NButton type="primary" @click="goToDetailPage">{{ $t('common.look') }}</NButton>
            <NButton @click="drawerVisible = false">{{ $t('common.close') }}</NButton>
          </div>
        </div>
      </NDrawerContent>
    </NDrawer>
  </NCard>
</template>

<style scoped>
.stat-card {
  min-height: 94px;
  cursor: pointer;
  transition:
    transform 0.2s ease,
    box-shadow 0.2s ease,
    filter 0.2s ease;
}

.stat-card:hover {
  filter: brightness(1.05);
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgb(0 0 0 / 18%);
}

.stat-card:focus-visible {
  outline: 2px solid rgb(255 255 255 / 85%);
  outline-offset: 2px;
}

.stat-title {
  margin: 0;
  font-size: clamp(14px, 1.25vw, 18px);
  line-height: 1.35;
  font-weight: 600;
}

.stat-hint {
  margin: 4px 0 0;
  font-size: 12px;
  line-height: 1.2;
  opacity: 0.88;
}

.stat-icon {
  font-size: clamp(22px, 2vw, 32px);
}

.stat-value {
  font-size: clamp(28px, 2.2vw, 34px);
  line-height: 1;
  font-weight: 700;
}

.drawer-body {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.drawer-value {
  display: flex;
  align-items: baseline;
  gap: 8px;
}

.drawer-value-num {
  font-size: 36px;
  font-weight: 700;
  line-height: 1;
}

.drawer-value-unit {
  color: var(--n-text-color-3);
}

.drawer-desc {
  margin: 0;
  line-height: 1.6;
  color: var(--n-text-color-2);
}

.drawer-list-wrap {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.drawer-list-title {
  font-size: 14px;
  font-weight: 600;
}

.drawer-list {
  max-height: 260px;
  overflow: auto;
}

.drawer-item-meta {
  color: var(--n-text-color-3);
  font-size: 12px;
}

.drawer-empty {
  min-height: 80px;
  color: var(--n-text-color-3);
  display: flex;
  align-items: center;
  justify-content: center;
}

.drawer-actions {
  display: flex;
  gap: 10px;
}

@media (width <= 768px) {
  .stat-grid {
    --n-x-gap: 12px !important;
    --n-y-gap: 12px !important;
  }

  .stat-card {
    min-height: 84px;
    padding-right: 12px;
    padding-bottom: 8px;
    padding-left: 12px;
  }

  .drawer-actions {
    flex-wrap: wrap;
  }
}
</style>
