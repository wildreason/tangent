package library

// Rocket is a launch animation for deployments and progress
// Perfect for CI/CD pipelines, deployment scripts, or progress indicators
//
// Frames: 4 (ignite, lift1, lift2, flight)
// Size: 7x7
// Author: Wildreason, Inc
// Use case: Deployment animations, launch sequences, progress indicators

const RocketName = "rocket"
const RocketWidth = 7
const RocketHeight = 7

var RocketFrames = map[string][]string{
	"ignite": {
		"___F___",
		"__5F6__",
		"__FFF__",
		"__FFF__",
		"___F___",
		"___:___",
		"___:___",
	},
	"lift1": {
		"___F___",
		"__5F6__",
		"__FFF__",
		"__FFF__",
		"___F___",
		"___#___",
		"___:___",
	},
	"lift2": {
		"___F___",
		"__5F6__",
		"__FFF__",
		"__FFF__",
		"___F___",
		"___F___",
		"___#___",
	},
	"flight": {
		"___F___",
		"__5F6__",
		"__FFF__",
		"__FFF__",
		"___F___",
		"_______",
		"_______",
	},
}

var RocketDescription = `Rocket - Launch Animation

A rocket launch sequence showing ignition, liftoff, and flight.
Perfect for deployment scripts, CI/CD pipelines, or any "launch" action.

Animation: 4 frames showing ignition through liftoff
Recommended FPS: 4-6 for dramatic launch effect

Use cases:
- Deployment animations
- CI/CD pipeline status
- Launch sequences
- Progress indicators
- Build/release celebrations`

func init() {
	frames := make([]Frame, 0, len(RocketFrames))
	for name, lines := range RocketFrames {
		frames = append(frames, Frame{Name: name, Lines: lines})
	}

	register(LibraryCharacter{
		Name:        RocketName,
		Description: RocketDescription,
		Author:      "Wildreason, Inc",
		Patterns:    frames,
		Width:       RocketWidth,
		Height:      RocketHeight,
	})
}

