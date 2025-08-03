package services

import (
	"fmt"
	"kansas-healthcare-api/data"
	"kansas-healthcare-api/models"
)

type AnalyticsService struct {
	repo data.Repository
}

func NewAnalyticsService(repo data.Repository) *AnalyticsService {
	return &AnalyticsService{repo: repo}
}

func (s *AnalyticsService) GetAllCountyData() ([]models.CountyStats, error) {
	return s.repo.GetCountyStats()
}

func (s *AnalyticsService) GetCountyData(county string) (*models.CountyStats, error) {
	return s.repo.GetCountyStatsByName(county)
}

func (s *AnalyticsService) GetRecommendations(county string) []models.Recommendation {
	// Generate dynamic recommendations based on county data
	recommendations := s.generateDynamicRecommendations(county)
	
	// Set county for all recommendations
	for i := range recommendations {
		recommendations[i].County = county
	}
	
	return recommendations
}

func (s *AnalyticsService) generateDynamicRecommendations(county string) []models.Recommendation {
	var recommendations []models.Recommendation
	idCounter := 1
	
	// Get county stats
	countyStats, err := s.repo.GetCountyStatsByName(county)
	if err != nil || countyStats == nil {
		return recommendations
	}
	
	// Get providers in county for detailed analysis
	providers, _ := s.repo.GetProvidersInCounty(county)
	terminatedCount := 0
	specialtyMap := make(map[string]int)
	
	for _, provider := range providers {
		if provider.Status == "Terminated" {
			terminatedCount++
		}
		specialtyMap[provider.ProviderType]++
	}
	
	claimsPerProvider := float64(countyStats.ClaimsCount) / float64(countyStats.ProviderCount)
	
	// 1. Critical provider shortage
	if countyStats.ProviderCount < 15 {
		recommendations = append(recommendations, models.Recommendation{
			ID:          idCounter,
			Type:        "EXPAND_NETWORK",
			Title:       "Critical Provider Shortage",
			Description: fmt.Sprintf("Only %d providers for %d claims - urgent expansion needed", countyStats.ProviderCount, countyStats.ClaimsCount),
			Priority:    "High",
			Icon:        "mdi-alert-circle",
		})
		idCounter++
	}
	
	// 2. High workload per provider
	if claimsPerProvider > 25 {
		recommendations = append(recommendations, models.Recommendation{
			ID:          idCounter,
			Type:        "EXPAND_NETWORK",
			Title:       "High Provider Workload",
			Description: fmt.Sprintf("%.0f claims per provider - consider network expansion", claimsPerProvider),
			Priority:    "Medium",
			Icon:        "mdi-chart-line",
		})
		idCounter++
	}
	
	// 3. High cost claims
	if countyStats.AvgClaimAmount > 1000 {
		recommendations = append(recommendations, models.Recommendation{
			ID:          idCounter,
			Type:        "COST_MANAGEMENT",
			Title:       "High Cost Claims",
			Description: fmt.Sprintf("Average claim $%.2f - review cost management strategies", countyStats.AvgClaimAmount),
			Priority:    "Medium",
			Icon:        "mdi-currency-usd",
		})
		idCounter++
	}
	
	// 4. Provider terminations
	if terminatedCount > 0 {
		recommendations = append(recommendations, models.Recommendation{
			ID:          idCounter,
			Type:        "CONTACT_FORMER",
			Title:       "Contact Former Providers",
			Description: fmt.Sprintf("%d providers left network - consider re-engagement", terminatedCount),
			Priority:    "Medium",
			Icon:        "mdi-phone",
		})
		idCounter++
	}
	
	// 5. Specialty gaps (if low specialty diversity)
	if len(specialtyMap) < 5 && countyStats.ProviderCount > 10 {
		recommendations = append(recommendations, models.Recommendation{
			ID:          idCounter,
			Type:        "EXPAND_SPECIALTIES",
			Title:       "Limited Specialty Coverage",
			Description: fmt.Sprintf("Only %d specialties available - expand specialty network", len(specialtyMap)),
			Priority:    "Medium",
			Icon:        "mdi-medical-bag",
		})
		idCounter++
	}
	
	// 6. Target out-of-network providers (for medium-high claim areas)
	if countyStats.ClaimsCount > 500 && countyStats.ProviderCount < 50 {
		potentialProviders := int(float64(countyStats.ClaimsCount) / 100) // Estimate
		recommendations = append(recommendations, models.Recommendation{
			ID:          idCounter,
			Type:        "TARGET_OON",
			Title:       "Target Out-of-Network Providers",
			Description: fmt.Sprintf("%d potential providers serving this area", potentialProviders),
			Priority:    "Medium",
			Icon:        "mdi-target",
		})
		idCounter++
	}
	
	// 7. Network optimization for high-density areas
	if countyStats.ProviderCount > 100 && claimsPerProvider < 20 {
		recommendations = append(recommendations, models.Recommendation{
			ID:          idCounter,
			Type:        "OPTIMIZE_NETWORK",
			Title:       "Network Optimization Opportunity",
			Description: "Low utilization per provider - consider network optimization",
			Priority:    "Low",
			Icon:        "mdi-tune",
		})
		idCounter++
	}
	
	return recommendations
}

func (s *AnalyticsService) GetActiveProviderCount() (int, error) {
	return s.repo.GetActiveProviderCount()
}

func (s *AnalyticsService) GetTerminatedNetworkAnalysis(networkId string) (*models.TerminatedAnalysisResult, error) {
	totalActive, err := s.repo.GetActiveProviderCount()
	if err != nil {
		return nil, err
	}

	termCount, err := s.repo.GetTerminatedNetworkCount(networkId)
	if err != nil {
		return nil, err
	}

	locationCount, err := s.repo.GetTerminatedServiceLocationCount(networkId)
	if err != nil {
		return nil, err
	}

	percentage := (float64(termCount) / float64(totalActive)) * 100

	return &models.TerminatedAnalysisResult{
		TermNetworkCount:     termCount,
		ServiceLocationCount: locationCount,
		PercentageTerminated: percentage,
		TotalActiveProviders: totalActive,
	}, nil
}

func (s *AnalyticsService) GetCountyTerminatedNetworkAnalysis(county, networkId string) (*models.TerminatedAnalysisResult, error) {
	totProvbyCounty, termNetworkCount, err := s.repo.GetCountyTerminatedNetworkCount(county, networkId)
	if err != nil {
		return nil, err
	}

	// Calculate percentage: (TotProvbyCounty - (TotProvbyCounty - TermNetworkCount))/100
	// This simplifies to: TermNetworkCount/100
	percentage := float64(termNetworkCount) / 100.0

	return &models.TerminatedAnalysisResult{
		TermNetworkCount:     termNetworkCount,
		ServiceLocationCount: 0, // Not used in county analysis
		PercentageTerminated: percentage,
		TotalActiveProviders: totProvbyCounty,
	}, nil
}

func (s *AnalyticsService) GetSpecialtyDensityAnalysis(county string) (map[string]interface{}, error) {
	providers, err := s.repo.GetProvidersInCounty(county)
	if err != nil {
		return nil, err
	}

	// Count providers by specialty
	specialtyCounts := make(map[string]int)
	for _, provider := range providers {
		if provider.Status == "Active" {
			specialtyCounts[provider.ProviderType]++
		}
	}

	// Get county area for density calculation
	countyArea := s.repo.GetCountyArea(county)

	// Get specialty density standards
	standards := s.repo.GetSpecialtyDensityStandards()

	// Convert to sorted list by gap (actual vs recommended)
	type SpecialtyDensity struct {
		Name  string  `json:"name"`
		Count int     `json:"count"`
		Gap   float64 `json:"gap"`
	}

	var densities []SpecialtyDensity
	for specialty, recommendedDensity := range standards {
		actualCount := specialtyCounts[specialty]
		actualDensity := float64(actualCount) / countyArea
		gap := recommendedDensity - actualDensity
		
		densities = append(densities, SpecialtyDensity{
			Name:  specialty,
			Count: actualCount,
			Gap:   gap,
		})
	}

	// Sort by gap (descending - highest gap first)
	for i := 0; i < len(densities)-1; i++ {
		for j := i + 1; j < len(densities); j++ {
			if densities[i].Gap < densities[j].Gap {
				densities[i], densities[j] = densities[j], densities[i]
			}
		}
	}

	return map[string]interface{}{
		"specialty_densities": densities,
	}, nil
}

func (s *AnalyticsService) GetRadiusAnalysis(county string, radius int, networkId string) (map[string]interface{}, error) {
	return s.repo.GetRadiusAnalysis(county, radius, networkId)
}
