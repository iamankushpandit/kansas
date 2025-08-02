package config

import "time"

// Business constants for terminated network analysis
const (
	// Time range for terminated provider analysis (2-5 years)
	TerminatedAnalysisMinYears = 2
	TerminatedAnalysisMaxYears = 5
	
	// Active service location indicator
	ActiveServiceLocationDate = "12/31/9999"
	
	// Network termination reason
	LeftNetworkReason = "Left Network"
)

// GetTerminatedAnalysisTimeRange returns the time range for terminated analysis
func GetTerminatedAnalysisTimeRange() (time.Time, time.Time) {
	now := time.Now()
	return now.AddDate(-TerminatedAnalysisMaxYears, 0, 0), now.AddDate(-TerminatedAnalysisMinYears, 0, 0)
}