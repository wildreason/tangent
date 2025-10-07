package characters

import (
	"fmt"
	"sync"
)

// Registry manages character collections
type Registry struct {
	characters map[string]*Character
	mutex      sync.RWMutex
}

// NewRegistry creates a new character registry
func NewRegistry() *Registry {
	return &Registry{
		characters: make(map[string]*Character),
	}
}

// Register adds a character to the registry
func (r *Registry) Register(char *Character) error {
	if char == nil {
		return fmt.Errorf("cannot register nil character")
	}

	if char.Name == "" {
		return fmt.Errorf("character must have a name")
	}

	if err := char.Validate(); err != nil {
		return fmt.Errorf("character %s validation failed: %w", char.Name, err)
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.characters[char.Name] = char
	return nil
}

// Get retrieves a character by name
func (r *Registry) Get(name string) (*Character, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	char, exists := r.characters[name]
	if !exists {
		return nil, fmt.Errorf("character %s not found", name)
	}

	return char, nil
}

// List returns all registered character names
func (r *Registry) List() []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	names := make([]string, 0, len(r.characters))
	for name := range r.characters {
		names = append(names, name)
	}

	return names
}

// Remove removes a character from the registry
func (r *Registry) Remove(name string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.characters, name)
}

// Clear removes all characters from the registry
func (r *Registry) Clear() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.characters = make(map[string]*Character)
}

// Count returns the number of registered characters
func (r *Registry) Count() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return len(r.characters)
}

// Global registry instance
var defaultRegistry = NewRegistry()

// Register registers a character in the global registry
func Register(char *Character) error {
	return defaultRegistry.Register(char)
}

// Get retrieves a character from the global registry
func Get(name string) (*Character, error) {
	return defaultRegistry.Get(name)
}

// List returns all character names in the global registry
func List() []string {
	return defaultRegistry.List()
}

// Remove removes a character from the global registry
func Remove(name string) {
	defaultRegistry.Remove(name)
}

// Clear clears the global registry
func Clear() {
	defaultRegistry.Clear()
}
