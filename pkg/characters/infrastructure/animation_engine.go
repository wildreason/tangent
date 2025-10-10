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

// Animate animates a character
func (e *SimpleAnimationEngine) Animate(character *domain.Character, fps int, loops int) error {
	if character == nil {
		return fmt.Errorf("character cannot be nil")
	}
	if fps <= 0 {
		return fmt.Errorf("fps must be positive")
	}
	if loops <= 0 {
		return fmt.Errorf("loops must be positive")
	}

	if len(character.Frames) == 0 {
		return fmt.Errorf("character %s has no animation frames", character.Name)
	}

	// Hide cursor and save position
	fmt.Print("\x1b[?25l\x1b[s")
	defer fmt.Print("\x1b[?25h") // Show cursor on exit

	frameDur := time.Second / time.Duration(fps)

	for loop := 0; loop < loops; loop++ {
		for _, frame := range character.Frames {
			e.renderFrame(frame)
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

// renderFrame renders a single frame in place
func (e *SimpleAnimationEngine) renderFrame(frame domain.Frame) {
	// Clear and print each line
	for _, line := range frame.Lines {
		fmt.Printf("\r\x1b[2K%s\n", line)
	}

	// Move cursor back up to overwrite the same area
	fmt.Printf("\x1b[%dA\x1b[u", len(frame.Lines))
}
