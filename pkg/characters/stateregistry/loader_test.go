package stateregistry

import (
	"testing"
)

func TestLoadEmbedded(t *testing.T) {
	registry, err := LoadEmbedded()
	if err != nil {
		t.Fatalf("Failed to load embedded states: %v", err)
	}

	states := registry.List()
	if len(states) == 0 {
		t.Fatal("No states loaded")
	}

	t.Logf("Loaded %d states: %v", len(states), states)

	// Test getting a specific state
	arise, ok := registry.Get("arise")
	if !ok {
		t.Fatal("Failed to get 'arise' state")
	}

	if arise.Name != "arise" {
		t.Errorf("Expected state name 'arise', got '%s'", arise.Name)
	}

	if len(arise.Frames) == 0 {
		t.Error("arise state has no frames")
	}

	t.Logf("arise state has %d frames", len(arise.Frames))
}

func TestDefaultRegistry(t *testing.T) {
	if DefaultRegistry == nil {
		t.Fatal("DefaultRegistry is nil")
	}

	states := List()
	if len(states) < 10 {
		t.Errorf("Expected at least 10 states, got %d", len(states))
	}

	t.Logf("DefaultRegistry has %d states", len(states))
}
