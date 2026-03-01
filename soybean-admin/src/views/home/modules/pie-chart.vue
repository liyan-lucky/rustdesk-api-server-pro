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
const isUltraNarrow = ref(false);
const chartData = ref<Api.Home.PieChart[]>([]);

function updateIsNarrow() {
  if (typeof window === 'undefined') return;
  isNarrow.value = window.innerWidth < 1400;
  isUltraNarrow.value = window.innerWidth < 420;
}

function ellipsisLabel(name: string) {
  if (!name) return '';
  return name.length > 30 ? `${name.slice(0, 30)}...` : name;
}

function compactSideLabel(name: string) {
  if (!name) return '';
  if (!isNarrow.value) return name;
  return name.length > 10 ? `${name.slice(0, 10)}...` : name;
}

const { domRef, updateOptions } = useEcharts(() => ({
  title: {
    text: $t('page.home.operatingSystem'),
    left: 'center',
    top: 14
  },
  tooltip: {
    trigger: 'item'
  },
  legend: {
    type: 'scroll',
    bottom: 8,
    left: 'center',
    right: 10,
    formatter: (name: string) => ellipsisLabel(name),
    itemStyle: {
      borderWidth: 0
    }
  },
  series: [
    {
      name: $t('page.home.operatingSystem'),
      type: 'pie',
      center: ['50%', '44%'],
      radius: ['42%', '68%'],
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
    const titleText = $t('page.home.operatingSystem');
    const centerX = isUltraNarrow.value ? '50%' : isNarrow.value ? '54%' : '50%';
    const centerY = isNarrow.value ? '40%' : '44%';
    const radius = isUltraNarrow.value
      ? (['32%', '54%'] as [string, string])
      : isNarrow.value
        ? (['30%', '52%'] as [string, string])
        : (['42%', '68%'] as [string, string]);
    const legendBottom = isNarrow.value ? 2 : 8;

    (opt as any).title = {
      ...(opt as any).title,
      text: titleText
    };

    (opt as any).tooltip = hasData ? { trigger: 'item' } : { show: false };
    (opt as any).legend = {
      ...(opt as any).legend,
      show: hasData,
      bottom: legendBottom,
      formatter: (name: string) => ellipsisLabel(name)
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
        name: titleText,
        type: 'pie',
        center: [centerX, centerY],
        radius,
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
          show: hasData && !isUltraNarrow.value,
          formatter: ({ name }: { name: string }) => compactSideLabel(name),
          width: isNarrow.value ? 72 : 120,
          overflow: 'truncate',
          fontSize: isNarrow.value ? 11 : 12,
          alignTo: isNarrow.value ? 'edge' : undefined,
          edgeDistance: isNarrow.value ? 6 : undefined,
          bleedMargin: isNarrow.value ? 6 : undefined
        },
        labelLine: {
          show: hasData && !isUltraNarrow.value,
          length: isNarrow.value ? 6 : 14,
          length2: isNarrow.value ? 6 : 12
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

watch(isUltraNarrow, () => {
  renderChart(chartData.value);
});
</script>

<template>
  <NCard :bordered="false" class="card-wrapper">
    <div ref="domRef" class="h-360px overflow-hidden"></div>
  </NCard>
</template>

<style scoped></style>
