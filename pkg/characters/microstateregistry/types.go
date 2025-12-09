// Package microstateregistry provides types and loading for micro (10x2) avatar definitions.
package microstateregistry

// MicroFrame represents a single frame in a micro avatar animation.
type MicroFrame struct {
	Name  string   `json:"name,omitempty"`
	Lines []string `json:"lines"`
}

// MicroState represents an animation state with multiple frames.
type MicroState struct {
	Name   string       `json:"name"`
	FPS    int          `json:"fps,omitempty"`
	Frames []MicroFrame `json:"frames"`
}

// MicroDefinition represents a complete micro avatar definition with base frame and states.
type MicroDefinition struct {
	Name      string       `json:"name"`
	Width     int          `json:"width"`
	Height    int          `json:"height"`
	BaseFrame MicroFrame   `json:"base_frame"`
	States    []MicroState `json:"states"`
}
