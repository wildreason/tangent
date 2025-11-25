package library

import (
	"testing"
)

func TestGenerateFromRegistry(t *testing.T) {
	// Test that all 7 characters are generated
	chars := AllCharacterNames()

	for _, name := range chars {
		char, err := Get(name)
		if err != nil {
			t.Fatalf("Failed to get character %s: %v", name, err)
		}

		if char.Name != name {
			t.Errorf("Expected character name %s, got %s", name, char.Name)
		}

		if char.Width != 11 {
			t.Errorf("Expected width 11, got %d", char.Width)
		}

		if char.Height != 4 {
			t.Errorf("Expected height 4, got %d", char.Height)
		}

		if len(char.Patterns) == 0 {
			t.Errorf("Character %s has no patterns", name)
		}

		t.Logf("Character %s: %d patterns, color %s", name, len(char.Patterns), char.Color)
	}
}

func TestAllCharactersSamePatterns(t *testing.T) {
	// All characters should have identical patterns (same states)
	sam, _ := Get("sam")
	rio, _ := Get("rio")

	if len(sam.Patterns) != len(rio.Patterns) {
		t.Errorf("sam and rio have different pattern counts: %d vs %d", len(sam.Patterns), len(rio.Patterns))
	}

	// Verify frame names are identical
	for i := range sam.Patterns {
		if sam.Patterns[i].Name != rio.Patterns[i].Name {
			t.Errorf("Frame %d has different names: %s vs %s", i, sam.Patterns[i].Name, rio.Patterns[i].Name)
		}

		// Verify frame content is identical
		if len(sam.Patterns[i].Lines) != len(rio.Patterns[i].Lines) {
			t.Errorf("Frame %d has different line counts", i)
		}

		for j := range sam.Patterns[i].Lines {
			if sam.Patterns[i].Lines[j] != rio.Patterns[i].Lines[j] {
				t.Errorf("Frame %d line %d differs: %s vs %s", i, j, sam.Patterns[i].Lines[j], rio.Patterns[i].Lines[j])
			}
		}
	}

	t.Logf("All characters have identical %d patterns", len(sam.Patterns))
}

func TestCharacterColors(t *testing.T) {
	expectedColors := map[string]string{
		CharacterSa: ColorSa,
		CharacterRi: ColorRi,
		CharacterGa: ColorGa,
		CharacterMa: ColorMa,
		CharacterPa: ColorPa,
		CharacterDa: ColorDa,
		CharacterNi: ColorNi,
	}

	for name, expectedColor := range expectedColors {
		char, err := Get(name)
		if err != nil {
			t.Fatalf("Failed to get character %s: %v", name, err)
		}

		if char.Color != expectedColor {
			t.Errorf("Character %s has color %s, expected %s", name, char.Color, expectedColor)
		}
	}
}
