package characters

import (
	"bytes"
	"testing"
)

func TestLibraryAgentMicro_AllCharacters(t *testing.T) {
	characterNames := []string{"sam", "rio", "ga", "ma", "pa", "da", "ni"}

	for _, name := range characterNames {
		t.Run(name, func(t *testing.T) {
			agent, err := LibraryAgentMicro(name)
			if err != nil {
				t.Fatalf("LibraryAgentMicro(%q) error = %v", name, err)
			}

			if agent == nil {
				t.Fatalf("LibraryAgentMicro(%q) returned nil", name)
			}

			// Verify dimensions
			char := agent.GetCharacter()
			if char.Width != 8 {
				t.Errorf("Width = %d, expected 8", char.Width)
			}
			if char.Height != 2 {
				t.Errorf("Height = %d, expected 2", char.Height)
			}

			t.Logf("%s-micro: width=%d, height=%d", name, char.Width, char.Height)
		})
	}
}

func TestLibraryAgentMicro_RequiredStates(t *testing.T) {
	requiredStates := []string{"resting", "arise", "wait", "read", "write", "search", "approval"}

	agent, err := LibraryAgentMicro("sam")
	if err != nil {
		t.Fatalf("LibraryAgentMicro(sam) error = %v", err)
	}

	for _, stateName := range requiredStates {
		if !agent.HasState(stateName) {
			t.Errorf("Missing required state: %s", stateName)
		}
	}

	t.Logf("sam-micro has states: %v", agent.ListStates())
}

func TestLibraryAgentMicro_StateRendering(t *testing.T) {
	agent, err := LibraryAgentMicro("sam")
	if err != nil {
		t.Fatalf("LibraryAgentMicro(sam) error = %v", err)
	}

	// Test rendering each state
	states := agent.ListStates()
	for _, stateName := range states {
		t.Run(stateName, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := agent.ShowState(buf, stateName)
			if err != nil {
				t.Errorf("ShowState(%q) error = %v", stateName, err)
				return
			}

			output := buf.String()
			if output == "" {
				t.Errorf("ShowState(%q) produced no output", stateName)
			}
		})
	}
}

func TestLibraryAgentMicro_NotFound(t *testing.T) {
	_, err := LibraryAgentMicro("nonexistent")
	if err == nil {
		t.Error("LibraryAgentMicro(nonexistent) should return error")
	}
}

func TestListMicroLibrary(t *testing.T) {
	names := ListMicroLibrary()

	if len(names) != 7 {
		t.Errorf("ListMicroLibrary() returned %d names, expected 7", len(names))
	}

	// Check all expected characters are present
	expected := map[string]bool{
		"sam-micro": true,
		"rio-micro": true,
		"ga-micro":  true,
		"ma-micro":  true,
		"pa-micro":  true,
		"da-micro":  true,
		"ni-micro":  true,
	}

	for _, name := range names {
		if !expected[name] {
			t.Errorf("Unexpected micro character: %s", name)
		}
		delete(expected, name)
	}

	for name := range expected {
		t.Errorf("Missing micro character: %s", name)
	}

	t.Logf("Micro library: %v", names)
}

func TestLibraryAgentMicro_FrameCache(t *testing.T) {
	agent, err := LibraryAgentMicro("sam")
	if err != nil {
		t.Fatalf("LibraryAgentMicro(sam) error = %v", err)
	}

	cache := agent.GetFrameCache()
	if cache == nil {
		t.Fatal("GetFrameCache() returned nil")
	}

	// Test base frame
	baseFrame := cache.GetBaseFrame()
	if len(baseFrame) != 2 {
		t.Errorf("Base frame has %d lines, expected 2", len(baseFrame))
	}

	// Test state frames
	restingFrames := cache.GetStateFrames("resting")
	if len(restingFrames) < 2 {
		t.Errorf("resting state has %d frames, expected at least 2", len(restingFrames))
	}

	// Each frame should have 2 lines
	for i, frame := range restingFrames {
		if len(frame) != 2 {
			t.Errorf("resting frame %d has %d lines, expected 2", i, len(frame))
		}
	}
}
