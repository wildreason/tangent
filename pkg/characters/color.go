package characters

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
)

// ColorizeFrame compiles pattern codes and applies ANSI RGB colors.
// Returns compiled lines with color codes applied.
// If hexColor is empty, returns compiled lines without color.
//
// This is useful for consumers who need pre-colored frames for custom
// animation systems (e.g., Bubble Tea TUI frameworks).
//
// Example:
//
//	agent, _ := characters.LibraryAgent("mercury")
//	char := agent.GetCharacter()
//	frame := char.States["wait"].Frames[0]
//	coloredLines := ColorizeFrame(frame, char.Color)
//	for _, line := range coloredLines {
//	    fmt.Println(line)  // Prints in silver color
//	}
func ColorizeFrame(frame domain.Frame, hexColor string) []string {
	compiler := infrastructure.NewPatternCompiler()
	result := make([]string, len(frame.Lines))

	for i, line := range frame.Lines {
		compiledLine := compiler.Compile(line)
		if hexColor != "" {
			compiledLine = ColorizeString(compiledLine, hexColor)
		}
		result[i] = compiledLine
	}

	return result
}

// ColorizeString wraps text with ANSI RGB color escape codes.
// Format: \x1b[38;2;R;G;Bm{text}\x1b[0m
// If hexColor is empty, returns text unchanged.
//
// This is useful for applying colors to arbitrary strings or
// for consumers who already have compiled frame lines.
//
// Example:
//
//	text := "  ▐▛███▜▌  "
//	colored := ColorizeString(text, "#C0C0C0")
//	fmt.Println(colored)  // Prints in silver color
func ColorizeString(text, hexColor string) string {
	if hexColor == "" {
		return text
	}
	r, g, b := HexToRGB(hexColor)
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, text)
}

// HexToRGB converts hex color string to RGB values.
// Accepts both "#RRGGBB" and "RRGGBB" formats.
// Returns (0, 0, 0) for invalid input.
//
// Example:
//
//	r, g, b := HexToRGB("#C0C0C0")  // r=192, g=192, b=192
//	r, g, b := HexToRGB("FF6B35")   // r=255, g=107, b=53
func HexToRGB(hex string) (r, g, b int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
		val, _ := strconv.ParseInt(hex, 16, 32)
		r = int((val >> 16) & 0xFF)
		g = int((val >> 8) & 0xFF)
		b = int(val & 0xFF)
	}
	return
}
