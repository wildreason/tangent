package infrastructure

import (
	"testing"
)

func TestPatternCompiler(t *testing.T) {
	compiler := NewPatternCompiler()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Single F", "F", "█"},
		{"Single R", "R", "▐"},
		{"Single L", "L", "▌"},
		{"Single T", "T", "▀"},
		{"Single B", "B", "▄"},
		{"Multiple chars", "FRF", "█▐█"},
		{"Mixed pattern", "F6RTL", "█▜▐▀▌"},
		{"With spaces", "F_R_F", "█ ▐ █"},
		{"Unknown chars", "FXF", "█X█"},
		{"Empty string", "", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := compiler.Compile(test.input)
			if result != test.expected {
				t.Errorf("expected %q, got %q", test.expected, result)
			}
		})
	}
}

func TestPatternValidation(t *testing.T) {
	compiler := NewPatternCompiler()

	tests := []struct {
		name    string
		pattern string
		wantErr bool
	}{
		{"Valid pattern", "FRF", false},
		{"Empty pattern", "", true},
		{"Valid single char", "F", false},
		{"Valid with spaces", "F_R_F", false},
		{"Valid with unknown chars", "FXF", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := compiler.Validate(test.pattern)
			if (err != nil) != test.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestLengthValidator(t *testing.T) {
	validator := &LengthValidator{}

	tests := []struct {
		name    string
		pattern string
		wantErr bool
	}{
		{"Non-empty pattern", "FRF", false},
		{"Empty pattern", "", true},
		{"Single char", "F", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validator.Validate(test.pattern)
			if (err != nil) != test.wantErr {
				t.Errorf("LengthValidator.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestCharacterValidator(t *testing.T) {
	validator := &CharacterValidator{}

	tests := []struct {
		name    string
		pattern string
		wantErr bool
	}{
		{"Valid pattern", "FRF", false},
		{"Empty pattern", "", false},
		{"Pattern with unknown chars", "FXF", false},
		{"Pattern with spaces", "F R F", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validator.Validate(test.pattern)
			if (err != nil) != test.wantErr {
				t.Errorf("CharacterValidator.Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
