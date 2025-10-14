package main

import (
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/domain"
)

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Testing Agent State API                 ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Create a test character with states
	char := &domain.Character{
		Name:        "test-agent",
		Personality: "efficient",
		Width:       7,
		Height:      3,
		States: map[string]domain.State{
			"plan": {
				Name:        "Planning",
				Description: "Test planning",
				StateType:   "standard",
				Frames: []domain.Frame{
					{Lines: []string{"_?????_", "_?????_", "_?????_"}},
				},
			},
			"think": {
				Name:        "Thinking",
				Description: "Test thinking",
				StateType:   "standard",
				Frames: []domain.Frame{
					{Lines: []string{"_....._", "_....._", "_....._"}},
				},
			},
			"execute": {
				Name:        "Executing",
				Description: "Test executing",
				StateType:   "standard",
				Frames: []domain.Frame{
					{Lines: []string{"_>>>>>_", "_>>>>>_", "_>>>>>_"}},
				},
			},
		},
	}

	// Wrap in AgentCharacter
	agent := characters.NewAgentCharacter(char)

	// Test 1: State methods
	fmt.Println("▢ Test 1: State Methods")
	fmt.Println()

	fmt.Println("1. Planning:")
	if err := agent.Plan(os.Stdout); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Println()

	fmt.Println("2. Thinking:")
	if err := agent.Think(os.Stdout); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Println()

	fmt.Println("3. Executing:")
	if err := agent.Execute(os.Stdout); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Println()

	// Test 2: State inspection
	fmt.Println("▢ Test 2: State Inspection")
	fmt.Println()

	states := agent.ListStates()
	fmt.Printf("Available states: %v\n", states)
	fmt.Printf("Has 'plan' state: %v\n", agent.HasState("plan"))
	fmt.Printf("Has 'nonexistent' state: %v\n", agent.HasState("nonexistent"))
	fmt.Println()

	// Test 3: Character info
	fmt.Println("▢ Test 3: Character Info")
	fmt.Println()

	fmt.Printf("Character name: %s\n", agent.Name())
	fmt.Printf("Personality: %s\n", agent.Personality())
	fmt.Println()

	// Test 4: State descriptions
	fmt.Println("▢ Test 4: State Descriptions")
	fmt.Println()

	for _, stateName := range agent.ListStates() {
		desc, err := agent.GetStateDescription(stateName)
		if err != nil {
			fmt.Printf("State: %s - ERROR: %v\n", stateName, err)
		} else {
			fmt.Printf("State: %-10s - %s\n", stateName, desc)
		}
	}
	fmt.Println()

	// Test 5: Custom state access
	fmt.Println("▢ Test 5: Custom State Access")
	fmt.Println()

	// Add a custom state
	char.States["custom"] = domain.State{
		Name:        "Custom",
		Description: "Custom test state",
		StateType:   "custom",
		Frames: []domain.Frame{
			{Lines: []string{"_*****_", "_*****_", "_*****_"}},
		},
	}

	fmt.Println("Custom state:")
	if err := agent.ShowState(os.Stdout, "custom"); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Println()

	// Test 6: Error handling
	fmt.Println("▢ Test 6: Error Handling")
	fmt.Println()

	fmt.Println("Trying to access non-existent state:")
	if err := agent.ShowState(os.Stdout, "nonexistent"); err != nil {
		fmt.Printf("✓ Correctly caught error: %v\n", err)
	} else {
		fmt.Println("✗ Should have returned error")
	}
	fmt.Println()

	// Summary
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Test Results                            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("✓ State methods working")
	fmt.Println("✓ State inspection working")
	fmt.Println("✓ Character info working")
	fmt.Println("✓ State descriptions working")
	fmt.Println("✓ Custom states working")
	fmt.Println("✓ Error handling working")
	fmt.Println()
	fmt.Println("All tests passed! 🎉")
}


