package library

// Pulse is a heartbeat/thinking indicator animation
// Perfect for showing AI processing, loading states, or heartbeat effects
//
// Frames: 3 (beat1, beat2, beat3)
// Size: 9x5
// Author: Wildreason, Inc
// Use case: AI thinking indicator, loading animation, heartbeat monitor

const PulseName = "pulse"
const PulseWidth = 9
const PulseHeight = 5

var PulseFrames = map[string][]string{
	"beat1": {
		"____F____",
		"___5F6___",
		"__5FFF6__",
		"___5F6___",
		"____F____",
	},
	"beat2": {
		"____F____",
		"___FFF___",
		"__FFFFF__",
		"___FFF___",
		"____F____",
	},
	"beat3": {
		"____F____",
		"___###___",
		"__#####__",
		"___###___",
		"____F____",
	},
}

var PulseDescription = `Pulse - Heartbeat/Thinking Indicator

A pulsing heart animation that cycles through expanding and contracting states.
Perfect for showing AI processing, loading states, or any continuous activity.

Animation: 3 frames showing pulse expansion
Recommended FPS: 5-8 for smooth heartbeat effect

Use cases:
- AI thinking/processing indicator
- Loading animation
- Heartbeat monitor
- Activity indicator
- Background processing visual`

func init() {
	frames := make([]Frame, 0, len(PulseFrames))
	for name, lines := range PulseFrames {
		frames = append(frames, Frame{Name: name, Lines: lines})
	}

	register(LibraryCharacter{
		Name:        PulseName,
		Description: PulseDescription,
		Author:      "Wildreason, Inc",
		Patterns:    frames,
		Width:       PulseWidth,
		Height:      PulseHeight,
	})
}

