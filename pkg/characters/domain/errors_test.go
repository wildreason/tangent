package domain

import (
	"errors"
	"strings"
	"testing"
)

func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *ValidationError
		expected string
	}{
		{
			name: "Error without cause",
			err: &ValidationError{
				Field:   "name",
				Value:   "test",
				Message: "invalid name",
			},
			expected: "validation error in field 'name': invalid name",
		},
		{
			name: "Error with cause",
			err: &ValidationError{
				Field:   "pattern",
				Value:   "FXF",
				Message: "invalid pattern",
				Cause:   errors.New("unknown character"),
			},
			expected: "validation error in field 'pattern': invalid pattern (caused by: unknown character)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			if result != tt.expected {
				t.Errorf("Error() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestValidationError_Unwrap(t *testing.T) {
	cause := errors.New("root cause")
	err := &ValidationError{
		Field:   "test",
		Message: "test error",
		Cause:   cause,
	}

	unwrapped := err.Unwrap()
	if unwrapped != cause {
		t.Errorf("Unwrap() = %v, want %v", unwrapped, cause)
	}

	// Test nil cause
	errNoCause := &ValidationError{
		Field:   "test",
		Message: "test error",
	}
	if errNoCause.Unwrap() != nil {
		t.Errorf("Unwrap() should return nil when no cause")
	}
}

func TestCharacterNotFoundError_Error(t *testing.T) {
	err := &CharacterNotFoundError{Name: "testChar"}
	expected := "character 'testChar' not found"

	result := err.Error()
	if result != expected {
		t.Errorf("Error() = %q, want %q", result, expected)
	}
}

func TestPatternCompilationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *PatternCompilationError
		expected string
	}{
		{
			name: "Error with position",
			err: &PatternCompilationError{
				Pattern:  "FXF",
				Position: 1,
				Message:  "invalid character",
			},
			expected: "pattern compilation error at position 1 in 'FXF': invalid character",
		},
		{
			name: "Error without position",
			err: &PatternCompilationError{
				Pattern:  "FXF",
				Position: -1,
				Message:  "invalid pattern",
			},
			expected: "pattern compilation error in 'FXF': invalid pattern",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			if result != tt.expected {
				t.Errorf("Error() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestPatternCompilationError_Unwrap(t *testing.T) {
	cause := errors.New("root cause")
	err := &PatternCompilationError{
		Pattern: "test",
		Message: "test error",
		Cause:   cause,
	}

	unwrapped := err.Unwrap()
	if unwrapped != cause {
		t.Errorf("Unwrap() = %v, want %v", unwrapped, cause)
	}
}

func TestAnimationError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *AnimationError
		expected string
	}{
		{
			name: "Error with cause",
			err: &AnimationError{
				CharacterName: "robot",
				Operation:     "start",
				Message:       "failed to start animation",
				Cause:         errors.New("no frames"),
			},
			expected: "animation error for character 'robot' during start: failed to start animation (caused by: no frames)",
		},
		{
			name: "Error without cause",
			err: &AnimationError{
				CharacterName: "robot",
				Operation:     "frame_display",
				Message:       "terminal too narrow",
			},
			expected: "animation error for character 'robot' during frame_display: terminal too narrow",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			if result != tt.expected {
				t.Errorf("Error() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestAnimationError_Unwrap(t *testing.T) {
	cause := errors.New("root cause")
	err := &AnimationError{
		CharacterName: "test",
		Operation:     "test",
		Message:       "test error",
		Cause:         cause,
	}

	unwrapped := err.Unwrap()
	if unwrapped != cause {
		t.Errorf("Unwrap() = %v, want %v", unwrapped, cause)
	}
}

func TestNewValidationError(t *testing.T) {
	err := NewValidationError("testField", "testValue", "test message")

	if err.Field != "testField" {
		t.Errorf("Field = %q, want %q", err.Field, "testField")
	}
	if err.Value != "testValue" {
		t.Errorf("Value = %v, want %v", err.Value, "testValue")
	}
	if err.Message != "test message" {
		t.Errorf("Message = %q, want %q", err.Message, "test message")
	}
	if err.Cause != nil {
		t.Errorf("Cause should be nil")
	}
}

func TestNewValidationErrorWithCause(t *testing.T) {
	cause := errors.New("root cause")
	err := NewValidationErrorWithCause("testField", "testValue", "test message", cause)

	if err.Field != "testField" {
		t.Errorf("Field = %q, want %q", err.Field, "testField")
	}
	if err.Value != "testValue" {
		t.Errorf("Value = %v, want %v", err.Value, "testValue")
	}
	if err.Message != "test message" {
		t.Errorf("Message = %q, want %q", err.Message, "test message")
	}
	if err.Cause != cause {
		t.Errorf("Cause = %v, want %v", err.Cause, cause)
	}
}

func TestNewCharacterNotFoundError(t *testing.T) {
	err := NewCharacterNotFoundError("myChar")

	if err.Name != "myChar" {
		t.Errorf("Name = %q, want %q", err.Name, "myChar")
	}
}

func TestNewPatternCompilationError(t *testing.T) {
	cause := errors.New("invalid char")
	err := NewPatternCompilationError("FXF", 1, "bad character", cause)

	if err.Pattern != "FXF" {
		t.Errorf("Pattern = %q, want %q", err.Pattern, "FXF")
	}
	if err.Position != 1 {
		t.Errorf("Position = %d, want %d", err.Position, 1)
	}
	if err.Message != "bad character" {
		t.Errorf("Message = %q, want %q", err.Message, "bad character")
	}
	if err.Cause != cause {
		t.Errorf("Cause = %v, want %v", err.Cause, cause)
	}
}

func TestNewAnimationError(t *testing.T) {
	cause := errors.New("no frames")
	err := NewAnimationError("robot", "start", "failed", cause)

	if err.CharacterName != "robot" {
		t.Errorf("CharacterName = %q, want %q", err.CharacterName, "robot")
	}
	if err.Operation != "start" {
		t.Errorf("Operation = %q, want %q", err.Operation, "start")
	}
	if err.Message != "failed" {
		t.Errorf("Message = %q, want %q", err.Message, "failed")
	}
	if err.Cause != cause {
		t.Errorf("Cause = %v, want %v", err.Cause, cause)
	}
}

func TestGetErrorSuggestion(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		contains string
	}{
		{
			name:     "ValidationError - name field",
			err:      &ValidationError{Field: "name", Message: "invalid"},
			contains: "letters, numbers, or underscores",
		},
		{
			name:     "ValidationError - dimensions field",
			err:      &ValidationError{Field: "dimensions", Message: "invalid"},
			contains: "positive integers",
		},
		{
			name:     "ValidationError - pattern field",
			err:      &ValidationError{Field: "pattern", Message: "invalid"},
			contains: "valid pattern characters",
		},
		{
			name:     "ValidationError - frames field",
			err:      &ValidationError{Field: "frames", Message: "invalid"},
			contains: "character height",
		},
		{
			name:     "ValidationError - frame_name field",
			err:      &ValidationError{Field: "frame_name", Message: "invalid"},
			contains: "descriptive frame names",
		},
		{
			name:     "ValidationError - unknown field",
			err:      &ValidationError{Field: "unknown", Message: "invalid"},
			contains: "Check the field value",
		},
		{
			name:     "CharacterNotFoundError",
			err:      &CharacterNotFoundError{Name: "testChar"},
			contains: "ListCharacters",
		},
		{
			name:     "PatternCompilationError with position",
			err:      &PatternCompilationError{Pattern: "FXF", Position: 1, Message: "bad"},
			contains: "position 1",
		},
		{
			name:     "PatternCompilationError without position",
			err:      &PatternCompilationError{Pattern: "FXF", Position: -1, Message: "bad"},
			contains: "valid pattern symbols",
		},
		{
			name:     "AnimationError - start operation",
			err:      &AnimationError{Operation: "start", Message: "failed"},
			contains: "frames",
		},
		{
			name:     "AnimationError - frame_display operation",
			err:      &AnimationError{Operation: "frame_display", Message: "failed"},
			contains: "terminal",
		},
		{
			name:     "AnimationError - timing operation",
			err:      &AnimationError{Operation: "timing", Message: "failed"},
			contains: "fps",
		},
		{
			name:     "AnimationError - unknown operation",
			err:      &AnimationError{Operation: "unknown", Message: "failed"},
			contains: "character data",
		},
		{
			name:     "Generic error",
			err:      errors.New("generic error"),
			contains: "Check the error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetErrorSuggestion(tt.err)
			if !strings.Contains(result, tt.contains) {
				t.Errorf("GetErrorSuggestion() = %q, should contain %q", result, tt.contains)
			}
		})
	}
}
