package infrastructure

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

func TestFileCharacterRepository(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()
	repo := NewFileCharacterRepository(tempDir)

	// Test character
	testChar := &domain.Character{
		Name:   "test-char",
		Width:  5,
		Height: 3,
		Frames: []domain.Frame{
			{
				Name:  "frame1",
				Lines: []string{"█▐█", "▀▄▀", "▌▐▌"},
			},
		},
	}

	t.Run("Save character", func(t *testing.T) {
		err := repo.Save(testChar)
		if err != nil {
			t.Errorf("Save() error = %v", err)
		}

		// Check if file was created
		expectedFile := filepath.Join(tempDir, "test-char.json")
		if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
			t.Errorf("Expected file %s was not created", expectedFile)
		}
	})

	t.Run("Load character", func(t *testing.T) {
		loadedChar, err := repo.Load("test-char")
		if err != nil {
			t.Errorf("Load() error = %v", err)
		}

		if loadedChar.Name != testChar.Name {
			t.Errorf("Loaded character name = %v, want %v", loadedChar.Name, testChar.Name)
		}
		if loadedChar.Width != testChar.Width {
			t.Errorf("Loaded character width = %v, want %v", loadedChar.Width, testChar.Width)
		}
		if loadedChar.Height != testChar.Height {
			t.Errorf("Loaded character height = %v, want %v", loadedChar.Height, testChar.Height)
		}
		if len(loadedChar.Frames) != len(testChar.Frames) {
			t.Errorf("Loaded character frames count = %v, want %v", len(loadedChar.Frames), len(testChar.Frames))
		}
	})

	t.Run("List characters", func(t *testing.T) {
		names, err := repo.List()
		if err != nil {
			t.Errorf("List() error = %v", err)
		}

		if len(names) != 1 {
			t.Errorf("List() returned %d names, want 1", len(names))
		}
		if names[0] != "test-char" {
			t.Errorf("List() returned %v, want [test-char]", names)
		}
	})

	t.Run("Delete character", func(t *testing.T) {
		err := repo.Delete("test-char")
		if err != nil {
			t.Errorf("Delete() error = %v", err)
		}

		// Check if file was deleted
		expectedFile := filepath.Join(tempDir, "test-char.json")
		if _, err := os.Stat(expectedFile); !os.IsNotExist(err) {
			t.Errorf("Expected file %s was not deleted", expectedFile)
		}
	})

	t.Run("Load non-existent character", func(t *testing.T) {
		_, err := repo.Load("non-existent")
		if err == nil {
			t.Errorf("Load() should return error for non-existent character")
		}
	})

	t.Run("Delete non-existent character", func(t *testing.T) {
		err := repo.Delete("non-existent")
		if err == nil {
			t.Errorf("Delete() should return error for non-existent character")
		}
	})

	t.Run("Save nil character", func(t *testing.T) {
		err := repo.Save(nil)
		if err == nil {
			t.Errorf("Save() should return error for nil character")
		}
	})

	t.Run("Load empty ID", func(t *testing.T) {
		_, err := repo.Load("")
		if err == nil {
			t.Errorf("Load() should return error for empty ID")
		}
	})

	t.Run("Delete empty ID", func(t *testing.T) {
		err := repo.Delete("")
		if err == nil {
			t.Errorf("Delete() should return error for empty ID")
		}
	})
}
