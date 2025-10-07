package main

import (
	"fmt"
	"os"
	"time"

	"local/characters/pkg/characters"
)

func main() {
	// Create exactly what you asked for:
	// Frame 1: rf, comp1, fb, fb, fb, comp2, lf
	// Frame 2: lf, comp1, fb, fb, fb, comp2, rf
	char, err := characters.NewBuilder("simple", 16, 16).
		// Frame 1: rf, comp1, fb, fb, fb, comp2, lf
		Pattern("▐▛███▜▌").
		// Just repeat the same pattern for all rows
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		Pattern("▐▛███▜▌").
		// Frame 2: lf, comp1, fb, fb, fb, comp2, rf
		NewFrame().
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Pattern("▌▛███▜▐").
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building character: %v\n", err)
		os.Exit(1)
	}

	// Register the character
	characters.Register(char)

	fmt.Println("Frame 1: rf, comp1, fb, fb, fb, comp2, lf = ▐▛███▜▌")
	fmt.Println("Frame 2: lf, comp1, fb, fb, fb, comp2, rf = ▌▛███▜▐")

	// Show static idle state
	fmt.Println("\n=== FRAME 1 (IDLE) ===")
	characters.ShowIdle(os.Stdout, char)

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Animate between the 2 frames
	fmt.Println("\n=== ANIMATION (2 frames) ===")
	characters.Animate(os.Stdout, char, 6, 3)

	fmt.Println("\nDone!")
}
