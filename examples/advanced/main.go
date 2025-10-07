package main

import (
	"fmt"
	"os"
	"time"

	"local/characters/pkg/characters"
)

func main() {
	// Create multiple characters using different approaches

	// Method 1: Using helper methods
	robot, err := characters.NewBuilder("robot", 12, 12).
		LeftCap().Comp1().FullBlocks(3).Comp2().RightCap().
		LeftCap().Comp1().FullBlocks(3).Comp2().RightCap().
		Pattern("  █  █  █  ").
		Block(12).
		Block(12).
		Block(12).
		Block(12).
		Pattern("  ██    ██  ").
		Pattern("  ██    ██  ").
		Pattern("  ██    ██  ").
		Pattern("  ██    ██  ").
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building robot: %v\n", err)
		os.Exit(1)
	}

	// Method 2: Using custom runes
	ghost, err := characters.NewBuilder("ghost", 10, 10).
		Custom('█', '█', '█', '█', '█', '█', '█', '█', '█', '█').
		Custom('█', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '█').
		Custom('█', ' ', '█', ' ', '█', ' ', '█', ' ', ' ', '█').
		Custom('█', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '█').
		Custom('█', '█', '█', '█', '█', '█', '█', '█', '█', '█').
		Custom(' ', ' ', ' ', '█', '█', '█', '█', ' ', ' ', ' ').
		Custom(' ', ' ', '█', '█', ' ', ' ', '█', '█', ' ', ' ').
		Custom(' ', '█', '█', ' ', ' ', ' ', ' ', '█', '█', ' ').
		Custom('█', '█', ' ', ' ', ' ', ' ', ' ', ' ', '█', '█').
		Custom('█', '█', ' ', ' ', ' ', ' ', ' ', ' ', '█', '█').
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building ghost: %v\n", err)
		os.Exit(1)
	}

	// Method 3: With animation frames
	spaceship, err := characters.NewBuilder("spaceship", 14, 8).
		Pattern("    ██████    ").
		Pattern("  ██████████  ").
		Pattern("██████████████").
		Pattern("  ██████████  ").
		Pattern("    ██████    ").
		Pattern("      ██      ").
		Pattern("      ██      ").
		Pattern("      ██      ").
		// Animation frame: engines firing
		NewFrame().
		Pattern("    ██████    ").
		Pattern("  ██████████  ").
		Pattern("██████████████").
		Pattern("  ██████████  ").
		Pattern("    ██████    ").
		Pattern("    ██  ██    ").
		Pattern("    ██  ██    ").
		Pattern("    ██  ██    ").
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building spaceship: %v\n", err)
		os.Exit(1)
	}

	// Register all characters
	characters.Register(robot)
	characters.Register(ghost)
	characters.Register(spaceship)

	fmt.Println("Created and registered 3 characters!")
	fmt.Printf("Available: %v\n", characters.List())

	// Demonstrate each character
	chars := []string{"robot", "ghost", "spaceship"}

	for _, name := range chars {
		char, err := characters.Get(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting %s: %v\n", name, err)
			continue
		}

		fmt.Printf("\n=== %s (IDLE) ===\n", name)
		characters.ShowIdle(os.Stdout, char)

		time.Sleep(1 * time.Second)

		if len(char.Frames) > 1 {
			fmt.Printf("\n=== %s (ANIMATED) ===\n", name)
			characters.Animate(os.Stdout, char, 8, 2)
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("\nAll demonstrations complete!")
}
