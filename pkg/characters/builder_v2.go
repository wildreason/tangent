package characters

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
	"github.com/wildreason/tangent/pkg/characters/service"
)

// CharacterBuilderV2 provides a fluent API for building characters using the new architecture
type CharacterBuilderV2 struct {
	spec    *domain.CharacterSpec
	service *service.CharacterService
}

// NewCharacterBuilderV2 creates a new character builder with the given name and dimensions
func NewCharacterBuilderV2(name string, width, height int) *CharacterBuilderV2 {
	if name == "" {
		panic("character name cannot be empty")
	}
	if width <= 0 || height <= 0 {
		panic("character dimensions must be positive")
	}

	// Create service with default infrastructure
	homeDir, _ := os.UserHomeDir()
	baseDir := filepath.Join(homeDir, ".tangent")
	os.MkdirAll(baseDir, 0755)

	repo := infrastructure.NewFileCharacterRepository(baseDir)
	compiler := infrastructure.NewPatternCompiler()
	animationEngine := infrastructure.NewAnimationEngine()
	characterService := service.NewCharacterService(repo, compiler, animationEngine)

	return &CharacterBuilderV2{
		spec:    domain.NewCharacterSpec(name, width, height),
		service: characterService,
	}
}

// AddFrame adds a new frame to the character specification with enhanced error handling
func (b *CharacterBuilderV2) AddFrame(name string, patterns []string) *CharacterBuilderV2 {
	if name == "" {
		panic(domain.NewValidationError("frame_name", name, "frame name cannot be empty"))
	}
	if len(patterns) != b.spec.Height {
		panic(domain.NewValidationError("patterns", len(patterns),
			fmt.Sprintf("frame %s has %d patterns, expected %d", name, len(patterns), b.spec.Height)))
	}

	// Validate each pattern
	for i, pattern := range patterns {
		if pattern == "" {
			panic(domain.NewValidationError("pattern", i,
				fmt.Sprintf("pattern %d in frame %s cannot be empty", i, name)))
		}
	}

	b.spec.AddFrame(name, patterns)
	return b
}

// AddFrameFromString adds a frame from a single string pattern (split by newlines) with enhanced error handling
func (b *CharacterBuilderV2) AddFrameFromString(name, pattern string) *CharacterBuilderV2 {
	if name == "" {
		panic(domain.NewValidationError("frame_name", name, "frame name cannot be empty"))
	}
	if pattern == "" {
		panic(domain.NewValidationError("pattern", pattern, "pattern cannot be empty"))
	}

	// Split by newlines to get individual line patterns
	lines := strings.Split(pattern, "\n")
	patterns := make([]string, 0, len(lines))

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			patterns = append(patterns, line)
		} else if i < len(lines)-1 { // Allow empty lines at the end
			panic(domain.NewValidationError("pattern", i,
				fmt.Sprintf("empty line %d in pattern for frame %s", i, name)))
		}
	}

	if len(patterns) != b.spec.Height {
		panic(domain.NewValidationError("patterns", len(patterns),
			fmt.Sprintf("frame %s has %d patterns after splitting, expected %d", name, len(patterns), b.spec.Height)))
	}

	return b.AddFrame(name, patterns)
}

// Build creates the character using the service layer with enhanced error handling
func (b *CharacterBuilderV2) Build() (*domain.Character, error) {
	if len(b.spec.Frames) == 0 {
		return nil, domain.NewValidationError("frames", len(b.spec.Frames), "character must have at least one frame")
	}

	return b.service.CreateCharacter(*b.spec)
}

// BuildAndSave creates the character and saves it to the repository
func (b *CharacterBuilderV2) BuildAndSave() (*domain.Character, error) {
	character, err := b.Build()
	if err != nil {
		return nil, err
	}

	err = b.service.SaveCharacter(character)
	if err != nil {
		return nil, fmt.Errorf("failed to save character: %w", err)
	}

	return character, nil
}

// Validate validates the current character specification
func (b *CharacterBuilderV2) Validate() error {
	if b.spec.Name == "" {
		return domain.ErrCharacterNameRequired
	}
	if b.spec.Width <= 0 || b.spec.Height <= 0 {
		return domain.ErrInvalidDimensions
	}
	if len(b.spec.Frames) == 0 {
		return fmt.Errorf("character must have at least one frame")
	}

	// Validate each frame
	for _, frame := range b.spec.Frames {
		if frame.Name == "" {
			return domain.ErrInvalidFrameName
		}
		if len(frame.Patterns) != b.spec.Height {
			return fmt.Errorf("frame %s has %d patterns, expected %d", frame.Name, len(frame.Patterns), b.spec.Height)
		}
		for i, pattern := range frame.Patterns {
			if pattern == "" {
				return fmt.Errorf("pattern %d in frame %s cannot be empty", i, frame.Name)
			}
		}
	}

	return nil
}

// GetSpec returns the current character specification
func (b *CharacterBuilderV2) GetSpec() *domain.CharacterSpec {
	return b.spec
}

// SetService allows setting a custom service (useful for testing)
func (b *CharacterBuilderV2) SetService(service *service.CharacterService) *CharacterBuilderV2 {
	b.service = service
	return b
}
