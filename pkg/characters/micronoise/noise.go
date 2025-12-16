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

// ApplyNoise injects random noise characters at FIXED slot positions.
// Slots are pre-selected positions that persist - only characters change.
// Lines are expected to be ANSI-colorized (format: \x1b[38;2;R;G;Bm{text}\x1b[0m).
// Noise characters get random color variations from the base color.
func ApplyNoise(lines []string, width, height int, slots []int) []string {
	if len(slots) == 0 || len(lines) == 0 {
		return lines
	}

	// Extract base color from first line
	baseR, baseG, baseB := extractColor(lines[0])

	// Parse lines into raw content (strip ANSI codes)
	rawLines := make([][]rune, len(lines))
	for i, line := range lines {
		rawLines[i] = []rune(stripANSI(line))
	}

	// Build noise for each slot position
	noiseChars := make(map[int]struct {
		char    rune
		r, g, b int
	})

	for _, pos := range slots {
		row := pos / width
		col := pos % width

		if row < len(rawLines) && col < len(rawLines[row]) {
			// Random color variation from base
			r, g, b := varyColor(baseR, baseG, baseB)
			noiseChars[pos] = struct {
				char    rune
				r, g, b int
			}{
				char: NoisePool[rand.Intn(len(NoisePool))],
				r:    r, g: g, b: b,
			}
		}
	}

	// Rebuild lines with noise injected at fixed slots
	result := make([]string, len(lines))
	for row := 0; row < len(rawLines); row++ {
		var sb strings.Builder
		for col := 0; col < len(rawLines[row]); col++ {
			pos := row*width + col
			if noise, ok := noiseChars[pos]; ok {
				// Noise character with varied color
				sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c\x1b[0m", noise.r, noise.g, noise.b, noise.char))
			} else {
				// Original character with base color
				sb.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c\x1b[0m", baseR, baseG, baseB, rawLines[row][col]))
			}
		}
		result[row] = sb.String()
	}

	return result
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

// stripANSI removes all ANSI escape codes from a string.
func stripANSI(s string) string {
	s = ansiColorRegex.ReplaceAllString(s, "")
	s = ansiResetRegex.ReplaceAllString(s, "")
	return s
}

// varyColor creates a random variation of the base color.
// Adds/subtracts up to 60 from each RGB component for dramatic effect.
func varyColor(r, g, b int) (int, int, int) {
	variation := 60
	r = clamp(r + rand.Intn(variation*2+1) - variation, 0, 255)
	g = clamp(g + rand.Intn(variation*2+1) - variation, 0, 255)
	b = clamp(b + rand.Intn(variation*2+1) - variation, 0, 255)
	return r, g, b
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
