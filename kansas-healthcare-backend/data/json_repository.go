package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kansas-healthcare-api/config"
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
	countyAreas              []models.CountyArea
	specialtyDensityStandards map[string]float64
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
	r.loadCountyAreas()
	r.loadSpecialtyDensityStandards()
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

func (r *JSONRepository) loadCountyAreas() {
	file, err := ioutil.ReadFile("data/county_areas.json")
	if err != nil {
		log.Fatal("Required file data/county_areas.json not found:", err)
	}
	if err := json.Unmarshal(file, &r.countyAreas); err != nil {
		log.Fatal("Error parsing county_areas.json:", err)
	}
}

func (r *JSONRepository) loadSpecialtyDensityStandards() {
	file, err := ioutil.ReadFile("data/specialty_density_standards.json")
	if err != nil {
		log.Fatal("Required file data/specialty_density_standards.json not found:", err)
	}
	if err := json.Unmarshal(file, &r.specialtyDensityStandards); err != nil {
		log.Fatal("Error parsing specialty_density_standards.json:", err)
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

// haversineDistance calculates the distance between two points on Earth using the Haversine formula
func (r *JSONRepository) haversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadiusMiles = 3959.0
	
	// Convert degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	lng1Rad := lng1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lng2Rad := lng2 * math.Pi / 180
	
	// Haversine formula
	dlat := lat2Rad - lat1Rad
	dlng := lng2Rad - lng1Rad
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusMiles * c
}

// calculateActualProviderDistances calculates average distance between providers using their actual locations
func (r *JSONRepository) calculateActualProviderDistances(county string) float64 {
	// Get active providers in county with locations
	var providerLocations []models.ProviderServiceLocation
	for _, provider := range r.providers {
		if provider.County == county && provider.Status == "Active" {
			// Find active service location for this provider
			for _, location := range r.providerServiceLocations {
				if location.ProviderID == provider.ProviderID && 
				   location.County == county && 
				   location.TerminationDate.Year() == 9999 {
					providerLocations = append(providerLocations, location)
					break
				}
			}
		}
	}
	
	if len(providerLocations) < 2 {
		return 0
	}
	
	// Calculate average distance to nearest provider for each provider
	var totalDistance float64
	for i, p1 := range providerLocations {
		minDistance := math.MaxFloat64
		for j, p2 := range providerLocations {
			if i != j {
				distance := r.haversineDistance(p1.Latitude, p1.Longitude, p2.Latitude, p2.Longitude)
				if distance < minDistance {
					minDistance = distance
				}
			}
		}
		totalDistance += minDistance
	}
	
	return totalDistance / float64(len(providerLocations))
}

func (r *JSONRepository) calculateProviderDensityMiles(providerCount int, county string) string {
	if providerCount == 0 {
		return "No providers"
	}
	
	// Try to calculate actual distances first
	actualDistance := r.calculateActualProviderDistances(county)
	if actualDistance > 0 {
		return fmt.Sprintf("~%.1f mi apart", actualDistance)
	}
	
	// Fallback to area-based calculation
	countyAreaSqMiles := r.getCountyArea(county)
	providersPerSqMile := float64(providerCount) / countyAreaSqMiles
	
	if providersPerSqMile >= 1 {
		return fmt.Sprintf("%.1f/sq mi", providersPerSqMile)
	} else {
		avgDistanceMiles := math.Sqrt(countyAreaSqMiles / float64(providerCount))
		return fmt.Sprintf("~%.1f mi apart", avgDistanceMiles)
	}
}

func (r *JSONRepository) getCountyArea(county string) float64 {
	for _, area := range r.countyAreas {
		if area.County == county {
			return area.AreaSqMiles
		}
	}
	// Fallback to average if county not found
	return 700.0
}

func (r *JSONRepository) GetCountyArea(county string) float64 {
	return r.getCountyArea(county)
}

func (r *JSONRepository) GetSpecialtyDensityStandards() map[string]float64 {
	return r.specialtyDensityStandards
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
			DensityMiles:   r.calculateProviderDensityMiles(providerCount, claims.County),
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
				DensityMiles:   r.calculateProviderDensityMiles(providerCount, county),
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

// GetCountyTerminatedNetworkCount implements the specific algorithm for county-based terminated network analysis
func (r *JSONRepository) GetCountyTerminatedNetworkCount(county, networkId string) (int, int, error) {
	fiveYearsAgo, twoYearsAgo := config.GetTerminatedAnalysisTimeRange()
	
	// Step 1: Get active providers in county with active service locations
	activeProviderIds := make(map[string]bool)
	for _, provider := range r.providers {
		if provider.Status == "Active" {
			// Check if provider has active service location in county
			for _, location := range r.providerServiceLocations {
				if location.ProviderID == provider.ProviderID &&
					location.County == county &&
					location.TerminationDate.Year() == 9999 { // Active service location
					activeProviderIds[provider.ProviderID] = true
					break
				}
			}
		}
	}
	
	totProvbyCounty := len(activeProviderIds)
	
	// Step 2: Check terminated network providers
	termNetworkCount := 0
	for providerId := range activeProviderIds {
		found := false
		for _, network := range r.providerNetwork {
			if network.ProviderID == providerId && network.NetworkID == networkId {
				found = true
				// Check if terminated in the specified timeframe
				if network.TerminationReason == config.LeftNetworkReason &&
					network.TerminationDate.After(fiveYearsAgo) &&
					network.TerminationDate.Before(twoYearsAgo) {
					termNetworkCount++
				}
				break
			}
		}
		// If no row found in Provider Network Table for the provider ID for that specific Network
		if !found {
			termNetworkCount++
		}
	}
	
	return totProvbyCounty, termNetworkCount, nil
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


