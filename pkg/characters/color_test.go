package characters

import (
	"strings"
	"testing"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

func TestHexToRGB(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		r, g, b int
	}{
		{"with hash prefix", "#C0C0C0", 192, 192, 192},
		{"without hash prefix", "C0C0C0", 192, 192, 192},
		{"orange", "#FF6B35", 255, 107, 53},
		{"blue", "#1E90FF", 30, 144, 255},
		{"crimson", "#DC143C", 220, 20, 60},
		{"gold", "#FFD700", 255, 215, 0},
		{"purple", "#9370DB", 147, 112, 219},
		{"teal", "#20B2AA", 32, 178, 170},
		{"black", "#000000", 0, 0, 0},
		{"white", "#FFFFFF", 255, 255, 255},
		{"invalid short", "FFF", 0, 0, 0},
		{"invalid long", "#FFFFFFF", 0, 0, 0},
		{"empty", "", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := HexToRGB(tt.hex)
			if r != tt.r || g != tt.g || b != tt.b {
				t.Errorf("HexToRGB(%q) = (%d, %d, %d), want (%d, %d, %d)",
					tt.hex, r, g, b, tt.r, tt.g, tt.b)
			}
		})
	}
}

func TestColorizeString(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		hexColor string
		want     string
	}{
		{
			name:     "silver text",
			text:     "test",
			hexColor: "#C0C0C0",
			want:     "\x1b[38;2;192;192;192mtest\x1b[0m",
		},
		{
			name:     "orange text",
			text:     "hello",
			hexColor: "#FF6B35",
			want:     "\x1b[38;2;255;107;53mhello\x1b[0m",
		},
		{
			name:     "empty color returns unchanged",
			text:     "test",
			hexColor: "",
			want:     "test",
		},
		{
			name:     "unicode characters",
			text:     "  ▐▛███▜▌  ",
			hexColor: "#1E90FF",
			want:     "\x1b[38;2;30;144;255m  ▐▛███▜▌  \x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ColorizeString(tt.text, tt.hexColor)
			if got != tt.want {
				t.Errorf("ColorizeString(%q, %q) = %q, want %q",
					tt.text, tt.hexColor, got, tt.want)
			}
		})
	}
}

func TestColorizeFrame(t *testing.T) {
	tests := []struct {
		name     string
		frame    domain.Frame
		hexColor string
		validate func(*testing.T, []string)
	}{
		{
			name: "simple frame with color",
			frame: domain.Frame{
				Name: "test",
				Lines: []string{
					"FFF",
					"___",
				},
			},
			hexColor: "#C0C0C0",
			validate: func(t *testing.T, result []string) {
				if len(result) != 2 {
					t.Errorf("expected 2 lines, got %d", len(result))
				}
				// Should contain ANSI escape codes
				for i, line := range result {
					if !strings.Contains(line, "\x1b[38;2;192;192;192m") {
						t.Errorf("line %d missing color code: %q", i, line)
					}
					if !strings.Contains(line, "\x1b[0m") {
						t.Errorf("line %d missing reset code: %q", i, line)
					}
				}
			},
		},
		{
			name: "frame without color",
			frame: domain.Frame{
				Name: "test",
				Lines: []string{
					"FFF",
				},
			},
			hexColor: "",
			validate: func(t *testing.T, result []string) {
				if len(result) != 1 {
					t.Errorf("expected 1 line, got %d", len(result))
				}
				// Should NOT contain ANSI escape codes
				if strings.Contains(result[0], "\x1b[") {
					t.Errorf("line should not have color codes: %q", result[0])
				}
			},
		},
		{
			name: "multi-line frame",
			frame: domain.Frame{
				Name: "test",
				Lines: []string{
					"__R5FFF6L__",
					"_26FFFFF51_",
					"___11_22___",
				},
			},
			hexColor: "#FF6B35",
			validate: func(t *testing.T, result []string) {
				if len(result) != 3 {
					t.Errorf("expected 3 lines, got %d", len(result))
				}
				// All lines should be colored
				for i, line := range result {
					if !strings.Contains(line, "\x1b[38;2;255;107;53m") {
						t.Errorf("line %d missing orange color code: %q", i, line)
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ColorizeFrame(tt.frame, tt.hexColor)
			tt.validate(t, result)
		})
	}
}

func TestColorizeFrameIntegration(t *testing.T) {
	// Test with real character from library
	agent, err := LibraryAgent("sa")
	if err != nil {
		t.Fatalf("failed to load sa: %v", err)
	}

	char := agent.GetCharacter()
	if char.Color == "" {
		t.Fatal("sa character has no color")
	}

	// Get a frame from wait state
	waitState, exists := char.States["wait"]
	if !exists {
		t.Fatal("sa has no wait state")
	}

	if len(waitState.Frames) == 0 {
		t.Fatal("wait state has no frames")
	}

	frame := waitState.Frames[0]

	// Colorize the frame
	coloredLines := ColorizeFrame(frame, char.Color)

	// Verify
	if len(coloredLines) != len(frame.Lines) {
		t.Errorf("expected %d colored lines, got %d", len(frame.Lines), len(coloredLines))
	}

	// All lines should contain color codes
	for i, line := range coloredLines {
		if !strings.Contains(line, "\x1b[38;2;") {
			t.Errorf("line %d missing ANSI RGB color code: %q", i, line)
		}
		if !strings.Contains(line, "\x1b[0m") {
			t.Errorf("line %d missing reset code: %q", i, line)
		}
	}
}
