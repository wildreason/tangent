package domain

import "strings"

// Character represents a terminal character with frames and agent states
type Character struct {
	Name        string
	Personality string // Optional: "efficient", "friendly", "analytical", "creative"
	Color       string // Hex color for the character (e.g., "#FF4500")
	Width       int
	Height      int
	BaseFrame   Frame            // Idle/immutable base character
	States      map[string]State // Agent states (plan, think, execute, etc.)
	Frames      []Frame          // Keep for backward compatibility
}

// State represents an agent behavioral state
type State struct {
	Name           string
	Description    string
	Frames         []Frame // Multiple frames for animation
	StateType      string  // "standard" or "custom"
	AnimationFPS   int     // FPS for this state (default: 5)
	AnimationLoops int     // Loop count for this state (default: 1)
}

// Frame represents a single frame of animation
type Frame struct {
	Name  string
	Lines []string
}

// CharacterSpec represents the specification for creating a character
type CharacterSpec struct {
	Name   string
	Width  int
	Height int
	Frames []FrameSpec
}

// FrameSpec represents the specification for a single frame
type FrameSpec struct {
	Name     string
	Patterns []string
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

// AddFrame adds a frame to a character specification
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
