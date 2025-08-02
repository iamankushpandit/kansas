package controllers

import (
	"bytes"
	"encoding/json"
	"kansas-healthcare-api/models"
	"kansas-healthcare-api/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAnalyticsService struct {
	mock.Mock
}

func (m *MockAnalyticsService) GetAllCountyData() ([]models.CountyStats, error) {
	args := m.Called()
	return args.Get(0).([]models.CountyStats), args.Error(1)
}

func (m *MockAnalyticsService) GetCountyData(county string) (*models.CountyStats, error) {
	args := m.Called(county)
	return args.Get(0).(*models.CountyStats), args.Error(1)
}

func (m *MockAnalyticsService) GetActiveProviderCount() (map[string]interface{}, error) {
	args := m.Called()
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (m *MockAnalyticsService) GetTerminatedNetworkAnalysis(networkId string) (map[string]interface{}, error) {
	args := m.Called(networkId)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (m *MockAnalyticsService) GetSpecialtyDensityAnalysis(county string) (map[string]interface{}, error) {
	args := m.Called(county)
	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func TestGetAllCountyData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockAnalyticsService)
	controller := NewAnalyticsController(mockService)
	
	expectedData := []models.CountyStats{
		{County: "Sedgwick", ProviderCount: 100, ClaimsCount: 5000, AvgClaimAmount: 250.50, Density: "high"},
	}
	
	mockService.On("GetAllCountyData").Return(expectedData, nil)
	
	router := gin.New()
	router.GET("/county-data", controller.GetAllCountyData)
	
	req, _ := http.NewRequest("GET", "/county-data", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []models.CountyStats
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, response)
	
	mockService.AssertExpectations(t)
}

func TestGetCountyData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockAnalyticsService)
	controller := NewAnalyticsController(mockService)
	
	expectedData := &models.CountyStats{
		County: "Sedgwick", 
		ProviderCount: 100, 
		ClaimsCount: 5000, 
		AvgClaimAmount: 250.50, 
		Density: "high",
	}
	
	mockService.On("GetCountyData", "Sedgwick").Return(expectedData, nil)
	
	router := gin.New()
	router.GET("/county-data/:county", controller.GetCountyData)
	
	req, _ := http.NewRequest("GET", "/county-data/Sedgwick", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response models.CountyStats
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, *expectedData, response)
	
	mockService.AssertExpectations(t)
}

func TestGetSpecialtyDensityAnalysis(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockAnalyticsService)
	controller := NewAnalyticsController(mockService)
	
	expectedData := map[string]interface{}{
		"specialty_densities": []map[string]interface{}{
			{"name": "Primary Care", "count": 10},
			{"name": "Cardiology", "count": 5},
		},
	}
	
	mockService.On("GetSpecialtyDensityAnalysis", "Sedgwick").Return(expectedData, nil)
	
	router := gin.New()
	router.GET("/specialty-density/:county", controller.GetSpecialtyDensityAnalysis)
	
	req, _ := http.NewRequest("GET", "/specialty-density/Sedgwick", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, response)
	
	mockService.AssertExpectations(t)
}