package services

import (
	"kansas-healthcare-api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetCountyStats() ([]models.CountyStats, error) {
	args := m.Called()
	return args.Get(0).([]models.CountyStats), args.Error(1)
}

func (m *MockRepository) GetCountyStatsByName(county string) (*models.CountyStats, error) {
	args := m.Called(county)
	return args.Get(0).(*models.CountyStats), args.Error(1)
}

func (m *MockRepository) GetActiveProviderCount() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) GetTerminatedNetworkCount(networkId string) (int, error) {
	args := m.Called(networkId)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) GetTerminatedServiceLocationCount(networkId string) (int, error) {
	args := m.Called(networkId)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) GetProvidersInCounty(county string) ([]models.Provider, error) {
	args := m.Called(county)
	return args.Get(0).([]models.Provider), args.Error(1)
}

func (m *MockRepository) GetRadiusAnalysis(county string, radius int, networkId string) (map[string]interface{}, error) {
	args := m.Called(county, radius, networkId)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func TestGetAllCountyData(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewAnalyticsService(mockRepo)
	
	expectedData := []models.CountyStats{
		{County: "Sedgwick", ProviderCount: 100, ClaimsCount: 5000, AvgClaimAmount: 250.50, Density: "high"},
		{County: "Johnson", ProviderCount: 80, ClaimsCount: 4000, AvgClaimAmount: 300.00, Density: "medium"},
	}
	
	mockRepo.On("GetCountyStats").Return(expectedData, nil)
	
	result, err := service.GetAllCountyData()
	
	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
	mockRepo.AssertExpectations(t)
}

func TestGetCountyData(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewAnalyticsService(mockRepo)
	
	expectedData := &models.CountyStats{
		County: "Sedgwick", 
		ProviderCount: 100, 
		ClaimsCount: 5000, 
		AvgClaimAmount: 250.50, 
		Density: "high",
	}
	
	mockRepo.On("GetCountyStatsByName", "Sedgwick").Return(expectedData, nil)
	
	result, err := service.GetCountyData("Sedgwick")
	
	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
	mockRepo.AssertExpectations(t)
}

func TestGetActiveProviderCount(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewAnalyticsService(mockRepo)
	
	expectedCount := 1000
	mockRepo.On("GetActiveProviderCount").Return(expectedCount, nil)
	
	result, err := service.GetActiveProviderCount()
	
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, result["total_active_providers"])
	mockRepo.AssertExpectations(t)
}

func TestGetSpecialtyDensityAnalysis(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewAnalyticsService(mockRepo)
	
	providers := []models.Provider{
		{ProviderID: "1", ProviderType: "Primary Care", Status: "Active"},
		{ProviderID: "2", ProviderType: "Primary Care", Status: "Active"},
		{ProviderID: "3", ProviderType: "Cardiology", Status: "Active"},
		{ProviderID: "4", ProviderType: "Cardiology", Status: "Terminated"},
	}
	
	mockRepo.On("GetProvidersInCounty", "Sedgwick").Return(providers, nil)
	
	result, err := service.GetSpecialtyDensityAnalysis("Sedgwick")
	
	assert.NoError(t, err)
	assert.Contains(t, result, "specialty_densities")
	
	densities := result["specialty_densities"].([]SpecialtyDensity)
	assert.Len(t, densities, 2)
	
	// Should be sorted by count (ascending)
	assert.Equal(t, "Cardiology", densities[0].Name)
	assert.Equal(t, 1, densities[0].Count)
	assert.Equal(t, "Primary Care", densities[1].Name)
	assert.Equal(t, 2, densities[1].Count)
	
	mockRepo.AssertExpectations(t)
}

func TestGetTerminatedNetworkAnalysis(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewAnalyticsService(mockRepo)
	
	mockRepo.On("GetActiveProviderCount").Return(1000, nil)
	mockRepo.On("GetTerminatedNetworkCount", "Commercial").Return(50, nil)
	mockRepo.On("GetTerminatedServiceLocationCount", "Commercial").Return(25, nil)
	
	result, err := service.GetTerminatedNetworkAnalysis("Commercial")
	
	assert.NoError(t, err)
	assert.Equal(t, 50, result["term_network_count"])
	assert.Equal(t, 25, result["service_location_count"])
	assert.Equal(t, 5.0, result["percentage_terminated"])
	assert.Equal(t, 1000, result["total_active_providers"])
	
	mockRepo.AssertExpectations(t)
}