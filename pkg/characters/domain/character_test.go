package domain

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
		t.Errorf("Frames length = %d, want 0", len(spec.Frames))
	}
}

func TestCharacterSpec_AddFrame(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)
	patterns := []string{"FFFF", "BBBB"}

	result := spec.AddFrame("idle", patterns)

	// Should return self for chaining
	if result != spec {
		t.Error("AddFrame should return self for chaining")
	}

	if len(spec.Frames) != 1 {
		t.Errorf("Frames length = %d, want 1", len(spec.Frames))
	}

	frame := spec.Frames[0]
	if frame.Name != "idle" {
		t.Errorf("Frame name = %q, want %q", frame.Name, "idle")
	}
	if len(frame.Patterns) != 2 {
		t.Errorf("Frame patterns length = %d, want 2", len(frame.Patterns))
	}
	if frame.Patterns[0] != "FFFF" {
		t.Errorf("Pattern[0] = %q, want %q", frame.Patterns[0], "FFFF")
	}
	if frame.Patterns[1] != "BBBB" {
		t.Errorf("Pattern[1] = %q, want %q", frame.Patterns[1], "BBBB")
	}
}

func TestCharacterSpec_AddFrame_Multiple(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)

	spec.
		AddFrame("idle", []string{"FFFF", "BBBB"}).
		AddFrame("wave", []string{"FRFF", "BBBB"})

	if len(spec.Frames) != 2 {
		t.Errorf("Frames length = %d, want 2", len(spec.Frames))
	}

	if spec.Frames[0].Name != "idle" {
		t.Errorf("Frame[0] name = %q, want %q", spec.Frames[0].Name, "idle")
	}
	if spec.Frames[1].Name != "wave" {
		t.Errorf("Frame[1] name = %q, want %q", spec.Frames[1].Name, "wave")
	}
}

func TestCharacterSpec_AddFrameFromString(t *testing.T) {
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
		t.Errorf("Frames length = %d, want 1", len(spec.Frames))
	}

	frame := spec.Frames[0]
	if frame.Name != "idle" {
		t.Errorf("Frame name = %q, want %q", frame.Name, "idle")
	}
	if len(frame.Patterns) != 3 {
		t.Errorf("Frame patterns length = %d, want 3", len(frame.Patterns))
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

func TestCharacterSpec_AddFrameFromString_WithWhitespace(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)
	pattern := `
	FFFF
	BBBB
	`

	spec.AddFrameFromString("idle", pattern)

	frame := spec.Frames[0]
	if len(frame.Patterns) != 2 {
		t.Errorf("Frame patterns length = %d, want 2 (whitespace should be trimmed)", len(frame.Patterns))
	}
	if frame.Patterns[0] != "FFFF" {
		t.Errorf("Pattern[0] = %q, want %q (should be trimmed)", frame.Patterns[0], "FFFF")
	}
}

func TestCharacterSpec_AddFrameFromString_EmptyLines(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)
	pattern := `FFFF

BBBB`

	spec.AddFrameFromString("idle", pattern)

	frame := spec.Frames[0]
	if len(frame.Patterns) != 2 {
		t.Errorf("Frame patterns length = %d, want 2 (empty lines should be skipped)", len(frame.Patterns))
	}
}

func TestCharacterSpec_AddFrameFromString_OnlyNewlines(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 2)
	pattern := "\n\n\n"

	spec.AddFrameFromString("empty", pattern)

	if len(spec.Frames) != 1 {
		t.Errorf("Should still add frame even with empty pattern")
	}

	frame := spec.Frames[0]
	if len(frame.Patterns) != 0 {
		t.Errorf("Pattern should be empty when only newlines provided, got %d patterns", len(frame.Patterns))
	}
}

func TestCharacterSpec_AddFrameFromString_SingleLine(t *testing.T) {
	spec := NewCharacterSpec("test", 8, 1)
	pattern := "FFFFFFFF"

	spec.AddFrameFromString("single", pattern)

	frame := spec.Frames[0]
	if len(frame.Patterns) != 1 {
		t.Errorf("Frame patterns length = %d, want 1", len(frame.Patterns))
	}
	if frame.Patterns[0] != "FFFFFFFF" {
		t.Errorf("Pattern = %q, want %q", frame.Patterns[0], "FFFFFFFF")
	}
}

func TestCharacterSpec_AddFrameFromString_MixedWhitespace(t *testing.T) {
	spec := NewCharacterSpec("test", 4, 3)
	pattern := "  FFFF  \n\tTTTT\t\n   BBBB   "

	spec.AddFrameFromString("mixed", pattern)

	frame := spec.Frames[0]
	if len(frame.Patterns) != 3 {
		t.Errorf("Frame patterns length = %d, want 3", len(frame.Patterns))
	}

	// All patterns should be trimmed
	for i, expected := range []string{"FFFF", "TTTT", "BBBB"} {
		if !strings.Contains(frame.Patterns[i], expected) {
			t.Errorf("Pattern[%d] = %q, should contain %q", i, frame.Patterns[i], expected)
		}
	}
}
