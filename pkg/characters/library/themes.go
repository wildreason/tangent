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

// Theme initialization - register all themes
func init() {
	// Theme 1: Bright (Original)
	registerTheme(ThemeDefinition{
		Name:        "bright",
		Description: "Original bright colors - 100% saturation for high impact and maximum distinction",
		Colors: map[string]string{
			CharacterSa: Theme1ColorSa,
			CharacterRi: Theme1ColorRi,
			CharacterGa: Theme1ColorGa,
			CharacterMa: Theme1ColorMa,
			CharacterPa: Theme1ColorPa,
			CharacterDa: Theme1ColorDa,
			CharacterNi: Theme1ColorNi,
		},
	})

	// Theme 2: Latte Window (Catppuccin-inspired - GUI user friendly)
	registerTheme(ThemeDefinition{
		Name:        "latte",
		Description: "Latte Window - Catppuccin-inspired warm pastels, GUI user friendly, approachable for non-developers",
		Colors: map[string]string{
			CharacterSa: Theme2ColorSa,
			CharacterRi: Theme2ColorRi,
			CharacterGa: Theme2ColorGa,
			CharacterMa: Theme2ColorMa,
			CharacterPa: Theme2ColorPa,
			CharacterDa: Theme2ColorDa,
			CharacterNi: Theme2ColorNi,
		},
	})

	// Theme 3: Garden Terminal (Earthy natural)
	registerTheme(ThemeDefinition{
		Name:        "garden",
		Description: "Garden Terminal - Earthy natural colors that reduce terminal intimidation, nature-inspired",
		Colors: map[string]string{
			CharacterSa: Theme3ColorSa,
			CharacterRi: Theme3ColorRi,
			CharacterGa: Theme3ColorGa,
			CharacterMa: Theme3ColorMa,
			CharacterPa: Theme3ColorPa,
			CharacterDa: Theme3ColorDa,
			CharacterNi: Theme3ColorNi,
		},
	})

	// Theme 4: Cozy Workspace (Modern GUI hybrid)
	registerTheme(ThemeDefinition{
		Name:        "cozy",
		Description: "Cozy Workspace - Modern GUI hybrid with professional warmth, Slack/Linear aesthetic",
		Colors: map[string]string{
			CharacterSa: Theme4ColorSa,
			CharacterRi: Theme4ColorRi,
			CharacterGa: Theme4ColorGa,
			CharacterMa: Theme4ColorMa,
			CharacterPa: Theme4ColorPa,
			CharacterDa: Theme4ColorDa,
			CharacterNi: Theme4ColorNi,
		},
	})
}
