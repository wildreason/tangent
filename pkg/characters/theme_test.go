package characters

import (
	"testing"

	"github.com/wildreason/tangent/pkg/characters/library"
)

func TestSetTheme(t *testing.T) {
	// Save original theme
	originalTheme := GetCurrentTheme()
	defer func() {
		SetTheme(originalTheme) // Restore original theme
	}()

	tests := []struct {
		name      string
		themeName string
		wantErr   bool
	}{
		{
			name:      "set bright theme",
			themeName: "bright",
			wantErr:   false,
		},
		{
			name:      "set latte theme",
			themeName: "latte",
			wantErr:   false,
		},
		{
			name:      "set garden theme",
			themeName: "garden",
			wantErr:   false,
		},
		{
			name:      "set cozy theme",
			themeName: "cozy",
			wantErr:   false,
		},
		{
			name:      "set nonexistent theme returns error",
			themeName: "nonexistent",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetTheme(tt.themeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetTheme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && GetCurrentTheme() != tt.themeName {
				t.Errorf("GetCurrentTheme() = %v, want %v", GetCurrentTheme(), tt.themeName)
			}
		})
	}
}

func TestGetCurrentTheme(t *testing.T) {
	// Save original theme
	originalTheme := GetCurrentTheme()
	defer func() {
		SetTheme(originalTheme) // Restore original theme
	}()

	// Default should be "latte"
	if GetCurrentTheme() != "latte" {
		t.Errorf("Default theme = %v, want latte", GetCurrentTheme())
	}

	// Change theme and verify
	SetTheme("bright")
	if GetCurrentTheme() != "bright" {
		t.Errorf("After SetTheme('bright'), GetCurrentTheme() = %v, want bright", GetCurrentTheme())
	}
}

func TestListThemes(t *testing.T) {
	themes := ListThemes()

	// Should have all 4 themes
	expectedCount := 4
	if len(themes) != expectedCount {
		t.Errorf("ListThemes() got %d themes, want %d", len(themes), expectedCount)
	}

	// Should include expected themes
	expectedThemes := map[string]bool{
		"bright": true,
		"latte":  true,
		"garden": true,
		"cozy":   true,
	}

	for _, theme := range themes {
		if !expectedThemes[theme] {
			t.Errorf("ListThemes() contains unexpected theme: %v", theme)
		}
	}
}

func TestLibraryAgentUsesTheme(t *testing.T) {
	// Save original theme
	originalTheme := GetCurrentTheme()
	defer func() {
		SetTheme(originalTheme) // Restore original theme
	}()

	// Test with latte theme (default)
	SetTheme("latte")
	agent, err := LibraryAgent("sa")
	if err != nil {
		t.Fatalf("LibraryAgent() error = %v", err)
	}

	char := agent.GetCharacter()
	if char.Color != library.Theme2ColorSa {
		t.Errorf("LibraryAgent with latte theme: color = %v, want %v", char.Color, library.Theme2ColorSa)
	}

	// Test with bright theme
	SetTheme("bright")
	agent2, err := LibraryAgent("sa")
	if err != nil {
		t.Fatalf("LibraryAgent() error = %v", err)
	}

	char2 := agent2.GetCharacter()
	if char2.Color != library.Theme1ColorSa {
		t.Errorf("LibraryAgent with bright theme: color = %v, want %v", char2.Color, library.Theme1ColorSa)
	}
}

func TestLibraryAgentThemeColors(t *testing.T) {
	// Save original theme
	originalTheme := GetCurrentTheme()
	defer func() {
		SetTheme(originalTheme) // Restore original theme
	}()

	themes := []struct {
		name      string
		themeName string
		character string
		wantColor string
	}{
		{
			name:      "bright theme sa",
			themeName: "bright",
			character: "sa",
			wantColor: library.Theme1ColorSa,
		},
		{
			name:      "latte theme sa",
			themeName: "latte",
			character: "sa",
			wantColor: library.Theme2ColorSa,
		},
		{
			name:      "garden theme ga",
			themeName: "garden",
			character: "ga",
			wantColor: library.Theme3ColorGa,
		},
		{
			name:      "cozy theme ma",
			themeName: "cozy",
			character: "ma",
			wantColor: library.Theme4ColorMa,
		},
	}

	for _, tt := range themes {
		t.Run(tt.name, func(t *testing.T) {
			err := SetTheme(tt.themeName)
			if err != nil {
				t.Fatalf("SetTheme() error = %v", err)
			}

			agent, err := LibraryAgent(tt.character)
			if err != nil {
				t.Fatalf("LibraryAgent() error = %v", err)
			}

			char := agent.GetCharacter()
			if char.Color != tt.wantColor {
				t.Errorf("LibraryAgent(%q) with %q theme: color = %v, want %v", tt.character, tt.themeName, char.Color, tt.wantColor)
			}
		})
	}
}

func TestThemeAPIMinimal(t *testing.T) {
	// Save original theme
	originalTheme := GetCurrentTheme()
	defer func() {
		SetTheme(originalTheme) // Restore original theme
	}()

	// Demonstrate minimal API usage
	// Step 1: Set theme
	err := SetTheme("latte")
	if err != nil {
		t.Fatalf("SetTheme() error = %v", err)
	}

	// Step 2: Create agent (automatically uses current theme)
	agent, err := LibraryAgent("sa")
	if err != nil {
		t.Fatalf("LibraryAgent() error = %v", err)
	}

	// Step 3: Verify color
	char := agent.GetCharacter()
	if char.Color != library.Theme2ColorSa {
		t.Errorf("Expected latte theme color %v, got %v", library.Theme2ColorSa, char.Color)
	}
}
