<template>
  <section class="mt-3" aria-labelledby="specialty-density-title">
    <v-row>
      <v-col cols="12">
          <h4 id="specialty-density-title" class="text-subtitle-2 text-primary">Specialty Density Analysis</h4>
        <p class="text-caption mb-2">Provider specialties sorted by density (least to most dense)</p>
        <div role="list" aria-label="Medical specialties ranked by provider density">
          <v-row>
            <v-col 
              v-for="rec in recommendations" 
              :key="rec.id" 
              cols="6" 
              class="py-1"
              role="listitem"
            >
              <div 
                class="d-flex align-center"
                :aria-label="`${rec.title} specialty: ${rec.description}. Priority level: ${rec.priority}`"
              >
                <v-icon 
                  :color="rec.priority === 'High' ? 'error' : rec.priority === 'Medium' ? 'warning' : 'success'" 
                  size="small" 
                  class="mr-2"
                  :aria-label="`${rec.priority} priority indicator`"
                >
                  {{ rec.icon }}
                </v-icon>
                <div class="text-body-2">
                  <span class="font-weight-medium">{{ rec.title }}</span>
                  <span class="text-caption ml-1">{{ rec.description }}</span>
                </div>
              </div>
            </v-col>
          </v-row>
        </div>
      </v-col>
    </v-row>
  </section>
</template>

<script>
export default {
  name: 'SpecialtyDensityAnalysis',
  props: {
    recommendations: Array
  }
}
</script>

<style scoped>
.tooltip-content {
  max-width: 300px;
  line-height: 1.4;
}
</style>