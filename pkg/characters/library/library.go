package library

import (
	"fmt"
	"sort"
)

// Frame represents a single animation frame
type Frame struct {
	Name  string
	Lines []string
}

// LibraryCharacter represents a pre-built character from the library
type LibraryCharacter struct {
	Name         string
	Description  string
	Author       string
	ColorPalette map[string]string // Pattern code → hex color mapping (optional)
	Patterns     []Frame
	Width        int
	Height       int
}

// characters holds all available library characters
var libraryCharacters = make(map[string]LibraryCharacter)

// register adds a character to the library
func register(char LibraryCharacter) {
	libraryCharacters[char.Name] = char
}

// Get retrieves a library character by name (returns patterns, not built character)
func Get(name string) (LibraryCharacter, error) {
	libChar, exists := libraryCharacters[name]
	if !exists {
		return LibraryCharacter{}, fmt.Errorf("library character %q not found", name)
	}
	return libChar, nil
}

// List returns all available library character names in alphabetical order
func List() []string {
	names := make([]string, 0, len(libraryCharacters))
	for name := range libraryCharacters {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// All returns all library characters with their metadata
func All() map[string]LibraryCharacter {
	result := make(map[string]LibraryCharacter, len(libraryCharacters))
	for name, char := range libraryCharacters {
		result[name] = char
	}
	return result
}
