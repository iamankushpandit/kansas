import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import MissingDataAlert from '@/components/MissingDataAlert.vue'

describe('MissingDataAlert', () => {
  it('renders missing counties correctly', () => {
    const missingCounties = ['County A', 'County B', 'County C']
    const wrapper = mount(MissingDataAlert, {
      props: { missingCounties }
    })

    expect(wrapper.text()).toContain('Missing Data (3 counties):')
    expect(wrapper.text()).toContain('County A, County B, County C')
  })

  it('handles single missing county', () => {
    const missingCounties = ['Single County']
    const wrapper = mount(MissingDataAlert, {
      props: { missingCounties }
    })

    expect(wrapper.text()).toContain('Missing Data (1 counties):')
    expect(wrapper.text()).toContain('Single County')
  })

  it('handles empty missing counties array', () => {
    const missingCounties = []
    const wrapper = mount(MissingDataAlert, {
      props: { missingCounties }
    })

    expect(wrapper.text()).toContain('Missing Data (0 counties):')
    expect(wrapper.text()).toContain('')
  })

  it('renders alert content', () => {
    const wrapper = mount(MissingDataAlert, {
      props: { missingCounties: ['Test County'] }
    })

    expect(wrapper.text()).toContain('Missing Data (1 counties):')
    expect(wrapper.text()).toContain('Test County')
  })

  it('renders with alert wrapper', () => {
    const wrapper = mount(MissingDataAlert, {
      props: { missingCounties: ['Test County'] }
    })

    expect(wrapper.find('div').exists()).toBe(true)
  })
})