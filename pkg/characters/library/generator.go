package library

import (
	"fmt"
	"sort"

	"github.com/wildreason/tangent/pkg/characters/microstateregistry"
	"github.com/wildreason/tangent/pkg/characters/stateregistry"
)

// CharacterMetadata holds the non-pattern metadata for a character
type CharacterMetadata struct {
	Name        string
	Description string
	Author      string
	Color       string
	Width       int
	Height      int
}

// GenerateFromRegistry creates a LibraryCharacter from state registry
func GenerateFromRegistry(metadata CharacterMetadata) LibraryCharacter {
	// Get all states from registry
	allStates := stateregistry.All()

	// Build frames from all states in alphabetical order
	var frames []Frame
	stateNames := make([]string, 0, len(allStates))
	for name := range allStates {
		stateNames = append(stateNames, name)
	}
	sort.Strings(stateNames)

	// Convert state registry frames to library frames
	for _, stateName := range stateNames {
		state := allStates[stateName]
		for i, stateFrame := range state.Frames {
			frame := Frame{
				Name:  fmt.Sprintf("%s_%d", stateName, i+1),
				Lines: stateFrame.Lines,
			}
			frames = append(frames, frame)
		}
	}

	return LibraryCharacter{
		Name:        metadata.Name,
		Description: metadata.Description,
		Author:      metadata.Author,
		Color:       metadata.Color,
		Width:       metadata.Width,
		Height:      metadata.Height,
		Patterns:    frames,
	}
}

// GenerateMicroFromRegistry creates a micro LibraryCharacter from the micro state registry
func GenerateMicroFromRegistry(metadata CharacterMetadata) LibraryCharacter {
	def := microstateregistry.Get()
	if def == nil {
		return LibraryCharacter{}
	}

	// Build frames from all states
	var frames []Frame

	// Add base frame first
	frames = append(frames, Frame{
		Name:  "base",
		Lines: def.BaseFrame.Lines,
	})

	// Convert micro state frames to library frames
	for _, state := range def.States {
		for i, stateFrame := range state.Frames {
			frame := Frame{
				Name:  fmt.Sprintf("%s_%d", state.Name, i+1),
				Lines: stateFrame.Lines,
			}
			frames = append(frames, frame)
		}
	}

	return LibraryCharacter{
		Name:        metadata.Name,
		Description: metadata.Description,
		Author:      metadata.Author,
		Color:       metadata.Color,
		Width:       def.Width,
		Height:      def.Height,
		Patterns:    frames,
	}
}
