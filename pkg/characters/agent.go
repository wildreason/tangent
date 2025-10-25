package characters

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
)

// hexToRGB converts hex color to RGB values
func hexToRGB(hex string) (r, g, b int) {
	// Remove # if present
	hex = strings.TrimPrefix(hex, "#")

	// Parse hex string
	if len(hex) == 6 {
		val, _ := strconv.ParseInt(hex, 16, 32)
		r = int((val >> 16) & 0xFF)
		g = int((val >> 8) & 0xFF)
		b = int(val & 0xFF)
	}
	return
}

// colorize wraps text with ANSI RGB color codes
func colorize(text string, hexColor string) string {
	if hexColor == "" {
		return text
	}
	r, g, b := hexToRGB(hexColor)
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, text)
}

// AgentCharacter wraps a Character with state-based API methods for AI agents
type AgentCharacter struct {
	character *domain.Character
}

// NewAgentCharacter creates a new AgentCharacter wrapper
func NewAgentCharacter(character *domain.Character) *AgentCharacter {
	return &AgentCharacter{
		character: character,
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

	// Create pattern compiler for frame compilation
	compiler := infrastructure.NewPatternCompiler()

	for _, line := range a.character.BaseFrame.Lines {
		compiledLine := compiler.Compile(line)
		coloredLine := colorize(compiledLine, a.character.Color)
		fmt.Fprintln(writer, coloredLine)
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

	// Use provided values if specified, otherwise use state's default values
	stateFPS := state.AnimationFPS
	if fps > 0 {
		stateFPS = fps
	}
	stateLoops := state.AnimationLoops
	if loops > 0 {
		stateLoops = loops
	}

	// Animate the state frames
	frameDur := time.Second / time.Duration(stateFPS)

	// Hide cursor
	fmt.Fprint(writer, "\x1b[?25l")
	defer fmt.Fprint(writer, "\x1b[?25h")

	// Create pattern compiler for frame compilation
	compiler := infrastructure.NewPatternCompiler()

	for loop := 0; loop < stateLoops; loop++ {
		for _, frame := range state.Frames {
			// Clear and print each line (compile pattern codes and apply color)
			for _, line := range frame.Lines {
				compiledLine := compiler.Compile(line)
				coloredLine := colorize(compiledLine, a.character.Color)
				fmt.Fprintf(writer, "\r\x1b[2K%s\n", coloredLine)
			}

			// Move cursor back up
			fmt.Fprintf(writer, "\x1b[%dA", len(frame.Lines))

			time.Sleep(frameDur)
		}
	}

	// Print final frame cleanly (compile pattern codes and apply color)
	finalFrame := state.Frames[len(state.Frames)-1]
	for _, line := range finalFrame.Lines {
		compiledLine := compiler.Compile(line)
		coloredLine := colorize(compiledLine, a.character.Color)
		fmt.Fprintln(writer, coloredLine)
	}

	return nil
}
