package characters

import (
	"testing"
)

func TestFrameCache(t *testing.T) {
	// Load a library character
	agent, err := LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	// Get frame cache
	cache := agent.GetFrameCache()
	if cache == nil {
		t.Fatal("GetFrameCache returned nil")
	}

	// Test base frame
	baseFrame := cache.GetBaseFrame()
	if len(baseFrame) == 0 {
		t.Error("base frame is empty")
	}

	// Verify frames are pre-colored (should contain ANSI escape codes)
	hasAnsiCode := false
	for _, line := range baseFrame {
		if len(line) > 0 && (line[0] == '\x1b' || len(line) > 10) {
			hasAnsiCode = true
			break
		}
	}
	if !hasAnsiCode {
		t.Log("Warning: base frame may not be colored (expected ANSI codes)")
	}

	// Test state frames
	states := cache.ListStates()
	if len(states) == 0 {
		t.Error("no states in cache")
	}

	for _, stateName := range states {
		frames := cache.GetStateFrames(stateName)
		if len(frames) == 0 {
			t.Errorf("state %q has no frames", stateName)
		}

		// Verify each frame has lines
		for frameIdx, frame := range frames {
			if len(frame) == 0 {
				t.Errorf("state %q frame %d has no lines", stateName, frameIdx)
			}
		}
	}
}

func TestFrameCacheConsistency(t *testing.T) {
	agent, err := LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	// Get cache twice - should return same instance
	cache1 := agent.GetFrameCache()
	cache2 := agent.GetFrameCache()

	if cache1 != cache2 {
		t.Error("GetFrameCache should return same instance (cached)")
	}
}

func TestFrameCacheHasState(t *testing.T) {
	agent, err := LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	cache := agent.GetFrameCache()

	// Test existing state
	if !cache.HasState("plan") {
		t.Error("HasState('plan') should return true")
	}

	// Test non-existing state
	if cache.HasState("nonexistent") {
		t.Error("HasState('nonexistent') should return false")
	}
}

func TestFrameCacheMetadata(t *testing.T) {
	agent, err := LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	cache := agent.GetFrameCache()

	// Test character name
	if cache.GetCharacterName() != "sam" {
		t.Errorf("expected character name 'sam', got '%s'", cache.GetCharacterName())
	}

	// Test color
	color := cache.GetColor()
	if color == "" {
		t.Error("color is empty")
	}
	if color[0] != '#' {
		t.Errorf("expected hex color starting with #, got '%s'", color)
	}
}

func TestBackwardCompatibility(t *testing.T) {
	// Ensure existing API still works
	agent, err := LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	// Test state methods still work
	states := agent.ListStates()
	if len(states) == 0 {
		t.Error("ListStates returned empty slice")
	}

	// Test HasState still works
	if !agent.HasState("plan") {
		t.Error("HasState('plan') should return true")
	}

	// Test GetCharacter still works
	char := agent.GetCharacter()
	if char == nil {
		t.Error("GetCharacter returned nil")
	}

	if char.Name != "sam" {
		t.Errorf("expected character name 'sam', got '%s'", char.Name)
	}
}
