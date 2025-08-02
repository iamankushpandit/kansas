// Vuetify 3 Configuration for Healthcare Analytics Platform
// 
// ARCHITECTURAL DECISION: Vuetify 3 Material Design Framework
// 
// Vuetify was selected over other UI frameworks for healthcare-specific reasons:
// 
// 1. ACCESSIBILITY COMPLIANCE:
//    - Built-in WCAG 2.1 AA compliance for healthcare equity
//    - Screen reader support and keyboard navigation out-of-the-box
//    - High contrast themes for visually impaired healthcare professionals
// 
// 2. HEALTHCARE UI CONSISTENCY:
//    - Material Design 3 provides consistent healthcare workflow patterns
//    - 80+ pre-built components reduce custom CSS and improve maintainability
//    - Standardized interaction patterns reduce training time for healthcare staff
// 
// 3. RESPONSIVE HEALTHCARE DESIGN:
//    - Mobile-first approach supports field representatives and rural clinics
//    - Touch-friendly interfaces for tablet-based healthcare applications
//    - Flexible grid system adapts to various healthcare device form factors
// 
// 4. THEME SYSTEM FOR HEALTHCARE BRANDING:
//    - Easy customization for different healthcare network branding
//    - Dark mode support for healthcare professionals working night shifts
//    - Color system supports medical data visualization requirements
//
import { createVuetify } from 'vuetify'        // Vuetify 3 framework for healthcare UI
import * as components from 'vuetify/components' // All Vuetify components for healthcare workflows
import * as directives from 'vuetify/directives' // Vue directives for enhanced healthcare UX
import 'vuetify/styles'                         // Material Design 3 styles
import '@mdi/font/css/materialdesignicons.css'  // Medical and healthcare icons

// Healthcare-optimized Vuetify configuration
export default createVuetify({
  components,
  directives,
  // Healthcare-optimized theme configuration
  theme: {
    defaultTheme: 'light', // Default to light theme for healthcare readability
    themes: {
      light: {
        colors: {
          primary: '#1976D2',   // Healthcare blue - trust and professionalism
          secondary: '#424242', // Neutral gray for secondary healthcare information
          accent: '#82B1FF',    // Light blue accent for healthcare highlights
          error: '#FF5252',     // Medical red for critical healthcare alerts
          info: '#2196F3',      // Information blue for healthcare notifications
          success: '#4CAF50',   // Medical green for positive healthcare outcomes
          warning: '#FFC107'    // Amber warning for healthcare cautions
        }
      },
      dark: {
        colors: {
          primary: '#2196F3',   // Brighter blue for dark mode healthcare visibility
          secondary: '#616161', // Lighter gray for dark mode healthcare readability
          accent: '#FF4081',    // Pink accent for dark mode healthcare highlights
          error: '#FF5252',     // Consistent medical red across themes
          info: '#2196F3',      // Consistent information blue
          success: '#4CAF50',   // Consistent medical green
          warning: '#FFC107'    // Consistent amber warning
        }
      }
    }
  }
})