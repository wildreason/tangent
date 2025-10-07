package main

import (
	"fmt"
	"os"
	"time"

	"local/characters/pkg/characters"
)

func main() {
	// Create your character using the builder pattern
	char, err := characters.NewBuilder("my_alien", 16, 16).
		// Head row 1: lf, comp1, fb, fb, fb, comp2, rf
		Pattern("▌▛███▜▐").
		// Head row 2: same pattern
		Pattern("▌▛███▜▐").
		// Eyes row
		Pattern("  █  █  █  ").
		// Body rows
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		// Legs
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		// Animation frame 1: left hand only
		NewFrame().
		Pattern("▌▛██  ").
		Pattern("▌▛██  ").
		Pattern("  █  █  █  ").
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		// Animation frame 2: no hands
		NewFrame().
		Pattern("   ███   ").
		Pattern("   ███   ").
		Pattern("  █  █  █  ").
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		// Animation frame 3: right hand only
		NewFrame().
		Pattern("  ███▜▐").
		Pattern("  ███▜▐").
		Pattern("  █  █  █  ").
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Block(16).
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Pattern("  ██        ██  ").
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building character: %v\n", err)
		os.Exit(1)
	}

	// Register the character
	if err := characters.Register(char); err != nil {
		fmt.Fprintf(os.Stderr, "Error registering character: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Character created and registered!")
	fmt.Printf("Available characters: %v\n", characters.List())

	// Show static idle state
	fmt.Println("\n=== IDLE STATE ===")
	characters.ShowIdle(os.Stdout, char)

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Animate the character
	fmt.Println("\n=== ANIMATION (3 loops) ===")
	if err := characters.Animate(os.Stdout, char, 6, 3); err != nil {
		fmt.Fprintf(os.Stderr, "Animation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nAnimation complete!")
}
