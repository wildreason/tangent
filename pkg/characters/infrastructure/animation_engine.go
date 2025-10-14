package infrastructure

import (
	"fmt"
	"time"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

// SimpleAnimationEngine implements the AnimationEngine interface
type SimpleAnimationEngine struct{}

// NewAnimationEngine creates a new animation engine
func NewAnimationEngine() domain.AnimationEngine {
	return &SimpleAnimationEngine{}
}

// Animate animates a character with enhanced error handling
func (e *SimpleAnimationEngine) Animate(character *domain.Character, fps int, loops int) error {
	if character == nil {
		return domain.NewValidationError("character", nil, "character cannot be nil")
	}
	if fps <= 0 {
		return domain.NewValidationError("fps", fps, "fps must be positive")
	}
	if loops <= 0 {
		return domain.NewValidationError("loops", loops, "loops must be positive")
	}

	if len(character.Frames) == 0 {
		return domain.NewAnimationError(character.Name, "start", "character has no animation frames", nil)
	}

	// Hide cursor
	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[?25h") // Show cursor on exit

	frameDur := time.Second / time.Duration(fps)

	for loop := 0; loop < loops; loop++ {
		for frameIndex, frame := range character.Frames {
			if err := e.renderFrame(frame, character.Name, frameIndex); err != nil {
				return err
			}
			time.Sleep(frameDur)
		}
	}

	// Render final frame cleanly
	finalFrame := character.Frames[len(character.Frames)-1]
	for _, line := range finalFrame.Lines {
		fmt.Println(line)
	}

	return nil
}

// renderFrame renders a single frame in place with error handling
func (e *SimpleAnimationEngine) renderFrame(frame domain.Frame, characterName string, frameIndex int) error {
	if len(frame.Lines) == 0 {
		return domain.NewAnimationError(characterName, "frame_display",
			fmt.Sprintf("frame %s has no lines", frame.Name), nil)
	}

	// Clear and print each line
	for lineIndex, line := range frame.Lines {
		if line == "" {
			return domain.NewAnimationError(characterName, "frame_display",
				fmt.Sprintf("empty line %d in frame %s", lineIndex, frame.Name), nil)
		}
		fmt.Printf("\r\x1b[2K%s\n", line)
	}

	// Move cursor back up to overwrite the same area
	fmt.Printf("\x1b[%dA", len(frame.Lines))

	return nil
}
