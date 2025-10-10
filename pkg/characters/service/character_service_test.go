package service

import (
	"errors"
	"testing"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

// Mock implementations for testing
type mockCharacterRepository struct {
	characters map[string]*domain.Character
}

func newMockCharacterRepository() *mockCharacterRepository {
	return &mockCharacterRepository{
		characters: make(map[string]*domain.Character),
	}
}

func (m *mockCharacterRepository) Save(character *domain.Character) error {
	if character == nil {
		return errors.New("character cannot be nil")
	}
	m.characters[character.Name] = character
	return nil
}

func (m *mockCharacterRepository) Load(id string) (*domain.Character, error) {
	if id == "" {
		return nil, errors.New("character ID cannot be empty")
	}
	char, exists := m.characters[id]
	if !exists {
		return nil, domain.ErrCharacterNotFound
	}
	return char, nil
}

func (m *mockCharacterRepository) List() ([]string, error) {
	names := make([]string, 0, len(m.characters))
	for name := range m.characters {
		names = append(names, name)
	}
	return names, nil
}

func (m *mockCharacterRepository) Delete(id string) error {
	if id == "" {
		return errors.New("character ID cannot be empty")
	}
	if _, exists := m.characters[id]; !exists {
		return domain.ErrCharacterNotFound
	}
	delete(m.characters, id)
	return nil
}

type mockPatternCompiler struct{}

func (m *mockPatternCompiler) Compile(pattern string) string {
	// Simple mock implementation
	result := ""
	for _, char := range pattern {
		switch char {
		case 'F':
			result += "█"
		case 'R':
			result += "▐"
		case 'L':
			result += "▌"
		default:
			result += string(char)
		}
	}
	return result
}

func (m *mockPatternCompiler) Validate(pattern string) error {
	if len(pattern) == 0 {
		return domain.ErrEmptyPattern
	}
	return nil
}

type mockAnimationEngine struct{}

func (m *mockAnimationEngine) Animate(character *domain.Character, fps int, loops int) error {
	if character == nil {
		return errors.New("character cannot be nil")
	}
	if fps <= 0 {
		return errors.New("fps must be positive")
	}
	if loops <= 0 {
		return errors.New("loops must be positive")
	}
	return nil
}

func TestCharacterService_CreateCharacter(t *testing.T) {
	repo := newMockCharacterRepository()
	compiler := &mockPatternCompiler{}
	animationEngine := &mockAnimationEngine{}
	service := NewCharacterService(repo, compiler, animationEngine)

	tests := []struct {
		name    string
		spec    domain.CharacterSpec
		wantErr bool
	}{
		{
			name: "Valid character spec",
			spec: domain.CharacterSpec{
				Name:   "test",
				Width:  5,
				Height: 3,
				Frames: []domain.FrameSpec{
					{
						Name:     "frame1",
						Patterns: []string{"FRF", "LRL", "FRF"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Empty name",
			spec: domain.CharacterSpec{
				Name:   "",
				Width:  5,
				Height: 3,
				Frames: []domain.FrameSpec{
					{
						Name:     "frame1",
						Patterns: []string{"FRF", "LRL", "FRF"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid dimensions",
			spec: domain.CharacterSpec{
				Name:   "test",
				Width:  0,
				Height: 3,
				Frames: []domain.FrameSpec{
					{
						Name:     "frame1",
						Patterns: []string{"FRF", "LRL", "FRF"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "No frames",
			spec: domain.CharacterSpec{
				Name:   "test",
				Width:  5,
				Height: 3,
				Frames: []domain.FrameSpec{},
			},
			wantErr: true,
		},
		{
			name: "Empty frame name",
			spec: domain.CharacterSpec{
				Name:   "test",
				Width:  5,
				Height: 3,
				Frames: []domain.FrameSpec{
					{
						Name:     "",
						Patterns: []string{"FRF", "LRL", "FRF"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Wrong frame pattern count",
			spec: domain.CharacterSpec{
				Name:   "test",
				Width:  5,
				Height: 3,
				Frames: []domain.FrameSpec{
					{
						Name:     "frame1",
						Patterns: []string{"FRF", "LRL"}, // Only 2 patterns, should be 3
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			character, err := service.CreateCharacter(tt.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && character == nil {
				t.Errorf("CreateCharacter() returned nil character when no error expected")
			}
			if !tt.wantErr && character != nil {
				if character.Name != tt.spec.Name {
					t.Errorf("CreateCharacter() character name = %v, want %v", character.Name, tt.spec.Name)
				}
				if character.Width != tt.spec.Width {
					t.Errorf("CreateCharacter() character width = %v, want %v", character.Width, tt.spec.Width)
				}
				if character.Height != tt.spec.Height {
					t.Errorf("CreateCharacter() character height = %v, want %v", character.Height, tt.spec.Height)
				}
			}
		})
	}
}

func TestCharacterService_SaveCharacter(t *testing.T) {
	repo := newMockCharacterRepository()
	compiler := &mockPatternCompiler{}
	animationEngine := &mockAnimationEngine{}
	service := NewCharacterService(repo, compiler, animationEngine)

	testChar := &domain.Character{
		Name:   "test",
		Width:  5,
		Height: 3,
		Frames: []domain.Frame{
			{
				Name:  "frame1",
				Lines: []string{"█▐█", "▌▐▌", "█▐█"},
			},
		},
	}

	t.Run("Save valid character", func(t *testing.T) {
		err := service.SaveCharacter(testChar)
		if err != nil {
			t.Errorf("SaveCharacter() error = %v", err)
		}
	})

	t.Run("Save nil character", func(t *testing.T) {
		err := service.SaveCharacter(nil)
		if err == nil {
			t.Errorf("SaveCharacter() should return error for nil character")
		}
	})
}

func TestCharacterService_LoadCharacter(t *testing.T) {
	repo := newMockCharacterRepository()
	compiler := &mockPatternCompiler{}
	animationEngine := &mockAnimationEngine{}
	service := NewCharacterService(repo, compiler, animationEngine)

	testChar := &domain.Character{
		Name:   "test",
		Width:  5,
		Height: 3,
		Frames: []domain.Frame{
			{
				Name:  "frame1",
				Lines: []string{"█▐█", "▌▐▌", "█▐█"},
			},
		},
	}

	// Save the character first
	repo.Save(testChar)

	t.Run("Load existing character", func(t *testing.T) {
		loadedChar, err := service.LoadCharacter("test")
		if err != nil {
			t.Errorf("LoadCharacter() error = %v", err)
		}
		if loadedChar.Name != testChar.Name {
			t.Errorf("LoadCharacter() character name = %v, want %v", loadedChar.Name, testChar.Name)
		}
	})

	t.Run("Load non-existent character", func(t *testing.T) {
		_, err := service.LoadCharacter("non-existent")
		if err == nil {
			t.Errorf("LoadCharacter() should return error for non-existent character")
		}
	})

	t.Run("Load with empty ID", func(t *testing.T) {
		_, err := service.LoadCharacter("")
		if err == nil {
			t.Errorf("LoadCharacter() should return error for empty ID")
		}
	})
}

func TestCharacterService_AnimateCharacter(t *testing.T) {
	repo := newMockCharacterRepository()
	compiler := &mockPatternCompiler{}
	animationEngine := &mockAnimationEngine{}
	service := NewCharacterService(repo, compiler, animationEngine)

	testChar := &domain.Character{
		Name:   "test",
		Width:  5,
		Height: 3,
		Frames: []domain.Frame{
			{
				Name:  "frame1",
				Lines: []string{"█▐█", "▌▐▌", "█▐█"},
			},
		},
	}

	t.Run("Animate valid character", func(t *testing.T) {
		err := service.AnimateCharacter(testChar, 5, 3)
		if err != nil {
			t.Errorf("AnimateCharacter() error = %v", err)
		}
	})

	t.Run("Animate nil character", func(t *testing.T) {
		err := service.AnimateCharacter(nil, 5, 3)
		if err == nil {
			t.Errorf("AnimateCharacter() should return error for nil character")
		}
	})

	t.Run("Animate with invalid fps", func(t *testing.T) {
		err := service.AnimateCharacter(testChar, 0, 3)
		if err == nil {
			t.Errorf("AnimateCharacter() should return error for invalid fps")
		}
	})

	t.Run("Animate with invalid loops", func(t *testing.T) {
		err := service.AnimateCharacter(testChar, 5, 0)
		if err == nil {
			t.Errorf("AnimateCharacter() should return error for invalid loops")
		}
	})
}
