package characters

import (
	"strings"
	"testing"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

func TestErrorHandler_HandleError(t *testing.T) {
	handler := NewErrorHandler()

	tests := []struct {
		name    string
		context string
		err     error
	}{
		{
			name:    "Validation error - name",
			context: "Create character",
			err:     domain.ErrCharacterNameRequired,
		},
		{
			name:    "Validation error - dimensions",
			context: "Create character",
			err:     domain.ErrInvalidDimensions,
		},
		{
			name:    "Character not found",
			context: "Load character",
			err:     domain.NewCharacterNotFoundError("missing-char"),
		},
		{
			name:    "Pattern compilation error",
			context: "Compile pattern",
			err:     domain.NewPatternCompilationError("FXZ", 1, "invalid character", nil),
		},
		{
			name:    "Animation error",
			context: "Animate character",
			err:     domain.NewAnimationError("test-char", "start", "no frames", nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test mainly ensures the function doesn't panic
			// In a real test, you'd capture output and verify it
			handler.HandleError(tt.context, tt.err)
		})
	}
}

func TestFormatError(t *testing.T) {
	tests := []struct {
		name    string
		context string
		err     error
		want    string
	}{
		{
			name:    "Validation error",
			context: "Test context",
			err:     domain.ErrCharacterNameRequired,
			want:    "✗ Test context: validation error in field 'name': character name is required and cannot be empty",
		},
		{
			name:    "Character not found",
			context: "Load character",
			err:     domain.NewCharacterNotFoundError("missing"),
			want:    "✗ Load character: character 'missing' not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatError(tt.context, tt.err)
			if !strings.Contains(result, tt.want) {
				t.Errorf("FormatError() = %v, want to contain %v", result, tt.want)
			}
		})
	}
}

func TestErrorTypeChecks(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		isValid    bool
		isNotFound bool
		isPattern  bool
		isAnim     bool
	}{
		{
			name:       "Validation error",
			err:        domain.ErrCharacterNameRequired,
			isValid:    true,
			isNotFound: false,
			isPattern:  false,
			isAnim:     false,
		},
		{
			name:       "Character not found",
			err:        domain.NewCharacterNotFoundError("test"),
			isValid:    false,
			isNotFound: true,
			isPattern:  false,
			isAnim:     false,
		},
		{
			name:       "Pattern compilation error",
			err:        domain.NewPatternCompilationError("FXZ", 1, "invalid", nil),
			isValid:    false,
			isNotFound: false,
			isPattern:  true,
			isAnim:     false,
		},
		{
			name:       "Animation error",
			err:        domain.NewAnimationError("test", "start", "error", nil),
			isValid:    false,
			isNotFound: false,
			isPattern:  false,
			isAnim:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if IsValidationError(tt.err) != tt.isValid {
				t.Errorf("IsValidationError() = %v, want %v", IsValidationError(tt.err), tt.isValid)
			}
			if IsCharacterNotFoundError(tt.err) != tt.isNotFound {
				t.Errorf("IsCharacterNotFoundError() = %v, want %v", IsCharacterNotFoundError(tt.err), tt.isNotFound)
			}
			if IsPatternCompilationError(tt.err) != tt.isPattern {
				t.Errorf("IsPatternCompilationError() = %v, want %v", IsPatternCompilationError(tt.err), tt.isPattern)
			}
			if IsAnimationError(tt.err) != tt.isAnim {
				t.Errorf("IsAnimationError() = %v, want %v", IsAnimationError(tt.err), tt.isAnim)
			}
		})
	}
}

func TestGetErrorField(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "Validation error with field",
			err:  domain.ErrCharacterNameRequired,
			want: "name",
		},
		{
			name: "Non-validation error",
			err:  domain.NewCharacterNotFoundError("test"),
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetErrorField(tt.err); got != tt.want {
				t.Errorf("GetErrorField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetErrorValue(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want interface{}
	}{
		{
			name: "Validation error with value",
			err:  domain.NewValidationError("test", "value", "message"),
			want: "value",
		},
		{
			name: "Non-validation error",
			err:  domain.NewCharacterNotFoundError("test"),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetErrorValue(tt.err); got != tt.want {
				t.Errorf("GetErrorValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
