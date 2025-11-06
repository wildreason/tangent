package characters

import (
	"strings"
	"testing"
)

func TestNewCharacterSpec(t *testing.T) {
	spec := NewCharacterSpec("robot", 8, 4)

	if spec.Name != "robot" {
		t.Errorf("Name = %q, want %q", spec.Name, "robot")
	}
	if spec.Width != 8 {
		t.Errorf("Width = %d, want %d", spec.Width, 8)
	}
	if spec.Height != 4 {
		t.Errorf("Height = %d, want %d", spec.Height, 4)
	}
	if spec.Frames == nil {
		t.Error("Frames should be initialized")
	}
	if len(spec.Frames) != 0 {
		t.Errorf("Frames length = %d, want %d", len(spec.Frames), 0)
	}
}

func TestAddFrame(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)
	patterns := []string{"FFFF", "BBBB"}

	result := spec.AddFrame("idle", patterns)

	// Should return self for chaining
	if result != spec {
		t.Error("AddFrame should return self for chaining")
	}

	if len(spec.Frames) != 1 {
		t.Errorf("Frames length = %d, want %d", len(spec.Frames), 1)
	}

	frame := spec.Frames[0]
	if frame.Name != "idle" {
		t.Errorf("Frame name = %q, want %q", frame.Name, "idle")
	}
	if len(frame.Patterns) != 2 {
		t.Errorf("Frame patterns length = %d, want %d", len(frame.Patterns), 2)
	}
	if frame.Patterns[0] != "FFFF" {
		t.Errorf("Pattern[0] = %q, want %q", frame.Patterns[0], "FFFF")
	}
}

func TestAddFrame_Chaining(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)

	spec.
		AddFrame("idle", []string{"FFFF", "BBBB"}).
		AddFrame("wave", []string{"FRFF", "BBBB"})

	if len(spec.Frames) != 2 {
		t.Errorf("Frames length = %d, want %d", len(spec.Frames), 2)
	}
	if spec.Frames[0].Name != "idle" {
		t.Errorf("Frame[0] name = %q, want %q", spec.Frames[0].Name, "idle")
	}
	if spec.Frames[1].Name != "wave" {
		t.Errorf("Frame[1] name = %q, want %q", spec.Frames[1].Name, "wave")
	}
}

func TestAddFrameFromString(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 3)
	pattern := `FFFF
TTTT
BBBB`

	result := spec.AddFrameFromString("idle", pattern)

	// Should return self for chaining
	if result != spec {
		t.Error("AddFrameFromString should return self for chaining")
	}

	if len(spec.Frames) != 1 {
		t.Errorf("Frames length = %d, want %d", len(spec.Frames), 1)
	}

	frame := spec.Frames[0]
	if frame.Name != "idle" {
		t.Errorf("Frame name = %q, want %q", frame.Name, "idle")
	}
	if len(frame.Patterns) != 3 {
		t.Errorf("Frame patterns length = %d, want %d", len(frame.Patterns), 3)
	}
	if frame.Patterns[0] != "FFFF" {
		t.Errorf("Pattern[0] = %q, want %q", frame.Patterns[0], "FFFF")
	}
	if frame.Patterns[1] != "TTTT" {
		t.Errorf("Pattern[1] = %q, want %q", frame.Patterns[1], "TTTT")
	}
	if frame.Patterns[2] != "BBBB" {
		t.Errorf("Pattern[2] = %q, want %q", frame.Patterns[2], "BBBB")
	}
}

func TestAddFrameFromString_WithWhitespace(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)
	pattern := `
	FFFF
	BBBB
	`

	spec.AddFrameFromString("idle", pattern)

	frame := spec.Frames[0]
	if len(frame.Patterns) != 2 {
		t.Errorf("Frame patterns length = %d, want %d (whitespace should be trimmed)", len(frame.Patterns), 2)
	}
}

func TestAddFrameFromString_EmptyLines(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)
	pattern := `FFFF

BBBB`

	spec.AddFrameFromString("idle", pattern)

	frame := spec.Frames[0]
	if len(frame.Patterns) != 2 {
		t.Errorf("Frame patterns length = %d, want %d (empty lines should be skipped)", len(frame.Patterns), 2)
	}
}

func TestBuild(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, 2)
	spec.AddFrame("idle", []string{"FFFF", "BBBB"})

	character, err := spec.Build()

	if err != nil {
		t.Fatalf("Build() error = %v, want nil", err)
	}

	if character.Name != "robot" {
		t.Errorf("Character name = %q, want %q", character.Name, "robot")
	}
	if character.Width != 4 {
		t.Errorf("Character width = %d, want %d", character.Width, 4)
	}
	if character.Height != 2 {
		t.Errorf("Character height = %d, want %d", character.Height, 2)
	}
	if len(character.Frames) != 1 {
		t.Errorf("Character frames length = %d, want %d", len(character.Frames), 1)
	}

	// Check that patterns were compiled
	frame := character.Frames[0]
	if frame.Name != "idle" {
		t.Errorf("Frame name = %q, want %q", frame.Name, "idle")
	}
	if len(frame.Lines) != 2 {
		t.Errorf("Frame lines length = %d, want %d", len(frame.Lines), 2)
	}
	// F should be compiled to █
	if !strings.Contains(frame.Lines[0], "█") {
		t.Errorf("Frame line should contain compiled pattern, got %q", frame.Lines[0])
	}
}

func TestBuild_MultipleFrames(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, 2)
	spec.
		AddFrame("idle", []string{"FFFF", "BBBB"}).
		AddFrame("wave", []string{"FRFF", "BBBB"})

	character, err := spec.Build()

	if err != nil {
		t.Fatalf("Build() error = %v, want nil", err)
	}

	if len(character.Frames) != 2 {
		t.Errorf("Character frames length = %d, want %d", len(character.Frames), 2)
	}
	if character.Frames[0].Name != "idle" {
		t.Errorf("Frame[0] name = %q, want %q", character.Frames[0].Name, "idle")
	}
	if character.Frames[1].Name != "wave" {
		t.Errorf("Frame[1] name = %q, want %q", character.Frames[1].Name, "wave")
	}
}

func TestString(t *testing.T) {
	spec := NewCharacterSpec("robot", 8, 4)
	spec.
		AddFrame("idle", []string{"FFFFFFFF", "LLRRRRRR", "LLRRRRRR", "BBBBBBBB"}).
		AddFrame("wave", []string{"FFFFRFFF", "LLRRRRRR", "LLRRRRRR", "BBBBBBBB"})

	result := spec.String()

	// Verify the output contains expected information
	expectedStrings := []string{
		"Character: robot (8x4)",
		"Frame 1 (idle):",
		"Frame 2 (wave):",
		"Line 1: FFFFFFFF",
		"Line 2: LLRRRRRR",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(result, expected) {
			t.Errorf("String() should contain %q, got:\n%s", expected, result)
		}
	}
}

func TestValidate_Success(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, 2)
	spec.AddFrame("idle", []string{"FFFF", "BBBB"})

	err := spec.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v, want nil", err)
	}
}

func TestValidate_EmptyName(t *testing.T) {
	spec := NewCharacterSpec("", 4, 2)
	spec.AddFrame("idle", []string{"FFFF", "BBBB"})

	err := spec.Validate()
	if err == nil {
		t.Error("Validate() should return error for empty name")
	}
	if !strings.Contains(err.Error(), "name cannot be empty") {
		t.Errorf("Error should mention empty name, got: %v", err)
	}
}

func TestValidate_InvalidWidth(t *testing.T) {
	spec := NewCharacterSpec("robot", 0, 2)
	spec.AddFrame("idle", []string{"FFFF", "BBBB"})

	err := spec.Validate()
	if err == nil {
		t.Error("Validate() should return error for invalid width")
	}
	if !strings.Contains(err.Error(), "width and height must be positive") {
		t.Errorf("Error should mention positive dimensions, got: %v", err)
	}
}

func TestValidate_InvalidHeight(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, -1)
	spec.AddFrame("idle", []string{"FFFF", "BBBB"})

	err := spec.Validate()
	if err == nil {
		t.Error("Validate() should return error for invalid height")
	}
	if !strings.Contains(err.Error(), "width and height must be positive") {
		t.Errorf("Error should mention positive dimensions, got: %v", err)
	}
}

func TestValidate_NoFrames(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, 2)

	err := spec.Validate()
	if err == nil {
		t.Error("Validate() should return error for no frames")
	}
	if !strings.Contains(err.Error(), "at least one frame") {
		t.Errorf("Error should mention missing frames, got: %v", err)
	}
}

func TestValidate_EmptyFrameName(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, 2)
	spec.Frames = append(spec.Frames, FrameSpec{
		Name:     "",
		Patterns: []string{"FFFF", "BBBB"},
	})

	err := spec.Validate()
	if err == nil {
		t.Error("Validate() should return error for empty frame name")
	}
	if !strings.Contains(err.Error(), "name cannot be empty") {
		t.Errorf("Error should mention empty frame name, got: %v", err)
	}
}

func TestValidate_NoPatterns(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, 2)
	spec.Frames = append(spec.Frames, FrameSpec{
		Name:     "idle",
		Patterns: []string{},
	})

	err := spec.Validate()
	if err == nil {
		t.Error("Validate() should return error for no patterns")
	}
	if !strings.Contains(err.Error(), "at least one pattern") {
		t.Errorf("Error should mention missing patterns, got: %v", err)
	}
}

func TestValidate_PatternHeightMismatch(t *testing.T) {
	spec := NewCharacterSpec("robot", 4, 3)
	spec.AddFrame("idle", []string{"FFFF", "BBBB"}) // Only 2 patterns, need 3

	err := spec.Validate()
	if err == nil {
		t.Error("Validate() should return error for pattern height mismatch")
	}
	if !strings.Contains(err.Error(), "has 2 patterns, expected 3") {
		t.Errorf("Error should mention pattern count mismatch, got: %v", err)
	}
}
