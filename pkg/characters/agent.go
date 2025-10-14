package characters

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

// AgentCharacter wraps a Character with state-based API methods for AI agents
type AgentCharacter struct {
	character *domain.Character
	animator  domain.AnimationEngine
}

// NewAgentCharacter creates a new AgentCharacter wrapper
func NewAgentCharacter(character *domain.Character) *AgentCharacter {
	return &AgentCharacter{
		character: character,
		animator:  nil, // Will be created on demand
	}
}

// Plan shows the planning state animation
func (a *AgentCharacter) Plan(writer io.Writer) error {
	return a.ShowState(writer, "plan")
}

// Think shows the thinking state animation
func (a *AgentCharacter) Think(writer io.Writer) error {
	return a.ShowState(writer, "think")
}

// Execute shows the executing state animation
func (a *AgentCharacter) Execute(writer io.Writer) error {
	return a.ShowState(writer, "execute")
}

// Wait shows the waiting state animation
func (a *AgentCharacter) Wait(writer io.Writer) error {
	return a.ShowState(writer, "wait")
}

// Error shows the error state animation
func (a *AgentCharacter) Error(writer io.Writer) error {
	return a.ShowState(writer, "error")
}

// Success shows the success state animation
func (a *AgentCharacter) Success(writer io.Writer) error {
	return a.ShowState(writer, "success")
}

// ShowState displays a specific agent state by name
func (a *AgentCharacter) ShowState(writer io.Writer, stateName string) error {
	if a.character == nil {
		return fmt.Errorf("agent character is nil")
	}

	if a.character.States == nil || len(a.character.States) == 0 {
		return fmt.Errorf("character %s has no states defined", a.character.Name)
	}

	state, exists := a.character.States[stateName]
	if !exists {
		return fmt.Errorf("state %q not found for character %s (available: %s)",
			stateName, a.character.Name, strings.Join(a.ListStates(), ", "))
	}

	if len(state.Frames) == 0 {
		return fmt.Errorf("state %q has no frames", stateName)
	}

	// Animate the state's frames
	for _, frame := range state.Frames {
		for _, line := range frame.Lines {
			fmt.Fprintln(writer, line)
		}

		// Brief pause between frames if multiple frames in state
		if len(state.Frames) > 1 {
			time.Sleep(200 * time.Millisecond)
		}
	}

	return nil
}

// ListStates returns all available state names for this character
func (a *AgentCharacter) ListStates() []string {
	if a.character == nil || a.character.States == nil {
		return []string{}
	}

	states := make([]string, 0, len(a.character.States))
	for stateName := range a.character.States {
		states = append(states, stateName)
	}

	// Sort for consistent output
	sort.Strings(states)
	return states
}

// GetCharacter returns the underlying domain.Character
func (a *AgentCharacter) GetCharacter() *domain.Character {
	return a.character
}

// Name returns the character's name
func (a *AgentCharacter) Name() string {
	if a.character == nil {
		return ""
	}
	return a.character.Name
}

// Personality returns the character's personality
func (a *AgentCharacter) Personality() string {
	if a.character == nil {
		return ""
	}
	return a.character.Personality
}

// HasState checks if a specific state exists
func (a *AgentCharacter) HasState(stateName string) bool {
	if a.character == nil || a.character.States == nil {
		return false
	}
	_, exists := a.character.States[stateName]
	return exists
}

// GetStateDescription returns the description of a specific state
func (a *AgentCharacter) GetStateDescription(stateName string) (string, error) {
	if a.character == nil || a.character.States == nil {
		return "", fmt.Errorf("character has no states")
	}

	state, exists := a.character.States[stateName]
	if !exists {
		return "", fmt.Errorf("state %q not found", stateName)
	}

	return state.Description, nil
}

// ShowBase displays the base (idle) character
func (a *AgentCharacter) ShowBase(writer io.Writer) error {
	if a.character == nil {
		return fmt.Errorf("agent character is nil")
	}

	if len(a.character.BaseFrame.Lines) == 0 {
		return fmt.Errorf("character %s has no base frame defined", a.character.Name)
	}

	for _, line := range a.character.BaseFrame.Lines {
		fmt.Fprintln(writer, line)
	}

	return nil
}

// AnimateState animates a specific state with proper frame animation
func (a *AgentCharacter) AnimateState(writer io.Writer, stateName string, fps int, loops int) error {
	if a.character == nil {
		return fmt.Errorf("agent character is nil")
	}

	if a.character.States == nil || len(a.character.States) == 0 {
		return fmt.Errorf("character %s has no states defined", a.character.Name)
	}

	state, exists := a.character.States[stateName]
	if !exists {
		return fmt.Errorf("state %q not found for character %s (available: %s)",
			stateName, a.character.Name, strings.Join(a.ListStates(), ", "))
	}

	if len(state.Frames) == 0 {
		return fmt.Errorf("state %q has no frames", stateName)
	}

	// Use state's FPS and loops if specified, otherwise use provided values
	stateFPS := fps
	if state.AnimationFPS > 0 {
		stateFPS = state.AnimationFPS
	}
	stateLoops := loops
	if state.AnimationLoops > 0 {
		stateLoops = state.AnimationLoops
	}

	// Animate the state frames
	frameDur := time.Second / time.Duration(stateFPS)

	// Hide cursor
	fmt.Fprint(writer, "\x1b[?25l")
	defer fmt.Fprint(writer, "\x1b[?25h")

	for loop := 0; loop < stateLoops; loop++ {
		for _, frame := range state.Frames {
			// Clear and print each line
			for _, line := range frame.Lines {
				fmt.Fprintf(writer, "\r\x1b[2K%s\n", line)
			}

			// Move cursor back up
			fmt.Fprintf(writer, "\x1b[%dA", len(frame.Lines))

			time.Sleep(frameDur)
		}
	}

	// Print final frame cleanly
	finalFrame := state.Frames[len(state.Frames)-1]
	for _, line := range finalFrame.Lines {
		fmt.Fprintln(writer, line)
	}

	return nil
}
