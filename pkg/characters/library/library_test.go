package library

import (
	"sort"
	"testing"
)

func TestGet_NotFound(t *testing.T) {
	_, err := Get("nonexistent_character")
	if err == nil {
		t.Error("Get() should return error for nonexistent character")
	}
	if err.Error() != `library character "nonexistent_character" not found` {
		t.Errorf("Error message = %q, want %q", err.Error(), `library character "nonexistent_character" not found`)
	}
}

func TestGet_Success(t *testing.T) {
	// Register a test character
	testChar := LibraryCharacter{
		Name:        "test_character",
		Description: "Test character",
		Author:      "Test Author",
		Color:       "#FF0000",
		Patterns: []Frame{
			{
				Name:  "idle",
				Lines: []string{"test", "line"},
			},
		},
		Width:  4,
		Height: 2,
	}
	register(testChar)

	char, err := Get("test_character")
	if err != nil {
		t.Errorf("Get() error = %v, want nil", err)
	}
	if char.Name != "test_character" {
		t.Errorf("Character name = %q, want %q", char.Name, "test_character")
	}
	if char.Description != "Test character" {
		t.Errorf("Character description = %q, want %q", char.Description, "Test character")
	}
	if char.Color != "#FF0000" {
		t.Errorf("Character color = %q, want %q", char.Color, "#FF0000")
	}
}

func TestList(t *testing.T) {
	// Clear and set up test characters
	libraryCharacters = make(map[string]LibraryCharacter)
	register(LibraryCharacter{Name: "zebra"})
	register(LibraryCharacter{Name: "apple"})
	register(LibraryCharacter{Name: "mango"})

	names := List()

	// Should return all character names
	if len(names) != 3 {
		t.Errorf("List() returned %d names, want 3", len(names))
	}

	// Should be sorted alphabetically
	if !sort.StringsAreSorted(names) {
		t.Errorf("List() should return sorted names, got %v", names)
	}

	// Verify the order
	expected := []string{"apple", "mango", "zebra"}
	for i, name := range expected {
		if names[i] != name {
			t.Errorf("List()[%d] = %q, want %q", i, names[i], name)
		}
	}
}

func TestList_Empty(t *testing.T) {
	// Clear all characters
	libraryCharacters = make(map[string]LibraryCharacter)

	names := List()

	if len(names) != 0 {
		t.Errorf("List() returned %d names, want 0", len(names))
	}
}

func TestAll(t *testing.T) {
	// Clear and set up test characters
	libraryCharacters = make(map[string]LibraryCharacter)

	char1 := LibraryCharacter{
		Name:        "char1",
		Description: "First character",
		Width:       4,
		Height:      2,
	}
	char2 := LibraryCharacter{
		Name:        "char2",
		Description: "Second character",
		Width:       6,
		Height:      3,
	}

	register(char1)
	register(char2)

	all := All()

	// Should return all characters
	if len(all) != 2 {
		t.Errorf("All() returned %d characters, want 2", len(all))
	}

	// Verify char1
	if c, exists := all["char1"]; !exists {
		t.Error("All() should contain char1")
	} else {
		if c.Description != "First character" {
			t.Errorf("char1.Description = %q, want %q", c.Description, "First character")
		}
		if c.Width != 4 {
			t.Errorf("char1.Width = %d, want %d", c.Width, 4)
		}
	}

	// Verify char2
	if c, exists := all["char2"]; !exists {
		t.Error("All() should contain char2")
	} else {
		if c.Description != "Second character" {
			t.Errorf("char2.Description = %q, want %q", c.Description, "Second character")
		}
		if c.Height != 3 {
			t.Errorf("char2.Height = %d, want %d", c.Height, 3)
		}
	}
}

func TestAll_Empty(t *testing.T) {
	// Clear all characters
	libraryCharacters = make(map[string]LibraryCharacter)

	all := All()

	if len(all) != 0 {
		t.Errorf("All() returned %d characters, want 0", len(all))
	}
}

func TestAll_ReturnsCopy(t *testing.T) {
	// Clear and set up test characters
	libraryCharacters = make(map[string]LibraryCharacter)
	register(LibraryCharacter{Name: "original"})

	all := All()

	// Modify the returned map
	all["new_char"] = LibraryCharacter{Name: "new_char"}

	// Original should not be modified
	if len(libraryCharacters) != 1 {
		t.Error("All() should return a copy, not the original map")
	}
	if _, exists := libraryCharacters["new_char"]; exists {
		t.Error("Modifying All() result should not affect the original library")
	}
}

func TestRegister(t *testing.T) {
	// Clear and test registration
	libraryCharacters = make(map[string]LibraryCharacter)

	char := LibraryCharacter{
		Name:        "registered_char",
		Description: "A registered character",
		Author:      "Test",
		Color:       "#00FF00",
		Width:       8,
		Height:      4,
	}

	register(char)

	// Verify it was registered
	if len(libraryCharacters) != 1 {
		t.Errorf("After register(), library should have 1 character, got %d", len(libraryCharacters))
	}

	retrieved, exists := libraryCharacters["registered_char"]
	if !exists {
		t.Error("Character should be registered in library")
	}
	if retrieved.Name != "registered_char" {
		t.Errorf("Retrieved name = %q, want %q", retrieved.Name, "registered_char")
	}
	if retrieved.Description != "A registered character" {
		t.Errorf("Retrieved description = %q, want %q", retrieved.Description, "A registered character")
	}
}
