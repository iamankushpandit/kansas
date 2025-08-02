import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const healthcareApi = {
  // County data
  async getAllCountyData() {
    const response = await api.get('/county-data')
    return response.data
  },

  async getCountyData(county) {
    const response = await api.get(`/county-data/${county}`)
    return response.data
  },

  // Provider data
  async getProviders() {
    const response = await api.get('/providers')
    return response.data
  },

  async getFilteredProviders(filter) {
    const response = await api.post('/filters', filter)
    return response.data
  },

  // Analytics
  async getActiveProviderCount() {
    const response = await api.get('/active-providers')
    return response.data
  },

  async getTerminatedNetworkAnalysis(networkId) {
    const response = await api.get(`/terminated-analysis?network_id=${networkId}`)
    return response.data
  },

  // Recommendations
  async getRecommendations(county) {
    const response = await api.get(`/recommendations/${county}`)
    return response.data
  }
}

export default api