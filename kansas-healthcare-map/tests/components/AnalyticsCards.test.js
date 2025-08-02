import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import AnalyticsCards from '@/components/AnalyticsCards.vue'

describe('AnalyticsCards', () => {
  const mockTerminatedAnalysis = {
    term_network_count: 25,
    service_location_count: 40,
    percentage_terminated: 12.5
  }

  it('renders active provider count', () => {
    const wrapper = mount(AnalyticsCards, {
      props: {
        activeProviderCount: 1500,
        terminatedAnalysis: null,
        selectedNetwork: 'Commercial'
      }
    })

    expect(wrapper.text()).toContain('1,500')
    expect(wrapper.text()).toContain('Total Active Providers Across Kansas')
  })

  it('renders terminated analysis when provided', () => {
    const wrapper = mount(AnalyticsCards, {
      props: {
        activeProviderCount: 1500,
        terminatedAnalysis: mockTerminatedAnalysis,
        selectedNetwork: 'Commercial'
      }
    })

    expect(wrapper.text()).toContain('Commercial Network Stability Analysis')
    expect(wrapper.text()).toContain('25')
    expect(wrapper.text()).toContain('40')
    expect(wrapper.text()).toContain('12.5%')
  })

  it('does not render terminated analysis when not provided', () => {
    const wrapper = mount(AnalyticsCards, {
      props: {
        activeProviderCount: 1500,
        terminatedAnalysis: null,
        selectedNetwork: 'Commercial'
      }
    })

    expect(wrapper.text()).not.toContain('Network Stability Analysis')
    expect(wrapper.text()).not.toContain('Providers Who Left Network')
  })

  it('formats numbers correctly', () => {
    const wrapper = mount(AnalyticsCards, {
      props: {
        activeProviderCount: 12345,
        terminatedAnalysis: mockTerminatedAnalysis,
        selectedNetwork: 'Medicare'
      }
    })

    expect(wrapper.text()).toContain('12,345')
  })

  it('has proper accessibility attributes', () => {
    const wrapper = mount(AnalyticsCards, {
      props: {
        activeProviderCount: 1500,
        terminatedAnalysis: mockTerminatedAnalysis,
        selectedNetwork: 'Commercial'
      }
    })

    expect(wrapper.find('[aria-labelledby="analytics-title"]').exists()).toBe(true)
    expect(wrapper.find('#analytics-title').exists()).toBe(true)
    expect(wrapper.find('[role="region"]').exists()).toBe(true)
    expect(wrapper.findAll('[aria-label*="providers"]')).toHaveLength(3)
  })

  it('shows correct network name in analysis', () => {
    const wrapper = mount(AnalyticsCards, {
      props: {
        activeProviderCount: 1500,
        terminatedAnalysis: mockTerminatedAnalysis,
        selectedNetwork: 'Tricare'
      }
    })

    expect(wrapper.text()).toContain('Tricare Network Stability Analysis')
  })

  it('handles zero values correctly', () => {
    const zeroAnalysis = {
      term_network_count: 0,
      service_location_count: 0,
      percentage_terminated: 0.0
    }

    const wrapper = mount(AnalyticsCards, {
      props: {
        activeProviderCount: 0,
        terminatedAnalysis: zeroAnalysis,
        selectedNetwork: 'Commercial'
      }
    })

    expect(wrapper.text()).toContain('0')
    expect(wrapper.text()).toContain('0.0%')
  })
})