package characters

import (
	"fmt"
	"strings"
	"time"
)

// Frame represents a single animation frame with metadata
type Frame struct {
	Name     string        // Frame identifier (e.g., "idle", "wave", "jump")
	Content  string        // Multi-line frame content
	Duration time.Duration // Optional duration for this frame
	Lines    []string      // Individual lines for easy access
}

// FrameMetadata contains optional metadata for frames
type FrameMetadata struct {
	States      map[string][]int // Named states mapping to frame indices
	Durations   []time.Duration  // Per-frame durations
	DefaultFPS  int              // Suggested frames per second
	LoopCount   int              // Suggested loop count (0 = infinite)
	Description string           // Frame sequence description
}

// GetFrames returns all frames from a character with metadata
// This is the primary API for frame extraction
func (c *Character) GetFrames() []Frame {
	frames := make([]Frame, len(c.Frames))

	for i, frameContent := range c.Frames {
		frames[i] = Frame{
			Name:    fmt.Sprintf("frame_%d", i),
			Content: frameContent,
			Lines:   strings.Split(frameContent, "\n"),
		}
	}

	return frames
}

// GetFrame returns a specific frame by index
func (c *Character) GetFrame(index int) (Frame, error) {
	if index < 0 || index >= len(c.Frames) {
		return Frame{}, fmt.Errorf("frame index %d out of range (0-%d)", index, len(c.Frames)-1)
	}

	frameContent := c.Frames[index]
	return Frame{
		Name:    fmt.Sprintf("frame_%d", index),
		Content: frameContent,
		Lines:   strings.Split(frameContent, "\n"),
	}, nil
}

// GetFrameLines returns all frames as [][]string for easy manipulation
func (c *Character) GetFrameLines() [][]string {
	result := make([][]string, len(c.Frames))

	for i, frameContent := range c.Frames {
		result[i] = strings.Split(frameContent, "\n")
	}

	return result
}

// GetIdleFrame returns the idle/default frame
func (c *Character) GetIdleFrame() Frame {
	content := c.Idle
	if content == "" && len(c.Frames) > 0 {
		content = c.Frames[0]
	}

	return Frame{
		Name:    "idle",
		Content: content,
		Lines:   strings.Split(content, "\n"),
	}
}

// FrameCount returns the number of frames in the character
func (c *Character) FrameCount() int {
	return len(c.Frames)
}

// Dimensions returns the width and height of the character
func (c *Character) Dimensions() (width, height int) {
	return c.Width, c.Height
}

// IsNormalized checks if all frames have consistent dimensions
func (c *Character) IsNormalized() bool {
	for _, frame := range c.Frames {
		lines := strings.Split(frame, "\n")

		// Check height
		if len(lines) != c.Height {
			return false
		}

		// Check width of each line
		for _, line := range lines {
			if len([]rune(line)) != c.Width {
				return false
			}
		}
	}

	return true
}

// Normalize ensures all frames have consistent width and height
// This prevents jitter during animation
func (c *Character) Normalize() *Character {
	if c.IsNormalized() {
		return c
	}

	normalized := &Character{
		Name:   c.Name,
		Width:  c.Width,
		Height: c.Height,
		Idle:   normalizeFrame(c.Idle, c.Width, c.Height),
		Frames: make([]string, len(c.Frames)),
	}

	for i, frame := range c.Frames {
		normalized.Frames[i] = normalizeFrame(frame, c.Width, c.Height)
	}

	return normalized
}

// normalizeFrame pads or trims a frame to exact dimensions
func normalizeFrame(frame string, width, height int) string {
	lines := strings.Split(frame, "\n")
	normalized := make([]string, height)

	for i := 0; i < height; i++ {
		if i < len(lines) {
			line := lines[i]
			runes := []rune(line)

			// Pad or trim to exact width
			if len(runes) < width {
				// Pad with spaces
				line = line + strings.Repeat(" ", width-len(runes))
			} else if len(runes) > width {
				// Trim to width
				line = string(runes[:width])
			}

			normalized[i] = line
		} else {
			// Add empty lines if needed
			normalized[i] = strings.Repeat(" ", width)
		}
	}

	return strings.Join(normalized, "\n")
}

// ToSpinnerFrames converts character frames to simple string slice
// Perfect for bubbles/spinner integration
func ToSpinnerFrames(c *Character) []string {
	// Return a copy to prevent modification
	frames := make([]string, len(c.Frames))
	copy(frames, c.Frames)
	return frames
}

// ExtractFrames is a convenience function for quick frame extraction
func ExtractFrames(c *Character) []string {
	return ToSpinnerFrames(c)
}

// FrameStats provides statistics about character frames
type FrameStats struct {
	TotalFrames  int
	Width        int
	Height       int
	IsNormalized bool
	TotalLines   int
	TotalRunes   int
}

// Stats returns statistics about the character's frames
func (c *Character) Stats() FrameStats {
	totalLines := 0
	totalRunes := 0

	for _, frame := range c.Frames {
		lines := strings.Split(frame, "\n")
		totalLines += len(lines)
		totalRunes += len([]rune(frame))
	}

	return FrameStats{
		TotalFrames:  len(c.Frames),
		Width:        c.Width,
		Height:       c.Height,
		IsNormalized: c.IsNormalized(),
		TotalLines:   totalLines,
		TotalRunes:   totalRunes,
	}
}
