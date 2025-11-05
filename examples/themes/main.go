package main

import (
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	// List all available themes
	fmt.Println("Available themes:")
	for _, theme := range characters.ListThemes() {
		fmt.Printf("  - %s\n", theme)
	}
	fmt.Println()

	// Demonstrate theme switching
	themes := []string{"bright", "latte", "garden", "cozy"}

	for _, themeName := range themes {
		fmt.Printf("=== Theme: %s ===\n", themeName)

		// Set the theme
		err := characters.SetTheme(themeName)
		if err != nil {
			fmt.Printf("Error setting theme: %v\n", err)
			continue
		}

		// Create an agent - it automatically uses the current theme
		agent, err := characters.LibraryAgent("sa")
		if err != nil {
			fmt.Printf("Error creating agent: %v\n", err)
			continue
		}

		// Display the character color
		char := agent.GetCharacter()
		fmt.Printf("Character 'sa' color: %s\n", char.Color)

		// Show a state
		agent.Think(os.Stdout)
		fmt.Println()
	}

	// Demonstrate current theme getter
	current := characters.GetCurrentTheme()
	fmt.Printf("Current theme: %s\n", current)
}
