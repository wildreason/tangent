package characters

import "local/characters/pkg/characters/library"

// Package characters provides a simple API for creating and animating
// Unicode Block Elements (U+2580–U+259F) characters in Go applications.
//
// This package focuses exclusively on Block Elements for pixel-perfect
// character design and animation.
//
// Block Elements used:
// - █ (U+2588) Full Block
// - ▌ (U+258C) Left Half Block
// - ▐ (U+2590) Right Half Block
// - ▛ (U+259B) Quadrant Upper Left + Upper Right + Lower Left
// - ▜ (U+259C) Quadrant Upper Left + Upper Right + Lower Right
// - ▙ (U+2599) Quadrant Upper Left + Lower Left + Lower Right
// - ▟ (U+259F) Quadrant Upper Right + Lower Left + Lower Right
// - ▘ (U+2598) Quadrant Upper Left
// - ▝ (U+259D) Quadrant Upper Right
// - ▖ (U+2596) Quadrant Lower Left
// - ▗ (U+2597) Quadrant Lower Right
//
// Example usage:
//
//	char, err := NewBuilder("alien", 16, 16).
//		Pattern("▌▛███▜▐").
//		Block(16).
//		NewFrame().
//		Pattern("▌▛██  ").
//		Block(16).
//		Build()
//
//	Register(char)
//	Animate(os.Stdout, char, 6, 3)
//
// Library characters:
//
//	alien, _ := Library("alien")
//	Animate(os.Stdout, alien, 4, 2)

// Library retrieves a pre-built character from the built-in library and builds it
func Library(name string) (*Character, error) {
	libChar, err := library.Get(name)
	if err != nil {
		return nil, err
	}

	// Build the character using the library patterns
	spec := NewCharacterSpec(libChar.Name, libChar.Width, libChar.Height)
	for _, frame := range libChar.Patterns {
		spec.AddFrame(frame.Name, frame.Lines)
	}

	return spec.Build()
}

// ListLibrary returns all available library character names
func ListLibrary() []string {
	return library.List()
}

// UseLibrary retrieves a library character and registers it in the global registry
func UseLibrary(name string) (*Character, error) {
	char, err := Library(name)
	if err != nil {
		return nil, err
	}
	Register(char)
	return char, nil
}

// LibraryInfo returns information about a library character
func LibraryInfo(name string) (string, error) {
	libChar, err := library.Get(name)
	if err != nil {
		return "", err
	}
	return libChar.Description, nil
}
