package main

import (
	"fmt"
	"os"

	"local/characters/pkg/characters"
)

func main() {
	fmt.Println("Hex-Style Pattern Creation (Variant B)")
	fmt.Println("======================================")
	fmt.Println()

	// Show the pattern legend
	compiler := characters.NewPatternCompiler()
	fmt.Println(compiler.GetPatternLegend())
	fmt.Println()

	// Create the alien character using hex-style patterns
	// Pattern format: "00R9FFF9L00" (no commas, like hex color codes!)
	alien := characters.NewCharacterSpec("alien", 11, 3).
		AddFrame("idle", []string{
			"00R9FFF9L00", // Line 1: space, space, right-half, 3-quad-9, full, full, full, 3-quad-9, left-half, space, space
			"0T9FFFFF7T0", // Line 2: space, top-half, 3-quad-9, full, full, full, full, full, 3-quad-7, top-half, space
			"00011000220", // Line 3: space, space, space, quad-1, quad-1, space, quad-2, quad-2, space, space, space
		}).
		AddFrame("left", []string{
			"00R9FFF9L00",
			"7T9FFFFF7T0",
			"0011000220_", // Using _ for space (for readability)
		}).
		AddFrame("right", []string{
			"00R9FFF9L00",
			"0T9FFFFF7T9",
			"_0011000220",
		})

	// Build and register
	char, err := alien.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	characters.Register(char)

	// Show the patterns in action
	fmt.Println("=== CHARACTER SPECIFICATION ===")
	fmt.Println("Frame 1 (idle):")
	fmt.Println("  Pattern: \"00R9FFF9L00\"")
	fmt.Println("  Pattern: \"0T9FFFFF7T0\"")
	fmt.Println("  Pattern: \"00011000220\"")
	fmt.Println()

	// Show static
	fmt.Println("=== IDLE STATE ===")
	characters.ShowIdle(os.Stdout, char)

	// Show animation
	fmt.Println("\n=== ANIMATION ===")
	characters.Animate(os.Stdout, char, 3, 2)

	fmt.Println("\nDone!")

	// Show how easy it is to edit
	fmt.Println("\n=== EASY EDITING EXAMPLE ===")
	fmt.Println("To make the head wider, just add more F's:")
	fmt.Println("  OLD: \"00R9FFF9L00\"")
	fmt.Println("  NEW: \"00R9FFFF9L0\"")
	fmt.Println()
	fmt.Println("To add shading, use . : or #:")
	fmt.Println("  OLD: \"0T9FFFFF7T0\"")
	fmt.Println("  NEW: \"0T9FF#FF7T0\"")
	fmt.Println()
	fmt.Println("It's just like editing hex colors: #F8394839")
	fmt.Println("But for block characters: 00R9FFF9L00")
}
