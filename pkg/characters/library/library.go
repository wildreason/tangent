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
	Name        string
	Description string
	Author      string
	Color       string // Single hex color for the entire character (e.g., "#FF4500")
	Patterns    []Frame
	Width       int
	Height      int
}

// characters holds all available library characters
var libraryCharacters = make(map[string]LibraryCharacter)

// microCharacters holds all available micro (10x2) library characters
var microCharacters = make(map[string]LibraryCharacter)

// register adds a character to the library
func register(char LibraryCharacter) {
	libraryCharacters[char.Name] = char
}

// registerMicro adds a micro character to the micro library
func registerMicro(char LibraryCharacter) {
	microCharacters[char.Name] = char
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

// GetMicro retrieves a micro library character by name
func GetMicro(name string) (LibraryCharacter, error) {
	// Try with -micro suffix first
	microName := name + "-micro"
	if libChar, exists := microCharacters[microName]; exists {
		return libChar, nil
	}
	// Try exact name
	if libChar, exists := microCharacters[name]; exists {
		return libChar, nil
	}
	return LibraryCharacter{}, fmt.Errorf("micro character %q not found", name)
}

// ListMicro returns all available micro character names in alphabetical order
func ListMicro() []string {
	names := make([]string, 0, len(microCharacters))
	for name := range microCharacters {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// AllMicro returns all micro library characters with their metadata
func AllMicro() map[string]LibraryCharacter {
	result := make(map[string]LibraryCharacter, len(microCharacters))
	for name, char := range microCharacters {
		result[name] = char
	}
	return result
}
