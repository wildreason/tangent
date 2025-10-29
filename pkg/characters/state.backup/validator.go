package state

import (
	"fmt"
	"strings"
)

// FrameValidator validates frame definitions
type FrameValidator struct {
	// Valid pattern codes
	validCodes map[rune]bool
}

// NewFrameValidator creates a new frame validator
func NewFrameValidator() *FrameValidator {
	// Define all valid pattern codes based on the pattern compiler
	// pkg/characters/infrastructure/pattern_compiler.go
	validCodes := make(map[rune]bool)

	// Basic blocks (uppercase and lowercase)
	for _, c := range "FTBLRftblr" {
		validCodes[c] = true
	}

	// Shading blocks
	for _, c := range ".:# " {
		validCodes[c] = true
	}

	// Quadrants (1-8)
	for _, c := range "12345678" {
		validCodes[c] = true
	}

	// Diagonals
	validCodes['/'] = true
	validCodes['\\'] = true

	// Special characters
	validCodes['_'] = true
	validCodes['X'] = true

	return &FrameValidator{
		validCodes: validCodes,
	}
}

// ValidateStateConfig validates the entire state configuration
func (fv *FrameValidator) ValidateStateConfig(config *StateConfig) error {
	// Validate state name
	if err := fv.validateStateName(config.Name); err != nil {
		return err
	}

	// Validate frame count
	if config.FrameCount <= 0 {
		return &ValidationError{
			Field:   "FrameCount",
			Message: "must be greater than 0",
			Value:   config.FrameCount,
		}
	}

	// Validate frames match frame count
	if len(config.Frames) != config.FrameCount {
		return &ValidationError{
			Field:   "Frames",
			Message: fmt.Sprintf("expected %d frames, got %d", config.FrameCount, len(config.Frames)),
		}
	}

	// Validate required height
	if config.RequiredHeight < 0 || config.RequiredHeight > 10 {
		return &ValidationError{
			Field:   "RequiredHeight",
			Message: "must be between 0 and 10",
			Value:   config.RequiredHeight,
		}
	}

	// Validate targets
	if len(config.Targets) == 0 {
		return &ValidationError{
			Field:   "Targets",
			Message: "at least one target character required",
		}
	}

	// Validate each frame
	for i, frame := range config.Frames {
		if err := fv.ValidateFrame(&frame, i); err != nil {
			return err
		}
	}

	return nil
}

// validateStateName validates the state name format
func (fv *FrameValidator) validateStateName(name string) error {
	if name == "" {
		return &ValidationError{
			Field:   "Name",
			Message: "cannot be empty",
		}
	}

	// Check for valid characters (alphanumeric + underscore)
	for _, c := range name {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return &ValidationError{
				Field:   "Name",
				Message: "can only contain alphanumeric characters and underscores",
				Value:   name,
			}
		}
	}

	// Check for reserved names
	reserved := []string{"base", "default", "init"}
	for _, r := range reserved {
		if strings.ToLower(name) == r {
			return &ValidationError{
				Field:   "Name",
				Message: "is a reserved name",
				Value:   name,
			}
		}
	}

	return nil
}

// ValidateFrame validates a single frame definition
func (fv *FrameValidator) ValidateFrame(frame *FrameDefinition, frameIndex int) error {
	if len(frame.Lines) == 0 {
		return &FrameValidationError{
			FrameIndex: frameIndex,
			LineIndex:  -1,
			Message:    "frame must have at least one line",
		}
	}

	// Check all lines have same width
	firstWidth := len(frame.Lines[0])
	for i, line := range frame.Lines {
		if len(line) != firstWidth {
			return &FrameValidationError{
				FrameIndex: frameIndex,
				LineIndex:  i,
				Message:    fmt.Sprintf("line width mismatch: expected %d, got %d", firstWidth, len(line)),
			}
		}

		// Validate pattern codes in line
		if err := fv.validatePatternCodes(line, frameIndex, i); err != nil {
			return err
		}
	}

	return nil
}

// validatePatternCodes validates all pattern codes in a line
func (fv *FrameValidator) validatePatternCodes(line string, frameIndex, lineIndex int) error {
	for i, c := range line {
		if !fv.validCodes[c] {
			return &FrameValidationError{
				FrameIndex: frameIndex,
				LineIndex:  lineIndex,
				Message:    fmt.Sprintf("invalid pattern code '%c' at position %d", c, i),
			}
		}
	}
	return nil
}

// ValidateDimensions validates frame dimensions match expected width/height
func (fv *FrameValidator) ValidateDimensions(frames []FrameDefinition, expectedWidth, expectedHeight int) error {
	for i, frame := range frames {
		// Check height
		if len(frame.Lines) != expectedHeight {
			return &FrameValidationError{
				FrameIndex: i,
				LineIndex:  -1,
				Message:    fmt.Sprintf("height mismatch: expected %d lines, got %d", expectedHeight, len(frame.Lines)),
			}
		}

		// Check width
		for j, line := range frame.Lines {
			if len(line) != expectedWidth {
				return &FrameValidationError{
					FrameIndex: i,
					LineIndex:  j,
					Message:    fmt.Sprintf("width mismatch: expected %d characters, got %d", expectedWidth, len(line)),
				}
			}
		}
	}

	return nil
}

// ValidateFramesConsistency checks that all frames in a set have consistent dimensions
func (fv *FrameValidator) ValidateFramesConsistency(frames []FrameDefinition) error {
	if len(frames) == 0 {
		return nil
	}

	// Use first frame as reference
	refHeight := len(frames[0].Lines)
	refWidth := len(frames[0].Lines[0])

	return fv.ValidateDimensions(frames, refWidth, refHeight)
}
