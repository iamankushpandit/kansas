<template>
  <section class="mt-4" v-if="selectedCounty" aria-labelledby="county-details-title">
    <v-card variant="outlined" role="region">
      <v-card-title class="text-h6">
        <h3 id="county-details-title">{{ selectedCounty.county }} County - Complete Details</h3>
      </v-card-title>
      <v-card-text>
        <CountyMetrics 
          :selected-county="selectedCounty"
          :county-terminated-analysis="countyTerminatedAnalysis"
        />
        
        <NetworkRecommendations 
          v-if="originalRecommendations.length > 0"
          :recommendations="originalRecommendations"
        />
        
        <SpecialtyDensityAnalysis 
          v-if="specialtyDensityRecommendations.length > 0"
          :recommendations="specialtyDensityRecommendations"
        />
      </v-card-text>
    </v-card>
  </section>
</template>

<script>
import CountyMetrics from './CountyMetrics.vue'
import NetworkRecommendations from './NetworkRecommendations.vue'
import SpecialtyDensityAnalysis from './SpecialtyDensityAnalysis.vue'

export default {
  name: 'CountyDetails',
  components: {
    CountyMetrics,
    NetworkRecommendations,
    SpecialtyDensityAnalysis
  },
  props: {
    selectedCounty: Object,
    countyTerminatedAnalysis: Object,
    originalRecommendations: Array,
    specialtyDensityRecommendations: Array
  }
}
</script>