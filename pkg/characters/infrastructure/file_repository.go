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

// Save saves a character to a JSON file
func (r *FileCharacterRepository) Save(character *domain.Character) error {
	if character == nil {
		return fmt.Errorf("character cannot be nil")
	}

	data, err := json.MarshalIndent(character, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal character: %w", err)
	}

	filename := filepath.Join(r.baseDir, character.Name+".json")
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write character file: %w", err)
	}

	return nil
}

// Load loads a character from a JSON file
func (r *FileCharacterRepository) Load(id string) (*domain.Character, error) {
	if id == "" {
		return nil, fmt.Errorf("character ID cannot be empty")
	}

	filename := filepath.Join(r.baseDir, id+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, domain.ErrCharacterNotFound
		}
		return nil, fmt.Errorf("failed to read character file: %w", err)
	}

	var character domain.Character
	if err := json.Unmarshal(data, &character); err != nil {
		return nil, fmt.Errorf("failed to unmarshal character: %w", err)
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
