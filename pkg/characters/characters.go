package characters

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
