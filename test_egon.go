package main

import (
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	fmt.Println("Testing Egon character...")

	// Load the egon character
	agent, err := characters.LibraryAgent("egon")
	if err != nil {
		fmt.Printf("Error loading egon: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n=== Egon Base Character ===")
	agent.ShowBase(os.Stdout)

	fmt.Println("\n=== Egon Plan State ===")
	agent.Plan(os.Stdout)

	fmt.Println("\n=== Egon Think State ===")
	agent.Think(os.Stdout)

	fmt.Println("\n=== Egon Execute State ===")
	agent.Execute(os.Stdout)

	fmt.Println("\n=== Egon State List ===")
	states := agent.ListStates()
	for _, state := range states {
		fmt.Printf("- %s\n", state)
	}

	fmt.Println("\n✅ Egon character test completed successfully!")
}
