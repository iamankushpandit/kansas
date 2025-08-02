<template>
  <div class="mt-3" role="region" aria-labelledby="legend-title">
    <h3 id="legend-title" class="text-h6">{{ getMetricLegendTitle() }}</h3>
    <div class="legend" role="list" aria-label="Map color legend">
      <div 
        class="legend-item" 
        v-for="item in getMetricLegendItems()" 
        :key="item.label"
        role="listitem"
      >
        <div 
          class="legend-color" 
          :style="{ backgroundColor: item.color }"
          :aria-label="`Color indicator for ${item.label}`"
        ></div>
        <span>{{ item.label }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MapLegend',
  props: {
    selectedMetric: String
  },
  methods: {
    getMetricLegendTitle() {
      return `${this.selectedMetric} Legend`
    },

    getMetricLegendItems() {
      if (this.selectedMetric === 'Provider Density') {
        return [
          { color: '#2E8B57', label: 'High (400+ providers)' },
          { color: '#90EE90', label: 'Medium (100-399 providers)' },
          { color: '#FFD700', label: 'Low (50-99 providers)' },
          { color: '#FF6347', label: 'Critical (<50 providers)' }
        ]
      } else if (this.selectedMetric === 'Claims Volume') {
        return [
          { color: '#2E8B57', label: 'High (7500+ claims)' },
          { color: '#90EE90', label: 'Medium-High (5000-7499 claims)' },
          { color: '#FFD700', label: 'Medium-Low (2500-4999 claims)' },
          { color: '#FF6347', label: 'Low (<2500 claims)' }
        ]
      } else {
        return [
          { color: '#2E8B57', label: 'Excellent (38+ ratio)' },
          { color: '#90EE90', label: 'Good (25-37 ratio)' },
          { color: '#FFD700', label: 'Fair (13-24 ratio)' },
          { color: '#FF6347', label: 'Poor (<13 ratio)' }
        ]
      }
    }
  }
}
</script>

<style scoped>
.legend {
  margin-top: 10px;
}

.legend-item {
  display: inline-flex;
  align-items: center;
  margin-right: 15px;
  margin-bottom: 8px;
  min-height: 44px;
}

.legend-color {
  width: 20px;
  height: 20px;
  margin-right: 10px;
  border: 1px solid #ccc;
  flex-shrink: 0;
}
</style>