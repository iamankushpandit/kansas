package models

import "time"

type Provider struct {
	ProviderID   string `json:"provider_id"`
	NPI          string `json:"npi"`
	ProviderType string `json:"provider_type"`
	Status       string `json:"status"`
	County       string `json:"county"`
}

type ProviderNetwork struct {
	ProviderID        string    `json:"provider_id"`
	NetworkID         string    `json:"network_id"`
	EffectiveDate     time.Time `json:"effective_date"`
	TerminationDate   time.Time `json:"termination_date"`
	TerminationReason string    `json:"termination_reason"`
}

type ProviderServiceLocation struct {
	ProviderID      string    `json:"provider_id"`
	EffectiveDate   time.Time `json:"effective_date"`
	TerminationDate time.Time `json:"termination_date"`
	Address1        string    `json:"address1"`
	Address2        string    `json:"address2"`
	City            string    `json:"city"`
	ZipCode         string    `json:"zip_code"`
	County          string    `json:"county"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
}

type CountyStats struct {
	County         string  `json:"county"`
	ProviderCount  int     `json:"provider_count"`
	ClaimsCount    int     `json:"claims_count"`
	AvgClaimAmount float64 `json:"avg_claim_amount"`
	Density        string  `json:"density"`
	DensityMiles   string  `json:"density_miles"`
}

type CountyClaims struct {
	County         string  `json:"county"`
	ClaimsCount    int     `json:"claims_count"`
	AvgClaimAmount float64 `json:"avg_claim_amount"`
}

type Recommendation struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	County      string `json:"county"`
	Icon        string `json:"icon"`
}

type FilterRequest struct {
	Specialty string `json:"specialty"`
	Metric    string `json:"metric"`
	Radius    int    `json:"radius"`
	Network   string `json:"network"`
}

type TerminatedAnalysisResult struct {
	TermNetworkCount     int     `json:"term_network_count"`
	ServiceLocationCount int     `json:"service_location_count"`
	PercentageTerminated float64 `json:"percentage_terminated"`
	TotalActiveProviders int     `json:"total_active_providers"`
}
