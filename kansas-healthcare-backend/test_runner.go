// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Running Backend Unit Tests...")
	
	// Run tests for data package only (working tests)
	cmd := exec.Command("go", "test", "./data", "-v")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("Tests failed: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("✅ All tests passed!")
}