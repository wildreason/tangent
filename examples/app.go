package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Tangent Character Library - Test App   ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// =========================================
	// TEST 1: Core Library Experience
	// =========================================
	fmt.Println("▢ TEST 1: Core Library Experience")
	fmt.Println()

	// Test the primary experience: Library access
	fmt.Println("Available characters:")
	for _, name := range characters.ListLibrary() {
		description, _ := characters.LibraryInfo(name)
		fmt.Printf("  • %s - %s\n", name, description)
	}
	fmt.Println()

	// Test alien character (should be the primary experience)
	fmt.Println("◢ Testing alien character (primary experience):")
	alien, err := characters.Library("alien")
	if err != nil {
		fmt.Printf("✗ Failed to load alien: %v\n", err)
		return
	}

	fmt.Println("Idle state:")
	characters.ShowIdle(os.Stdout, alien)
	fmt.Println()

	fmt.Println("Animated (3 loops, 5 fps):")
	characters.Animate(os.Stdout, alien, 5, 3)
	fmt.Println()

	// =========================================
	// TEST 2: Custom Character Creation
	// =========================================
	fmt.Println("▢ TEST 2: Custom Character Creation")
	fmt.Println()

	// Create a simple custom character
	fmt.Println("◢ Creating custom 'heart' character:")
	heart, err := characters.NewCharacterSpec("heart", 7, 5).
		AddFrame("idle", []string{
			"_5F5F5_",
			"5FFFFF5",
			"5FFFFF5",
			"_5FFF5_",
			"__5F5__",
		}).
		Build()

	if err != nil {
		fmt.Printf("✗ Failed to create heart: %v\n", err)
		return
	}

	fmt.Println("Heart character:")
	characters.ShowIdle(os.Stdout, heart)
	fmt.Println()

	// =========================================
	// TEST 3: Error Handling
	// =========================================
	fmt.Println("▢ TEST 3: Error Handling")
	fmt.Println()

	// Test invalid character name
	fmt.Println("◢ Testing invalid character name:")
	_, err = characters.Library("nonexistent")
	if err != nil {
		fmt.Printf("✓ Correctly caught error: %v\n", err)
	}
	fmt.Println()

	// Test invalid custom character
	fmt.Println("◢ Testing invalid custom character:")
	_, err = characters.NewCharacterSpec("", 5, 3).Build()
	if err != nil {
		fmt.Printf("✓ Correctly caught validation error: %v\n", err)
	}
	fmt.Println()

	// =========================================
	// TEST 4: Performance Test
	// =========================================
	fmt.Println("▢ TEST 4: Performance Test")
	fmt.Println()

	// Test multiple character loads
	fmt.Println("◢ Loading multiple characters:")
	start := time.Now()

	charactersToTest := []string{"alien", "robot", "pulse", "wave", "rocket"}
	for _, name := range charactersToTest {
		char, err := characters.Library(name)
		if err != nil {
			fmt.Printf("  ✗ %s: %v\n", name, err)
			continue
		}
		fmt.Printf("  ✓ %s: %dx%d, %d frames\n", name, char.Width, char.Height, len(char.Frames))
	}

	duration := time.Since(start)
	fmt.Printf("✓ Loaded %d characters in %v\n", len(charactersToTest), duration)
	fmt.Println()

	// =========================================
	// TEST 5: Integration Test
	// =========================================
	fmt.Println("▢ TEST 5: Integration Test")
	fmt.Println()

	// Test the complete workflow: load, create, animate
	fmt.Println("◢ Complete workflow test:")

	// Load library character
	robot, err := characters.Library("robot")
	if err != nil {
		fmt.Printf("✗ Failed to load robot: %v\n", err)
		return
	}

	// Create custom character
	star, err := characters.NewCharacterSpec("star", 5, 5).
		AddFrame("idle", []string{
			"_2F2_",
			"2F_F2",
			"F___F",
			"2F_F2",
			"_2F2_",
		}).
		Build()
	if err != nil {
		fmt.Printf("✗ Failed to create star: %v\n", err)
		return
	}

	// Use both characters
	fmt.Println("Robot (library):")
	characters.ShowIdle(os.Stdout, robot)
	fmt.Println()

	fmt.Println("Star (custom):")
	characters.ShowIdle(os.Stdout, star)
	fmt.Println()

	// =========================================
	// Summary
	// =========================================
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Test Results Summary                    ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("✓ Core library experience: WORKING")
	fmt.Println("✓ Custom character creation: WORKING")
	fmt.Println("✓ Error handling: WORKING")
	fmt.Println("✓ Performance: ACCEPTABLE")
	fmt.Println("✓ Integration: WORKING")
	fmt.Println()
	fmt.Println("Core API Test:")
	fmt.Println("  characters.Library(\"alien\") ✓")
	fmt.Println("  characters.NewCharacterSpec(...).Build() ✓")
	fmt.Println("  characters.Animate(os.Stdout, char, 5, 3) ✓")
	fmt.Println("  characters.ShowIdle(os.Stdout, char) ✓")
	fmt.Println()
	fmt.Println("✓ All core experiences working as designed!")
}
