package client

import (
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	c, err := New("sam")
	if err != nil {
		t.Fatalf("New(sam) failed: %v", err)
	}
	if c == nil {
		t.Fatal("New(sam) returned nil")
	}

	if c.GetState() != "resting" {
		t.Errorf("initial state = %q, want resting", c.GetState())
	}
}

func TestNewMicro(t *testing.T) {
	c, err := NewMicro("sam")
	if err != nil {
		t.Fatalf("NewMicro(sam) failed: %v", err)
	}

	w, h := c.GetDimensions()
	if w != 8 || h != 2 {
		t.Errorf("dimensions = %dx%d, want 8x2", w, h)
	}
}

func TestNewInvalidCharacter(t *testing.T) {
	_, err := New("nonexistent")
	if err == nil {
		t.Error("New(nonexistent) should return error")
	}

	_, err = NewMicro("nonexistent")
	if err == nil {
		t.Error("NewMicro(nonexistent) should return error")
	}
}

func TestSetState(t *testing.T) {
	c, _ := NewMicro("sam")

	c.SetState("write")
	if c.GetState() != "write" {
		t.Errorf("state = %q, want write", c.GetState())
	}

	c.SetState("resting")
	if c.GetState() != "resting" {
		t.Errorf("state = %q, want resting", c.GetState())
	}
}

func TestSetStateResetsFrame(t *testing.T) {
	c, _ := NewMicro("sam")

	// Advance some frames
	c.Tick()
	c.Tick()
	c.Tick()

	if c.GetFrameIndex() == 0 {
		t.Skip("not enough frames to test")
	}

	c.SetState("write")
	if c.GetFrameIndex() != 0 {
		t.Errorf("frame index = %d, want 0 after state change", c.GetFrameIndex())
	}
}

func TestAliasResolution(t *testing.T) {
	c, _ := NewMicro("sam")

	// Test default aliases
	c.SetState("thinking")
	if c.GetState() != "wait" {
		t.Errorf("thinking resolved to %q, want wait", c.GetState())
	}

	c.SetState("bash")
	if c.GetState() != "search" {
		t.Errorf("bash resolved to %q, want search", c.GetState())
	}

	c.SetState("success")
	if c.GetState() != "approval" {
		t.Errorf("success resolved to %q, want approval", c.GetState())
	}
}

func TestCustomAlias(t *testing.T) {
	c, _ := NewMicro("sam")

	// Set custom alias
	c.SetAlias("busy", "write")
	c.SetState("busy")
	if c.GetState() != "write" {
		t.Errorf("busy resolved to %q, want write", c.GetState())
	}

	// Custom alias overrides default
	c.SetAlias("thinking", "search")
	c.SetState("thinking")
	if c.GetState() != "search" {
		t.Errorf("thinking resolved to %q after custom alias, want search", c.GetState())
	}

	// Remove custom alias, falls back to default
	c.RemoveAlias("thinking")
	c.SetState("thinking")
	if c.GetState() != "wait" {
		t.Errorf("thinking resolved to %q after removing alias, want wait", c.GetState())
	}
}

func TestFPSConfig(t *testing.T) {
	c, _ := NewMicro("sam")

	// Default FPS
	if c.GetFPS() != 5 {
		t.Errorf("default FPS = %d, want 5", c.GetFPS())
	}

	// Set default FPS
	c.SetDefaultFPS(10)
	if c.GetFPS() != 10 {
		t.Errorf("FPS after SetDefaultFPS = %d, want 10", c.GetFPS())
	}

	// State-specific FPS
	c.SetStateFPS("write", 8)
	c.SetState("write")
	if c.GetFPS() != 8 {
		t.Errorf("FPS for write = %d, want 8", c.GetFPS())
	}

	// Override FPS
	c.SetFPS(15)
	if c.GetFPS() != 15 {
		t.Errorf("FPS after override = %d, want 15", c.GetFPS())
	}

	// State change clears override
	c.SetState("resting")
	if c.GetFPS() != 10 {
		t.Errorf("FPS after state change = %d, want 10 (default)", c.GetFPS())
	}
}

func TestSetStateWithFPS(t *testing.T) {
	c, _ := NewMicro("sam")
	c.SetDefaultFPS(5)

	c.SetStateWithFPS("write", 20)
	if c.GetFPS() != 20 {
		t.Errorf("FPS = %d, want 20", c.GetFPS())
	}

	// State change clears the override
	c.SetState("resting")
	if c.GetFPS() != 5 {
		t.Errorf("FPS after state change = %d, want 5", c.GetFPS())
	}
}

func TestGetFrame(t *testing.T) {
	c, _ := NewMicro("sam")

	frame := c.GetFrame()
	if len(frame) != 2 {
		t.Errorf("frame has %d lines, want 2", len(frame))
	}

	for i, line := range frame {
		if len(line) == 0 {
			t.Errorf("frame line %d is empty", i)
		}
	}
}

func TestTick(t *testing.T) {
	c, _ := NewMicro("sam")

	initial := c.GetFrameIndex()
	c.Tick()
	after := c.GetFrameIndex()

	if after != initial+1 {
		t.Errorf("frame index = %d, want %d", after, initial+1)
	}
}

func TestLoopCount(t *testing.T) {
	c, _ := NewMicro("sam")

	// Get frame count for state
	frames := c.cache.GetStateFrames("resting")
	if len(frames) < 2 {
		t.Skip("not enough frames to test loop")
	}

	// Advance through all frames
	for i := 0; i < len(frames); i++ {
		c.Tick()
	}

	if c.GetLoopCount() != 1 {
		t.Errorf("loop count = %d, want 1", c.GetLoopCount())
	}

	// Complete another loop
	for i := 0; i < len(frames); i++ {
		c.Tick()
	}

	if c.GetLoopCount() != 2 {
		t.Errorf("loop count = %d, want 2", c.GetLoopCount())
	}
}

func TestQueueState(t *testing.T) {
	c, _ := NewMicro("sam")

	frames := c.cache.GetStateFrames("resting")
	if len(frames) < 2 {
		t.Skip("not enough frames")
	}

	// Queue state after 1 loop
	c.QueueState("write", AfterLoops(1))

	// Advance through first loop
	for i := 0; i < len(frames); i++ {
		c.Tick()
	}

	if c.GetState() != "write" {
		t.Errorf("state = %q, want write after queue trigger", c.GetState())
	}
}

func TestQueueStateAfterFrames(t *testing.T) {
	c, _ := NewMicro("sam")

	c.QueueState("write", AfterFrames(3))

	c.Tick()
	c.Tick()
	if c.GetState() == "write" {
		t.Error("state changed too early")
	}

	c.Tick()
	if c.GetState() != "write" {
		t.Errorf("state = %q, want write after 3 frames", c.GetState())
	}
}

func TestQueueStateImmediate(t *testing.T) {
	c, _ := NewMicro("sam")

	c.QueueState("write", Immediate())
	c.Tick()

	if c.GetState() != "write" {
		t.Errorf("state = %q, want write after immediate queue", c.GetState())
	}
}

func TestClearQueue(t *testing.T) {
	c, _ := NewMicro("sam")

	c.QueueState("write", AfterFrames(1))
	c.ClearQueue()
	c.Tick()

	if c.GetState() != "resting" {
		t.Errorf("state = %q, want resting after clear queue", c.GetState())
	}
}

func TestOnStateChangeCallback(t *testing.T) {
	c, _ := NewMicro("sam")

	var from, to string
	done := make(chan struct{})

	c.OnStateChange(func(f, t string) {
		from, to = f, t
		close(done)
	})

	c.SetState("write")

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("callback not called")
	}

	if from != "resting" || to != "write" {
		t.Errorf("callback got from=%q, to=%q, want resting, write", from, to)
	}
}

func TestOnLoopCompleteCallback(t *testing.T) {
	c, _ := NewMicro("sam")

	frames := c.cache.GetStateFrames("resting")
	if len(frames) < 2 {
		t.Skip("not enough frames")
	}

	var state string
	var loop int
	done := make(chan struct{})

	c.OnLoopComplete(func(s string, l int) {
		state, loop = s, l
		close(done)
	})

	for i := 0; i < len(frames); i++ {
		c.Tick()
	}

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("callback not called")
	}

	if state != "resting" || loop != 1 {
		t.Errorf("callback got state=%q, loop=%d, want resting, 1", state, loop)
	}
}

func TestStartStop(t *testing.T) {
	c, _ := NewMicro("sam")
	c.SetDefaultFPS(10) // 100ms per frame

	c.Start()
	if !c.IsRunning() {
		t.Error("IsRunning() = false after Start()")
	}

	time.Sleep(250 * time.Millisecond) // Let a few frames tick

	c.Stop()
	if c.IsRunning() {
		t.Error("IsRunning() = true after Stop()")
	}

	// Give goroutine time to exit
	time.Sleep(50 * time.Millisecond)

	frameAfterStop := c.GetFrameIndex()
	time.Sleep(200 * time.Millisecond)

	if c.GetFrameIndex() != frameAfterStop {
		t.Error("frame index changed after Stop()")
	}
}

func TestConcurrentAccess(t *testing.T) {
	c, _ := NewMicro("sam")
	c.Start()
	defer c.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.GetFrame()
				c.GetState()
				c.GetFPS()
			}
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			states := []string{"resting", "write", "search", "wait"}
			for j := 0; j < 20; j++ {
				c.SetState(states[j%len(states)])
			}
		}(i)
	}

	wg.Wait()
}

func TestListStates(t *testing.T) {
	c, _ := NewMicro("sam")

	states := c.ListStates()
	if len(states) < 5 {
		t.Errorf("ListStates() returned %d states, expected at least 5", len(states))
	}

	// Check for expected states
	expected := map[string]bool{
		"resting":  false,
		"write":    false,
		"search":   false,
		"wait":     false,
		"approval": false,
	}

	for _, s := range states {
		if _, ok := expected[s]; ok {
			expected[s] = true
		}
	}

	for s, found := range expected {
		if !found {
			t.Errorf("missing expected state: %s", s)
		}
	}
}

func TestHasState(t *testing.T) {
	c, _ := NewMicro("sam")

	if !c.HasState("resting") {
		t.Error("HasState(resting) = false")
	}

	if !c.HasState("thinking") { // alias
		t.Error("HasState(thinking) = false")
	}

	if c.HasState("nonexistent_state_xyz") {
		t.Error("HasState(nonexistent) = true")
	}
}

func TestGetColor(t *testing.T) {
	c, _ := NewMicro("sam")

	color := c.GetColor()
	if color == "" {
		t.Error("GetColor() returned empty string")
	}

	if color[0] != '#' {
		t.Errorf("GetColor() = %q, want hex color starting with #", color)
	}
}

func TestGetCharacterName(t *testing.T) {
	c, _ := NewMicro("sam")

	name := c.GetCharacterName()
	if name == "" {
		t.Error("GetCharacterName() returned empty string")
	}
}

func TestStripANSI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"\x1b[38;2;255;0;0mred\x1b[0m", "red"},
		{"plain", "plain"},
		{"\x1b[1m\x1b[32mgreen\x1b[0m", "green"},
		{"", ""},
	}

	for _, tt := range tests {
		got := stripANSI(tt.input)
		if got != tt.want {
			t.Errorf("stripANSI(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestGetFrameRaw(t *testing.T) {
	c, _ := NewMicro("sam")

	raw := c.GetFrameRaw()
	if len(raw) != 2 {
		t.Errorf("GetFrameRaw() returned %d lines, want 2", len(raw))
	}

	// Raw should not contain ANSI codes
	for i, line := range raw {
		if containsANSI(line) {
			t.Errorf("GetFrameRaw() line %d contains ANSI codes", i)
		}
	}
}

func containsANSI(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '\x1b' {
			return true
		}
	}
	return false
}

func TestFallbackToResting(t *testing.T) {
	c, _ := NewMicro("sam")

	// Set to a non-existent state
	c.SetState("completely_invalid_state_name")

	// Should fall back to resting
	if c.GetState() != "resting" {
		t.Errorf("state = %q, want resting (fallback)", c.GetState())
	}
}
