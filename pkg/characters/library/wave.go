package library

// Wave is a friendly greeting bot with waving hands
// Perfect for welcome messages, greetings, or friendly interactions
//
// Frames: 5 (wave1-wave5)
// Size: 13x5
// Author: Wildreason, Inc
// Use case: Welcome screens, greeting animations, friendly AI interactions

const WaveName = "wave"
const WaveWidth = 13
const WaveHeight = 5

var WaveFrames = map[string][]string{
	"wave1": {
		"___R6FFF6L___",
		"__T5FFFFF6T__",
		"_1_________2_",
		"_____BBB_____",
		"___________",
	},
	"wave2": {
		"___R6FFF6L___",
		"__T5FFFFF6T__",
		"_1_________2_",
		"_____BBB_____",
		"__L________",
	},
	"wave3": {
		"___R6FFF6L___",
		"__T5FFFFF6T__",
		"_1_________2_",
		"L____BBB_____",
		"___________",
	},
	"wave4": {
		"___R6FFF6L___",
		"__T5FFFFF6T__",
		"_1_________2_",
		"_____BBB_____",
		"_________R_",
	},
	"wave5": {
		"___R6FFF6L___",
		"__T5FFFFF6T__",
		"_1_________2_",
		"_____BBB____R",
		"___________",
	},
}

var WaveDescription = `Wave - Friendly Greeting Bot

A friendly bot character that waves hello with alternating hands.
Perfect for welcome messages, greetings, or any friendly AI interaction.

Animation: 5 frames showing left and right hand waves
Recommended FPS: 5-7 for natural waving motion

Use cases:
- Welcome screens
- Greeting animations
- Friendly AI interactions
- Onboarding sequences
- Happy status indicators`

func init() {
	frames := make([]Frame, 0, len(WaveFrames))
	for name, lines := range WaveFrames {
		frames = append(frames, Frame{Name: name, Lines: lines})
	}

	register(LibraryCharacter{
		Name:        WaveName,
		Description: WaveDescription,
		Author:      "Wildreason, Inc",
		Patterns:    frames,
		Width:       WaveWidth,
		Height:      WaveHeight,
	})
}

