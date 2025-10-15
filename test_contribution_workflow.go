package main

import (
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	fmt.Println("🔄 Testing Contribution Workflow")
	fmt.Println("================================")

	// Test 1: Load monsterra from JSON file
	fmt.Println("\n1️⃣ Loading Monsterra from JSON file:")
	monsterra, err := characters.LoadCharacter("monsterra", characters.CharacterSource{
		Type: characters.SourceJSON,
		Path: "monsterra.json",
	})
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✅ Loaded: %s (%s)\n", monsterra.Name(), monsterra.Personality())

	// Test 2: Use the character
	fmt.Println("\n2️⃣ Using Monsterra character:")
	fmt.Println("Base character:")
	monsterra.ShowBase(os.Stdout)

	fmt.Println("\nPlan state:")
	monsterra.Plan(os.Stdout)

	fmt.Println("\nThink state:")
	monsterra.Think(os.Stdout)

	fmt.Println("\nExecute state:")
	monsterra.Execute(os.Stdout)

	// Test 3: List all available characters
	fmt.Println("\n3️⃣ Available characters:")
	allChars := characters.ListCharacters()
	for i, name := range allChars {
		fmt.Printf("  %d. %s\n", i+1, name)
	}

	// Test 4: Test contribution validation
	fmt.Println("\n4️⃣ Validating contribution:")
	if err := characters.ValidateContribution("monsterra.json"); err != nil {
		fmt.Printf("❌ Validation failed: %v\n", err)
	} else {
		fmt.Println("✅ Contribution is valid")
	}

	// Test 5: Install contribution
	fmt.Println("\n5️⃣ Installing contribution:")
	if err := characters.InstallContribution("monsterra"); err != nil {
		fmt.Printf("❌ Installation failed: %v\n", err)
	} else {
		fmt.Println("✅ Contribution installed successfully")
	}

	// Test 6: Load installed character
	fmt.Println("\n6️⃣ Loading installed character:")
	installedMonsterra, err := characters.GetCharacter("monsterra")
	if err != nil {
		fmt.Printf("❌ Error loading installed character: %v\n", err)
	} else {
		fmt.Printf("✅ Loaded installed character: %s\n", installedMonsterra.Name())
	}

	// Test 7: List contributions
	fmt.Println("\n7️⃣ Available contributions:")
	contributions, err := characters.ListContributions()
	if err != nil {
		fmt.Printf("❌ Error listing contributions: %v\n", err)
	} else {
		for i, name := range contributions {
			fmt.Printf("  %d. %s\n", i+1, name)
		}
	}

	fmt.Println("\n✅ Contribution workflow test completed!")
}
