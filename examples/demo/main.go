package main

import (
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
	"github.com/wildreason/tangent/pkg/characters/service"
)

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Wildreason Characters - Demo           ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// =========================================
	// PART 1: Using Library Characters
	// =========================================
	fmt.Println("▢ PART 1: Library Characters")
	fmt.Println()

	// List all available library characters
	fmt.Println("Available in library:")
	for _, name := range characters.ListLibrary() {
		description, _ := characters.LibraryInfo(name)
		fmt.Printf("  • %s - %s\n", name, description)
	}
	fmt.Println()

	// Load and display alien character
	alien, err := characters.Library("alien")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("◢ Alien character (idle state):")
	characters.ShowIdle(os.Stdout, alien)
	fmt.Println()

	fmt.Println("◢ Alien character (animated):")
	fmt.Println("Press Ctrl+C to stop...")
	characters.Animate(os.Stdout, alien, 4, 2)
	fmt.Println()

	// =========================================
	// PART 2: New Architecture (Service Layer)
	// =========================================
	fmt.Println("▢ PART 2: New Architecture (Service Layer)")
	fmt.Println()
	fmt.Println("Using the new layered architecture with proper separation of concerns!")
	fmt.Println()

	// Create service with infrastructure
	tempDir := "/tmp/tangent-demo"
	os.MkdirAll(tempDir, 0755)
	repo := infrastructure.NewFileCharacterRepository(tempDir)
	compiler := infrastructure.NewPatternCompiler()
	animationEngine := infrastructure.NewAnimationEngine()
	service := service.NewCharacterService(repo, compiler, animationEngine)

	// Create character using domain specification
	spec := domain.NewCharacterSpec("demo-robot", 9, 4).
		AddFrame("idle", []string{
			"_L5FFF5R_",
			"_6FFFFF6_",
			"__FF_FF__",
			"__FF_FF__",
		}).
		AddFrame("wave", []string{
			"_L5FFF5R_",
			"56FFFFF6_",
			"_FF__FF__",
			"__FF_FF__",
		})

	// Create character using service
	domainRobot, err := service.CreateCharacter(*spec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating domain robot: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("◢ Domain Robot character (created via service):")
	for _, frame := range domainRobot.Frames {
		fmt.Printf("Frame: %s\n", frame.Name)
		for _, line := range frame.Lines {
			fmt.Printf("  %s\n", line)
		}
		fmt.Println()
	}

	// Save character using service
	err = service.SaveCharacter(domainRobot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving domain robot: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Domain robot saved to repository")

	// Load character using service
	loadedRobot, err := service.LoadCharacter("demo-robot")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading domain robot: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✓ Domain robot loaded from repository: %s\n", loadedRobot.Name)
	fmt.Println()

	// =========================================
	// PART 4: New Builder Pattern (V2)
	// =========================================
	fmt.Println("▢ PART 4: New Builder Pattern (V2)")
	fmt.Println()
	fmt.Println("Using the new CharacterBuilderV2 with fluent API and service integration!")
	fmt.Println()

	// Create character using new builder
	domainRobotV2, err := characters.NewCharacterBuilderV2("robot-v2", 9, 4).
		AddFrame("idle", []string{
			"_L5FFF5R_",
			"_6FFFFF6_",
			"__FF_FF__",
			"__FF_FF__",
		}).
		AddFrame("wave", []string{
			"_L5FFF5R_",
			"56FFFFF6_",
			"_FF__FF__",
			"__FF_FF__",
		}).
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building robot V2: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("◢ Robot V2 character (created via new builder):")
	for _, frame := range domainRobotV2.Frames {
		fmt.Printf("Frame: %s\n", frame.Name)
		for _, line := range frame.Lines {
			fmt.Printf("  %s\n", line)
		}
		fmt.Println()
	}

	// Test fluent API with validation
	fmt.Println("◢ Testing fluent API validation:")
	invalidBuilder := characters.NewCharacterBuilderV2("invalid", 3, 2)
	err = invalidBuilder.Validate()
	if err != nil {
		fmt.Printf("✓ Validation correctly caught error: %v\n", err)
	}

	// Add frame and validate again
	invalidBuilder.AddFrame("idle", []string{"FRF", "LRL"})
	err = invalidBuilder.Validate()
	if err == nil {
		fmt.Println("✓ Validation passed after adding frame")
	}
	fmt.Println()

	// =========================================
	// PART 5: Legacy API (Backward Compatibility)
	// =========================================
	fmt.Println("▢ PART 5: Legacy API (Backward Compatibility)")
	fmt.Println()
	fmt.Println("The legacy API still works - backward compatibility maintained!")
	fmt.Println()

	// Create a simple robot character using legacy API
	robot, err := characters.NewCharacterSpec("robot", 9, 4).
		AddFrame("idle", []string{
			"_L5FFF5R_",
			"_6FFFFF6_",
			"__FF_FF__",
			"__FF_FF__",
		}).
		AddFrame("wave", []string{
			"_L5FFF5R_",
			"56FFFFF6_",
			"_FF__FF__",
			"__FF_FF__",
		}).
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building robot: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("◢ Legacy Robot character (idle):")
	characters.ShowIdle(os.Stdout, robot)
	fmt.Println()

	fmt.Println("◢ Legacy Robot character (animated):")
	characters.Animate(os.Stdout, robot, 3, 2)
	fmt.Println()

	// =========================================
	// PART 6: Registry Management
	// =========================================
	fmt.Println("▢ PART 6: Character Registry")
	fmt.Println()

	// Register characters
	characters.Register(robot)
	characters.UseLibrary("alien") // Register alien from library

	fmt.Println("Registered characters:", characters.List())
	fmt.Println()

	// Retrieve and use
	retrievedRobot, err := characters.Get("robot")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("◢ Retrieved 'robot' from registry:")
	characters.ShowIdle(os.Stdout, retrievedRobot)
	fmt.Println()

	// =========================================
	// Summary
	// =========================================
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Quick Reference                         ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Architecture Options:")
	fmt.Println("  • New: Use service layer with dependency injection")
	fmt.Println("  • Builder V2: Use new fluent API with validation")
	fmt.Println("  • Legacy: Use existing API (backward compatible)")
	fmt.Println()
	fmt.Println("Pattern Codes:")
	fmt.Println("  F=█  T=▀  B=▄  L=▌  R=▐")
	fmt.Println("  1=▘  2=▝  3=▖  4=▗")
	fmt.Println("  5=▛  6=▜  7=▙  8=▟")
	fmt.Println("  \\=▚  /=▞  _=Space")
	fmt.Println()
	fmt.Println("Create your own:")
	fmt.Println("  1. Use Tangent CLI: tangent")
	fmt.Println("  2. Or code directly with NewCharacterSpec()")
	fmt.Println("  3. Or use new service layer for advanced features")
	fmt.Println("  4. Or use CharacterBuilderV2 for fluent API")
	fmt.Println()
	fmt.Println("✓ Demo complete!")
}
