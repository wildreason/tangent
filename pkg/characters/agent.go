package characters

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
	"github.com/wildreason/tangent/pkg/characters/micronoise"
)

// colorize wraps text with ANSI RGB color codes
// This is a convenience wrapper around the exported ColorizeString function
func colorize(text string, hexColor string) string {
	return ColorizeString(text, hexColor)
}

// AgentCharacter wraps a Character with state-based API methods for AI agents
type AgentCharacter struct {
	character *domain.Character
	frameCache *FrameCache // Pre-rendered colored frames for performance
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

	// Check if this is a micro avatar (8x2)
	isMicro := a.character.Width == 8 && a.character.Height == 2
	noiseConfig := micronoise.GetConfig(stateName)
	noiseCounter := 0

	// Select fixed noise slots once (positions persist, characters change)
	var noiseSlots []int
	if isMicro && noiseConfig != nil {
		noiseSlots = micronoise.SelectSlots(noiseConfig.Count)
	}

	for loop := 0; loop < stateLoops; loop++ {
		for _, frame := range state.Frames {
			// Compile and colorize lines
			lines := make([]string, len(frame.Lines))
			for i, line := range frame.Lines {
				compiledLine := compiler.Compile(line)
				lines[i] = colorize(compiledLine, a.character.Color)
			}

			// Apply micro noise with hybrid breathing pattern
			if isMicro && len(noiseSlots) > 0 {
				if micronoise.ShouldRefresh(noiseCounter, noiseConfig.Intensity) {
					// Calculate dynamic noise count based on frame (breathing pattern)
					activeCount := micronoise.CalculateNoiseCount(noiseConfig.Count, noiseCounter)
					if activeCount > 0 {
						lines = micronoise.ApplyNoise(lines, a.character.Width, a.character.Height, noiseSlots, activeCount)
					}
				}
				noiseCounter++
			}

			// Clear and print each line
			for _, line := range lines {
				fmt.Fprintf(writer, "\r\x1b[2K%s\n", line)
			}

			// Move cursor back up
			fmt.Fprintf(writer, "\x1b[%dA", len(frame.Lines))

			time.Sleep(frameDur)
		}
	}

	// Print final frame cleanly
	finalFrame := state.Frames[len(state.Frames)-1]
	lines := make([]string, len(finalFrame.Lines))
	for i, line := range finalFrame.Lines {
		compiledLine := compiler.Compile(line)
		lines[i] = colorize(compiledLine, a.character.Color)
	}

	// Apply noise to final frame (use current count from breathing pattern)
	if isMicro && len(noiseSlots) > 0 {
		activeCount := micronoise.CalculateNoiseCount(noiseConfig.Count, noiseCounter)
		if activeCount > 0 {
			lines = micronoise.ApplyNoise(lines, a.character.Width, a.character.Height, noiseSlots, activeCount)
		}
	}

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	return nil
}

// FrameCache provides O(1) access to pre-rendered, pre-colored frames
// This is useful for TUI frameworks that need fast frame access during
// 60 FPS animations without repeated pattern compilation and colorization.
type FrameCache struct {
	baseFrame     []string              // Pre-rendered base frame
	stateFrames   map[string][][]string // state -> []frames (each frame is []lines)
	characterName string
	color         string
}

// GetFrameCache returns a pre-rendered frame cache for this character.
// All frames are compiled from patterns and colorized with the character's color.
// This provides O(1) frame access for high-performance animation systems.
//
// Example:
//
//	agent, _ := characters.LibraryAgent("sam")
//	cache := agent.GetFrameCache()
//
//	// Get base frame
//	baseLines := cache.GetBaseFrame()
//	for _, line := range baseLines {
//	    fmt.Println(line)  // Already compiled and colored
//	}
//
//	// Get state frames
//	planFrames := cache.GetStateFrames("plan")
//	for _, frame := range planFrames {
//	    for _, line := range frame {
//	        fmt.Println(line)  // Already compiled and colored
//	    }
//	}
func (a *AgentCharacter) GetFrameCache() *FrameCache {
	if a.frameCache != nil {
		return a.frameCache
	}

	// Build cache on first access
	compiler := infrastructure.NewPatternCompiler()

	// Pre-render base frame
	baseLines := make([]string, len(a.character.BaseFrame.Lines))
	for i, line := range a.character.BaseFrame.Lines {
		compiled := compiler.Compile(line)
		baseLines[i] = colorize(compiled, a.character.Color)
	}

	// Pre-render all state frames
	stateFrames := make(map[string][][]string)
	for stateName, state := range a.character.States {
		frames := make([][]string, len(state.Frames))
		for frameIdx, frame := range state.Frames {
			lines := make([]string, len(frame.Lines))
			for lineIdx, line := range frame.Lines {
				compiled := compiler.Compile(line)
				lines[lineIdx] = colorize(compiled, a.character.Color)
			}
			frames[frameIdx] = lines
		}
		stateFrames[stateName] = frames
	}

	a.frameCache = &FrameCache{
		baseFrame:     baseLines,
		stateFrames:   stateFrames,
		characterName: a.character.Name,
		color:         a.character.Color,
	}

	return a.frameCache
}

// GetBaseFrame returns the pre-rendered base (idle) frame
func (fc *FrameCache) GetBaseFrame() []string {
	return fc.baseFrame
}

// GetStateFrames returns all pre-rendered frames for a given state.
// Returns nil if the state doesn't exist.
// Each element in the outer slice is a frame ([]string lines).
func (fc *FrameCache) GetStateFrames(stateName string) [][]string {
	return fc.stateFrames[stateName]
}

// HasState checks if a state exists in the cache
func (fc *FrameCache) HasState(stateName string) bool {
	_, exists := fc.stateFrames[stateName]
	return exists
}

// ListStates returns all available state names in the cache
func (fc *FrameCache) ListStates() []string {
	states := make([]string, 0, len(fc.stateFrames))
	for stateName := range fc.stateFrames {
		states = append(states, stateName)
	}
	sort.Strings(states)
	return states
}

// GetCharacterName returns the character name
func (fc *FrameCache) GetCharacterName() string {
	return fc.characterName
}

// GetColor returns the character's hex color
func (fc *FrameCache) GetColor() string {
	return fc.color
}
