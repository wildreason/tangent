package characters

import (
	"fmt"
	"strings"
)

// Character represents a single character sprite
type Character struct {
	Name   string
	Idle   string
	Frames []string
	Width  int
	Height int
}

// Builder provides a fluent API for building characters
type Builder struct {
	name   string
	width  int
	height int
	rows   []string
	frames [][]string
}

// NewBuilder creates a new character builder
func NewBuilder(name string, width, height int) *Builder {
	return &Builder{
		name:   name,
		width:  width,
		height: height,
		rows:   make([]string, 0, height),
		frames: make([][]string, 0),
	}
}

// Row adds a row to the current frame
func (b *Builder) Row(pattern string) *Builder {
	padded := b.padToWidth(pattern)
	b.rows = append(b.rows, padded)
	return b
}

// Block creates a row of n full blocks
func (b *Builder) Block(n int) *Builder {
	return b.Row(strings.Repeat(string(FullBlock), n))
}

// Space creates a row of n spaces
func (b *Builder) Space(n int) *Builder {
	return b.Row(strings.Repeat(" ", n))
}

// Pattern creates a row from a pattern string
func (b *Builder) Pattern(pattern string) *Builder {
	return b.Row(pattern)
}

// Custom creates a row from custom runes
func (b *Builder) Custom(runes ...rune) *Builder {
	return b.Row(string(runes))
}

// Helper methods for common patterns
func (b *Builder) LeftCap() *Builder {
	return b.Custom(LeftHalfBlock)
}

func (b *Builder) RightCap() *Builder {
	return b.Custom(RightHalfBlock)
}

func (b *Builder) ThreeQuad5() *Builder {
	return b.Custom(ThreeQuad5)
}

func (b *Builder) ThreeQuad6() *Builder {
	return b.Custom(ThreeQuad6)
}

func (b *Builder) FullBlocks(n int) *Builder {
	return b.Custom([]rune(strings.Repeat(string(FullBlock), n))...)
}

func (b *Builder) Spaces(n int) *Builder {
	return b.Custom([]rune(strings.Repeat(" ", n))...)
}

// NewFrame starts a new animation frame
func (b *Builder) NewFrame() *Builder {
	if len(b.rows) > 0 {
		b.frames = append(b.frames, b.rows)
		b.rows = make([]string, 0, b.height)
	}
	return b
}

// Build creates the final Character
func (b *Builder) Build() (*Character, error) {
	// Add current frame if it exists
	if len(b.rows) > 0 {
		b.frames = append(b.frames, b.rows)
	}

	if len(b.frames) == 0 {
		return nil, fmt.Errorf("no frames defined for character %s", b.name)
	}

	// Use first frame as idle
	idle := strings.Join(b.frames[0], "\n")

	// Convert frames to strings
	frameStrings := make([]string, len(b.frames))
	for i, frame := range b.frames {
		frameStrings[i] = strings.Join(frame, "\n")
	}

	return &Character{
		Name:   b.name,
		Idle:   idle,
		Frames: frameStrings,
		Width:  b.width,
		Height: b.height,
	}, nil
}

// padToWidth centers and pads a string to the target width
func (b *Builder) padToWidth(s string) string {
	runes := []rune(s)
	if len(runes) >= b.width {
		return string(runes[:b.width])
	}

	left := (b.width - len(runes)) / 2
	right := b.width - len(runes) - left

	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

// Validate checks if the character meets requirements
func (c *Character) Validate() error {
	// Check idle frame
	if err := c.validateFrame(c.Idle, "idle"); err != nil {
		return err
	}

	// Check animation frames
	for i, frame := range c.Frames {
		if err := c.validateFrame(frame, fmt.Sprintf("frame %d", i)); err != nil {
			return err
		}
	}

	return nil
}

func (c *Character) validateFrame(frame, name string) error {
	lines := strings.Split(frame, "\n")
	if len(lines) != c.Height {
		return fmt.Errorf("%s: expected %d lines, got %d", name, c.Height, len(lines))
	}

	for i, line := range lines {
		if len([]rune(line)) != c.Width {
			return fmt.Errorf("%s line %d: expected width %d, got %d", name, i+1, c.Width, len([]rune(line)))
		}
	}

	return nil
}
