package client

// QueueCondition defines when a queued state should be activated.
type QueueCondition interface {
	// shouldActivate returns true if the queued state should now become active.
	// loopCount is the number of completed loops for the current state.
	// frameCount is the total frames displayed for the current state.
	shouldActivate(loopCount, frameCount int) bool
}

// afterLoops activates after n complete animation loops.
type afterLoops struct {
	n int
}

func (a afterLoops) shouldActivate(loopCount, frameCount int) bool {
	return loopCount >= a.n
}

// AfterLoops creates a condition that activates after n complete loops.
// Example: AfterLoops(2) activates after the animation loops twice.
func AfterLoops(n int) QueueCondition {
	if n < 1 {
		n = 1
	}
	return afterLoops{n: n}
}

// afterFrames activates after n frames have been displayed.
type afterFrames struct {
	n int
}

func (a afterFrames) shouldActivate(loopCount, frameCount int) bool {
	return frameCount >= a.n
}

// AfterFrames creates a condition that activates after n frames.
// Example: AfterFrames(10) activates after 10 frames have been shown.
func AfterFrames(n int) QueueCondition {
	if n < 1 {
		n = 1
	}
	return afterFrames{n: n}
}

// immediate activates on the next tick.
type immediate struct{}

func (i immediate) shouldActivate(loopCount, frameCount int) bool {
	return true
}

// Immediate creates a condition that activates on the next tick.
// Useful for soft transitions that wait for the current frame to complete.
func Immediate() QueueCondition {
	return immediate{}
}

// queuedStateEntry holds a queued state transition.
type queuedStateEntry struct {
	state     string
	condition QueueCondition
	fps       int // 0 = use default/state FPS
}
