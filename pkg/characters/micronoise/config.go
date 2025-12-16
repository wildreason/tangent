package micronoise

import (
	"math"
	"math/rand"
)

// Config defines noise behavior for a micro avatar state
type Config struct {
	Count     int // Max noise slots at peak (picked from interior positions)
	Intensity int // Change frequency: 1=slow, 10=every frame
}

// StateConfigs maps active state names to their noise configuration.
// States not in this map (resting, wait) will have no noise.
// Count: max noise slots at peak (max 12 for 8x2)
// Intensity: 1=slow refresh, 10=every frame
var StateConfigs = map[string]Config{
	"arise":    {Count: 5, Intensity: 8},
	"read":     {Count: 6, Intensity: 10}, // Most used - max chaos
	"write":    {Count: 6, Intensity: 10},
	"search":   {Count: 5, Intensity: 9},
	"approval": {Count: 4, Intensity: 6},
}

// Phase timing constants (at 20 FPS)
const (
	PhaseRecognitionEnd = 15  // Frames 0-15: no noise (~0.75s)
	PhaseAwakeningEnd   = 40  // Frames 16-40: ramp up (~1.2s)
	BreathingFrequency  = 0.15 // Sin wave frequency for breathing
	BreathingAmplitude  = 0.5  // Amplitude as fraction of max count
)

// Fibonacci sequence for awakening phase
var fibonacciRamp = []int{0, 1, 1, 2, 3, 5, 8, 13}

// CalculateNoiseCount returns the effective noise count based on frame counter.
// Implements the hybrid breathing pattern:
//   - Phase 1 (Recognition): frames 0-15, returns 0
//   - Phase 2 (Awakening): frames 16-40, Fibonacci ramp-up
//   - Phase 3 (Breathing): frames 40+, sine wave oscillation
func CalculateNoiseCount(maxCount, frameCounter int) int {
	// Phase 1: Recognition - no noise, let user see shape
	if frameCounter <= PhaseRecognitionEnd {
		return 0
	}

	// Phase 2: Awakening - Fibonacci ramp-up
	if frameCounter <= PhaseAwakeningEnd {
		awakeningFrame := frameCounter - PhaseRecognitionEnd
		// Map awakening frames to fibonacci sequence
		fibIndex := awakeningFrame / 3 // Advance every 3 frames
		if fibIndex >= len(fibonacciRamp) {
			fibIndex = len(fibonacciRamp) - 1
		}
		count := fibonacciRamp[fibIndex]
		if count > maxCount {
			count = maxCount
		}
		return count
	}

	// Phase 3: Breathing - sine wave oscillation
	// count = base + amplitude * sin(frame * frequency)
	// base is half of max, amplitude is half of max
	// This gives range of 0 to maxCount
	base := float64(maxCount) * 0.5
	amplitude := float64(maxCount) * BreathingAmplitude
	breathingFrame := float64(frameCounter - PhaseAwakeningEnd)

	count := base + amplitude*math.Sin(breathingFrame*BreathingFrequency)

	// Clamp to valid range
	result := int(math.Round(count))
	if result < 0 {
		result = 0
	}
	if result > maxCount {
		result = maxCount
	}

	return result
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
