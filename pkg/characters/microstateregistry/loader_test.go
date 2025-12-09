package microstateregistry

import (
	"testing"
)

func TestLoadEmbedded(t *testing.T) {
	def, err := LoadEmbedded()
	if err != nil {
		t.Fatalf("Failed to load embedded micro definition: %v", err)
	}

	if def == nil {
		t.Fatal("Loaded definition is nil")
	}

	// Test dimensions
	if def.Width != 10 {
		t.Errorf("Expected width 10, got %d", def.Width)
	}
	if def.Height != 2 {
		t.Errorf("Expected height 2, got %d", def.Height)
	}

	t.Logf("Loaded micro definition: width=%d, height=%d", def.Width, def.Height)
}

func TestDefaultDefinition(t *testing.T) {
	if DefaultDefinition == nil {
		t.Fatal("DefaultDefinition is nil")
	}

	states := ListStates()
	if len(states) < 7 {
		t.Errorf("Expected at least 7 states, got %d", len(states))
	}

	t.Logf("DefaultDefinition has %d states: %v", len(states), states)
}

func TestRequiredStates(t *testing.T) {
	requiredStates := []string{"resting", "arise", "wait", "read", "write", "search", "approval"}

	for _, name := range requiredStates {
		state := GetState(name)
		if state == nil {
			t.Errorf("Required state %q not found", name)
			continue
		}

		// Each state must have at least 2 frames
		if len(state.Frames) < 2 {
			t.Errorf("State %q has %d frames, expected at least 2", name, len(state.Frames))
		}

		// Each frame must have exactly 2 lines
		for i, frame := range state.Frames {
			if len(frame.Lines) != 2 {
				t.Errorf("State %q frame %d has %d lines, expected 2", name, i, len(frame.Lines))
			}
		}
	}
}

func TestBaseFrame(t *testing.T) {
	base := GetBaseFrame()
	if len(base.Lines) != 2 {
		t.Errorf("Base frame has %d lines, expected 2", len(base.Lines))
	}

	// Each line should be 10 chars wide
	for i, line := range base.Lines {
		if len(line) != 10 {
			t.Errorf("Base frame line %d has %d chars, expected 10", i, len(line))
		}
	}
}

func TestDimensions(t *testing.T) {
	if Width() != 10 {
		t.Errorf("Width() returned %d, expected 10", Width())
	}
	if Height() != 2 {
		t.Errorf("Height() returned %d, expected 2", Height())
	}
}
