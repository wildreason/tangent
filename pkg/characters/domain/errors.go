package domain

import (
	"fmt"
)

// Error types for better error handling
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
	Cause   error
}

func (e *ValidationError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("validation error in field '%s': %s (caused by: %v)", e.Field, e.Message, e.Cause)
	}
	return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

func (e *ValidationError) Unwrap() error {
	return e.Cause
}

type CharacterNotFoundError struct {
	Name string
}

func (e *CharacterNotFoundError) Error() string {
	return fmt.Sprintf("character '%s' not found", e.Name)
}

type PatternCompilationError struct {
	Pattern  string
	Position int
	Message  string
	Cause    error
}

func (e *PatternCompilationError) Error() string {
	if e.Position >= 0 {
		return fmt.Sprintf("pattern compilation error at position %d in '%s': %s", e.Position, e.Pattern, e.Message)
	}
	return fmt.Sprintf("pattern compilation error in '%s': %s", e.Pattern, e.Message)
}

func (e *PatternCompilationError) Unwrap() error {
	return e.Cause
}

type AnimationError struct {
	CharacterName string
	Operation     string
	Message       string
	Cause         error
}

func (e *AnimationError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("animation error for character '%s' during %s: %s (caused by: %v)", e.CharacterName, e.Operation, e.Message, e.Cause)
	}
	return fmt.Sprintf("animation error for character '%s' during %s: %s", e.CharacterName, e.Operation, e.Message)
}

func (e *AnimationError) Unwrap() error {
	return e.Cause
}

// Domain-specific errors with helpful context
var (
	ErrCharacterNameRequired = &ValidationError{
		Field:   "name",
		Message: "character name is required and cannot be empty",
	}

	ErrCharacterNameInvalid = &ValidationError{
		Field:   "name",
		Message: "character name contains invalid characters (only letters, numbers, hyphens, and underscores allowed)",
	}

	ErrInvalidDimensions = &ValidationError{
		Field:   "dimensions",
		Message: "character dimensions must be positive integers (width > 0, height > 0)",
	}

	ErrEmptyPattern = &ValidationError{
		Field:   "pattern",
		Message: "pattern cannot be empty",
	}

	ErrInvalidFrameCount = &ValidationError{
		Field:   "frames",
		Message: "frame count does not match character height",
	}

	ErrInvalidFrameName = &ValidationError{
		Field:   "frame_name",
		Message: "frame name cannot be empty",
	}

	ErrCharacterNotFound = &CharacterNotFoundError{}

	ErrPatternTooLong = &ValidationError{
		Field:   "pattern",
		Message: "pattern exceeds maximum length",
	}

	ErrInvalidPatternCharacter = &ValidationError{
		Field:   "pattern",
		Message: "pattern contains invalid characters",
	}
)

// Helper functions to create contextual errors
func NewValidationError(field string, value interface{}, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
	}
}

func NewValidationErrorWithCause(field string, value interface{}, message string, cause error) *ValidationError {
	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
		Cause:   cause,
	}
}

func NewCharacterNotFoundError(name string) *CharacterNotFoundError {
	return &CharacterNotFoundError{Name: name}
}

func NewPatternCompilationError(pattern string, position int, message string, cause error) *PatternCompilationError {
	return &PatternCompilationError{
		Pattern:  pattern,
		Position: position,
		Message:  message,
		Cause:    cause,
	}
}

func NewAnimationError(characterName, operation, message string, cause error) *AnimationError {
	return &AnimationError{
		CharacterName: characterName,
		Operation:     operation,
		Message:       message,
		Cause:         cause,
	}
}

// Error suggestion helpers
func GetErrorSuggestion(err error) string {
	switch e := err.(type) {
	case *ValidationError:
		return getValidationSuggestion(e)
	case *CharacterNotFoundError:
		return fmt.Sprintf("Try using ListCharacters() to see available characters, or create a new character with name '%s'", e.Name)
	case *PatternCompilationError:
		return getPatternSuggestion(e)
	case *AnimationError:
		return getAnimationSuggestion(e)
	default:
		return "Check the error message for details and ensure all inputs are valid"
	}
}

func getValidationSuggestion(err *ValidationError) string {
	switch err.Field {
	case "name":
		return "Use a descriptive name with letters, numbers, hyphens, or underscores (e.g., 'my-character', 'robot_v2')"
	case "dimensions":
		return "Ensure width and height are positive integers (e.g., width: 8, height: 4)"
	case "pattern":
		return "Use valid pattern characters: F=█, R=▐, L=▌, T=▀, B=▄, 1-8=quadrants, _=space"
	case "frames":
		return "Ensure the number of frame patterns matches the character height"
	case "frame_name":
		return "Use descriptive frame names (e.g., 'idle', 'wave', 'jump')"
	default:
		return "Check the field value and ensure it meets the requirements"
	}
}

func getPatternSuggestion(err *PatternCompilationError) string {
	if err.Position >= 0 {
		return fmt.Sprintf("Check character at position %d in pattern '%s' - use valid pattern characters", err.Position, err.Pattern)
	}
	return "Review the pattern and ensure all characters are valid pattern symbols"
}

func getAnimationSuggestion(err *AnimationError) string {
	switch err.Operation {
	case "start":
		return "Ensure the character has frames and the terminal supports ANSI escape codes"
	case "frame_display":
		return "Check that character frames are properly formatted and terminal is wide enough"
	case "timing":
		return "Verify fps and loops parameters are positive integers"
	default:
		return "Check character data and animation parameters"
	}
}
