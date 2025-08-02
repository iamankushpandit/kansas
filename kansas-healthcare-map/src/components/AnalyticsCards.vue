<template>
  <section aria-labelledby="analytics-title">
    <h3 id="analytics-title" class="sr-only">Network Analytics</h3>
    
    <v-card variant="outlined" class="mb-4" role="region" aria-labelledby="statewide-overview-title">
      <v-card-text>
        <div class="d-flex align-center mb-2">
          <h4 id="statewide-overview-title" class="mr-2 text-subtitle-1">Statewide Provider Network Overview</h4>
          <v-tooltip location="top">
            <template v-slot:activator="{ props }">
              <v-btn 
                v-bind="props" 
                icon 
                size="x-small" 
                variant="text"
                aria-label="More information about statewide provider count"
              >
                <v-icon size="small" color="info">mdi-information</v-icon>
              </v-btn>
            </template>
            <span>Total number of healthcare providers currently active across all Kansas counties in our network database</span>
          </v-tooltip>
        </div>
        <div class="text-body-2 mb-1">Total Active Providers Across Kansas</div>
        <div 
          class="text-h5 text-primary" 
          :aria-label="`${activeProviderCount?.toLocaleString()} total active providers statewide`"
        >
          {{ activeProviderCount?.toLocaleString() }}
        </div>
      </v-card-text>
    </v-card>

    <v-card variant="outlined" class="mb-4" v-if="terminatedAnalysis" role="region" aria-labelledby="network-stability-title">
      <v-card-text>
        <div class="d-flex align-center mb-2">
          <h4 id="network-stability-title" class="mr-2 text-subtitle-1">{{ selectedNetwork }} Network Stability Analysis</h4>
          <v-tooltip location="top">
            <template v-slot:activator="{ props }">
              <v-btn 
                v-bind="props" 
                icon 
                size="x-small" 
                variant="text"
                aria-label="More information about network stability analysis"
              >
                <v-icon size="small" color="info">mdi-information</v-icon>
              </v-btn>
            </template>
            <span>Analysis of provider departures from the {{ selectedNetwork }} network over the past 2-5 years to assess network stability and retention patterns</span>
          </v-tooltip>
        </div>
        <dl class="text-body-2 mb-3">
          <div class="d-flex align-center mb-1">
            <dt class="mr-2">Providers Who Left Network (2-5 years ago):</dt>
            <v-tooltip location="top">
              <template v-slot:activator="{ props }">
                <v-btn 
                  v-bind="props" 
                  icon 
                  size="x-small" 
                  variant="text"
                  aria-label="More information about terminated providers"
                >
                  <v-icon size="x-small" color="grey">mdi-help-circle</v-icon>
                </v-btn>
              </template>
              <span>Number of providers who terminated their contract with {{ selectedNetwork }} between 2-5 years ago. This timeframe excludes recent departures (less than 2 years) and very old ones (more than 5 years)</span>
            </v-tooltip>
            <dd class="ml-2">
              <strong :aria-label="`${terminatedAnalysis.term_network_count} providers left the network`">
                {{ terminatedAnalysis.term_network_count }}
              </strong>
            </dd>
          </div>
          <div class="d-flex align-center mb-1">
            <dt class="mr-2">Associated Service Locations:</dt>
            <v-tooltip location="top">
              <template v-slot:activator="{ props }">
                <v-btn 
                  v-bind="props" 
                  icon 
                  size="x-small" 
                  variant="text"
                  aria-label="More information about service locations"
                >
                  <v-icon size="x-small" color="grey">mdi-help-circle</v-icon>
                </v-btn>
              </template>
              <span>Number of physical practice locations that were closed or terminated when providers left the network. Each provider may have multiple service locations</span>
            </v-tooltip>
            <dd class="ml-2">
              <strong :aria-label="`${terminatedAnalysis.service_location_count} service locations affected`">
                {{ terminatedAnalysis.service_location_count }}
              </strong>
            </dd>
          </div>
          <div class="d-flex align-center">
            <dt class="mr-2">Historical Termination Rate:</dt>
            <v-tooltip location="top">
              <template v-slot:activator="{ props }">
                <v-btn 
                  v-bind="props" 
                  icon 
                  size="x-small" 
                  variant="text"
                  aria-label="More information about termination rate"
                >
                  <v-icon size="x-small" color="grey">mdi-help-circle</v-icon>
                </v-btn>
              </template>
              <span>Percentage of total active providers who left the {{ selectedNetwork }} network in the 2-5 year timeframe. Higher rates may indicate network retention challenges</span>
            </v-tooltip>
            <dd class="ml-2">
              <strong 
                class="text-warning" 
                :aria-label="`${terminatedAnalysis.percentage_terminated.toFixed(1)} percent termination rate`"
              >
                {{ terminatedAnalysis.percentage_terminated.toFixed(1) }}%
              </strong>
            </dd>
          </div>
        </dl>
      </v-card-text>
    </v-card>
  </section>
</template>

<script>
export default {
  name: 'AnalyticsCards',
  props: {
    activeProviderCount: Number,
    terminatedAnalysis: Object,
    selectedNetwork: String
  }
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

.text-primary {
  color: #1565c0 !important;
}

.text-warning {
  color: #ef6c00 !important;
}
</style>