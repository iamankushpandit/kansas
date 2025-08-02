<template>
  <v-alert type="warning" class="mt-3" density="compact">
    <div class="text-caption">
      <strong>Missing Data ({{ missingCounties.length }} counties):</strong><br>
      {{ missingCounties.join(', ') }}
    </div>
  </v-alert>
</template>

<script>
export default {
  name: 'MissingDataAlert',
  props: {
    missingCounties: Array
  }
}
</script>