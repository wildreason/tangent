package micronoise

import "math/rand"

// Config defines noise behavior for a micro avatar state
type Config struct {
	Count     int // Number of noise slots (picked from interior positions)
	Intensity int // Change frequency: 1=slow, 10=fast
}

// StateConfigs maps active state names to their noise configuration.
// States not in this map (resting, wait) will have no noise.
// Count: noise slots from interior positions (max 12 for 8x2)
// Intensity: 1=slow refresh, 10=every frame
var StateConfigs = map[string]Config{
	"arise":    {Count: 5, Intensity: 8},
	"read":     {Count: 6, Intensity: 10}, // Most used - max chaos
	"write":    {Count: 6, Intensity: 10},
	"search":   {Count: 5, Intensity: 9},
	"approval": {Count: 4, Intensity: 6},
}

// NoisePool contains block characters for dynamic noise effect.
// These create the "Wall Street rush" visual.
var NoisePool = []rune{
	'▘', '▝', '▖', '▗', // Single quadrants
	'▛', '▜', '▙', '▟', // Three-quadrant composites
	'░', '▒', '▓', // Shading blocks
	'▚', '▞', // Diagonals
}

// InteriorPositions defines positions that can have noise (excludes frame).
// For 8x2 grid: positions 0,7 (row 0) and 8,15 (row 1) are the frame.
// Interior: positions 1-6 on row 0, positions 9-14 on row 1.
var InteriorPositions = []int{1, 2, 3, 4, 5, 6, 9, 10, 11, 12, 13, 14}

// SelectSlots picks fixed noise slot positions from interior positions.
// These slots persist - only the characters change, not the positions.
func SelectSlots(count int) []int {
	if count <= 0 {
		return nil
	}
	if count > len(InteriorPositions) {
		count = len(InteriorPositions)
	}

	// Copy and shuffle interior positions
	positions := make([]int, len(InteriorPositions))
	copy(positions, InteriorPositions)
	rand.Shuffle(len(positions), func(i, j int) {
		positions[i], positions[j] = positions[j], positions[i]
	})

	// Return first 'count' positions
	return positions[:count]
}

// GetConfig returns the noise config for a state, or nil if no noise.
// Calm states (resting, wait) return nil.
func GetConfig(stateName string) *Config {
	if cfg, ok := StateConfigs[stateName]; ok {
		return &cfg
	}
	return nil
}

// ShouldRefresh determines if noise should change this tick
// based on intensity and a counter.
// Higher intensity = refresh more often.
func ShouldRefresh(counter, intensity int) bool {
	if intensity <= 0 {
		return false
	}
	// Intensity 10: every tick
	// Intensity 5: every 2 ticks
	// Intensity 1: every 10 ticks
	divisor := 11 - intensity
	if divisor < 1 {
		divisor = 1
	}
	return counter%divisor == 0
}
