import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import CountyMetrics from '@/components/CountyMetrics.vue'

describe('CountyMetrics', () => {
  const mockCounty = {
    county: 'Sedgwick',
    provider_count: 100,
    claims_count: 5000,
    avg_claim_amount: 250.75,
    density: 'high'
  }

  const mockTerminatedAnalysis = {
    percentage_terminated: 15.5,
    term_network_count: 12
  }

  it('renders county metrics correctly', () => {
    const wrapper = mount(CountyMetrics, {
      props: {
        selectedCounty: mockCounty,
        countyTerminatedAnalysis: mockTerminatedAnalysis
      }
    })

    expect(wrapper.text()).toContain('Active Providers: 100')
    expect(wrapper.text()).toContain('Density Classification: high')
    expect(wrapper.text()).toContain('Total Claims: 5,000')
    expect(wrapper.text()).toContain('Avg Claim Amount: $250.75')
  })

  it('renders termination analysis when provided', () => {
    const wrapper = mount(CountyMetrics, {
      props: {
        selectedCounty: mockCounty,
        countyTerminatedAnalysis: mockTerminatedAnalysis
      }
    })

    expect(wrapper.text()).toContain('Termination Rate: 15.5%')
    expect(wrapper.text()).toContain('Providers Terminated: 12')
  })

  it('calculates claims per provider correctly', () => {
    const wrapper = mount(CountyMetrics, {
      props: {
        selectedCounty: mockCounty,
        countyTerminatedAnalysis: mockTerminatedAnalysis
      }
    })

    expect(wrapper.vm.getClaimsPerProvider()).toBe(50) // 5000/100
    expect(wrapper.text()).toContain('Claims per Provider: 50')
  })

  it('handles zero provider count', () => {
    const zeroProviderCounty = { ...mockCounty, provider_count: 0 }
    const wrapper = mount(CountyMetrics, {
      props: {
        selectedCounty: zeroProviderCounty,
        countyTerminatedAnalysis: mockTerminatedAnalysis
      }
    })

    expect(wrapper.vm.getClaimsPerProvider()).toBe('N/A')
  })

  it('calculates network coverage ratio correctly', () => {
    const wrapper = mount(CountyMetrics, {
      props: {
        selectedCounty: mockCounty,
        countyTerminatedAnalysis: mockTerminatedAnalysis
      }
    })

    expect(wrapper.vm.getNetworkCoverageRatio()).toBe(20) // Math.floor(100/5000*1000)
  })

  it('handles missing termination analysis', () => {
    const wrapper = mount(CountyMetrics, {
      props: {
        selectedCounty: mockCounty,
        countyTerminatedAnalysis: null
      }
    })

    expect(wrapper.text()).not.toContain('Network Termination')
    expect(wrapper.text()).not.toContain('Termination Rate')
  })
})