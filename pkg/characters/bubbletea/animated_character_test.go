package bubbletea

import (
	"testing"
	"time"

	"github.com/wildreason/tangent/pkg/characters"
)

func TestNewAnimatedCharacter(t *testing.T) {
	// Load a library character
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	// Create animated character
	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	if char == nil {
		t.Fatal("NewAnimatedCharacter returned nil")
	}

	if char.agent == nil {
		t.Error("agent is nil")
	}

	if char.cache == nil {
		t.Error("cache is nil")
	}

	if char.tickInterval != 500*time.Millisecond {
		t.Errorf("expected tick interval 500ms, got %v", char.tickInterval)
	}

	if !char.playing {
		t.Error("expected playing to be true by default")
	}
}

func TestSetState(t *testing.T) {
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	// Test valid state
	err = char.SetState("plan")
	if err != nil {
		t.Errorf("failed to set valid state: %v", err)
	}

	if char.GetState() != "plan" {
		t.Errorf("expected state 'plan', got '%s'", char.GetState())
	}

	// Test invalid state
	err = char.SetState("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent state")
	}
}

func TestPlayPause(t *testing.T) {
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	// Test initial state
	if !char.IsPlaying() {
		t.Error("expected IsPlaying to be true initially")
	}

	// Test pause
	char.Pause()
	if char.IsPlaying() {
		t.Error("expected IsPlaying to be false after Pause")
	}

	// Test play
	char.Play()
	if !char.IsPlaying() {
		t.Error("expected IsPlaying to be true after Play")
	}
}

func TestListStates(t *testing.T) {
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	states := char.ListStates()
	if len(states) == 0 {
		t.Error("expected non-empty states list")
	}

	// Verify states are sorted
	for i := 1; i < len(states); i++ {
		if states[i-1] > states[i] {
			t.Errorf("states not sorted: %v", states)
			break
		}
	}
}

func TestSetTickInterval(t *testing.T) {
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	newInterval := 100 * time.Millisecond
	char.SetTickInterval(newInterval)

	if char.GetTickInterval() != newInterval {
		t.Errorf("expected tick interval %v, got %v", newInterval, char.GetTickInterval())
	}
}

func TestView(t *testing.T) {
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	// Test view returns non-empty string
	view := char.View()
	if view == "" {
		t.Error("expected non-empty view")
	}

	// Test view after setting state
	if err := char.SetState("plan"); err != nil {
		t.Fatalf("failed to set state: %v", err)
	}

	view = char.View()
	if view == "" {
		t.Error("expected non-empty view after setting state")
	}
}

func TestReset(t *testing.T) {
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	// Advance frame manually
	char.currentFrame = 5

	// Reset
	char.Reset()

	if char.currentFrame != 0 {
		t.Errorf("expected currentFrame to be 0 after Reset, got %d", char.currentFrame)
	}
}

func TestGetDimensions(t *testing.T) {
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		t.Fatalf("failed to load library agent: %v", err)
	}

	char := NewAnimatedCharacter(agent, 500*time.Millisecond)

	width := char.GetWidth()
	height := char.GetHeight()

	if width <= 0 {
		t.Errorf("expected positive width, got %d", width)
	}

	if height <= 0 {
		t.Errorf("expected positive height, got %d", height)
	}
}
