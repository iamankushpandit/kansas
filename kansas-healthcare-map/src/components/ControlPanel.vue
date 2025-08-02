<template>
  <aside id="map-controls" role="complementary" aria-labelledby="controls-title">
    <v-card>
      <v-card-title>
        <h2 id="controls-title" class="text-h6">Map Controls</h2>
      </v-card-title>
      <v-card-text>
        <FilterControls 
          :selected-specialty="selectedSpecialty"
          :selected-network="selectedNetwork"
          :selected-metric="selectedMetric"
          :specialties="specialties"
          :networks="networks"
          :metrics="metrics"
          :loading="loading"
          @update-specialty="$emit('update-specialty', $event)"
          @update-network="$emit('update-network', $event)"
          @update-metric="$emit('update-metric', $event)"
        />

        <v-divider class="my-4" role="separator" aria-label="Section divider"></v-divider>

        <AnalyticsCards 
          :active-provider-count="activeProviderCount"
          :terminated-analysis="terminatedAnalysis"
          :selected-network="selectedNetwork"
        />

        <MissingDataAlert 
          v-if="missingCounties.length > 0"
          :missing-counties="missingCounties"
        />
      </v-card-text>
    </v-card>
  </aside>
</template>

<script>
import FilterControls from './FilterControls.vue'
import AnalyticsCards from './AnalyticsCards.vue'
import MissingDataAlert from './MissingDataAlert.vue'

export default {
  name: 'ControlPanel',
  components: {
    FilterControls,
    AnalyticsCards,
    MissingDataAlert
  },
  props: {
    selectedSpecialty: String,
    selectedNetwork: String,
    selectedMetric: String,
    specialties: Array,
    networks: Array,
    metrics: Array,
    loading: Boolean,
    activeProviderCount: Number,
    terminatedAnalysis: Object,
    missingCounties: Array
  },
  emits: ['update-specialty', 'update-network', 'update-metric']
}
</script>