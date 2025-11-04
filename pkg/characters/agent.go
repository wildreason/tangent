package characters

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
)

// colorize wraps text with ANSI RGB color codes
// This is a convenience wrapper around the exported ColorizeString function
func colorize(text string, hexColor string) string {
	return ColorizeString(text, hexColor)
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

// Plan shows the planning state animation.
// This state typically represents the agent analyzing and planning its next actions.
// It's a convenience wrapper around ShowState(writer, "plan").
func (a *AgentCharacter) Plan(writer io.Writer) error {
	return a.ShowState(writer, "plan")
}

// Think shows the thinking state animation.
// This state typically represents the agent processing information or contemplating.
// It's a convenience wrapper around ShowState(writer, "think").
func (a *AgentCharacter) Think(writer io.Writer) error {
	return a.ShowState(writer, "think")
}

// Execute shows the executing state animation.
// This state typically represents the agent actively performing actions or tasks.
// It's a convenience wrapper around ShowState(writer, "execute").
func (a *AgentCharacter) Execute(writer io.Writer) error {
	return a.ShowState(writer, "execute")
}

// Wait shows the waiting state animation.
// This state typically represents the agent in an idle state, waiting for input or events.
// It's a convenience wrapper around ShowState(writer, "wait").
func (a *AgentCharacter) Wait(writer io.Writer) error {
	return a.ShowState(writer, "wait")
}

// Error shows the error state animation.
// This state typically represents the agent encountering or handling an error condition.
// It's a convenience wrapper around ShowState(writer, "error").
func (a *AgentCharacter) Error(writer io.Writer) error {
	return a.ShowState(writer, "error")
}

// Success shows the success state animation.
// This state typically represents the agent celebrating or acknowledging successful completion.
// It's a convenience wrapper around ShowState(writer, "success").
func (a *AgentCharacter) Success(writer io.Writer) error {
	return a.ShowState(writer, "success")
}

// ShowState displays a specific agent state by name.
// It renders all frames in the state sequentially with a brief pause between frames.
//
// Parameters:
//   - writer: The io.Writer to output the state animation to (typically os.Stdout)
//   - stateName: The name of the state to display (e.g., "plan", "think", "execute")
//
// Returns an error if the state doesn't exist or the character has no states defined.
//
// Example:
//
//	agent.ShowState(os.Stdout, "plan")
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

// ShowBase displays the base (idle) character.
// The base frame represents the character's default resting state.
//
// Parameters:
//   - writer: The io.Writer to output the base frame to (typically os.Stdout)
//
// Returns an error if the character is nil or has no base frame defined.
//
// The function compiles pattern codes to Unicode blocks and applies the character's
// color to all rendered output.
//
// Example:
//
//	agent.ShowBase(os.Stdout)
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

// AnimateState animates a specific state with proper frame animation.
// It displays the state's frames in sequence with smooth transitions, hiding the cursor
// during animation and showing it again when complete.
//
// Parameters:
//   - writer: The io.Writer to output the animation to (typically os.Stdout)
//   - stateName: The name of the state to animate (e.g., "plan", "think", "execute")
//   - fps: Frames per second for the animation. If 0 or negative, uses the state's default AnimationFPS
//   - loops: Number of times to loop the animation. If 0 or negative, uses the state's default AnimationLoops
//
// Returns an error if the state doesn't exist, has no frames, or the character has no states defined.
//
// The function compiles pattern codes to Unicode blocks and applies the character's color
// to all rendered output. The cursor is hidden during animation for a smoother experience.
//
// Example:
//
//	// Animate the "plan" state at 5 FPS for 3 loops
//	err := agent.AnimateState(os.Stdout, "plan", 5, 3)
//	if err != nil {
//	    log.Fatal(err)
//	}
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
