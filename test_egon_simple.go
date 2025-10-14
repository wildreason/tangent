package main

import (
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	fmt.Println("Testing Egon character (simple)...")

	// Test basic library access
	availableChars := characters.ListLibrary()
	fmt.Printf("Available characters: %v\n", availableChars)

	// Check if egon is in the list
	found := false
	for _, char := range availableChars {
		if char == "egon" {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("❌ Egon not found in library")
		os.Exit(1)
	}

	fmt.Println("✅ Egon found in library")

	// Test library info
	info, err := characters.LibraryInfo("egon")
	if err != nil {
		fmt.Printf("Error getting egon info: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Egon info: %s\n", info)

	// Test loading the character
	agent, err := characters.LibraryAgent("egon")
	if err != nil {
		fmt.Printf("Error loading egon: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Egon loaded successfully")

	// Get the underlying character
	domainChar := agent.GetCharacter()
	fmt.Printf("Character name: %s\n", domainChar.Name)
	fmt.Printf("Character width: %d\n", domainChar.Width)
	fmt.Printf("Character height: %d\n", domainChar.Height)
	fmt.Printf("Number of frames: %d\n", len(domainChar.Frames))

	// Show first frame
	if len(domainChar.Frames) > 0 {
		fmt.Println("\nFirst frame:")
		for _, line := range domainChar.Frames[0].Lines {
			fmt.Println(line)
		}
	}

	fmt.Println("\n✅ Egon character test completed successfully!")
}
