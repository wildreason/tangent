package state

import (
	"strings"
	"testing"
)

func TestValidateStateName(t *testing.T) {
	validator := NewFrameValidator()

	tests := []struct {
		name      string
		stateName string
		wantError bool
	}{
		{"valid lowercase", "arise", false},
		{"valid uppercase", "ARISE", false},
		{"valid mixed", "AriseState", false},
		{"valid with underscore", "arise_state", false},
		{"valid with numbers", "state123", false},
		{"empty name", "", true},
		{"reserved name base", "base", true},
		{"reserved name BASE", "BASE", true},
		{"invalid char hyphen", "arise-state", true},
		{"invalid char space", "arise state", true},
		{"invalid char special", "arise!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.validateStateName(tt.stateName)
			if (err != nil) != tt.wantError {
				t.Errorf("validateStateName(%q) error = %v, wantError %v", tt.stateName, err, tt.wantError)
			}
		})
	}
}

func TestValidateFrame(t *testing.T) {
	validator := NewFrameValidator()

	tests := []struct {
		name      string
		frame     FrameDefinition
		wantError bool
	}{
		{
			name: "valid frame",
			frame: FrameDefinition{
				Lines: []string{
					"___________",
					"__R5FFF6L__",
					"_26FFFFF51_",
					"___11_22___",
				},
			},
			wantError: false,
		},
		{
			name: "valid with lowercase",
			frame: FrameDefinition{
				Lines: []string{
					"___________",
					"__rfffffl__",
					"_26fffff51_",
					"___11_22___",
				},
			},
			wantError: false,
		},
		{
			name: "empty lines",
			frame: FrameDefinition{
				Lines: []string{},
			},
			wantError: true,
		},
		{
			name: "width mismatch",
			frame: FrameDefinition{
				Lines: []string{
					"___________",
					"__R5FFF6L__",
					"_26FFFFF51_",
					"___11_22_",  // Missing one char
				},
			},
			wantError: true,
		},
		{
			name: "invalid pattern code",
			frame: FrameDefinition{
				Lines: []string{
					"___________",
					"__R5ZZZ6L__",  // Z is invalid
					"_26FFFFF51_",
					"___11_22___",
				},
			},
			wantError: true,
		},
		{
			name: "valid with special chars",
			frame: FrameDefinition{
				Lines: []string{
					"___________",
					"__R5FFF6L__",
					"_26:...:51_",  // : and . are valid
					"__22___11__",
				},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateFrame(&tt.frame, 0)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateFrame() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestValidateDimensions(t *testing.T) {
	validator := NewFrameValidator()

	tests := []struct {
		name      string
		frames    []FrameDefinition
		width     int
		height    int
		wantError bool
	}{
		{
			name: "correct dimensions 11x4",
			frames: []FrameDefinition{
				{
					Lines: []string{
						"___________",
						"__R5FFF6L__",
						"_26FFFFF51_",
						"___11_22___",
					},
				},
			},
			width:     11,
			height:    4,
			wantError: false,
		},
		{
			name: "height mismatch",
			frames: []FrameDefinition{
				{
					Lines: []string{
						"___________",
						"__R5FFF6L__",
						"_26FFFFF51_",
					},
				},
			},
			width:     11,
			height:    4,
			wantError: true,
		},
		{
			name: "width mismatch",
			frames: []FrameDefinition{
				{
					Lines: []string{
						"__________",  // 10 chars
						"__R5FFF6L_",
						"_26FFFFF51",
						"___11_22__",
					},
				},
			},
			width:     11,
			height:    4,
			wantError: true,
		},
		{
			name: "multiple frames correct",
			frames: []FrameDefinition{
				{
					Lines: []string{
						"___________",
						"__R5FFF6L__",
						"_26FFFFF51_",
						"___11_22___",
					},
				},
				{
					Lines: []string{
						"___________",
						"__rfffffl__",
						"_26fffff51_",
						"___11_22___",
					},
				},
			},
			width:     11,
			height:    4,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateDimensions(tt.frames, tt.width, tt.height)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateDimensions() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestValidateStateConfig(t *testing.T) {
	validator := NewFrameValidator()

	tests := []struct {
		name      string
		config    StateConfig
		wantError bool
	}{
		{
			name: "valid config",
			config: StateConfig{
				Name:           "arise",
				Description:    "Awakening state",
				FrameCount:     2,
				RequiredHeight: 4,
				Targets:        []string{"fire"},
				Frames: []FrameDefinition{
					{
						Lines: []string{
							"___________",
							"__rfffffl__",
							"_26fffff51_",
							"___11_22___",
						},
					},
					{
						Lines: []string{
							"___________",
							"__r5ffffl__",
							"_26fffff51_",
							"___11_22___",
						},
					},
				},
			},
			wantError: false,
		},
		{
			name: "empty state name",
			config: StateConfig{
				Name:           "",
				FrameCount:     1,
				RequiredHeight: 4,
				Targets:        []string{"fire"},
				Frames: []FrameDefinition{
					{Lines: []string{"___________", "__R5FFF6L__", "_26FFFFF51_", "___11_22___"}},
				},
			},
			wantError: true,
		},
		{
			name: "zero frame count",
			config: StateConfig{
				Name:           "test",
				FrameCount:     0,
				RequiredHeight: 4,
				Targets:        []string{"fire"},
				Frames:         []FrameDefinition{},
			},
			wantError: true,
		},
		{
			name: "frame count mismatch",
			config: StateConfig{
				Name:           "test",
				FrameCount:     2,
				RequiredHeight: 4,
				Targets:        []string{"fire"},
				Frames: []FrameDefinition{
					{Lines: []string{"___________", "__R5FFF6L__", "_26FFFFF51_", "___11_22___"}},
				},
			},
			wantError: true,
		},
		{
			name: "no targets",
			config: StateConfig{
				Name:           "test",
				FrameCount:     1,
				RequiredHeight: 4,
				Targets:        []string{},
				Frames: []FrameDefinition{
					{Lines: []string{"___________", "__R5FFF6L__", "_26FFFFF51_", "___11_22___"}},
				},
			},
			wantError: true,
		},
		{
			name: "reserved name",
			config: StateConfig{
				Name:           "base",
				FrameCount:     1,
				RequiredHeight: 4,
				Targets:        []string{"fire"},
				Frames: []FrameDefinition{
					{Lines: []string{"___________", "__R5FFF6L__", "_26FFFFF51_", "___11_22___"}},
				},
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateStateConfig(&tt.config)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateStateConfig() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestValidateFramesConsistency(t *testing.T) {
	validator := NewFrameValidator()

	tests := []struct {
		name      string
		frames    []FrameDefinition
		wantError bool
	}{
		{
			name: "consistent frames",
			frames: []FrameDefinition{
				{
					Lines: []string{
						"___________",
						"__R5FFF6L__",
						"_26FFFFF51_",
						"___11_22___",
					},
				},
				{
					Lines: []string{
						"___________",
						"__rfffffl__",
						"_26fffff51_",
						"___11_22___",
					},
				},
			},
			wantError: false,
		},
		{
			name: "inconsistent height",
			frames: []FrameDefinition{
				{
					Lines: []string{
						"___________",
						"__R5FFF6L__",
						"_26FFFFF51_",
						"___11_22___",
					},
				},
				{
					Lines: []string{
						"___________",
						"__rfffffl__",
						"_26fffff51_",
					},
				},
			},
			wantError: true,
		},
		{
			name: "inconsistent width",
			frames: []FrameDefinition{
				{
					Lines: []string{
						"___________",
						"__R5FFF6L__",
						"_26FFFFF51_",
						"___11_22___",
					},
				},
				{
					Lines: []string{
						"__________",  // 10 chars
						"__rfffffl_",
						"_26fffff51",
						"___11_22__",
					},
				},
			},
			wantError: true,
		},
		{
			name:      "empty frames",
			frames:    []FrameDefinition{},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateFramesConsistency(tt.frames)
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateFramesConsistency() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestValidationErrorMessages(t *testing.T) {
	validator := NewFrameValidator()

	// Test that error messages are informative
	config := StateConfig{
		Name:           "test!invalid",
		FrameCount:     1,
		RequiredHeight: 4,
		Targets:        []string{"fire"},
		Frames: []FrameDefinition{
			{Lines: []string{"___________", "__R5FFF6L__", "_26FFFFF51_", "___11_22___"}},
		},
	}

	err := validator.ValidateStateConfig(&config)
	if err == nil {
		t.Fatal("expected error for invalid state name")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "Name") {
		t.Errorf("error message should mention field name, got: %s", errMsg)
	}
}

func TestPatternCodeValidation(t *testing.T) {
	validator := NewFrameValidator()

	validCodes := []string{
		"F", "R", "L", "B", "T",          // Shape codes (uppercase)
		"f", "r", "l", "b", "t",          // Shape codes (lowercase)
		"1", "2", "3", "4", "5", "6", "7", "8", // Quadrants
		"_", " ", ".", ":", "#",          // Special codes
		"/", "\\",                        // Diagonals
		"X",                              // Mirror
	}

	invalidCodes := []string{
		"0", "9", "Z", "z", "!", "@", "$", "%", "^", "&", "*", "(", ")", "-", "=",
	}

	// Test valid codes
	for _, code := range validCodes {
		line := "__" + code + code + code + code + code + code + code + "__"
		frame := FrameDefinition{Lines: []string{line}}
		err := validator.ValidateFrame(&frame, 0)
		if err != nil {
			t.Errorf("Valid code %q should not produce error, got: %v", code, err)
		}
	}

	// Test invalid codes
	for _, code := range invalidCodes {
		line := "__" + code + code + code + code + code + code + code + "__"
		frame := FrameDefinition{Lines: []string{line}}
		err := validator.ValidateFrame(&frame, 0)
		if err == nil {
			t.Errorf("Invalid code %q should produce error", code)
		}
	}
}
