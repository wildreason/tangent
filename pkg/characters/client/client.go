package client

import (
	"fmt"
	"sync"
	"time"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/micronoise"
)

// TangentClient is a framework-agnostic animation controller for tangent characters.
// It manages animation state, frame timing, and state transitions in a thread-safe manner.
//
// Usage with tcell/tview:
//
//	tc := client.NewMicro("sam")
//	tc.SetStateFPS("resting", 2)
//	tc.SetStateFPS("write", 8)
//	tc.Start()
//
//	// In your render loop
//	frame := tc.GetFrame()
//	avatarView.SetText(strings.Join(frame, "\n"))
//
//	// On state change events
//	tc.SetState("write")
type TangentClient struct {
	mu sync.RWMutex

	// Character data
	cache *characters.FrameCache
	agent *characters.AgentCharacter

	// Current animation state
	currentState string
	frameIndex   int
	loopCount    int
	frameCount   int // total frames shown for current state

	// FPS configuration
	defaultFPS  int            // fallback FPS (default: 5)
	stateFPS    map[string]int // per-state FPS overrides
	overrideFPS int            // temporary FPS override (0 = disabled)

	// State aliases
	aliases map[string]string

	// State queue
	queuedState *queuedStateEntry

	// Callbacks
	onStateChange  func(from, to string)
	onLoopComplete func(state string, loop int)

	// Auto-tick
	ticker     *time.Ticker
	tickerDone chan struct{}
	running    bool

	// Micro noise support
	isMicro      bool
	noiseCounter int
	noiseSlots   []int
	width        int
	height       int
}

// New creates a TangentClient for a regular (11x4) character.
func New(name string) (*TangentClient, error) {
	agent, err := characters.LibraryAgent(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load character %q: %w", name, err)
	}

	return newClient(agent), nil
}

// NewMicro creates a TangentClient for a micro (8x2) character.
func NewMicro(name string) (*TangentClient, error) {
	agent, err := characters.LibraryAgentMicro(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load micro character %q: %w", name, err)
	}

	return newClient(agent), nil
}

func newClient(agent *characters.AgentCharacter) *TangentClient {
	cache := agent.GetFrameCache()
	char := agent.GetCharacter()
	isMicro := char.Width == 8 && char.Height == 2

	// Initialize noise slots for micro avatars
	var noiseSlots []int
	if isMicro {
		if cfg := micronoise.GetConfig("resting"); cfg != nil {
			noiseSlots = micronoise.SelectSlots(cfg.Count)
		}
	}

	c := &TangentClient{
		cache:        cache,
		agent:        agent,
		currentState: "resting",
		defaultFPS:   5,
		stateFPS:     make(map[string]int),
		aliases:      make(map[string]string),
		// Noise fields
		isMicro:      isMicro,
		noiseCounter: 0,
		noiseSlots:   noiseSlots,
		width:        char.Width,
		height:       char.Height,
	}

	return c
}

// --- State Control ---

// SetState changes the current animation state immediately.
// The state name is resolved through aliases (custom first, then defaults).
// Resets frame index and loop count. Triggers OnStateChange callback if set.
func (c *TangentClient) SetState(state string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	resolved := c.resolveState(state)
	if resolved == c.currentState {
		return
	}

	oldState := c.currentState
	c.currentState = resolved
	c.frameIndex = 0
	c.loopCount = 0
	c.frameCount = 0
	c.overrideFPS = 0 // clear any FPS override
	c.queuedState = nil

	// Reset noise for recognition phase
	if c.isMicro {
		c.noiseCounter = 0
		if cfg := micronoise.GetConfig(resolved); cfg != nil {
			c.noiseSlots = micronoise.SelectSlots(cfg.Count)
		} else {
			c.noiseSlots = nil
		}
	}

	if c.onStateChange != nil {
		go c.onStateChange(oldState, resolved)
	}

	// Restart ticker if running (FPS may have changed)
	if c.running {
		c.restartTicker()
	}
}

// SetStateWithFPS changes state and temporarily overrides FPS for this state activation.
func (c *TangentClient) SetStateWithFPS(state string, fps int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	resolved := c.resolveState(state)
	oldState := c.currentState

	c.currentState = resolved
	c.frameIndex = 0
	c.loopCount = 0
	c.frameCount = 0
	c.overrideFPS = fps
	c.queuedState = nil

	// Reset noise for recognition phase
	if c.isMicro {
		c.noiseCounter = 0
		if cfg := micronoise.GetConfig(resolved); cfg != nil {
			c.noiseSlots = micronoise.SelectSlots(cfg.Count)
		} else {
			c.noiseSlots = nil
		}
	}

	if oldState != resolved && c.onStateChange != nil {
		go c.onStateChange(oldState, resolved)
	}

	if c.running {
		c.restartTicker()
	}
}

// QueueState queues a state transition that will occur when the condition is met.
// Only one state can be queued at a time; calling again replaces the queue.
func (c *TangentClient) QueueState(state string, condition QueueCondition) {
	c.mu.Lock()
	defer c.mu.Unlock()

	resolved := c.resolveState(state)
	c.queuedState = &queuedStateEntry{
		state:     resolved,
		condition: condition,
	}
}

// QueueStateWithFPS queues a state transition with a specific FPS override.
func (c *TangentClient) QueueStateWithFPS(state string, condition QueueCondition, fps int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	resolved := c.resolveState(state)
	c.queuedState = &queuedStateEntry{
		state:     resolved,
		condition: condition,
		fps:       fps,
	}
}

// ClearQueue removes any queued state transition.
func (c *TangentClient) ClearQueue() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.queuedState = nil
}

// --- FPS Control ---

// SetDefaultFPS sets the fallback FPS used when no state-specific FPS is configured.
func (c *TangentClient) SetDefaultFPS(fps int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if fps < 1 {
		fps = 1
	}
	c.defaultFPS = fps

	if c.running {
		c.restartTicker()
	}
}

// SetStateFPS configures the FPS for a specific state.
func (c *TangentClient) SetStateFPS(state string, fps int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	resolved := c.resolveState(state)
	if fps < 1 {
		delete(c.stateFPS, resolved)
	} else {
		c.stateFPS[resolved] = fps
	}

	if c.running && c.currentState == resolved {
		c.restartTicker()
	}
}

// SetFPS temporarily overrides the current FPS (cleared on state change).
func (c *TangentClient) SetFPS(fps int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if fps < 1 {
		c.overrideFPS = 0
	} else {
		c.overrideFPS = fps
	}

	if c.running {
		c.restartTicker()
	}
}

// GetFPS returns the current effective FPS.
// Priority: override > state-specific > default.
func (c *TangentClient) GetFPS() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.effectiveFPS()
}

func (c *TangentClient) effectiveFPS() int {
	if c.overrideFPS > 0 {
		return c.overrideFPS
	}
	if fps, ok := c.stateFPS[c.currentState]; ok {
		return fps
	}
	return c.defaultFPS
}

// --- Aliases ---

// SetAlias adds a custom state alias. Custom aliases take precedence over defaults.
func (c *TangentClient) SetAlias(from, to string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.aliases[from] = to
}

// RemoveAlias removes a custom alias.
func (c *TangentClient) RemoveAlias(from string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.aliases, from)
}

func (c *TangentClient) resolveState(name string) string {
	// Check custom aliases first
	if resolved, ok := c.aliases[name]; ok {
		if c.cache.HasState(resolved) {
			return resolved
		}
	}
	// Check default aliases
	if resolved, ok := DefaultAliases[name]; ok {
		if c.cache.HasState(resolved) {
			return resolved
		}
	}
	// Check if state exists directly
	if c.cache.HasState(name) {
		return name
	}
	// Fallback to resting if state doesn't exist
	if c.cache.HasState("resting") {
		return "resting"
	}
	return name
}

// --- Frame Retrieval ---

// GetFrame returns the current animation frame as pre-colored lines.
// This is safe to call from any goroutine.
// For micro avatars (8x2), applies noise with breathing pattern.
func (c *TangentClient) GetFrame() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	frames := c.cache.GetStateFrames(c.currentState)
	if len(frames) == 0 {
		return c.cache.GetBaseFrame()
	}
	lines := frames[c.frameIndex%len(frames)]

	// Apply micro noise (breathing pattern)
	if c.isMicro && len(c.noiseSlots) > 0 {
		if cfg := micronoise.GetConfig(c.currentState); cfg != nil {
			if micronoise.ShouldRefresh(c.noiseCounter, cfg.Intensity) {
				activeCount := micronoise.CalculateNoiseCount(cfg.Count, c.noiseCounter)
				if activeCount > 0 {
					lines = micronoise.ApplyNoise(lines, c.width, c.height, c.noiseSlots, activeCount)
				}
			}
		}
	}

	return lines
}

// GetFrameRaw returns the current frame without color codes.
// Useful when applying custom colors.
func (c *TangentClient) GetFrameRaw() []string {
	frame := c.GetFrame()
	result := make([]string, len(frame))
	for i, line := range frame {
		result[i] = stripANSI(line)
	}
	return result
}

// Tick advances the animation by one frame.
// Call this at your desired frame rate when not using Start().
// Processes queued state transitions and triggers callbacks.
func (c *TangentClient) Tick() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Increment noise counter for micro avatars
	if c.isMicro {
		c.noiseCounter++
	}

	frames := c.cache.GetStateFrames(c.currentState)
	if len(frames) == 0 {
		return
	}

	c.frameIndex++
	c.frameCount++

	if c.frameIndex >= len(frames) {
		c.frameIndex = 0
		c.loopCount++

		if c.onLoopComplete != nil {
			go c.onLoopComplete(c.currentState, c.loopCount)
		}
	}

	// Check queue
	c.processQueue()
}

func (c *TangentClient) processQueue() {
	if c.queuedState == nil {
		return
	}

	if c.queuedState.condition.shouldActivate(c.loopCount, c.frameCount) {
		entry := c.queuedState
		c.queuedState = nil

		oldState := c.currentState
		c.currentState = entry.state
		c.frameIndex = 0
		c.loopCount = 0
		c.frameCount = 0
		c.overrideFPS = entry.fps

		if c.onStateChange != nil {
			go c.onStateChange(oldState, entry.state)
		}

		if c.running {
			c.restartTicker()
		}
	}
}

// --- Auto-tick ---

// Start begins automatic frame advancement using an internal ticker.
// The ticker rate is based on the current effective FPS.
func (c *TangentClient) Start() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.running {
		return
	}

	c.running = true
	c.startTicker()
}

// Stop halts the internal ticker.
func (c *TangentClient) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.running {
		return
	}

	c.running = false
	c.stopTicker()
}

// IsRunning returns true if the internal ticker is active.
func (c *TangentClient) IsRunning() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.running
}

func (c *TangentClient) startTicker() {
	fps := c.effectiveFPS()
	interval := time.Second / time.Duration(fps)

	c.tickerDone = make(chan struct{})
	c.ticker = time.NewTicker(interval)

	// Capture locally to avoid race with stopTicker
	ticker := c.ticker
	done := c.tickerDone

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				c.Tick()
			}
		}
	}()
}

func (c *TangentClient) stopTicker() {
	if c.ticker != nil {
		c.ticker.Stop()
		c.ticker = nil
	}
	if c.tickerDone != nil {
		close(c.tickerDone)
		c.tickerDone = nil
	}
}

func (c *TangentClient) restartTicker() {
	c.stopTicker()
	c.startTicker()
}

// --- Callbacks ---

// OnStateChange sets a callback invoked when the state changes.
// The callback receives the old and new state names.
// Called in a separate goroutine to avoid blocking.
func (c *TangentClient) OnStateChange(fn func(from, to string)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.onStateChange = fn
}

// OnLoopComplete sets a callback invoked when an animation loop completes.
// Called in a separate goroutine to avoid blocking.
func (c *TangentClient) OnLoopComplete(fn func(state string, loop int)) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.onLoopComplete = fn
}

// --- Query Methods ---

// GetState returns the current animation state name.
func (c *TangentClient) GetState() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.currentState
}

// GetLoopCount returns the number of completed loops for the current state.
func (c *TangentClient) GetLoopCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.loopCount
}

// GetFrameIndex returns the current frame index within the animation.
func (c *TangentClient) GetFrameIndex() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.frameIndex
}

// ListStates returns all available state names.
func (c *TangentClient) ListStates() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache.ListStates()
}

// HasState returns true if the state exists (checking aliases).
func (c *TangentClient) HasState(state string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Check direct state
	if c.cache.HasState(state) {
		return true
	}
	// Check custom aliases
	if resolved, ok := c.aliases[state]; ok {
		return c.cache.HasState(resolved)
	}
	// Check default aliases
	if resolved, ok := DefaultAliases[state]; ok {
		return c.cache.HasState(resolved)
	}
	return false
}

// GetColor returns the character's hex color code.
func (c *TangentClient) GetColor() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache.GetColor()
}

// GetDimensions returns the character's width and height.
func (c *TangentClient) GetDimensions() (width, height int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	char := c.agent.GetCharacter()
	return char.Width, char.Height
}

// GetCharacterName returns the character's name.
func (c *TangentClient) GetCharacterName() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache.GetCharacterName()
}

// --- Utility ---

// stripANSI removes ANSI escape codes from a string.
func stripANSI(s string) string {
	result := make([]byte, 0, len(s))
	inEscape := false
	for i := 0; i < len(s); i++ {
		if s[i] == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if s[i] == 'm' {
				inEscape = false
			}
			continue
		}
		result = append(result, s[i])
	}
	return string(result)
}
