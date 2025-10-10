package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

// FileCharacterRepository implements the CharacterRepository interface
type FileCharacterRepository struct {
	baseDir string
}

// NewFileCharacterRepository creates a new file-based character repository
func NewFileCharacterRepository(baseDir string) domain.CharacterRepository {
	return &FileCharacterRepository{
		baseDir: baseDir,
	}
}

// Save saves a character to a JSON file with enhanced error handling
func (r *FileCharacterRepository) Save(character *domain.Character) error {
	if character == nil {
		return domain.NewValidationError("character", nil, "character cannot be nil")
	}

	// Validate character before saving
	if character.Name == "" {
		return domain.ErrCharacterNameRequired
	}

	if character.Width <= 0 || character.Height <= 0 {
		return domain.ErrInvalidDimensions
	}

	if len(character.Frames) == 0 {
		return domain.NewValidationError("frames", len(character.Frames), "character must have at least one frame")
	}

	data, err := json.MarshalIndent(character, "", "  ")
	if err != nil {
		return domain.NewValidationErrorWithCause("character", character.Name, "failed to marshal character to JSON", err)
	}

	filename := filepath.Join(r.baseDir, character.Name+".json")
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return domain.NewValidationErrorWithCause("file", filename, "failed to write character file", err)
	}

	return nil
}

// Load loads a character from a JSON file with enhanced error handling
func (r *FileCharacterRepository) Load(id string) (*domain.Character, error) {
	if id == "" {
		return nil, domain.NewValidationError("id", id, "character ID cannot be empty")
	}

	filename := filepath.Join(r.baseDir, id+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.NewCharacterNotFoundError(id)
		}
		return nil, domain.NewValidationErrorWithCause("file", filename, "failed to read character file", err)
	}

	var character domain.Character
	if err := json.Unmarshal(data, &character); err != nil {
		return nil, domain.NewValidationErrorWithCause("character", id, "failed to unmarshal character from JSON", err)
	}

	// Validate loaded character
	if character.Name == "" {
		return nil, domain.NewValidationError("character", id, "loaded character has empty name")
	}

	if character.Width <= 0 || character.Height <= 0 {
		return nil, domain.NewValidationError("character", id, "loaded character has invalid dimensions")
	}

	return &character, nil
}

// List returns a list of all character IDs
func (r *FileCharacterRepository) List() ([]string, error) {
	files, err := filepath.Glob(filepath.Join(r.baseDir, "*.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to list character files: %w", err)
	}

	var names []string
	for _, file := range files {
		name := filepath.Base(file)
		if len(name) > 5 && name[len(name)-5:] == ".json" {
			names = append(names, name[:len(name)-5])
		}
	}

	return names, nil
}

// Delete deletes a character file
func (r *FileCharacterRepository) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("character ID cannot be empty")
	}

	filename := filepath.Join(r.baseDir, id+".json")
	if err := os.Remove(filename); err != nil {
		if os.IsNotExist(err) {
			return domain.ErrCharacterNotFound
		}
		return fmt.Errorf("failed to delete character file: %w", err)
	}

	return nil
}
