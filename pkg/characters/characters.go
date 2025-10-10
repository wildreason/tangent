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

// Package characters provides a simple API for creating and animating
// Unicode Block Elements (U+2580–U+259F) characters in Go applications.
//
// This package focuses exclusively on Block Elements for pixel-perfect
// character design and animation.
//
// Block Elements used:
// - █ (U+2588) Full Block
// - ▌ (U+258C) Left Half Block
// - ▐ (U+2590) Right Half Block
// - ▛ (U+259B) Quadrant Upper Left + Upper Right + Lower Left
// - ▜ (U+259C) Quadrant Upper Left + Upper Right + Lower Right
// - ▙ (U+2599) Quadrant Upper Left + Lower Left + Lower Right
// - ▟ (U+259F) Quadrant Upper Right + Lower Left + Lower Right
// - ▘ (U+2598) Quadrant Upper Left
// - ▝ (U+259D) Quadrant Upper Right
// - ▖ (U+2596) Quadrant Lower Left
// - ▗ (U+2597) Quadrant Lower Right
//
// Example usage:
//
//	char, err := NewBuilder("alien", 16, 16).
//		Pattern("▌▛███▜▐").
//		Block(16).
//		NewFrame().
//		Pattern("▌▛██  ").
//		Block(16).
//		Build()
//
//	Register(char)
//	Animate(os.Stdout, char, 6, 3)
//
// Library characters:
//
//	alien, _ := Library("alien")
//	Animate(os.Stdout, alien, 4, 2)

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
