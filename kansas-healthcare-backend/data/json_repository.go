package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kansas-healthcare-api/models"
	"log"
	"math"
	"time"
)

type JSONRepository struct {
	providers                []models.Provider
	providerNetwork          []models.ProviderNetwork
	providerServiceLocations []models.ProviderServiceLocation
	countyClaims             []models.CountyClaims
}

func NewJSONRepository() *JSONRepository {
	repo := &JSONRepository{}
	repo.loadData()
	return repo
}

func (r *JSONRepository) loadData() {
	r.loadProviders()
	r.loadProviderNetworks()
	r.loadProviderServiceLocations()
	r.loadCountyClaims()
}

func (r *JSONRepository) loadProviders() {
	file, err := ioutil.ReadFile("data/providers.json")
	if err != nil {
		log.Fatal("Required file data/providers.json not found:", err)
	}
	if err := json.Unmarshal(file, &r.providers); err != nil {
		log.Fatal("Error parsing providers.json:", err)
	}
}

func (r *JSONRepository) loadProviderNetworks() {
	file, err := ioutil.ReadFile("data/provider_networks.json")
	if err != nil {
		log.Fatal("Required file data/provider_networks.json not found:", err)
	}
	if err := json.Unmarshal(file, &r.providerNetwork); err != nil {
		log.Fatal("Error parsing provider_networks.json:", err)
	}
}

func (r *JSONRepository) loadProviderServiceLocations() {
	file, err := ioutil.ReadFile("data/provider_service_locations.json")
	if err != nil {
		log.Fatal("Required file data/provider_service_locations.json not found:", err)
	}
	if err := json.Unmarshal(file, &r.providerServiceLocations); err != nil {
		log.Fatal("Error parsing provider_service_locations.json:", err)
	}
}

func (r *JSONRepository) loadCountyClaims() {
	file, err := ioutil.ReadFile("data/claims.json")
	if err != nil {
		log.Fatal("Required file data/claims.json not found:", err)
	}
	if err := json.Unmarshal(file, &r.countyClaims); err != nil {
		log.Fatal("Error parsing claims.json:", err)
	}
}

func (r *JSONRepository) GetProviders() ([]models.Provider, error) {
	return r.providers, nil
}

func (r *JSONRepository) GetProviderNetworks() ([]models.ProviderNetwork, error) {
	return r.providerNetwork, nil
}

func (r *JSONRepository) GetProviderServiceLocations() ([]models.ProviderServiceLocation, error) {
	return r.providerServiceLocations, nil
}

func (r *JSONRepository) calculateProviderDensity(providerCount int) string {
	if providerCount >= 200 {
		return "high"
	} else if providerCount >= 50 {
		return "medium"
	} else if providerCount >= 30 {
		return "low"
	}
	return "critical"
}

func (r *JSONRepository) calculateProviderDensityMiles(providerCount int) string {
	if providerCount == 0 {
		return "No providers"
	}
	
	// Kansas counties average ~700 square miles
	avgCountyAreaSqMiles := 700.0
	
	// Calculate providers per square mile
	providersPerSqMile := float64(providerCount) / avgCountyAreaSqMiles
	
	if providersPerSqMile >= 1 {
		return fmt.Sprintf("%.1f/sq mi", providersPerSqMile)
	} else {
		// For sparse areas, show average distance between providers
		avgDistanceMiles := math.Sqrt(avgCountyAreaSqMiles / float64(providerCount))
		return fmt.Sprintf("~%.1f mi apart", avgDistanceMiles)
	}
}

func (r *JSONRepository) GetCountyStats() ([]models.CountyStats, error) {
	var countyStats []models.CountyStats
	
	// Create a map to count active providers by county
	providerCounts := make(map[string]int)
	for _, provider := range r.providers {
		if provider.Status == "Active" {
			providerCounts[provider.County]++
		}
	}
	
	// Combine with claims data
	for _, claims := range r.countyClaims {
		providerCount := providerCounts[claims.County]
		density := r.calculateProviderDensity(providerCount)
		
		countyStats = append(countyStats, models.CountyStats{
			County:         claims.County,
			ProviderCount:  providerCount,
			ClaimsCount:    claims.ClaimsCount,
			AvgClaimAmount: claims.AvgClaimAmount,
			Density:        density,
			DensityMiles:   r.calculateProviderDensityMiles(providerCount),
		})
	}
	
	return countyStats, nil
}

func (r *JSONRepository) GetCountyStatsByName(county string) (*models.CountyStats, error) {
	// Count active providers in the county
	providerCount := 0
	for _, provider := range r.providers {
		if provider.County == county && provider.Status == "Active" {
			providerCount++
		}
	}
	
	// Find claims data for the county
	for _, claims := range r.countyClaims {
		if claims.County == county {
			density := r.calculateProviderDensity(providerCount)
			return &models.CountyStats{
				County:         county,
				ProviderCount:  providerCount,
				ClaimsCount:    claims.ClaimsCount,
				AvgClaimAmount: claims.AvgClaimAmount,
				Density:        density,
				DensityMiles:   r.calculateProviderDensityMiles(providerCount),
			}, nil
		}
	}
	
	return nil, nil
}

func (r *JSONRepository) GetProvidersInCounty(county string) ([]models.Provider, error) {
	var countyProviders []models.Provider
	for _, provider := range r.providers {
		if provider.County == county {
			countyProviders = append(countyProviders, provider)
		}
	}
	return countyProviders, nil
}

func (r *JSONRepository) GetFilteredProviders(filter models.FilterRequest) ([]models.Provider, error) {
	// Get active providers in the specified network
	activeNetworkProviders := make(map[string]bool)
	for _, network := range r.providerNetwork {
		if network.NetworkID == filter.Network && network.TerminationReason == "" {
			activeNetworkProviders[network.ProviderID] = true
		}
	}

	var filtered []models.Provider
	for _, provider := range r.providers {
		// Only include active providers
		if provider.Status != "Active" {
			continue
		}
		
		// Filter by network
		if !activeNetworkProviders[provider.ProviderID] {
			continue
		}
		
		// Filter by specialty (skip if "All")
		if filter.Specialty != "All" && provider.ProviderType != filter.Specialty {
			continue
		}
		
		filtered = append(filtered, provider)
	}
	return filtered, nil
}

func (r *JSONRepository) GetActiveProviderCount() (int, error) {
	count := 0
	for _, provider := range r.providers {
		if provider.Status == "Active" {
			count++
		}
	}
	return count, nil
}

func (r *JSONRepository) GetTerminatedNetworkCount(networkId string) (int, error) {
	now := time.Now()
	twoYearsAgo := now.AddDate(-2, 0, 0)
	fiveYearsAgo := now.AddDate(-5, 0, 0)

	count := 0
	for _, network := range r.providerNetwork {
		if network.NetworkID == networkId &&
			network.TerminationReason == "Left Network" &&
			network.TerminationDate.After(fiveYearsAgo) &&
			network.TerminationDate.Before(twoYearsAgo) {
			count++
		}
	}
	return count, nil
}

func (r *JSONRepository) GetTerminatedServiceLocationCount(networkId string) (int, error) {
	now := time.Now()
	twoYearsAgo := now.AddDate(-2, 0, 0)
	fiveYearsAgo := now.AddDate(-5, 0, 0)

	// Get provider IDs that terminated from the network
	terminatedProviderIds := make(map[string]bool)
	for _, network := range r.providerNetwork {
		if network.NetworkID == networkId &&
			network.TerminationReason == "Left Network" &&
			network.TerminationDate.After(fiveYearsAgo) &&
			network.TerminationDate.Before(twoYearsAgo) {
			terminatedProviderIds[network.ProviderID] = true
		}
	}

	// Count service locations for those providers
	count := 0
	for _, location := range r.providerServiceLocations {
		if terminatedProviderIds[location.ProviderID] &&
			location.TerminationDate.After(fiveYearsAgo) &&
			location.TerminationDate.Before(twoYearsAgo) {
			count++
		}
	}
	return count, nil
}

func (r *JSONRepository) GetRadiusAnalysis(county string, radius int, networkId string) (map[string]interface{}, error) {
	// Get active providers in the network for the county
	activeNetworkProviders := make(map[string]bool)
	for _, network := range r.providerNetwork {
		if network.NetworkID == networkId && network.TerminationReason == "" {
			activeNetworkProviders[network.ProviderID] = true
		}
	}

	// Count providers in the county
	countyProviders := 0
	specialtyCount := make(map[string]int)
	for _, provider := range r.providers {
		if provider.County == county && provider.Status == "Active" && activeNetworkProviders[provider.ProviderID] {
			countyProviders++
			specialtyCount[provider.ProviderType]++
		}
	}

	// Get claims data for the county
	var claimsData *models.CountyClaims
	for _, claims := range r.countyClaims {
		if claims.County == county {
			claimsData = &claims
			break
		}
	}

	result := map[string]interface{}{
		"county":           county,
		"radius":          radius,
		"network":         networkId,
		"provider_count":  countyProviders,
		"specialty_count": len(specialtyCount),
		"specialties":     specialtyCount,
	}

	if claimsData != nil {
		result["claims_count"] = claimsData.ClaimsCount
		result["avg_claim_amount"] = claimsData.AvgClaimAmount
	}

	return result, nil
}


