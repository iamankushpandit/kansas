package services

import "kansas-healthcare-api/models"

type AnalyticsServiceInterface interface {
	GetAllCountyData() ([]models.CountyStats, error)
	GetCountyData(county string) (*models.CountyStats, error)
	GetRecommendations(county string) []models.Recommendation
	GetActiveProviderCount() (int, error)
	GetTerminatedNetworkAnalysis(networkId string) (*models.TerminatedAnalysisResult, error)
	GetCountyTerminatedNetworkAnalysis(county, networkId string) (*models.TerminatedAnalysisResult, error)
	GetSpecialtyDensityAnalysis(county string) (map[string]interface{}, error)
	GetRadiusAnalysis(county string, radius int, networkId string) (map[string]interface{}, error)
}

type ProviderServiceInterface interface {
	GetAllProviders() ([]models.Provider, error)
	GetProviderNetworks() ([]models.ProviderNetwork, error)
	GetFilteredProviders(filter models.FilterRequest) ([]models.Provider, error)
}