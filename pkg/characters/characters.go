package characters

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
	"github.com/wildreason/tangent/pkg/characters/library"
	"github.com/wildreason/tangent/pkg/characters/service"
)

// Package characters provides a simple API for terminal character animation.
//
// Core API (3 functions only):
//
//	// 1. Get library character
//	alien, _ := characters.Library("alien")
//
//	// 2. Create custom character
//	robot := characters.NewCharacterSpec("my-robot", 8, 4).
//		AddFrame("idle", []string{"FRF", "LRL", "FRF", "LRL"}).
//		Build()
//
//	// 3. Use character
//	characters.Animate(os.Stdout, alien, 5, 3)
//	characters.ShowIdle(os.Stdout, robot)
//
// Library characters available:
// - alien: Animated character (3 frames)
// - robot: Static character (1 frame)
// - pulse: Loading indicator (3 frames)
// - wave: Progress indicator (5 frames)
// - rocket: Launch sequence (4 frames)
//
// Pattern characters: F=█ T=▀ B=▄ L=▌ R=▐ 1-8=quadrants _=space

// NewCharacterService creates a new character service with default implementations
func NewCharacterService() *service.CharacterService {
	homeDir, _ := os.UserHomeDir()
	baseDir := filepath.Join(homeDir, ".tangent")
	os.MkdirAll(baseDir, 0755)

	repo := infrastructure.NewFileCharacterRepository(baseDir)
	compiler := infrastructure.NewPatternCompiler()
	animationEngine := infrastructure.NewAnimationEngine()

	return service.NewCharacterService(repo, compiler, animationEngine)
}

// ConvertDomainToLegacy converts a domain.Character to the legacy Character type
func ConvertDomainToLegacy(domainChar *domain.Character) *Character {
	if domainChar == nil {
		return nil
	}

	// Convert frames to strings
	frameStrings := make([]string, len(domainChar.Frames))
	for i, frame := range domainChar.Frames {
		frameStrings[i] = strings.Join(frame.Lines, "\n")
	}

	// Use first frame as idle, or empty string if no frames
	idle := ""
	if len(frameStrings) > 0 {
		idle = frameStrings[0]
	}

	return &Character{
		Name:   domainChar.Name,
		Idle:   idle,
		Frames: frameStrings,
		Width:  domainChar.Width,
		Height: domainChar.Height,
	}
}

// ConvertLegacyToDomain converts a legacy Character to the domain.Character type
func ConvertLegacyToDomain(legacyChar *Character) *domain.Character {
	if legacyChar == nil {
		return nil
	}

	// Convert frame strings to domain frames
	frames := make([]domain.Frame, len(legacyChar.Frames))
	for i, frameString := range legacyChar.Frames {
		lines := strings.Split(frameString, "\n")
		frames[i] = domain.Frame{
			Name:  fmt.Sprintf("frame_%d", i),
			Lines: lines,
		}
	}

	return &domain.Character{
		Name:   legacyChar.Name,
		Width:  legacyChar.Width,
		Height: legacyChar.Height,
		Frames: frames,
	}
}

// ConvertLegacySpecToDomain converts a legacy CharacterSpec to domain.CharacterSpec
func ConvertLegacySpecToDomain(legacySpec *CharacterSpec) *domain.CharacterSpec {
	if legacySpec == nil {
		return nil
	}

	domainSpec := domain.NewCharacterSpec(legacySpec.Name, legacySpec.Width, legacySpec.Height)
	for _, frame := range legacySpec.Frames {
		domainSpec.AddFrame(frame.Name, frame.Patterns)
	}
	return domainSpec
}

// Library retrieves a pre-built character from the built-in library and builds it
// Deprecated: Use LibraryAgent() for state-based API
func Library(name string) (*Character, error) {
	libChar, err := library.Get(name)
	if err != nil {
		return nil, err
	}

	// Build the character using the library patterns
	spec := NewCharacterSpec(libChar.Name, libChar.Width, libChar.Height)
	for _, frame := range libChar.Patterns {
		spec.AddFrame(frame.Name, frame.Lines)
	}

	// Convert to domain spec and use the new service layer
	domainSpec := ConvertLegacySpecToDomain(spec)
	characterService := NewCharacterService()
	domainChar, err := characterService.CreateCharacter(*domainSpec)
	if err != nil {
		return nil, err
	}

	// Convert to legacy format for backward compatibility
	return ConvertDomainToLegacy(domainChar), nil
}

// LibraryAgent retrieves a pre-built character from the built-in library with state-based API
func LibraryAgent(name string) (*AgentCharacter, error) {
	libChar, err := library.Get(name)
	if err != nil {
		return nil, err
	}

	// Build the character using the library patterns
	spec := NewCharacterSpec(libChar.Name, libChar.Width, libChar.Height)
	for _, frame := range libChar.Patterns {
		spec.AddFrame(frame.Name, frame.Lines)
	}

	// Convert to domain spec and use the new service layer
	domainSpec := ConvertLegacySpecToDomain(spec)
	characterService := NewCharacterService()
	domainChar, err := characterService.CreateCharacter(*domainSpec)
	if err != nil {
		return nil, err
	}

	// Wrap in AgentCharacter for state-based API
	return NewAgentCharacter(domainChar), nil
}

// ListLibrary returns all available library character names
func ListLibrary() []string {
	return library.List()
}

// UseLibrary retrieves a library character and registers it in the global registry
func UseLibrary(name string) (*Character, error) {
	char, err := Library(name)
	if err != nil {
		return nil, err
	}
	Register(char)
	return char, nil
}

// LibraryInfo returns information about a library character
func LibraryInfo(name string) (string, error) {
	libChar, err := library.Get(name)
	if err != nil {
		return "", err
	}
	return libChar.Description, nil
}

// Core API functions are already implemented in animator.go
// Animate() and ShowIdle() functions are available for immediate use
