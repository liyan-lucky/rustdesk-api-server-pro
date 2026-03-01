<script setup lang="ts">
import { onMounted, onUnmounted, ref, watch } from 'vue';
import { $t } from '@/locales';
import { useEcharts } from '@/hooks/common/echarts';
import { fetchPieCharts } from '@/service/api/home';
import { useAppStore } from '@/store/modules/app';

defineOptions({
  name: 'PieChart'
});

const appStore = useAppStore();
const isNarrow = ref(false);
const chartData = ref<Api.Home.PieChart[]>([]);

function updateIsNarrow() {
  if (typeof window === 'undefined') return;
  isNarrow.value = window.innerWidth < 1400;
}

function ellipsisLabel(name: string) {
  if (!name) return '';
  return name.length > 30 ? `${name.slice(0, 30)}...` : name;
}

const { domRef, updateOptions } = useEcharts(() => ({
  title: {
    text: $t('page.home.operatingSystem'),
    left: 'center'
  },
  tooltip: {
    trigger: 'item'
  },
  legend: {
    type: 'scroll',
    bottom: '1%',
    left: 'center',
    right: 12,
    formatter: (name: string) => ellipsisLabel(name),
    itemStyle: {
      borderWidth: 0
    }
  },
  series: [
    {
      name: $t('page.home.operatingSystem'),
      type: 'pie',
      radius: ['45%', '75%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 10,
        borderColor: '#fff',
        borderWidth: 1
      },
      emphasis: {
        label: {
          show: true,
          fontSize: 12
        }
      },
      data: []
    }
  ]
}));

function renderChart(dataList: Api.Home.PieChart[] = []) {
  updateOptions(opt => {
    const data = (dataList || []).filter(item => Number(item?.value || 0) > 0);
    const hasData = data.length > 0;

    (opt as any).tooltip = hasData ? { trigger: 'item' } : { show: false };
    (opt as any).legend = {
      ...(opt as any).legend,
      show: hasData
    };
    (opt as any).graphic = hasData
      ? []
      : [
          {
            type: 'text',
            left: 'center',
            top: 'middle',
            style: {
              text: $t('common.noData'),
              fill: '#999',
              fontSize: 14,
              fontWeight: 500
            }
          }
        ];
    opt.series = [
      {
        name: $t('page.home.operatingSystem'),
        type: 'pie',
        radius: ['45%', '75%'],
        avoidLabelOverlap: false,
        silent: !hasData,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 1
        },
        emphasis: {
          label: {
            show: hasData,
            fontSize: 12
          }
        },
        label: {
          show: hasData && !isNarrow.value,
          formatter: ({ name }: { name: string }) => ellipsisLabel(name)
        },
        data: (hasData ? data : [{ name: $t('common.noData'), value: 1, itemStyle: { color: '#e5e7eb' } }]) as []
      }
    ] as any;
    return opt;
  });
}

async function fetchChartsData() {
  const pie = await fetchPieCharts();
  chartData.value = pie.data || [];
  renderChart(chartData.value);
}

async function init() {
  updateIsNarrow();
  await fetchChartsData();
}
// init
init();

onMounted(() => {
  updateIsNarrow();
  window.addEventListener('resize', updateIsNarrow);
});

onUnmounted(() => {
  window.removeEventListener('resize', updateIsNarrow);
});

watch(
  () => appStore.locale,
  () => {
    renderChart(chartData.value);
  }
);

watch(isNarrow, () => {
  renderChart(chartData.value);
});
</script>

<template>
  <NCard :bordered="false" class="card-wrapper">
    <div ref="domRef" class="h-360px overflow-hidden"></div>
  </NCard>
</template>

<style scoped></style>
