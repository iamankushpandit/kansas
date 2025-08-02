<template>
  <v-app-bar color="primary" dark role="banner">
    <v-app-bar-title>
      <h1 class="text-h6 ma-0">Kansas Healthcare Provider Network</h1>
    </v-app-bar-title>
    <v-spacer></v-spacer>
    <v-btn 
      icon 
      @click="$emit('export-pdf')" 
      :aria-label="`Export ${selectedCounty ? selectedCounty + ' county' : 'Kansas'} healthcare data to PDF`"
      :disabled="loading"
    >
      <v-icon>mdi-file-pdf-box</v-icon>
    </v-btn>
    <v-btn 
      icon 
      @click="$emit('toggle-theme')"
      :aria-label="`Switch to ${isDark ? 'light' : 'dark'} theme`"
    >
      <v-icon>{{ isDark ? 'mdi-weather-sunny' : 'mdi-weather-night' }}</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
export default {
  name: 'AppHeader',
  props: {
    selectedCounty: String,
    loading: Boolean,
    isDark: Boolean
  },
  emits: ['export-pdf', 'toggle-theme']
}
</script>