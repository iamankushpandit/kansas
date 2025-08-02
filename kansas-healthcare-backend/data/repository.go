package data

import "kansas-healthcare-api/models"

type Repository interface {
	GetProviders() ([]models.Provider, error)
	GetProviderNetworks() ([]models.ProviderNetwork, error)
	GetProviderServiceLocations() ([]models.ProviderServiceLocation, error)
	GetCountyStats() ([]models.CountyStats, error)
	GetCountyStatsByName(county string) (*models.CountyStats, error)
	GetFilteredProviders(filter models.FilterRequest) ([]models.Provider, error)
	GetActiveProviderCount() (int, error)
	GetTerminatedNetworkCount(networkId string) (int, error)
	GetTerminatedServiceLocationCount(networkId string) (int, error)
	GetProvidersInCounty(county string) ([]models.Provider, error)
	GetRadiusAnalysis(county string, radius int, networkId string) (map[string]interface{}, error)
}
