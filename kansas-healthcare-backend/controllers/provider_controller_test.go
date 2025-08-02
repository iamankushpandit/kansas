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

type MockProviderService struct {
	mock.Mock
}

func (m *MockProviderService) GetAllProviders() ([]models.Provider, error) {
	args := m.Called()
	return args.Get(0).([]models.Provider), args.Error(1)
}

func (m *MockProviderService) GetProviderNetworks() ([]models.ProviderNetwork, error) {
	args := m.Called()
	return args.Get(0).([]models.ProviderNetwork), args.Error(1)
}

func (m *MockProviderService) GetFilteredProviders(filter models.FilterRequest) ([]models.Provider, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.Provider), args.Error(1)
}

func TestGetProviders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockProviderService)
	controller := NewProviderController(mockService)
	
	expectedProviders := []models.Provider{
		{ProviderID: "1", NPI: "123456789", ProviderType: "Primary Care", Status: "Active", County: "Sedgwick"},
		{ProviderID: "2", NPI: "987654321", ProviderType: "Cardiology", Status: "Active", County: "Johnson"},
	}
	
	mockService.On("GetAllProviders").Return(expectedProviders, nil)
	
	router := gin.New()
	router.GET("/providers", controller.GetProviders)
	
	req, _ := http.NewRequest("GET", "/providers", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []models.Provider
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedProviders, response)
	
	mockService.AssertExpectations(t)
}

func TestGetFilteredData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockProviderService)
	controller := NewProviderController(mockService)
	
	filter := models.FilterRequest{
		Specialty: "Primary Care",
		Network:   "Commercial",
		Metric:    "Provider Density",
	}
	
	expectedProviders := []models.Provider{
		{ProviderID: "1", NPI: "123456789", ProviderType: "Primary Care", Status: "Active", County: "Sedgwick"},
	}
	
	mockService.On("GetFilteredProviders", filter).Return(expectedProviders, nil)
	
	router := gin.New()
	router.POST("/filters", controller.GetFilteredData)
	
	jsonData, _ := json.Marshal(filter)
	req, _ := http.NewRequest("POST", "/filters", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response []models.Provider
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedProviders, response)
	
	mockService.AssertExpectations(t)
}

func TestGetFilteredDataBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockService := new(MockProviderService)
	controller := NewProviderController(mockService)
	
	router := gin.New()
	router.POST("/filters", controller.GetFilteredData)
	
	req, _ := http.NewRequest("POST", "/filters", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
}