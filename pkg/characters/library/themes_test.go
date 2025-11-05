package library

import (
	"testing"
)

func TestGetTheme(t *testing.T) {
	tests := []struct {
		name      string
		themeName string
		wantErr   bool
	}{
		{
			name:      "bright theme exists",
			themeName: "bright",
			wantErr:   false,
		},
		{
			name:      "latte theme exists",
			themeName: "latte",
			wantErr:   false,
		},
		{
			name:      "garden theme exists",
			themeName: "garden",
			wantErr:   false,
		},
		{
			name:      "cozy theme exists",
			themeName: "cozy",
			wantErr:   false,
		},
		{
			name:      "nonexistent theme returns error",
			themeName: "nonexistent",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			theme, err := GetTheme(tt.themeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTheme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && theme.Name != tt.themeName {
				t.Errorf("GetTheme() got name = %v, want %v", theme.Name, tt.themeName)
			}
		})
	}
}

func TestListThemes(t *testing.T) {
	themes := ListThemes()

	// Should have all 4 themes
	expectedCount := 4
	if len(themes) != expectedCount {
		t.Errorf("ListThemes() got %d themes, want %d", len(themes), expectedCount)
	}

	// Should be sorted
	expectedOrder := []string{"bright", "cozy", "garden", "latte"}
	for i, expected := range expectedOrder {
		if i >= len(themes) {
			t.Errorf("ListThemes() missing theme at index %d", i)
			continue
		}
		if themes[i] != expected {
			t.Errorf("ListThemes()[%d] = %v, want %v", i, themes[i], expected)
		}
	}
}

func TestThemeDefinition_GetColor(t *testing.T) {
	theme, err := GetTheme("bright")
	if err != nil {
		t.Fatalf("Failed to get bright theme: %v", err)
	}

	tests := []struct {
		name          string
		characterName string
		wantColor     string
		wantErr       bool
	}{
		{
			name:          "sa character has color",
			characterName: CharacterSa,
			wantColor:     Theme1ColorSa,
			wantErr:       false,
		},
		{
			name:          "ri character has color",
			characterName: CharacterRi,
			wantColor:     Theme1ColorRi,
			wantErr:       false,
		},
		{
			name:          "nonexistent character returns error",
			characterName: "nonexistent",
			wantColor:     "",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color, err := theme.GetColor(tt.characterName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && color != tt.wantColor {
				t.Errorf("GetColor() = %v, want %v", color, tt.wantColor)
			}
		})
	}
}

func TestThemeColors(t *testing.T) {
	// Test that all themes have colors for all characters
	themes := ListThemes()
	characters := AllCharacterNames()

	for _, themeName := range themes {
		theme, err := GetTheme(themeName)
		if err != nil {
			t.Errorf("Failed to get theme %q: %v", themeName, err)
			continue
		}

		t.Run(themeName, func(t *testing.T) {
			for _, charName := range characters {
				color, err := theme.GetColor(charName)
				if err != nil {
					t.Errorf("Theme %q missing color for character %q", themeName, charName)
					continue
				}
				if color == "" {
					t.Errorf("Theme %q has empty color for character %q", themeName, charName)
				}
				// Validate hex color format
				if len(color) != 7 || color[0] != '#' {
					t.Errorf("Theme %q character %q has invalid hex color: %q", themeName, charName, color)
				}
			}
		})
	}
}

func TestBrightThemeColors(t *testing.T) {
	theme, err := GetTheme("bright")
	if err != nil {
		t.Fatalf("Failed to get bright theme: %v", err)
	}

	// Verify bright theme uses Theme1 constants
	expectedColors := map[string]string{
		CharacterSa: Theme1ColorSa,
		CharacterRi: Theme1ColorRi,
		CharacterGa: Theme1ColorGa,
		CharacterMa: Theme1ColorMa,
		CharacterPa: Theme1ColorPa,
		CharacterDa: Theme1ColorDa,
		CharacterNi: Theme1ColorNi,
	}

	for charName, expectedColor := range expectedColors {
		color, err := theme.GetColor(charName)
		if err != nil {
			t.Errorf("Failed to get color for %q: %v", charName, err)
			continue
		}
		if color != expectedColor {
			t.Errorf("Bright theme %q color = %v, want %v", charName, color, expectedColor)
		}
	}
}

func TestLatteThemeColors(t *testing.T) {
	theme, err := GetTheme("latte")
	if err != nil {
		t.Fatalf("Failed to get latte theme: %v", err)
	}

	// Verify latte theme uses Theme2 constants
	expectedColors := map[string]string{
		CharacterSa: Theme2ColorSa,
		CharacterRi: Theme2ColorRi,
		CharacterGa: Theme2ColorGa,
		CharacterMa: Theme2ColorMa,
		CharacterPa: Theme2ColorPa,
		CharacterDa: Theme2ColorDa,
		CharacterNi: Theme2ColorNi,
	}

	for charName, expectedColor := range expectedColors {
		color, err := theme.GetColor(charName)
		if err != nil {
			t.Errorf("Failed to get color for %q: %v", charName, err)
			continue
		}
		if color != expectedColor {
			t.Errorf("Latte theme %q color = %v, want %v", charName, color, expectedColor)
		}
	}
}
