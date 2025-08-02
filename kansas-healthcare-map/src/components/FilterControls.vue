<template>
  <div role="group" aria-labelledby="filter-controls-title">
    <h3 id="filter-controls-title" class="sr-only">Data Filters</h3>
    
    <v-select
      :model-value="selectedSpecialty"
      :items="specialties"
      label="Provider Specialty"
      variant="outlined"
      @update:model-value="$emit('update-specialty', $event)"
      aria-describedby="specialty-help"
      :disabled="loading"
    ></v-select>
    <div id="specialty-help" class="sr-only">Filter providers by medical specialty type</div>

    <v-select
      :model-value="selectedNetwork"
      :items="networks"
      label="Network Type"
      variant="outlined"
      @update:model-value="$emit('update-network', $event)"
      aria-describedby="network-help"
      :disabled="loading"
    ></v-select>
    <div id="network-help" class="sr-only">Filter by insurance network type</div>

    <v-select
      :model-value="selectedMetric"
      :items="metrics"
      label="Display Metric"
      variant="outlined"
      @update:model-value="$emit('update-metric', $event)"
      aria-describedby="metric-help"
      :disabled="loading"
    ></v-select>
    <div id="metric-help" class="sr-only">Choose which data metric to visualize on the map</div>
  </div>
</template>

<script>
export default {
  name: 'FilterControls',
  props: {
    selectedSpecialty: String,
    selectedNetwork: String,
    selectedMetric: String,
    specialties: Array,
    networks: Array,
    metrics: Array,
    loading: Boolean
  },
  emits: ['update-specialty', 'update-network', 'update-metric']
}
</script>

<style scoped>
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border: 0;
}

.v-select:focus-within {
  outline: 3px solid #1976d2;
  outline-offset: 2px;
}
</style>