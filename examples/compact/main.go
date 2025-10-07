package main

import (
	"fmt"
	"os"

	"local/characters/pkg/characters"
)

func main() {
	fmt.Println("Compact Hex-Style Patterns")
	fmt.Println("===========================")
	fmt.Println()
	fmt.Println("Like hex colors (#F8394839), but for characters!")
	fmt.Println()

	// Create alien using compact patterns - no commas!
	alien := characters.NewCharacterSpec("alien", 11, 3).
		AddFrame("idle", []string{
			"00R9FFF9L00",
			"0T9FFFFF7T0",
			"00011000220",
		}).
		AddFrame("left", []string{
			"00R9FFF9L00",
			"7T9FFFFF7T0",
			"00011000220",
		}).
		AddFrame("right", []string{
			"00R9FFF9L00",
			"0T9FFFFF7T9",
			"00011000220",
		})

	char, err := alien.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	characters.Register(char)

	// Show it
	fmt.Println("Pattern: \"00R9FFF9L00\"")
	fmt.Println("Result:")
	characters.ShowIdle(os.Stdout, char)

	// Animate
	fmt.Println("\nAnimating...")
	characters.Animate(os.Stdout, char, 4, 2)

	fmt.Println("\nDone!")

	// Show the mapping
	fmt.Println("\nQuick Reference:")
	fmt.Println("  F=█ T=▀ B=▄ L=▌ R=▐")
	fmt.Println("  7=▛ 9=▜ 6=▙ 8=▟")
	fmt.Println("  1=▘ 2=▝ 3=▖ 4=▗")
	fmt.Println("  .=░ :=▒ #=▓ 0=Space")
}
