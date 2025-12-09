package microstateregistry

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed states/micro.json
var microFS embed.FS

// DefaultDefinition is the global micro avatar definition loaded from embedded JSON
var DefaultDefinition *MicroDefinition

func init() {
	var err error
	DefaultDefinition, err = LoadEmbedded()
	if err != nil {
		panic(fmt.Sprintf("Failed to load micro state registry: %v", err))
	}
}

// LoadEmbedded loads the micro definition from the embedded JSON file
func LoadEmbedded() (*MicroDefinition, error) {
	data, err := microFS.ReadFile("states/micro.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read micro.json: %w", err)
	}

	var def MicroDefinition
	if err := json.Unmarshal(data, &def); err != nil {
		return nil, fmt.Errorf("failed to parse micro.json: %w", err)
	}

	return &def, nil
}

// Get returns the default micro definition
func Get() *MicroDefinition {
	return DefaultDefinition
}

// GetState retrieves a specific state by name
func GetState(name string) *MicroState {
	if DefaultDefinition == nil {
		return nil
	}
	for i := range DefaultDefinition.States {
		if DefaultDefinition.States[i].Name == name {
			return &DefaultDefinition.States[i]
		}
	}
	return nil
}

// GetBaseFrame returns the base frame from the default definition
func GetBaseFrame() MicroFrame {
	if DefaultDefinition == nil {
		return MicroFrame{}
	}
	return DefaultDefinition.BaseFrame
}

// ListStates returns all state names
func ListStates() []string {
	if DefaultDefinition == nil {
		return nil
	}
	names := make([]string, len(DefaultDefinition.States))
	for i, s := range DefaultDefinition.States {
		names[i] = s.Name
	}
	return names
}

// Width returns the micro avatar width
func Width() int {
	if DefaultDefinition == nil {
		return 0
	}
	return DefaultDefinition.Width
}

// Height returns the micro avatar height
func Height() int {
	if DefaultDefinition == nil {
		return 0
	}
	return DefaultDefinition.Height
}
