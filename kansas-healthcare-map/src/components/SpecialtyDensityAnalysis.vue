<template>
  <section class="mt-3" aria-labelledby="specialty-density-title">
    <v-row>
      <v-col cols="12">
        <h4 id="specialty-density-title" class="text-subtitle-2 text-primary mb-3">Specialty Density Analysis</h4>
        <v-data-table
          :headers="headers"
          :items="tableData"
          :items-per-page="-1"
          class="elevation-1"
          density="compact"
        >
          <template v-slot:item.current_density="{ item }">
            <span :class="getPriorityColor(item.gap)">{{ item.current_density }}</span>
          </template>
          <template v-slot:item.gap="{ item }">
            <v-chip 
              :color="getPriorityChipColor(item.gap)" 
              size="small" 
              variant="flat"
            >
              {{ item.gap === null ? 'N/A' : (item.gap > 0 ? `+${item.gap.toFixed(2)}` : item.gap.toFixed(2)) }}
            </v-chip>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </section>
</template>

<script>
export default {
  name: 'SpecialtyDensityAnalysis',
  props: {
    recommendations: Array,
    specialtyDensityAnalysis: Object
  },
  computed: {
    headers() {
      return [
        { title: 'Specialty', key: 'specialty', sortable: false },
        { title: 'Current Density', key: 'current_density', sortable: false },
        { title: 'Recommended', key: 'recommended_density', sortable: false },
        { title: 'Gap', key: 'gap', sortable: true }
      ]
    },
    tableData() {
      if (!this.specialtyDensityAnalysis?.specialty_densities) return []
      
      return this.specialtyDensityAnalysis.specialty_densities.map(item => ({
        specialty: item.name,
        current_density: item.count === 0 ? 'No providers' : `${(item.count / 700).toFixed(3)}/sq mi`,
        recommended_density: `${(item.recommended || 0).toFixed(2)}/sq mi`,
        gap: item.gap !== undefined ? item.gap : null
      }))
    }
  },
  methods: {
    getPriorityColor(gap) {
      if (gap === null || gap === undefined) return 'text-error'
      if (gap > 1.0) return 'text-error'
      if (gap > 0.3) return 'text-warning'
      return 'text-success'
    },
    getPriorityChipColor(gap) {
      if (gap === null || gap === undefined) return 'error'
      if (gap > 1.0) return 'error'
      if (gap > 0.3) return 'warning'
      return 'success'
    }
  }
}
</script>

<style scoped>
.tooltip-content {
  max-width: 300px;
  line-height: 1.4;
}
</style>