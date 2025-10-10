package service

import (
	"fmt"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

// CharacterService orchestrates character operations
type CharacterService struct {
	repo            domain.CharacterRepository
	compiler        domain.PatternCompiler
	animationEngine domain.AnimationEngine
}

// NewCharacterService creates a new character service
func NewCharacterService(repo domain.CharacterRepository, compiler domain.PatternCompiler, animationEngine domain.AnimationEngine) *CharacterService {
	return &CharacterService{
		repo:            repo,
		compiler:        compiler,
		animationEngine: animationEngine,
	}
}

// CreateCharacter creates a new character from a specification
func (s *CharacterService) CreateCharacter(spec domain.CharacterSpec) (*domain.Character, error) {
	// Validate character specification
	if err := s.validateCharacterSpec(spec); err != nil {
		return nil, fmt.Errorf("invalid character specification: %w", err)
	}

	// Compile patterns and create frames
	frames := make([]domain.Frame, len(spec.Frames))
	for i, frameSpec := range spec.Frames {
		if err := s.validateFrameSpec(frameSpec, spec.Height); err != nil {
			return nil, fmt.Errorf("invalid frame specification: %w", err)
		}

		// Compile patterns
		compiledLines := make([]string, len(frameSpec.Patterns))
		for j, pattern := range frameSpec.Patterns {
			if err := s.compiler.Validate(pattern); err != nil {
				return nil, fmt.Errorf("invalid pattern in frame %s: %w", frameSpec.Name, err)
			}
			compiledLines[j] = s.compiler.Compile(pattern)
		}

		frames[i] = domain.Frame{
			Name:  frameSpec.Name,
			Lines: compiledLines,
		}
	}

	return &domain.Character{
		Name:   spec.Name,
		Width:  spec.Width,
		Height: spec.Height,
		Frames: frames,
	}, nil
}

// SaveCharacter saves a character to the repository
func (s *CharacterService) SaveCharacter(character *domain.Character) error {
	if character == nil {
		return fmt.Errorf("character cannot be nil")
	}
	return s.repo.Save(character)
}

// LoadCharacter loads a character from the repository
func (s *CharacterService) LoadCharacter(id string) (*domain.Character, error) {
	if id == "" {
		return nil, fmt.Errorf("character ID cannot be empty")
	}
	return s.repo.Load(id)
}

// ListCharacters returns a list of all character IDs
func (s *CharacterService) ListCharacters() ([]string, error) {
	return s.repo.List()
}

// DeleteCharacter deletes a character from the repository
func (s *CharacterService) DeleteCharacter(id string) error {
	if id == "" {
		return fmt.Errorf("character ID cannot be empty")
	}
	return s.repo.Delete(id)
}

// AnimateCharacter animates a character
func (s *CharacterService) AnimateCharacter(character *domain.Character, fps int, loops int) error {
	if character == nil {
		return fmt.Errorf("character cannot be nil")
	}
	if fps <= 0 {
		return fmt.Errorf("fps must be positive")
	}
	if loops <= 0 {
		return fmt.Errorf("loops must be positive")
	}
	return s.animationEngine.Animate(character, fps, loops)
}

// validateCharacterSpec validates a character specification
func (s *CharacterService) validateCharacterSpec(spec domain.CharacterSpec) error {
	if spec.Name == "" {
		return domain.ErrCharacterNameRequired
	}
	if spec.Width <= 0 || spec.Height <= 0 {
		return domain.ErrInvalidDimensions
	}
	if len(spec.Frames) == 0 {
		return fmt.Errorf("character must have at least one frame")
	}
	return nil
}

// validateFrameSpec validates a frame specification
func (s *CharacterService) validateFrameSpec(frameSpec domain.FrameSpec, expectedHeight int) error {
	if frameSpec.Name == "" {
		return domain.ErrInvalidFrameName
	}
	if len(frameSpec.Patterns) != expectedHeight {
		return fmt.Errorf("frame %s has %d patterns, expected %d", frameSpec.Name, len(frameSpec.Patterns), expectedHeight)
	}
	for i, pattern := range frameSpec.Patterns {
		if pattern == "" {
			return fmt.Errorf("pattern %d in frame %s cannot be empty", i, frameSpec.Name)
		}
	}
	return nil
}
