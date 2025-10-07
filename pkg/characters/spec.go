package characters

import (
	"fmt"
	"strings"
)

// CharacterSpec defines a character using simple text patterns
type CharacterSpec struct {
	Name   string      `json:"name"`
	Width  int         `json:"width"`
	Height int         `json:"height"`
	Frames []FrameSpec `json:"frames"`
}

// FrameSpec defines a single frame using text patterns
type FrameSpec struct {
	Name     string   `json:"name"`
	Patterns []string `json:"patterns"`
}

// NewCharacterSpec creates a new character specification
func NewCharacterSpec(name string, width, height int) *CharacterSpec {
	return &CharacterSpec{
		Name:   name,
		Width:  width,
		Height: height,
		Frames: make([]FrameSpec, 0),
	}
}

// AddFrame adds a frame to the character specification
func (cs *CharacterSpec) AddFrame(name string, patterns []string) *CharacterSpec {
	cs.Frames = append(cs.Frames, FrameSpec{
		Name:     name,
		Patterns: patterns,
	})
	return cs
}

// AddFrameFromString adds a frame from a single string pattern
func (cs *CharacterSpec) AddFrameFromString(name, pattern string) *CharacterSpec {
	// Split by newlines to get individual line patterns
	lines := strings.Split(pattern, "\n")
	patterns := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			patterns = append(patterns, line)
		}
	}

	return cs.AddFrame(name, patterns)
}

// Build compiles the character specification into a Character
func (cs *CharacterSpec) Build() (*Character, error) {
	compiler := NewPatternCompiler()

	// Convert frame patterns to the format expected by CompileCharacter
	framePatterns := make([][]string, len(cs.Frames))
	for i, frame := range cs.Frames {
		framePatterns[i] = frame.Patterns
	}

	return compiler.CompileCharacter(cs.Name, cs.Width, cs.Height, framePatterns)
}

// String returns a string representation of the character specification
func (cs *CharacterSpec) String() string {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("Character: %s (%dx%d)\n", cs.Name, cs.Width, cs.Height))

	for i, frame := range cs.Frames {
		result.WriteString(fmt.Sprintf("Frame %d (%s):\n", i+1, frame.Name))
		for j, pattern := range frame.Patterns {
			result.WriteString(fmt.Sprintf("  Line %d: %s\n", j+1, pattern))
		}
		result.WriteString("\n")
	}

	return result.String()
}

// Validate checks if the character specification is valid
func (cs *CharacterSpec) Validate() error {
	if cs.Name == "" {
		return fmt.Errorf("character name cannot be empty")
	}

	if cs.Width <= 0 || cs.Height <= 0 {
		return fmt.Errorf("width and height must be positive")
	}

	if len(cs.Frames) == 0 {
		return fmt.Errorf("character must have at least one frame")
	}

	// Validate each frame
	for i, frame := range cs.Frames {
		if frame.Name == "" {
			return fmt.Errorf("frame %d name cannot be empty", i+1)
		}

		if len(frame.Patterns) == 0 {
			return fmt.Errorf("frame %d must have at least one pattern", i+1)
		}

		if len(frame.Patterns) != cs.Height {
			return fmt.Errorf("frame %d has %d patterns, expected %d", i+1, len(frame.Patterns), cs.Height)
		}
	}

	return nil
}
