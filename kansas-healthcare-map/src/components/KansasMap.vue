<template>
  <section aria-labelledby="map-section-title">
    <h2 id="map-section-title" class="sr-only">Interactive Kansas County Map</h2>
    <div 
      id="kansas-map" 
      style="height: 500px; width: 100%;" 
      role="img" 
      :aria-label="`Interactive map of Kansas showing ${selectedMetric.toLowerCase()} by county. ${selectedCounty ? selectedCounty.county + ' county is currently selected.' : 'Click on a county to view details.'}`"
      tabindex="0"
      @keydown="handleMapKeydown"
    ></div>
    
    <MapLegend 
      :selected-metric="selectedMetric"
      @legend-ready="$emit('legend-ready')"
    />
  </section>
</template>

<script>
import MapLegend from './MapLegend.vue'

export default {
  name: 'KansasMap',
  components: {
    MapLegend
  },
  props: {
    selectedMetric: String,
    selectedCounty: Object
  },
  emits: ['county-click', 'legend-ready'],
  methods: {
    handleMapKeydown(event) {
      if (event.key === 'Enter' || event.key === ' ') {
        event.preventDefault()
        this.announceToScreenReader('Use Tab to navigate to county selection controls, or click on the map to select a county.')
      }
    },

    announceToScreenReader(message) {
      const announcement = document.createElement('div')
      announcement.setAttribute('aria-live', 'polite')
      announcement.setAttribute('aria-atomic', 'true')
      announcement.className = 'sr-only'
      announcement.textContent = message
      document.body.appendChild(announcement)
      
      setTimeout(() => {
        document.body.removeChild(announcement)
      }, 1000)
    }
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

#kansas-map {
  border: 1px solid #e0e0e0;
  border-radius: 4px;
  min-height: 500px;
}

#kansas-map:focus {
  outline: 3px solid #1976d2;
  outline-offset: 2px;
}
</style>