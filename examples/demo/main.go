package main

import (
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
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
	// PART 2: Creating Custom Characters
	// =========================================
	fmt.Println("▢ PART 2: Custom Characters (Pattern-Based)")
	fmt.Println()
	fmt.Println("This is what Tangent generates - clean pattern codes!")
	fmt.Println()

	// Create a simple robot character
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

	fmt.Println("◢ Robot character (idle):")
	characters.ShowIdle(os.Stdout, robot)
	fmt.Println()

	fmt.Println("◢ Robot character (animated):")
	characters.Animate(os.Stdout, robot, 3, 2)
	fmt.Println()

	// =========================================
	// PART 3: Registry Management
	// =========================================
	fmt.Println("▢ PART 3: Character Registry")
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
	fmt.Println("Pattern Codes:")
	fmt.Println("  F=█  T=▀  B=▄  L=▌  R=▐")
	fmt.Println("  1=▘  2=▝  3=▖  4=▗")
	fmt.Println("  5=▛  6=▜  7=▙  8=▟")
	fmt.Println("  \\=▚  /=▞  _=Space")
	fmt.Println()
	fmt.Println("Create your own:")
	fmt.Println("  1. Use Tangent CLI: tangent")
	fmt.Println("  2. Or code directly with NewCharacterSpec()")
	fmt.Println()
	fmt.Println("✓ Demo complete!")
}
