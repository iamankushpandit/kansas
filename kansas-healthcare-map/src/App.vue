<template>
  <v-app>
    <!-- Skip Links -->
    <div class="skip-links">
      <a href="#main-content" class="skip-link">Skip to main content</a>
      <a href="#map-controls" class="skip-link">Skip to map controls</a>
    </div>

    <AppHeader 
      :selected-county="selectedCounty?.county"
      :loading="loading"
      :is-dark="$vuetify.theme.global.name === 'dark'"
      @export-pdf="exportToPDF"
      @toggle-theme="toggleTheme"
    />

    <v-main id="main-content" role="main">
      <!-- Loading State -->
      <v-overlay :model-value="loading" class="align-center justify-center">
        <v-progress-circular 
          indeterminate 
          size="64" 
          color="primary"
          aria-label="Loading healthcare data"
        ></v-progress-circular>
        <div class="mt-4 text-center">
          <p class="text-h6">Loading Healthcare Data...</p>
          <p class="text-body-2">Please wait while we fetch the latest provider information</p>
        </div>
      </v-overlay>

      <v-container fluid>
        <v-row>
          <!-- Map Container -->
          <v-col cols="12" md="8">
            <v-card>
              <v-card-text>
                <KansasMap 
                  :selected-metric="selectedMetric"
                  :selected-county="selectedCounty"
                  @county-click="onCountyClick"
                />
                
                <CountyDetails 
                  :selected-county="selectedCounty"
                  :county-terminated-analysis="countyTerminatedAnalysis"
                  :original-recommendations="originalRecommendations"
                  :specialty-density-recommendations="specialtyDensityRecommendations"
                />
              </v-card-text>
            </v-card>
          </v-col>

          <!-- Controls Panel -->
          <v-col cols="12" md="4">
            <ControlPanel 
              :selected-specialty="selectedSpecialty"
              :selected-network="selectedNetwork"
              :selected-metric="selectedMetric"
              :specialties="specialties"
              :networks="networks"
              :metrics="metrics"
              :loading="loading"
              :active-provider-count="activeProviderCount"
              :terminated-analysis="terminatedAnalysis"
              :missing-counties="missingCounties"
              @update-specialty="updateSpecialty"
              @update-network="updateNetwork"
              @update-metric="updateMetric"
            />
          </v-col>
        </v-row>
        
        <!-- Error State -->
        <v-row v-if="error && !loading">
          <v-col cols="12">
            <v-alert 
              type="error" 
              variant="tonal" 
              closable
              @click:close="error = null"
              role="alert"
              aria-live="assertive"
            >
              <v-alert-title>Error Loading Data</v-alert-title>
              {{ error }}
            </v-alert>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { healthcareApi } from './services/api.js'
import configData from './data/config.json'
import countyMapping from './data/countyMapping.json'
import { useTheme } from 'vuetify'

// Components
import AppHeader from './components/AppHeader.vue'
import KansasMap from './components/KansasMap.vue'
import CountyDetails from './components/CountyDetails.vue'
import ControlPanel from './components/ControlPanel.vue'

export default {
  name: 'KansasHealthcareMap',
  components: {
    AppHeader,
    KansasMap,
    CountyDetails,
    ControlPanel
  },
  setup() {
    const theme = useTheme()
    return { theme }
  },
  data() {
    return {
      chart: null,
      selectedSpecialty: 'All',
      selectedNetwork: 'Commercial',
      selectedMetric: 'Provider Density',
      selectedCounty: null,
      loading: false,
      kansasCountyData: [],
      activeProviderCount: 0,
      terminatedAnalysis: null,
      countyTerminatedAnalysis: null,
      originalRecommendations: [],
      specialtyDensityRecommendations: [],
      specialtyDensityAnalysis: null,
      missingCounties: [],
      error: null,
      specialties: configData.specialties,
      networks: configData.networks,
      metrics: configData.metrics
    }
  },
  async mounted() {
    await this.loadData()
    this.initializeMap()
  },
  methods: {
    toggleTheme() {
      this.theme.global.name.value = this.theme.global.name.value === 'light' ? 'dark' : 'light'
      this.updateChartTheme()
    },

    updateSpecialty(value) {
      this.selectedSpecialty = value
      this.updateMap()
    },

    updateNetwork(value) {
      this.selectedNetwork = value
      this.updateMap()
    },

    updateMetric(value) {
      this.selectedMetric = value
      this.updateMap()
    },

    async loadData() {
      try {
        console.log('[INFO] Starting data load process')
        this.loading = true
        
        const filter = {
          specialty: this.selectedSpecialty,
          metric: this.selectedMetric,
          network: this.selectedNetwork
        }
        
        try {
          console.log('[INFO] Loading filtered providers with filter:', filter)
          const filteredProviders = await healthcareApi.getFilteredProviders(filter)
          const baseCountyData = await healthcareApi.getAllCountyData()
          this.kansasCountyData = this.recalculateCountyStatsWithBase(filteredProviders, baseCountyData)
          console.log('[INFO] Successfully loaded filtered data for', this.kansasCountyData.length, 'counties')
        } catch (error) {
          console.error('[ERROR] Error loading filtered data:', error)
          console.log('[WARN] Falling back to base county data')
          this.kansasCountyData = await healthcareApi.getAllCountyData()
        }
        
        console.log('[INFO] Loading analytics data')
        this.activeProviderCount = (await healthcareApi.getActiveProviderCount()).total_active_providers
        this.terminatedAnalysis = await healthcareApi.getTerminatedNetworkAnalysis(this.selectedNetwork)
        console.log('[INFO] Data load completed successfully')
        this.error = null
        
      } catch (error) {
        console.error('[ERROR] Critical error loading data:', error)
        this.error = 'Failed to load healthcare data. Please refresh the page and try again.'
      } finally {
        this.loading = false
      }
    },

    async initializeMap() {
      try {
        await this.loadScript('https://code.highcharts.com/highcharts.js')
        await this.loadScript('https://code.highcharts.com/maps/modules/map.js')
        await this.loadScript('https://code.highcharts.com/mapdata/countries/us/us-ks-all.js')

        window.Highcharts.setOptions({
          lang: {
            thousandsSep: ',',
            decimalPoint: '.',
            locale: 'en-US'
          }
        })

        this.setHighchartsTheme()
        this.createMap()
      } catch (error) {
        console.error('Failed to load map dependencies:', error)
      }
    },

    loadScript(src) {
      return new Promise((resolve, reject) => {
        if (document.querySelector(`script[src="${src}"]`)) {
          resolve()
          return
        }
        
        const script = document.createElement('script')
        script.src = src
        script.onload = resolve
        script.onerror = reject
        document.head.appendChild(script)
      })
    },

    createMap() {
      if (!this.kansasCountyData.length) return

      this.identifyMissingCounties()

      const mapData = this.kansasCountyData.map(county => {
        let value = county.provider_count
        
        if (this.selectedMetric === 'Claims Volume') {
          value = county.claims_count
        } else if (this.selectedMetric === 'Network Coverage') {
          value = Math.floor(county.provider_count / county.claims_count * 1000)
        }

        return {
          'hc-key': this.getCountyCode(county.county),
          value: value,
          name: county.county,
          fip: 'FIP: ' + this.getCountyCode(county.county).replace('us-ks-', ''),
          provider_count: county.provider_count,
          claims_count: county.claims_count,
          avg_claim_amount: county.avg_claim_amount,
          density: county.density
        }
      })

      this.chart = window.Highcharts.mapChart('kansas-map', {
        accessibility: { enabled: false },
        chart: { map: window.Highcharts.maps['countries/us/us-ks-all'] },
        title: { text: null },
        subtitle: { text: null },
        mapNavigation: {
          enabled: true,
          buttonOptions: { verticalAlign: 'bottom' }
        },
        colorAxis: this.getColorAxisConfig(mapData),
        legend: { enabled: false },
        series: [{
          data: mapData,
          name: 'Provider Count',
          states: { hover: { color: '#BADA55' } },
          dataLabels: {
            enabled: true,
            format: '{point.name}<br/>{point.fip}',
            style: {
              fontSize: '10px',
              fontWeight: 'bold',
              textOutline: '1px contrast'
            }
          },
          point: {
            events: {
              click: (event) => {
                this.onCountyClick(event.point)
              }
            }
          }
        }],
        tooltip: {
          headerFormat: '',
          pointFormat: '<b>{point.name} County</b><br/>' +
                      this.getTooltipFormat() +
                      'Click for details'
        }
      })
    },

    async onCountyClick(point) {
      try {
        console.log('[INFO] County selected:', point.name)
        const countyData = await healthcareApi.getCountyData(point.name)
        this.selectedCounty = countyData
        
        console.log('[INFO] Loading recommendations for county:', point.name)
        this.originalRecommendations = await healthcareApi.getRecommendations(point.name)
        this.specialtyDensityAnalysis = await healthcareApi.getSpecialtyDensityAnalysis(point.name)
        this.specialtyDensityRecommendations = this.generateSpecialtyDensityRecommendations()
        this.countyTerminatedAnalysis = await healthcareApi.getCountyTerminatedNetworkAnalysis(point.name, this.selectedNetwork)
        console.log('[INFO] Successfully loaded county data for:', point.name)
      } catch (error) {
        console.error('[ERROR] Error loading county data for', point.name, ':', error)
        console.log('[WARN] Using fallback county data')
        this.selectedCounty = {
          county: point.name,
          provider_count: point.provider_count,
          claims_count: point.claims_count,
          avg_claim_amount: point.avg_claim_amount
        }
        this.originalRecommendations = []
        this.specialtyDensityRecommendations = []
        this.countyTerminatedAnalysis = null
        this.specialtyDensityAnalysis = null
      }
    },

    async updateMap() {
      if (!this.chart) return

      try {
        if (this.selectedNetwork) {
          this.terminatedAnalysis = await healthcareApi.getTerminatedNetworkAnalysis(this.selectedNetwork)
        }

        let countyData = [...this.kansasCountyData]
        
        const filter = {
          specialty: this.selectedSpecialty,
          metric: this.selectedMetric,
          network: this.selectedNetwork
        }
        
        try {
          const filteredProviders = await healthcareApi.getFilteredProviders(filter)
          console.log(`Filtered providers for ${this.selectedSpecialty} + ${this.selectedNetwork}:`, filteredProviders.length)
          countyData = this.recalculateCountyStats(filteredProviders)
        } catch (error) {
          console.error('Error getting filtered providers:', error)
        }

        const mapData = countyData.map(county => {
          let value = county.provider_count
          
          if (this.selectedMetric === 'Claims Volume') {
            value = county.claims_count
          } else if (this.selectedMetric === 'Network Coverage') {
            value = county.provider_count > 0 ? Math.floor(county.provider_count / county.claims_count * 1000) : 0
          }

          return {
            'hc-key': this.getCountyCode(county.county),
            value: value,
            name: county.county,
            fip: 'FIP: ' + this.getCountyCode(county.county).replace('us-ks-', ''),
            provider_count: county.provider_count,
            claims_count: county.claims_count,
            avg_claim_amount: county.avg_claim_amount,
            density: county.density
          }
        })

        this.chart.update({ colorAxis: this.getColorAxisConfig(mapData) })
        this.chart.series[0].setData(mapData)

        if (this.selectedCounty) {
          this.specialtyDensityAnalysis = await healthcareApi.getSpecialtyDensityAnalysis(this.selectedCounty.county)
          this.specialtyDensityRecommendations = this.generateSpecialtyDensityRecommendations()
          this.countyTerminatedAnalysis = await healthcareApi.getCountyTerminatedNetworkAnalysis(this.selectedCounty.county, this.selectedNetwork)
        }

      } catch (error) {
        console.error('Error updating map:', error)
      }
    },

    generateSpecialtyDensityRecommendations() {
      if (!this.specialtyDensityAnalysis || !this.specialtyDensityAnalysis.specialty_densities) {
        return []
      }

      const densities = this.specialtyDensityAnalysis.specialty_densities
      const recommendations = []
      let idCounter = 1

      densities.forEach((specialty, index) => {
        const densityValue = this.calculateSpecialtyDensityMiles(specialty.count)
        let priority = 'Low'
        
        if (index < 3) {
          priority = 'High'
        } else if (index < 6) {
          priority = 'Medium'
        }

        recommendations.push({
          id: idCounter++,
          type: 'SPECIALTY_DENSITY',
          title: specialty.name,
          description: `${specialty.count} (${densityValue})`,
          priority: priority,
          icon: this.getSpecialtyIcon(specialty.name)
        })
      })

      return recommendations
    },

    getSpecialtyIcon(specialty) {
      const iconMap = {
        'Primary Care': 'mdi-stethoscope',
        'Cardiology': 'mdi-heart',
        'Orthopedics': 'mdi-bone',
        'Pediatrics': 'mdi-baby-face',
        'Psychiatry': 'mdi-brain',
        'Emergency Medicine': 'mdi-ambulance',
        'Neurology': 'mdi-head-outline',
        'Dermatology': 'mdi-hand-back-left',
        'Gastroenterology': 'mdi-stomach',
        'Endocrinology': 'mdi-needle',
        'Oncology': 'mdi-ribbon',
        'Radiology': 'mdi-radioactive',
        'Anesthesiology': 'mdi-sleep',
        'Pathology': 'mdi-microscope',
        'Ophthalmology': 'mdi-eye',
        'Urology': 'mdi-human-male',
        'Nephrology': 'mdi-kidney',
        'Pulmonology': 'mdi-lungs',
        'Rheumatology': 'mdi-human-handsup',
        'Infectious Disease': 'mdi-virus',
        'Hematology': 'mdi-water',
        'General Surgery': 'mdi-medical-bag',
        'Plastic Surgery': 'mdi-face-recognition',
        'Otolaryngology': 'mdi-ear-hearing',
        'Gynecology': 'mdi-gender-female',
        'Obstetrics': 'mdi-baby-carriage',
        'Family Medicine': 'mdi-home-heart',
        'Internal Medicine': 'mdi-medical-bag',
        'Sports Medicine': 'mdi-run',
        'Pain Management': 'mdi-lightning-bolt',
        'Physical Medicine': 'mdi-dumbbell',
        'Rehabilitation': 'mdi-wheelchair-accessibility',
        'Geriatrics': 'mdi-account-supervisor',
        'Allergy/Immunology': 'mdi-flower-pollen',
        'Critical Care': 'mdi-heart-pulse',
        'Hospitalist': 'mdi-hospital-building'
      }
      return iconMap[specialty] || 'mdi-medical-bag'
    },

    calculateSpecialtyDensityMiles(providerCount) {
      if (providerCount === 0) {
        return 'No providers'
      }
      
      const avgCountyAreaSqMiles = 700
      const providersPerSqMile = providerCount / avgCountyAreaSqMiles
      
      if (providersPerSqMile >= 1) {
        return `${providersPerSqMile.toFixed(1)}/sq mi`
      } else {
        const avgDistanceMiles = Math.sqrt(avgCountyAreaSqMiles / providerCount)
        return `~${avgDistanceMiles.toFixed(1)} mi apart`
      }
    },

    getCountyCode(countyName) {
      return countyMapping[countyName] || `us-ks-${countyName.toLowerCase()}`
    },

    getColorAxisConfig(mapData = null) {
      let maxValue = 400
      if (mapData && mapData.length > 0) {
        const values = mapData.map(d => d.value).filter(v => v > 0)
        if (values.length > 0) {
          maxValue = Math.max(...values)
        }
      }

      if (this.selectedMetric === 'Provider Density') {
        const dynamicMax = Math.max(maxValue, 10)
        return {
          min: 0,
          max: dynamicMax,
          stops: [
            [0, '#FF6347'],
            [0.25, '#FFD700'],
            [0.5, '#90EE90'],
            [1, '#2E8B57']
          ],
          labels: { format: '{value} providers' }
        }
      } else if (this.selectedMetric === 'Claims Volume') {
        const dynamicMax = Math.max(maxValue, 100)
        return {
          min: 0,
          max: dynamicMax,
          stops: [
            [0, '#FF6347'],
            [0.25, '#FFD700'],
            [0.5, '#90EE90'],
            [1, '#2E8B57']
          ],
          labels: { format: '{value} claims' }
        }
      } else {
        const dynamicMax = Math.max(maxValue, 10)
        return {
          min: 0,
          max: dynamicMax,
          stops: [
            [0, '#FF6347'],
            [0.25, '#FFD700'],
            [0.5, '#90EE90'],
            [1, '#2E8B57']
          ],
          labels: { format: '{value} ratio' }
        }
      }
    },

    recalculateCountyStats(filteredProviders) {
      const providersByCounty = {}
      filteredProviders.forEach(provider => {
        if (!providersByCounty[provider.county]) {
          providersByCounty[provider.county] = []
        }
        providersByCounty[provider.county].push(provider)
      })

      return this.kansasCountyData.map(county => {
        const countyProviders = providersByCounty[county.county] || []
        
        return {
          ...county,
          provider_count: countyProviders.length
        }
      })
    },

    recalculateCountyStatsWithBase(filteredProviders, baseCountyData) {
      const providersByCounty = {}
      filteredProviders.forEach(provider => {
        if (!providersByCounty[provider.county]) {
          providersByCounty[provider.county] = []
        }
        providersByCounty[provider.county].push(provider)
      })

      return baseCountyData.map(county => {
        const countyProviders = providersByCounty[county.county] || []
        
        return {
          ...county,
          provider_count: countyProviders.length
        }
      })
    },

    setHighchartsTheme() {
      const isDark = this.theme.global.name.value === 'dark'
      window.Highcharts.setOptions({
        chart: {
          backgroundColor: isDark ? '#1e1e1e' : '#ffffff',
          style: { fontFamily: 'Roboto, sans-serif' }
        },
        title: { style: { color: isDark ? '#ffffff' : '#333333' } },
        subtitle: { style: { color: isDark ? '#cccccc' : '#666666' } },
        legend: { itemStyle: { color: isDark ? '#ffffff' : '#333333' } },
        tooltip: {
          backgroundColor: isDark ? '#2d2d2d' : '#ffffff',
          style: { color: isDark ? '#ffffff' : '#333333' },
          borderColor: isDark ? '#555555' : '#cccccc'
        }
      })
    },

    updateChartTheme() {
      if (!this.chart) return
      
      this.setHighchartsTheme()
      const isDark = this.theme.global.name.value === 'dark'
      
      this.chart.update({
        chart: { backgroundColor: isDark ? '#1e1e1e' : '#ffffff' },
        title: { style: { color: isDark ? '#ffffff' : '#333333' } },
        subtitle: { style: { color: isDark ? '#cccccc' : '#666666' } }
      })
    },

    identifyMissingCounties() {
      const allKansasCounties = Object.keys(countyMapping)
      const dataCounties = this.kansasCountyData.map(c => c.county)
      this.missingCounties = allKansasCounties.filter(county => !dataCounties.includes(county))
    },

    getTooltipFormat() {
      if (this.selectedMetric === 'Provider Density') {
        return 'Providers: <b>{point.provider_count}</b><br/>' +
               'Density: <b>{point.density}</b><br/>' +
               'Claims: <b>{point.claims_count:,.0f}</b><br/>' +
               'Avg Claim: <b>${point.avg_claim_amount:,.2f}</b><br/>'
      } else if (this.selectedMetric === 'Claims Volume') {
        return 'Claims: <b>{point.claims_count:,.0f}</b><br/>' +
               'Providers: <b>{point.provider_count}</b><br/>' +
               'Density: <b>{point.density}</b><br/>' +
               'Avg Claim: <b>${point.avg_claim_amount:,.2f}</b><br/>'
      } else {
        return 'Coverage Ratio: <b>{point.value}</b><br/>' +
               'Providers: <b>{point.provider_count}</b><br/>' +
               'Claims: <b>{point.claims_count:,.0f}</b><br/>' +
               'Density: <b>{point.density}</b><br/>'
      }
    },

    async exportToPDF() {
      try {
        console.log('[INFO] Starting PDF export process')
        const html2canvas = (await import('html2canvas')).default
        const { jsPDF } = await import('jspdf')
        
        const countyName = this.selectedCounty ? this.selectedCounty.county : 'Kansas'
        const date = new Date().toISOString().split('T')[0]
        const filename = `${countyName}-County-Healthcare-Provider-Recommendations-${date}.pdf`
        
        console.log('[INFO] Generating PDF for:', filename)
        const element = document.querySelector('.v-main')
        const canvas = await html2canvas(element, {
          scale: 0.8,
          useCORS: true,
          allowTaint: true,
          backgroundColor: '#ffffff'
        })
        
        const imgData = canvas.toDataURL('image/png')
        const pdf = new jsPDF('p', 'mm', 'a4')
        const imgWidth = 210
        const pageHeight = 297
        const imgHeight = (canvas.height * imgWidth) / canvas.width
        let heightLeft = imgHeight
        
        let position = 0
        pdf.addImage(imgData, 'PNG', 0, position, imgWidth, imgHeight)
        heightLeft -= pageHeight
        
        while (heightLeft >= 0) {
          position = heightLeft - imgHeight
          pdf.addPage()
          pdf.addImage(imgData, 'PNG', 0, position, imgWidth, imgHeight)
          heightLeft -= pageHeight
        }
        
        pdf.save(filename)
        console.log('[INFO] PDF export completed successfully:', filename)
      } catch (error) {
        console.error('[ERROR] PDF export failed:', error)
        alert('PDF export failed. Please try again.')
      }
    }
  }
}
</script>

<style scoped>
.skip-links {
  position: absolute;
  top: -40px;
  left: 6px;
  z-index: 1000;
}

.skip-link {
  position: absolute;
  top: -40px;
  left: 6px;
  background: #000;
  color: #fff;
  padding: 8px;
  text-decoration: none;
  border-radius: 4px;
  font-weight: bold;
  z-index: 1001;
}

.skip-link:focus {
  top: 6px;
}
</style>