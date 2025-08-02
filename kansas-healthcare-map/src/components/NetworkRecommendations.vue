<template>
  <section class="mt-3" aria-labelledby="network-recommendations-title">
    <v-row>
      <v-col cols="12">
        <h4 id="network-recommendations-title" class="text-subtitle-2 text-primary">Network Recommendations</h4>
        <v-list density="compact" class="pa-0" role="list" aria-label="Network improvement recommendations">
          <v-list-item 
            v-for="rec in recommendations" 
            :key="rec.id" 
            class="px-0"
            role="listitem"
            :aria-label="`${rec.priority} priority: ${rec.title}. ${rec.description}`"
          >
            <template v-slot:prepend>
              <v-icon 
                :color="rec.priority === 'High' ? 'error' : 'warning'" 
                size="small"
                :aria-label="`${rec.priority} priority indicator`"
              >
                {{ rec.icon }}
              </v-icon>
            </template>
            <v-list-item-title class="text-body-2">{{ rec.title }}</v-list-item-title>
            <v-list-item-subtitle class="text-caption">{{ rec.description }}</v-list-item-subtitle>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>
  </section>
</template>

<script>
export default {
  name: 'NetworkRecommendations',
  props: {
    recommendations: Array
  }
}
</script>