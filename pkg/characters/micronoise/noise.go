package micronoise

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ANSI escape code pattern: \x1b[38;2;R;G;Bm
var ansiColorRegex = regexp.MustCompile(`\x1b\[38;2;(\d+);(\d+);(\d+)m`)
var ansiResetRegex = regexp.MustCompile(`\x1b\[0m`)

// ApplyShiftingGradient applies a dark-to-light gradient that shifts left each frame.
// Creates a "marquee" / "Wall Street ticker" effect where brightness sweeps across.
//
// Parameters:
//   - lines: The avatar lines (ANSI colorized)
//   - width, height: Avatar dimensions
//   - frameCounter: Current frame number (controls gradient shift)
//   - cfg: Flicker configuration for the current state
//
// Each column has a brightness level from BrightnessLevels.
// Every frame, the pattern shifts left by 1 position (wrapping around).
func ApplyShiftingGradient(lines []string, width, height int, frameCounter int, cfg *FlickerConfig) []string {
	if cfg == nil || !cfg.Enabled || len(lines) == 0 {
		return lines
	}

	// Extract base color from first line
	baseR, baseG, baseB := extractColor(lines[0])

	// Parse lines into raw content (strip ANSI codes)
	rawLines := make([][]rune, len(lines))
	for i, line := range lines {
		rawLines[i] = []rune(stripANSI(line))
	}

	numLevels := len(BrightnessLevels)

	// Rebuild with shifted gradient colors
	result := make([]string, len(lines))
	for row := 0; row < len(rawLines); row++ {
		var sb strings.Builder
		lineLen := len(rawLines[row])

		for col := 0; col < lineLen; col++ {
			char := rawLines[row][col]

			// Calculate which brightness level this column gets
			// Shift pattern left by frameCounter positions (wrapping)
			brightnessIdx := (col + frameCounter) % numLevels
			brightness := BrightnessLevels[brightnessIdx]

			// Apply brightness to base color
			r := clamp(int(float64(baseR)*brightness), 0, 255)
			g := clamp(int(float64(baseG)*brightness), 0, 255)
			b := clamp(int(float64(baseB)*brightness), 0, 255)

			// Write character with brightness-adjusted color
			sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c\x1b[0m", r, g, b, char))
		}
		result[row] = sb.String()
	}

	return result
}

// Legacy function names for compatibility

// ApplyRandomFlicker redirects to ApplyShiftingGradient
func ApplyRandomFlicker(lines []string, width, height int, cfg *FlickerConfig) []string {
	// Without frameCounter, use static gradient (no animation)
	return ApplyShiftingGradient(lines, width, height, 0, cfg)
}

// ApplyColorWave redirects to ApplyShiftingGradient
func ApplyColorWave(lines []string, width, height int, frameCounter int, cfg *FlickerConfig) []string {
	return ApplyShiftingGradient(lines, width, height, frameCounter, cfg)
}

// extractColor extracts RGB values from ANSI colorized string.
// Returns (255, 255, 255) if no color found.
func extractColor(s string) (r, g, b int) {
	matches := ansiColorRegex.FindStringSubmatch(s)
	if len(matches) == 4 {
		r, _ = strconv.Atoi(matches[1])
		g, _ = strconv.Atoi(matches[2])
		b, _ = strconv.Atoi(matches[3])
		return
	}
	return 255, 255, 255
}

// HexToRGB converts hex color string to RGB values.
func HexToRGB(hex string) (int, int, int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 255, 255, 255
	}
	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return int(r), int(g), int(b)
}

// stripANSI removes all ANSI escape codes from a string.
func stripANSI(s string) string {
	s = ansiColorRegex.ReplaceAllString(s, "")
	s = ansiResetRegex.ReplaceAllString(s, "")
	return s
}

// clamp restricts a value to a range.
func clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// ApplyNoise is deprecated, kept for compatibility
func ApplyNoise(lines []string, width, height int, slots []int, activeCount int) []string {
	return lines
}
