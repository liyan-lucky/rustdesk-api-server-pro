<script setup lang="ts">
import { $t } from '@/locales';
import { useEcharts } from '@/hooks/common/echarts';
import { fetchPieCharts } from '@/service/api/home';

defineOptions({
  name: 'PieChart'
});

const { domRef, updateOptions } = useEcharts(() => ({
  title: {
    text: $t('page.home.operatingSystem'),
    left: 'center'
  },
  tooltip: {
    trigger: 'item'
  },
  legend: {
    bottom: '1%',
    left: 'center',
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
          fontSize: '12'
        }
      },
      data: []
    }
  ]
}));
async function fetchChartsData() {
  const pie = await fetchPieCharts();
  updateOptions(opt => {
    const data = (pie.data || []).filter(item => Number(item?.value || 0) > 0);
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
            fontSize: '12'
          }
        },
        label: {
          show: hasData
        },
        data: (hasData ? data : [{ name: $t('common.noData'), value: 1, itemStyle: { color: '#e5e7eb' } }]) as []
      }
    ] as any;
    return opt;
  });
}
async function init() {
  fetchChartsData();
}
// init
init();
</script>

<template>
  <NCard :bordered="false" class="card-wrapper">
    <div ref="domRef" class="h-360px overflow-hidden"></div>
  </NCard>
</template>

<style scoped></style>
