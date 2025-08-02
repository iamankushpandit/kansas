import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import MapLegend from '@/components/MapLegend.vue'

describe('MapLegend', () => {
  it('renders provider density legend correctly', () => {
    const wrapper = mount(MapLegend, {
      props: { selectedMetric: 'Provider Density' }
    })

    expect(wrapper.text()).toContain('Provider Density Legend')
    expect(wrapper.text()).toContain('High (400+ providers)')
    expect(wrapper.text()).toContain('Medium (100-399 providers)')
    expect(wrapper.text()).toContain('Low (50-99 providers)')
    expect(wrapper.text()).toContain('Critical (<50 providers)')
  })

  it('renders claims volume legend correctly', () => {
    const wrapper = mount(MapLegend, {
      props: { selectedMetric: 'Claims Volume' }
    })

    expect(wrapper.text()).toContain('Claims Volume Legend')
    expect(wrapper.text()).toContain('High (7500+ claims)')
    expect(wrapper.text()).toContain('Medium-High (5000-7499 claims)')
    expect(wrapper.text()).toContain('Medium-Low (2500-4999 claims)')
    expect(wrapper.text()).toContain('Low (<2500 claims)')
  })

  it('renders network coverage legend correctly', () => {
    const wrapper = mount(MapLegend, {
      props: { selectedMetric: 'Network Coverage' }
    })

    expect(wrapper.text()).toContain('Network Coverage Legend')
    expect(wrapper.text()).toContain('Excellent (38+ ratio)')
    expect(wrapper.text()).toContain('Good (25-37 ratio)')
    expect(wrapper.text()).toContain('Fair (13-24 ratio)')
    expect(wrapper.text()).toContain('Poor (<13 ratio)')
  })

  it('has proper accessibility attributes', () => {
    const wrapper = mount(MapLegend, {
      props: { selectedMetric: 'Provider Density' }
    })

    expect(wrapper.find('[role="region"]').exists()).toBe(true)
    expect(wrapper.find('#legend-title').exists()).toBe(true)
    expect(wrapper.find('[role="list"]').exists()).toBe(true)
    expect(wrapper.findAll('[role="listitem"]')).toHaveLength(4)
  })

  it('displays correct colors for legend items', () => {
    const wrapper = mount(MapLegend, {
      props: { selectedMetric: 'Provider Density' }
    })

    const legendColors = wrapper.findAll('.legend-color')
    expect(legendColors[0].attributes('style')).toContain('background-color: rgb(46, 139, 87)')
    expect(legendColors[1].attributes('style')).toContain('background-color: rgb(144, 238, 144)')
    expect(legendColors[2].attributes('style')).toContain('background-color: rgb(255, 215, 0)')
    expect(legendColors[3].attributes('style')).toContain('background-color: rgb(255, 99, 71)')
  })

  it('generates correct legend title', () => {
    const wrapper = mount(MapLegend, {
      props: { selectedMetric: 'Test Metric' }
    })

    expect(wrapper.vm.getMetricLegendTitle()).toBe('Test Metric Legend')
  })
})