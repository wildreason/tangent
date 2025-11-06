package patterns

import (
	"strings"
	"testing"
)

func TestDefaultPatternCodes(t *testing.T) {
	codes := DefaultPatternCodes()

	tests := []struct {
		name     string
		value    rune
		expected rune
	}{
		{"FullBlock", codes.FullBlock, '█'},
		{"TopHalf", codes.TopHalf, '▀'},
		{"BottomHalf", codes.BottomHalf, '▄'},
		{"LeftHalf", codes.LeftHalf, '▌'},
		{"RightHalf", codes.RightHalf, '▐'},
		{"LightShade", codes.LightShade, '░'},
		{"MediumShade", codes.MediumShade, '▒'},
		{"DarkShade", codes.DarkShade, '▓'},
		{"Quad1", codes.Quad1, '▘'},
		{"Quad2", codes.Quad2, '▝'},
		{"Quad3", codes.Quad3, '▖'},
		{"Quad4", codes.Quad4, '▗'},
		{"Quad5", codes.Quad5, '▛'},
		{"Quad6", codes.Quad6, '▜'},
		{"Quad7", codes.Quad7, '▙'},
		{"Quad8", codes.Quad8, '▟'},
		{"DiagonalBackward", codes.DiagonalBackward, '▚'},
		{"DiagonalForward", codes.DiagonalForward, '▞'},
		{"Space", codes.Space, ' '},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("%s = %c, want %c", tt.name, tt.value, tt.expected)
			}
		})
	}
}

func TestGetPatternHelp(t *testing.T) {
	help := GetPatternHelp()

	// Verify the help string contains expected patterns
	expectedPatterns := []string{
		"Pattern codes:",
		"F=█", "T=▀", "B=▄", "L=▌", "R=▐",
		"1-8=quads",
		".=░", ":=▒", "#=▓",
		"_= ", // Space
	}

	for _, expected := range expectedPatterns {
		if !strings.Contains(help, expected) {
			t.Errorf("GetPatternHelp() should contain %q, got %q", expected, help)
		}
	}

	// Verify it's a single line format
	if strings.Count(help, "\n") > 0 {
		t.Errorf("GetPatternHelp() should be single line, got:\n%s", help)
	}
}

func TestGetPatternDescription(t *testing.T) {
	desc := GetPatternDescription()

	// Verify the description contains expected patterns
	expectedPatterns := []string{
		"Pattern codes:",
		"F=█", "T=▀", "B=▄", "L=▌", "R=▐",
		"(basic blocks)",
		"1-8=quads:",
		"▘▝▖▗", // Quad1-4
		"▛▜▙▟", // Quad5-8
		".=░", ":=▒", "#=▓",
		"(shades)",
		"_= ", // Space
		"(special)",
	}

	for _, expected := range expectedPatterns {
		if !strings.Contains(desc, expected) {
			t.Errorf("GetPatternDescription() should contain %q, got:\n%s", expected, desc)
		}
	}

	// Verify it's a multi-line format
	if strings.Count(desc, "\n") < 2 {
		t.Errorf("GetPatternDescription() should be multi-line, got: %s", desc)
	}
}
