package characters

import (
	"fmt"
	"strings"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
	"github.com/wildreason/tangent/pkg/characters/library"
)

// Package characters provides a simple API for terminal character animation.
//
// Core API:
//
//	// Get character with agent states
//	agent, _ := characters.LibraryAgent("mercury")
//
//	// Use agent states
//	agent.Plan(os.Stdout)
//	agent.Think(os.Stdout)
//	agent.Execute(os.Stdout)
//
// Library characters available:
// - alien: Animated character (3 frames)
// - robot: Static character (1 frame)
// - pulse: Loading indicator (3 frames)
// - wave: Progress indicator (5 frames)
// - rocket: Launch sequence (4 frames)
//
// Pattern characters: F=█ T=▀ B=▄ L=▌ R=▐ 1-8=quadrants _=space

// CharacterService provides character creation functionality
type CharacterService struct {
	compiler domain.PatternCompiler
}

// NewCharacterService creates a new character service with default implementations
func NewCharacterService() *CharacterService {
	compiler := infrastructure.NewPatternCompiler()

	return &CharacterService{
		compiler: compiler,
	}
}

// CreateCharacter creates a character from a domain spec
func (cs *CharacterService) CreateCharacter(spec domain.CharacterSpec) (*domain.Character, error) {
	// Convert patterns to frames
	frames := make([]domain.Frame, len(spec.Frames))
	for i, frameSpec := range spec.Frames {
		// Compile patterns to actual character lines
		lines := make([]string, len(frameSpec.Patterns))
		for j, pattern := range frameSpec.Patterns {
			compiled := cs.compiler.Compile(pattern)
			lines[j] = compiled
		}

		frames[i] = domain.Frame{
			Name:  frameSpec.Name,
			Lines: lines,
		}
	}

	// Create character
	character := &domain.Character{
		Name:   spec.Name,
		Width:  spec.Width,
		Height: spec.Height,
		Frames: frames,
		States: make(map[string]domain.State),
	}

	return character, nil
}

// LibraryAgent retrieves a pre-built character from the built-in library with state-based API
func LibraryAgent(name string) (*AgentCharacter, error) {
	libChar, err := library.Get(name)
	if err != nil {
		return nil, err
	}

	// Convert library character to domain character
	compiler := infrastructure.NewPatternCompiler()

	// Convert frames
	frames := make([]domain.Frame, len(libChar.Patterns))
	for i, frame := range libChar.Patterns {
		// Compile patterns to actual character lines
		lines := make([]string, len(frame.Lines))
		for j, pattern := range frame.Lines {
			compiled := compiler.Compile(pattern)
			lines[j] = compiled
		}

		frames[i] = domain.Frame{
			Name:  frame.Name,
			Lines: lines,
		}
	}

	// Separate base frame from state frames
	var baseFrame domain.Frame
	states := make(map[string]domain.State)

	// Find base frame (first frame named "base" or first frame if no "base")
	baseFrameFound := false
	for _, frame := range frames {
		if frame.Name == "base" {
			baseFrame = frame
			baseFrameFound = true
			break
		}
	}

	// If no "base" frame found, use the first frame as base
	if !baseFrameFound && len(frames) > 0 {
		baseFrame = frames[0]
	}

	// Group remaining frames by state name
	stateFrames := make(map[string][]domain.Frame)
	for _, frame := range frames {
		if frame.Name != "base" && frame.Name != baseFrame.Name {
			// Extract state name from pattern name
			stateName := frame.Name
			if strings.Contains(frame.Name, "_") {
				// Handle patterns like "plan_1", "think_2" -> "plan", "think"
				parts := strings.Split(frame.Name, "_")
				if len(parts) >= 2 {
					stateName = parts[0]
				}
			}
			stateFrames[stateName] = append(stateFrames[stateName], frame)
		}
	}

	// Create states from grouped frames
	for stateName, stateFramesList := range stateFrames {
		states[stateName] = domain.State{
			Name:           stateName,
			Description:    fmt.Sprintf("%s state", stateName),
			Frames:         stateFramesList,
			StateType:      "standard",
			AnimationFPS:   5,
			AnimationLoops: 1,
		}
	}

	// Create domain character
	domainChar := &domain.Character{
		Name:        libChar.Name,
		Personality: "", // No personality for library characters
		Color:       libChar.Color,
		Width:       libChar.Width,
		Height:      libChar.Height,
		BaseFrame:   baseFrame,
		States:      states,
		Frames:      frames, // Keep for backward compatibility
	}

	// Wrap in AgentCharacter for state-based API
	return NewAgentCharacter(domainChar), nil
}

// ListLibrary returns all available library character names
func ListLibrary() []string {
	return library.List()
}

// LibraryInfo returns information about a library character
func LibraryInfo(name string) (string, error) {
	libChar, err := library.Get(name)
	if err != nil {
		return "", err
	}
	return libChar.Description, nil
}


// ShowIdle displays the idle state of a character
func ShowIdle(writer interface{}, character *domain.Character) error {
	if len(character.Frames) == 0 {
		return fmt.Errorf("character %s has no frames", character.Name)
	}

	// Show first frame as idle
	for _, line := range character.Frames[0].Lines {
		fmt.Println(line)
	}

	return nil
}
