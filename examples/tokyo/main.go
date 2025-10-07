package main

import (
	"fmt"
	"os"

	"local/characters/pkg/characters"
)

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Paris Character - Created with Tangent ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Pattern codes exported from tangent:
	// Frame: head
	// Line 1: __R6FFF6L__
	// Line 2: _T6FFFFF5T_
	// Line 3: ___11_22___
	//
	// Frame: wave
	// Line 1: __R6FFF6L__
	// Line 2: __6FFFFF5T1
	// Line 3: ___11_22___
	//
	// Frame: wave-2
	// Line 1: __R6FFF6L__
	// Line 2: 2T6FFFFF5__
	// Line 3: ___11_22___

	spec := characters.NewCharacterSpec("paris", 11, 3).
		AddFrame("head", []string{
			"__R6FFF6L__",
			"_T6FFFFF5T_",
			"___11_22___",
		}).
		AddFrame("wave", []string{
			"__R6FFF6L__",
			"__6FFFFF5T1",
			"___11_22___",
		}).
		AddFrame("wave-2", []string{
			"__R6FFF6L__",
			"2T6FFFFF5__",
			"___11_22___",
		})

	char, err := spec.Build()
	if err != nil {
		fmt.Printf("✗ Error building character: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("◢ Showing idle frame:")
	fmt.Println()
	characters.ShowIdle(os.Stdout, char)

	fmt.Println()
	fmt.Println("◢ Animating character (5 cycles at 5 FPS):")
	fmt.Println()
	characters.Animate(os.Stdout, char, 5, 5)

	fmt.Println()
	fmt.Println("✓ Animation complete!")
}
