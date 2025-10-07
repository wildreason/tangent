// Package bubbletea provides integration helpers for using Tangent characters
// with the Bubble Tea TUI framework and Charmbracelet ecosystem.
//
// This adapter makes it trivial to use Tangent-designed characters as
// Bubble Tea spinner frames while maintaining full control over your
// event loop, styling, and layout.
package bubbletea

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/wildreason/tangent/pkg/characters"
)

// SpinnerFromCharacter creates a Bubble Tea spinner from a Tangent character.
// This is the primary integration point for Bubble Tea apps.
//
// Example:
//
//	alien, _ := characters.Library("alien")
//	s := bubbletea.SpinnerFromCharacter(alien, 6)
//	s.Start()  // Use in your Bubble Tea model
func SpinnerFromCharacter(char *characters.Character, fps int) spinner.Model {
	// Extract frames from character
	frames := characters.ToSpinnerFrames(char)

	// Create spinner with Tangent frames
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: frames,
		FPS:    time.Second / time.Duration(fps),
	}

	return s
}

// FramesFromCharacter extracts frame strings for custom integration.
// Use this if you need more control than SpinnerFromCharacter provides.
//
// Example:
//
//	frames := bubbletea.FramesFromCharacter(alien)
//	// Custom spinner configuration
//	s := spinner.New()
//	s.Spinner = spinner.Spinner{
//		Frames: frames,
//		FPS:    time.Second / 10,  // Custom FPS
//	}
func FramesFromCharacter(char *characters.Character) []string {
	return characters.ExtractFrames(char)
}

// NormalizedSpinner creates a spinner with normalized frames (no jitter).
// This ensures all frames have consistent width and height.
//
// Example:
//
//	s := bubbletea.NormalizedSpinner(myChar, 5)
func NormalizedSpinner(char *characters.Character, fps int) spinner.Model {
	// Normalize frames to prevent jitter
	normalized := char.Normalize()
	return SpinnerFromCharacter(normalized, fps)
}

// MultiCharacterSpinners creates multiple spinners from multiple characters.
// Useful for apps with multiple animated elements.
//
// Example:
//
//	chars := map[string]*characters.Character{
//		"agent":   agentChar,
//		"loading": loadingChar,
//	}
//	spinners := bubbletea.MultiCharacterSpinners(chars, 5)
//	agent := spinners["agent"]
//	loading := spinners["loading"]
func MultiCharacterSpinners(chars map[string]*characters.Character, fps int) map[string]spinner.Model {
	spinners := make(map[string]spinner.Model, len(chars))

	for name, char := range chars {
		spinners[name] = SpinnerFromCharacter(char, fps)
	}

	return spinners
}

// LibrarySpinner loads a library character and creates a spinner in one call.
// This is the fastest way to get started with Tangent library characters.
//
// Example:
//
//	s, err := bubbletea.LibrarySpinner("alien", 5)
//	if err != nil {
//		// handle error
//	}
func LibrarySpinner(name string, fps int) (spinner.Model, error) {
	char, err := characters.Library(name)
	if err != nil {
		return spinner.Model{}, err
	}

	return SpinnerFromCharacter(char, fps), nil
}

// SpinnerConfig provides advanced configuration options
type SpinnerConfig struct {
	FPS       int
	Normalize bool // Normalize frames to prevent jitter
	Style     any  // Optional lipgloss.Style (interface{} to avoid import)
}

// ConfiguredSpinner creates a spinner with advanced configuration
func ConfiguredSpinner(char *characters.Character, config SpinnerConfig) spinner.Model {
	// Apply normalization if requested
	if config.Normalize {
		char = char.Normalize()
	}

	// Create spinner
	frames := characters.ToSpinnerFrames(char)
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: frames,
		FPS:    time.Second / time.Duration(config.FPS),
	}

	// Note: Styling should be done in the View() method with lipgloss
	// We don't apply styles here to maintain separation of concerns

	return s
}
