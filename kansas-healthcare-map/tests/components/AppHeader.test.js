import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import AppHeader from '@/components/AppHeader.vue'

describe('AppHeader', () => {
  it('renders title correctly', () => {
    const wrapper = mount(AppHeader, {
      props: {
        selectedCounty: null,
        loading: false,
        isDark: false
      }
    })

    expect(wrapper.text()).toContain('Kansas Healthcare Provider Network')
    expect(wrapper.find('h1').exists()).toBe(true)
  })

  it('emits export-pdf event when PDF button clicked', async () => {
    const wrapper = mount(AppHeader, {
      props: {
        selectedCounty: 'Sedgwick',
        loading: false,
        isDark: false
      }
    })

    const pdfButton = wrapper.find('[aria-label*="Export"]')
    await pdfButton.trigger('click')

    expect(wrapper.emitted('export-pdf')).toBeTruthy()
  })

  it('emits toggle-theme event when theme button clicked', async () => {
    const wrapper = mount(AppHeader, {
      props: {
        selectedCounty: null,
        loading: false,
        isDark: false
      }
    })

    const themeButton = wrapper.find('[aria-label*="Switch to"]')
    await themeButton.trigger('click')

    expect(wrapper.emitted('toggle-theme')).toBeTruthy()
  })

  it('disables PDF button when loading', () => {
    const wrapper = mount(AppHeader, {
      props: {
        selectedCounty: null,
        loading: true,
        isDark: false
      }
    })

    const pdfButton = wrapper.find('[aria-label*="Export"]')
    expect(pdfButton.attributes('disabled')).toBeDefined()
  })

  it('shows theme toggle button', () => {
    const wrapper = mount(AppHeader, {
      props: {
        selectedCounty: null,
        loading: false,
        isDark: false
      }
    })

    expect(wrapper.find('[aria-label*="Switch to"]').exists()).toBe(true)
  })

  it('shows different theme label for dark mode', () => {
    const wrapper = mount(AppHeader, {
      props: {
        selectedCounty: null,
        loading: false,
        isDark: true
      }
    })

    expect(wrapper.find('[aria-label*="Switch to light"]').exists()).toBe(true)
  })

  it('has proper accessibility attributes', () => {
    const wrapper = mount(AppHeader, {
      props: {
        selectedCounty: 'Sedgwick',
        loading: false,
        isDark: false
      }
    })

    expect(wrapper.find('[role="banner"]').exists()).toBe(true)
    expect(wrapper.find('[aria-label*="Export Sedgwick county"]').exists()).toBe(true)
    expect(wrapper.find('[aria-label*="Switch to dark theme"]').exists()).toBe(true)
  })
})