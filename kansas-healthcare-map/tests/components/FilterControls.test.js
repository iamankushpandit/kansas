import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import FilterControls from '@/components/FilterControls.vue'

describe('FilterControls', () => {
  const defaultProps = {
    selectedSpecialty: 'All',
    selectedNetwork: 'All',
    selectedMetric: 'provider_density',
    specialties: ['All', 'Cardiology', 'Dermatology'],
    networks: ['All', 'Commercial', 'Medicare'],
    metrics: [
      { title: 'Provider Density', value: 'provider_density' },
      { title: 'Claims Volume', value: 'claims_volume' }
    ],
    loading: false
  }

  it('renders all filter controls', () => {
    const wrapper = mount(FilterControls, {
      props: defaultProps
    })

    expect(wrapper.find('[label="Provider Specialty"]').exists()).toBe(true)
    expect(wrapper.find('[label="Network Type"]').exists()).toBe(true)
    expect(wrapper.find('[label="Display Metric"]').exists()).toBe(true)
  })

  it('emits update-specialty event', async () => {
    const wrapper = mount(FilterControls, {
      props: defaultProps
    })

    await wrapper.vm.$emit('update-specialty', 'Cardiology')

    expect(wrapper.emitted('update-specialty')).toBeTruthy()
    expect(wrapper.emitted('update-specialty')[0]).toEqual(['Cardiology'])
  })

  it('emits update-network event', async () => {
    const wrapper = mount(FilterControls, {
      props: defaultProps
    })

    await wrapper.vm.$emit('update-network', 'Medicare')

    expect(wrapper.emitted('update-network')).toBeTruthy()
    expect(wrapper.emitted('update-network')[0]).toEqual(['Medicare'])
  })

  it('emits update-metric event', async () => {
    const wrapper = mount(FilterControls, {
      props: defaultProps
    })

    await wrapper.vm.$emit('update-metric', 'claims_volume')

    expect(wrapper.emitted('update-metric')).toBeTruthy()
    expect(wrapper.emitted('update-metric')[0]).toEqual(['claims_volume'])
  })

  it('disables controls when loading', () => {
    const wrapper = mount(FilterControls, {
      props: { ...defaultProps, loading: true }
    })

    const selects = wrapper.findAll('.v-select')
    selects.forEach(select => {
      expect(select.props('disabled')).toBe(true)
    })
  })

  it('has proper accessibility attributes', () => {
    const wrapper = mount(FilterControls, {
      props: defaultProps
    })

    expect(wrapper.find('[role="group"]').exists()).toBe(true)
    expect(wrapper.find('#filter-controls-title').exists()).toBe(true)
    expect(wrapper.find('[aria-describedby="specialty-help"]').exists()).toBe(true)
    expect(wrapper.find('[aria-describedby="network-help"]').exists()).toBe(true)
    expect(wrapper.find('[aria-describedby="metric-help"]').exists()).toBe(true)
  })
})