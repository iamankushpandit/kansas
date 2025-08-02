package data

import (
	"kansas-healthcare-api/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateProviderDensity(t *testing.T) {
	repo := &JSONRepository{}
	
	tests := []struct {
		providerCount int
		expected      string
	}{
		{250, "high"},
		{150, "medium"},
		{75, "medium"},
		{25, "critical"},
	}
	
	for _, test := range tests {
		result := repo.calculateProviderDensity(test.providerCount)
		assert.Equal(t, test.expected, result)
	}
}

func TestCalculateProviderDensityMiles(t *testing.T) {
	repo := &JSONRepository{}
	
	tests := []struct {
		providerCount int
		expected      string
	}{
		{0, "No providers"},
		{1000, "1.4/sq mi"},
		{50, "~3.7 mi apart"},
	}
	
	for _, test := range tests {
		result := repo.calculateProviderDensityMiles(test.providerCount)
		assert.Equal(t, test.expected, result)
	}
}

func TestGetProvidersInCounty(t *testing.T) {
	repo := &JSONRepository{
		providers: []models.Provider{
			{ProviderID: "1", County: "Sedgwick", Status: "Active"},
			{ProviderID: "2", County: "Sedgwick", Status: "Terminated"},
			{ProviderID: "3", County: "Johnson", Status: "Active"},
		},
	}
	
	result, err := repo.GetProvidersInCounty("Sedgwick")
	
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Sedgwick", result[0].County)
	assert.Equal(t, "Sedgwick", result[1].County)
}

func TestGetActiveProviderCount(t *testing.T) {
	repo := &JSONRepository{
		providers: []models.Provider{
			{ProviderID: "1", Status: "Active"},
			{ProviderID: "2", Status: "Active"},
			{ProviderID: "3", Status: "Terminated"},
		},
	}
	
	result, err := repo.GetActiveProviderCount()
	
	assert.NoError(t, err)
	assert.Equal(t, 2, result)
}

func TestGetFilteredProviders(t *testing.T) {
	repo := &JSONRepository{
		providers: []models.Provider{
			{ProviderID: "1", ProviderType: "Primary Care", Status: "Active"},
			{ProviderID: "2", ProviderType: "Cardiology", Status: "Active"},
			{ProviderID: "3", ProviderType: "Primary Care", Status: "Terminated"},
		},
		providerNetwork: []models.ProviderNetwork{
			{ProviderID: "1", NetworkID: "Commercial", TerminationReason: ""},
			{ProviderID: "2", NetworkID: "Commercial", TerminationReason: ""},
			{ProviderID: "3", NetworkID: "Commercial", TerminationReason: "Left Network"},
		},
	}
	
	filter := models.FilterRequest{
		Specialty: "Primary Care",
		Network:   "Commercial",
		Metric:    "Provider Density",
	}
	
	result, err := repo.GetFilteredProviders(filter)
	
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].ProviderID)
	assert.Equal(t, "Primary Care", result[0].ProviderType)
	assert.Equal(t, "Active", result[0].Status)
}

func TestGetTerminatedNetworkCount(t *testing.T) {
	now := time.Now()
	threeYearsAgo := now.AddDate(-3, 0, 0)
	
	repo := &JSONRepository{
		providerNetwork: []models.ProviderNetwork{
			{ProviderID: "1", NetworkID: "Commercial", TerminationReason: "Left Network", TerminationDate: threeYearsAgo},
			{ProviderID: "2", NetworkID: "Commercial", TerminationReason: "Left Network", TerminationDate: now.AddDate(-1, 0, 0)}, // Too recent
			{ProviderID: "3", NetworkID: "Medicare", TerminationReason: "Left Network", TerminationDate: threeYearsAgo}, // Wrong network
		},
	}
	
	result, err := repo.GetTerminatedNetworkCount("Commercial")
	
	assert.NoError(t, err)
	assert.Equal(t, 1, result)
}