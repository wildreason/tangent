package characters

import (
	"testing"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

func TestCharacterBuilderV2_NewCharacterBuilderV2(t *testing.T) {
	tests := []struct {
		name        string
		charName    string
		width       int
		height      int
		shouldPanic bool
	}{
		{"Valid builder", "test-char", 5, 3, false},
		{"Empty name", "", 5, 3, true},
		{"Zero width", "test-char", 0, 3, true},
		{"Zero height", "test-char", 5, 0, true},
		{"Negative width", "test-char", -1, 3, true},
		{"Negative height", "test-char", 5, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.shouldPanic {
						t.Errorf("NewCharacterBuilder() panicked unexpectedly: %v", r)
					}
				} else if tt.shouldPanic {
					t.Errorf("NewCharacterBuilder() should have panicked but didn't")
				}
			}()

			builder := NewCharacterBuilderV2(tt.charName, tt.width, tt.height)
			if !tt.shouldPanic {
				if builder.spec.Name != tt.charName {
					t.Errorf("Expected name %s, got %s", tt.charName, builder.spec.Name)
				}
				if builder.spec.Width != tt.width {
					t.Errorf("Expected width %d, got %d", tt.width, builder.spec.Width)
				}
				if builder.spec.Height != tt.height {
					t.Errorf("Expected height %d, got %d", tt.height, builder.spec.Height)
				}
			}
		})
	}
}

func TestCharacterBuilderV2_AddFrame(t *testing.T) {
	builder := NewCharacterBuilderV2("test", 5, 3)

	tests := []struct {
		name        string
		frameName   string
		patterns    []string
		shouldPanic bool
	}{
		{"Valid frame", "idle", []string{"FRF", "LRL", "FRF"}, false},
		{"Empty frame name", "", []string{"FRF", "LRL", "FRF"}, true},
		{"Wrong pattern count", "idle", []string{"FRF", "LRL"}, true},
		{"Too many patterns", "idle", []string{"FRF", "LRL", "FRF", "EXTRA"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.shouldPanic {
						t.Errorf("AddFrame() panicked unexpectedly: %v", r)
					}
				} else if tt.shouldPanic {
					t.Errorf("AddFrame() should have panicked but didn't")
				}
			}()

			result := builder.AddFrame(tt.frameName, tt.patterns)
			if !tt.shouldPanic {
				if result != builder {
					t.Errorf("AddFrame() should return the same builder instance")
				}
				if len(builder.spec.Frames) == 0 {
					t.Errorf("Frame should have been added")
				}
			}
		})
	}
}

func TestCharacterBuilderV2_AddFrameFromString(t *testing.T) {
	builder := NewCharacterBuilderV2("test", 5, 3)

	tests := []struct {
		name        string
		frameName   string
		pattern     string
		shouldPanic bool
	}{
		{"Valid pattern", "idle", "FRF\nLRL\nFRF", false},
		{"Empty frame name", "", "FRF\nLRL\nFRF", true},
		{"Empty pattern", "idle", "", true},
		{"Wrong line count", "idle", "FRF\nLRL", true},
		{"Too many lines", "idle", "FRF\nLRL\nFRF\nEXTRA", true},
		{"With empty lines", "idle", "FRF\n\nLRL\nFRF", true}, // Should panic on empty lines
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.shouldPanic {
						t.Errorf("AddFrameFromString() panicked unexpectedly: %v", r)
					}
				} else if tt.shouldPanic {
					t.Errorf("AddFrameFromString() should have panicked but didn't")
				}
			}()

			result := builder.AddFrameFromString(tt.frameName, tt.pattern)
			if !tt.shouldPanic {
				if result != builder {
					t.Errorf("AddFrameFromString() should return the same builder instance")
				}
			}
		})
	}
}

func TestCharacterBuilderV2_Build(t *testing.T) {
	// Create a builder and test basic functionality
	builder := NewCharacterBuilderV2("test", 5, 3)
	builder.AddFrame("idle", []string{"FRF", "LRL", "FRF"})

	// Build the character (this will use the default service)
	character, err := builder.Build()
	if err != nil {
		t.Errorf("Build() error = %v", err)
	}

	if character == nil {
		t.Errorf("Build() returned nil character")
		return
	}

	if character.Name != "test" {
		t.Errorf("Expected character name 'test', got %s", character.Name)
	}
	if character.Width != 5 {
		t.Errorf("Expected character width 5, got %d", character.Width)
	}
	if character.Height != 3 {
		t.Errorf("Expected character height 3, got %d", character.Height)
	}
	if len(character.Frames) != 1 {
		t.Errorf("Expected 1 frame, got %d", len(character.Frames))
	}
	if character.Frames[0].Name != "idle" {
		t.Errorf("Expected frame name 'idle', got %s", character.Frames[0].Name)
	}
}

func TestCharacterBuilderV2_Build_NoFrames(t *testing.T) {
	builder := NewCharacterBuilderV2("test", 5, 3)
	// Don't add any frames
	_, err := builder.Build()
	if err == nil {
		t.Errorf("Build() should return error when no frames are added")
	}
}

func TestCharacterBuilderV2_BuildAndSave(t *testing.T) {
	builder := NewCharacterBuilderV2("test", 5, 3)
	builder.AddFrame("idle", []string{"FRF", "LRL", "FRF"})

	// Build and save the character
	character, err := builder.BuildAndSave()
	if err != nil {
		t.Errorf("BuildAndSave() error = %v", err)
	}

	if character == nil {
		t.Errorf("BuildAndSave() returned nil character")
		return
	}

	if character.Name != "test" {
		t.Errorf("Expected character name 'test', got %s", character.Name)
	}
}

func TestCharacterBuilderV2_Validate(t *testing.T) {
	tests := []struct {
		name      string
		setupFunc func(*CharacterBuilderV2)
		wantErr   bool
	}{
		{
			name: "Valid character",
			setupFunc: func(b *CharacterBuilderV2) {
				b.AddFrame("idle", []string{"FRF", "LRL", "FRF"})
			},
			wantErr: false,
		},
		{
			name: "No frames",
			setupFunc: func(b *CharacterBuilderV2) {
				// Don't add any frames
			},
			wantErr: true,
		},
		{
			name: "Empty frame name",
			setupFunc: func(b *CharacterBuilderV2) {
				b.spec.Frames = append(b.spec.Frames, domain.FrameSpec{
					Name:     "",
					Patterns: []string{"FRF", "LRL", "FRF"},
				})
			},
			wantErr: true,
		},
		{
			name: "Wrong pattern count",
			setupFunc: func(b *CharacterBuilderV2) {
				b.spec.Frames = append(b.spec.Frames, domain.FrameSpec{
					Name:     "idle",
					Patterns: []string{"FRF", "LRL"}, // Only 2 patterns, should be 3
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewCharacterBuilderV2("test", 5, 3)
			tt.setupFunc(builder)

			err := builder.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCharacterBuilderV2_GetSpec(t *testing.T) {
	builder := NewCharacterBuilderV2("test", 5, 3)
	builder.AddFrame("idle", []string{"FRF", "LRL", "FRF"})

	spec := builder.GetSpec()
	if spec == nil {
		t.Errorf("GetSpec() returned nil")
		return
	}

	if spec.Name != "test" {
		t.Errorf("Expected spec name 'test', got %s", spec.Name)
	}
	if spec.Width != 5 {
		t.Errorf("Expected spec width 5, got %d", spec.Width)
	}
	if spec.Height != 3 {
		t.Errorf("Expected spec height 3, got %d", spec.Height)
	}
	if len(spec.Frames) != 1 {
		t.Errorf("Expected 1 frame in spec, got %d", len(spec.Frames))
	}
}

func TestCharacterBuilderV2_SetService(t *testing.T) {
	builder := NewCharacterBuilderV2("test", 5, 3)
	// Test that SetService returns the same builder instance
	result := builder.SetService(builder.service)
	if result != builder {
		t.Errorf("SetService() should return the same builder instance")
	}
}

func TestCharacterBuilderV2_FluentAPI(t *testing.T) {
	// Test fluent API chaining
	character, err := NewCharacterBuilderV2("robot", 7, 4).
		AddFrame("idle", []string{"FRF", "LRL", "FRF", "LRL"}).
		AddFrame("wave", []string{"FRF", "LRL", "FRF", "LRL"}).
		Build()

	if err != nil {
		t.Errorf("Fluent API Build() error = %v", err)
	}

	if character == nil {
		t.Errorf("Fluent API Build() returned nil character")
		return
	}

	if character.Name != "robot" {
		t.Errorf("Expected character name 'robot', got %s", character.Name)
	}
	if len(character.Frames) != 2 {
		t.Errorf("Expected 2 frames, got %d", len(character.Frames))
	}
}
