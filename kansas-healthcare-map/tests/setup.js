import { config } from '@vue/test-utils'

// Mock Vuetify components for tests
config.global.stubs = {
  'v-app': { template: '<div><slot /></div>' },
  'v-main': { template: '<div><slot /></div>' },
  'v-container': { template: '<div><slot /></div>' },
  'v-row': { template: '<div><slot /></div>' },
  'v-col': { template: '<div><slot /></div>' },
  'v-card': { template: '<div><slot /></div>' },
  'v-card-text': { template: '<div><slot /></div>' },
  'v-card-title': { template: '<div><slot /></div>' },
  'v-app-bar': { template: '<div><slot /></div>' },
  'v-app-bar-title': { template: '<div><slot /></div>' },
  'v-spacer': { template: '<div></div>' },
  'v-btn': { template: '<button><slot /></button>' },
  'v-icon': { template: '<i><slot /></i>' },
  'v-select': { template: '<select><slot /></select>' },
  'v-alert': { template: '<div><slot /></div>' },
  'v-alert-title': { template: '<div><slot /></div>' },
  'v-overlay': { template: '<div><slot /></div>' },
  'v-progress-circular': { template: '<div></div>' },
  'v-tooltip': { template: '<div><slot name="activator" /><slot /></div>' }
}

// Mock window.matchMedia
Object.defineProperty(window, 'matchMedia', {
  writable: true,
  value: vi.fn().mockImplementation(query => ({
    matches: false,
    media: query,
    onchange: null,
    addListener: vi.fn(),
    removeListener: vi.fn(),
    addEventListener: vi.fn(),
    removeEventListener: vi.fn(),
    dispatchEvent: vi.fn(),
  })),
})

// Mock ResizeObserver
global.ResizeObserver = vi.fn().mockImplementation(() => ({
  observe: vi.fn(),
  unobserve: vi.fn(),
  disconnect: vi.fn(),
}))

// Mock import.meta.env
vi.stubGlobal('import', {
  meta: {
    env: {
      VITE_API_BASE_URL: 'http://localhost:8080/api/v1'
    }
  }
})