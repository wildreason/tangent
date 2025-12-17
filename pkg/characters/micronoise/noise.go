package micronoise

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// ANSI escape code pattern: \x1b[38;2;R;G;Bm
var ansiColorRegex = regexp.MustCompile(`\x1b\[38;2;(\d+);(\d+);(\d+)m`)
var ansiResetRegex = regexp.MustCompile(`\x1b\[0m`)

// ApplyRandomFlicker assigns random theme colors to each block.
// Creates a "Wall Street rush" / stock ticker effect where each character
// independently flickers between the 7 theme colors.
//
// Parameters:
//   - lines: The avatar lines (ANSI colorized)
//   - width, height: Avatar dimensions (unused, kept for API compatibility)
//   - cfg: Flicker configuration for the current state
//
// Each frame tick, all 16 blocks (8x2) get a random color from ThemePalette.
// FPS controls chaos level: higher FPS = faster flicker = more chaos.
func ApplyRandomFlicker(lines []string, width, height int, cfg *FlickerConfig) []string {
	if cfg == nil || !cfg.Enabled || len(lines) == 0 {
		return lines
	}

	// Parse lines into raw content (strip ANSI codes)
	rawLines := make([][]rune, len(lines))
	for i, line := range lines {
		rawLines[i] = []rune(stripANSI(line))
	}

	// Rebuild with random colors per character
	result := make([]string, len(lines))
	for row := 0; row < len(rawLines); row++ {
		var sb strings.Builder
		for col := 0; col < len(rawLines[row]); col++ {
			char := rawLines[row][col]

			// Pick random color from palette
			colorIdx := rand.Intn(len(ThemePalette))
			hex := ThemePalette[colorIdx]
			r, g, b := HexToRGB(hex)

			// Write character with random color
			sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c\x1b[0m", r, g, b, char))
		}
		result[row] = sb.String()
	}

	return result
}

// HexToRGB converts hex color string to RGB values.
// Returns (255, 255, 255) if parsing fails.
func HexToRGB(hex string) (int, int, int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 255, 255, 255
	}
	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return 255, 255, 255
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return 255, 255, 255
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return 255, 255, 255
	}
	return int(r), int(g), int(b)
}

// stripANSI removes all ANSI escape codes from a string.
func stripANSI(s string) string {
	s = ansiColorRegex.ReplaceAllString(s, "")
	s = ansiResetRegex.ReplaceAllString(s, "")
	return s
}

// Legacy compatibility shims

// ApplyColorWave is deprecated, redirects to ApplyRandomFlicker
func ApplyColorWave(lines []string, width, height int, frameCounter int, cfg *FlickerConfig) []string {
	return ApplyRandomFlicker(lines, width, height, cfg)
}

// ApplyNoise is deprecated, kept for compatibility
func ApplyNoise(lines []string, width, height int, slots []int, activeCount int) []string {
	return lines
}
