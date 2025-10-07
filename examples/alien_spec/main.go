package main

import (
	"fmt"
	"os"
	"time"

	"local/characters/pkg/characters"
)

func main() {
	// Create alien character based on your corrected specification
	// Frame 1: 0,0,rf,comp1,fb,fb,fb,comp2,lf,0,0
	// Frame 2: 0,uh,comp2,fb,fb,fb,fb,fb,comp1,uh,0
	// Frame 3: 0,0,0,ul,ul,0,ur,ur,0,0,0

	char, err := characters.NewBuilder("alien_spec", 11, 3).
		// Frame 1: 0,0,rf,comp1,fb,fb,fb,comp2,lf,0,0
		Pattern("  ▐▛███▜▌  ").
		// Frame 2: 0,uh,comp2,fb,fb,fb,fb,fb,comp1,uh,0
		Pattern(" ▀▜█████▛▀ ").
		// Frame 3: 0,0,0,ul,ul,0,ur,ur,0,0
		Pattern("   ▘▘ ▝▝  ").
		// Frame 2
		NewFrame().
		Pattern("  ▐▛███▜▌  ").
		Pattern("▛▀▜█████▛▀ ").
		Pattern("   ▘▘ ▝▝  ").
		// Frame 3
		NewFrame().
		Pattern("  ▐▛███▜▌  ").
		Pattern(" ▀▜█████▛▀▜").
		Pattern("   ▘▘ ▝▝  ").
		Build()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building character: %v\n", err)
		os.Exit(1)
	}

	// Register the character
	characters.Register(char)

	fmt.Println("Alien character based on your corrected specification:")
	fmt.Println("Frame 1: 0,0,rf,comp1,fb,fb,fb,comp2,lf,0,0")
	fmt.Println("Frame 2: 0,uh,comp2,fb,fb,fb,fb,fb,comp1,uh,0")
	fmt.Println("Frame 3: 0,0,ul,ul,0,ur,ur,0,0")

	// Show static idle state
	fmt.Println("\n=== FRAME 1 (IDLE) ===")
	characters.ShowIdle(os.Stdout, char)

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Animate between the 3 frames
	fmt.Println("\n=== ANIMATION (3 frames) ===")
	characters.Animate(os.Stdout, char, 6, 3)

	fmt.Println("\nDone!")
}
