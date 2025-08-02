package services

import (
	"kansas-healthcare-api/data"
	"kansas-healthcare-api/models"
)

type ProviderService struct {
	repo data.Repository
}

func NewProviderService(repo data.Repository) *ProviderService {
	return &ProviderService{repo: repo}
}

func (s *ProviderService) GetAllProviders() ([]models.Provider, error) {
	return s.repo.GetProviders()
}

func (s *ProviderService) GetProviderNetworks() ([]models.ProviderNetwork, error) {
	return s.repo.GetProviderNetworks()
}

func (s *ProviderService) GetFilteredProviders(filter models.FilterRequest) ([]models.Provider, error) {
	return s.repo.GetFilteredProviders(filter)
}
