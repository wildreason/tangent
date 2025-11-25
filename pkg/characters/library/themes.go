package library

import (
	"fmt"
	"sort"
)

// ThemeDefinition represents a color theme for characters
type ThemeDefinition struct {
	Name        string
	Description string
	Colors      map[string]string // character name -> hex color
}

// Theme registry (private)
var themeRegistry = make(map[string]ThemeDefinition)

// registerTheme adds a theme to the registry (private)
func registerTheme(theme ThemeDefinition) {
	themeRegistry[theme.Name] = theme
}

// GetTheme returns a theme by name
func GetTheme(name string) (ThemeDefinition, error) {
	theme, exists := themeRegistry[name]
	if !exists {
		return ThemeDefinition{}, fmt.Errorf("theme %q not found", name)
	}
	return theme, nil
}

// ListThemes returns all available theme names in sorted order
func ListThemes() []string {
	names := make([]string, 0, len(themeRegistry))
	for name := range themeRegistry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// GetColor returns the color for a character in this theme
func (t ThemeDefinition) GetColor(characterName string) (string, error) {
	color, exists := t.Colors[characterName]
	if !exists {
		return "", fmt.Errorf("character %q not found in theme %q", characterName, t.Name)
	}
	return color, nil
}

// Theme initialization is now in themes_generated.go (auto-generated from constants.go)
