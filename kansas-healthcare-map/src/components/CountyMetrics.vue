<template>
  <div>
    <v-row>
      <v-col cols="6">
        <div class="text-subtitle-2 text-primary">Provider Metrics</div>
        <div class="text-body-2">Active Providers: <strong>{{ selectedCounty.provider_count }}</strong></div>
        <div class="text-body-2">Density Classification: <strong class="text-capitalize">{{ selectedCounty.density }}</strong></div>
      </v-col>
      <v-col cols="6">
        <div class="text-subtitle-2 text-primary">Claims Metrics</div>
        <div class="text-body-2">Total Claims: <strong>{{ selectedCounty.claims_count?.toLocaleString() }}</strong></div>
        <div class="text-body-2">Avg Claim Amount: <strong>${{ selectedCounty.avg_claim_amount?.toFixed(2) }}</strong></div>
      </v-col>
    </v-row>
    <v-row class="mt-2" v-if="countyTerminatedAnalysis">
      <v-col cols="6">
        <div class="text-subtitle-2 text-primary">Network Termination</div>
        <div class="text-body-2">Termination Rate: <strong>{{ countyTerminatedAnalysis.percentage_terminated?.toFixed(1) }}%</strong></div>
        <div class="text-body-2">Providers Terminated: <strong>{{ countyTerminatedAnalysis.term_network_count }}</strong></div>
      </v-col>
      <v-col cols="6">
        <div class="text-subtitle-2 text-primary">Additional Metrics</div>
        <div class="text-body-2">Claims per Provider: <strong>{{ getClaimsPerProvider() }}</strong></div>
        <div class="text-body-2">Network Coverage Ratio: <strong>{{ getNetworkCoverageRatio() }}</strong></div>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: 'CountyMetrics',
  props: {
    selectedCounty: Object,
    countyTerminatedAnalysis: Object
  },
  methods: {
    getClaimsPerProvider() {
      if (!this.selectedCounty || this.selectedCounty.provider_count === 0) return 'N/A'
      return Math.round(this.selectedCounty.claims_count / this.selectedCounty.provider_count)
    },

    getNetworkCoverageRatio() {
      if (!this.selectedCounty || this.selectedCounty.claims_count === 0) return 'N/A'
      return Math.floor(this.selectedCounty.provider_count / this.selectedCounty.claims_count * 1000)
    }
  }
}
</script>