package characters

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

// Animator handles character animation
type Animator struct {
	writer io.Writer
	fps    int
	loops  int
}

// NewAnimator creates a new animator
func NewAnimator(writer io.Writer, fps int) *Animator {
	if fps <= 0 {
		fps = 6
	}
	return &Animator{
		writer: writer,
		fps:    fps,
		loops:  1,
	}
}

// SetLoops sets the number of animation loops (0 = infinite)
func (a *Animator) SetLoops(loops int) *Animator {
	a.loops = loops
	return a
}

// Animate plays the character's animation
func (a *Animator) Animate(char *Character) error {
	if len(char.Frames) == 0 {
		return fmt.Errorf("character %s has no animation frames", char.Name)
	}

	// Hide cursor and save position
	fmt.Fprint(a.writer, "\x1b[?25l\x1b[s")
	defer fmt.Fprint(a.writer, "\x1b[?25h") // Show cursor on exit

	frameDur := time.Second / time.Duration(a.fps)
	totalLoops := a.loops
	if totalLoops <= 0 {
		totalLoops = math.MaxInt
	}

	for i := 0; i < totalLoops; i++ {
		for _, frame := range char.Frames {
			a.renderFrame(frame)
			time.Sleep(frameDur)
		}
	}

	// Render final frame cleanly
	finalFrame := char.Frames[len(char.Frames)-1]
	lines := strings.Split(finalFrame, "\n")
	for _, line := range lines {
		fmt.Fprintln(a.writer, line)
	}

	return nil
}

// ShowIdle displays the character's idle state
func (a *Animator) ShowIdle(char *Character) {
	fmt.Fprint(a.writer, char.Idle)
}

// renderFrame renders a single frame in place
func (a *Animator) renderFrame(frame string) {
	lines := strings.Split(frame, "\n")

	// Clear and print each line
	for _, line := range lines {
		fmt.Fprintf(a.writer, "\r\x1b[2K%s\n", line)
	}

	// Move cursor back up to overwrite the same area
	fmt.Fprintf(a.writer, "\x1b[%dA\x1b[u", len(lines))
}

// Simple animation function for quick use
func Animate(writer io.Writer, char *Character, fps int, loops int) error {
	animator := NewAnimator(writer, fps).SetLoops(loops)
	return animator.Animate(char)
}

// ShowIdle is a simple function to display idle state
func ShowIdle(writer io.Writer, char *Character) {
	animator := NewAnimator(writer, 0)
	animator.ShowIdle(char)
}
