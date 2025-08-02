// +build ignore

package main

import (
	"fmt"
	"kansas-healthcare-api/data"
)

func main() {
	fmt.Println("Testing Terminated Network Analysis...")
	
	repo := data.NewJSONRepository()
	
	// Test the new county terminated network analysis
	totProviders, termCount, err := repo.GetCountyTerminatedNetworkCount("Sedgwick", "Commercial")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("County: Sedgwick\n")
	fmt.Printf("Total Active Providers: %d\n", totProviders)
	fmt.Printf("Terminated Network Count: %d\n", termCount)
	
	// Calculate percentage as per your formula: (TotProvbyCounty - (TotProvbyCounty - TermNetworkCount))/100
	// This simplifies to: TermNetworkCount/100
	percentage := float64(termCount) / 100.0
	fmt.Printf("Percentage Terminated: %.2f%%\n", percentage)
}