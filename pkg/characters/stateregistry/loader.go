package stateregistry

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

//go:embed states
var statesFS embed.FS

// DefaultRegistry is the global state registry loaded from embedded JSON files
var DefaultRegistry *Registry

func init() {
	var err error
	DefaultRegistry, err = LoadEmbedded()
	if err != nil {
		panic(fmt.Sprintf("Failed to load state registry: %v", err))
	}
}

// LoadEmbedded loads all states from embedded JSON files
func LoadEmbedded() (*Registry, error) {
	registry := NewRegistry()

	// Read all JSON files from embedded filesystem
	entries, err := statesFS.ReadDir("states")
	if err != nil {
		return nil, fmt.Errorf("failed to read states directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		// Read file
		path := filepath.Join("states", entry.Name())
		data, err := statesFS.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", entry.Name(), err)
		}

		// Parse JSON
		var state StateDefinition
		if err := json.Unmarshal(data, &state); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", entry.Name(), err)
		}

		// Register state
		registry.Register(state)
	}

	return registry, nil
}

// Get retrieves a state from the default registry
func Get(name string) (StateDefinition, bool) {
	return DefaultRegistry.Get(name)
}

// List returns all state names from the default registry
func List() []string {
	return DefaultRegistry.List()
}

// All returns all states from the default registry
func All() map[string]StateDefinition {
	return DefaultRegistry.All()
}
