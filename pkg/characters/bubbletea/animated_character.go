package bubbletea

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wildreason/tangent/pkg/characters"
)

// AnimatedCharacter is a Bubble Tea component that provides
// ready-to-use character animation with state management.
//
// This eliminates the need for custom adapters and provides
// plug-and-play character animation for Bubble Tea applications.
//
// Example:
//
//	agent, _ := characters.LibraryAgent("sam")
//	char := bubbletea.NewAnimatedCharacter(agent, 500*time.Millisecond)
//	char.SetState("plan")
//
//	// Use in Bubble Tea program:
//	program := tea.NewProgram(char)
//	program.Run()
type AnimatedCharacter struct {
	agent         *characters.AgentCharacter
	cache         *characters.FrameCache
	currentState  string
	currentFrame  int
	tickInterval  time.Duration
	playing       bool
	width         int
	height        int
}

// NewAnimatedCharacter creates a new Bubble Tea component from an AgentCharacter.
// The tickInterval determines the animation speed (e.g., 500ms = 2 FPS).
//
// Example:
//
//	agent, _ := characters.LibraryAgent("sam")
//	char := bubbletea.NewAnimatedCharacter(agent, 500*time.Millisecond)
func NewAnimatedCharacter(agent *characters.AgentCharacter, tickInterval time.Duration) *AnimatedCharacter {
	cache := agent.GetFrameCache()

	// Default to base frame state
	initialState := ""
	if len(cache.ListStates()) > 0 {
		initialState = cache.ListStates()[0]
	}

	char := agent.GetCharacter()

	return &AnimatedCharacter{
		agent:        agent,
		cache:        cache,
		currentState: initialState,
		currentFrame: 0,
		tickInterval: tickInterval,
		playing:      true,
		width:        char.Width,
		height:       char.Height,
	}
}

// TickMsg is sent on each animation tick
type TickMsg time.Time

// Init implements tea.Model
func (m *AnimatedCharacter) Init() tea.Cmd {
	if m.playing {
		return m.tick()
	}
	return nil
}

// tick returns a command that sends a TickMsg after the tick interval
func (m *AnimatedCharacter) tick() tea.Cmd {
	return tea.Tick(m.tickInterval, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Update implements tea.Model
func (m *AnimatedCharacter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		if !m.playing {
			return m, nil
		}

		// Advance to next frame
		frames := m.cache.GetStateFrames(m.currentState)
		if len(frames) > 0 {
			m.currentFrame = (m.currentFrame + 1) % len(frames)
		}

		return m, m.tick()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case " ":
			// Toggle play/pause with spacebar
			m.playing = !m.playing
			if m.playing {
				return m, m.tick()
			}
		}
	}

	return m, nil
}

// View implements tea.Model
func (m *AnimatedCharacter) View() string {
	var lines []string

	if m.currentState == "" {
		// Show base frame if no state set
		lines = m.cache.GetBaseFrame()
	} else {
		frames := m.cache.GetStateFrames(m.currentState)
		if len(frames) == 0 {
			// Fallback to base frame if state has no frames
			lines = m.cache.GetBaseFrame()
		} else {
			// Show current frame from state animation
			lines = frames[m.currentFrame]
		}
	}

	return strings.Join(lines, "\n")
}

// SetState changes the current animation state
// Returns error if state doesn't exist
func (m *AnimatedCharacter) SetState(stateName string) error {
	if !m.cache.HasState(stateName) {
		return fmt.Errorf("state %q not found (available: %s)",
			stateName, strings.Join(m.cache.ListStates(), ", "))
	}

	m.currentState = stateName
	m.currentFrame = 0
	return nil
}

// GetState returns the current state name
func (m *AnimatedCharacter) GetState() string {
	return m.currentState
}

// Play starts or resumes animation
func (m *AnimatedCharacter) Play() {
	m.playing = true
}

// Pause pauses animation
func (m *AnimatedCharacter) Pause() {
	m.playing = false
}

// IsPlaying returns whether animation is currently playing
func (m *AnimatedCharacter) IsPlaying() bool {
	return m.playing
}

// SetTickInterval changes the animation speed
func (m *AnimatedCharacter) SetTickInterval(interval time.Duration) {
	m.tickInterval = interval
}

// GetTickInterval returns the current tick interval
func (m *AnimatedCharacter) GetTickInterval() time.Duration {
	return m.tickInterval
}

// ListStates returns all available state names
func (m *AnimatedCharacter) ListStates() []string {
	return m.cache.ListStates()
}

// GetWidth returns the character width
func (m *AnimatedCharacter) GetWidth() int {
	return m.width
}

// GetHeight returns the character height
func (m *AnimatedCharacter) GetHeight() int {
	return m.height
}

// Reset resets animation to first frame
func (m *AnimatedCharacter) Reset() {
	m.currentFrame = 0
}
