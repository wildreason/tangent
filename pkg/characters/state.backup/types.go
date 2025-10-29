package state

// StateConfig defines the configuration for adding a new state to characters
type StateConfig struct {
	// Name of the state (e.g., "arise", "sleeping")
	Name string

	// Description of what the state represents
	Description string

	// Number of frames in the animation
	FrameCount int

	// Required height for the character (0 means no change)
	RequiredHeight int

	// Frame definitions (pattern code lines)
	Frames []FrameDefinition

	// Target characters ("all" or specific names)
	Targets []string

	// Template name (optional)
	Template string

	// Skip confirmation prompts
	SkipConfirmation bool

	// Skip documentation updates
	SkipDocs bool

	// Preview only (no changes)
	Preview bool

	// Dry run (show what would change)
	DryRun bool

	// Allow height updates
	AllowHeightUpdate bool

	// Custom CHANGELOG message
	ChangelogMessage string

	// Position where to insert frames ("after_base" or specific frame name)
	InsertPosition string
}

// FrameDefinition defines a single animation frame
type FrameDefinition struct {
	// Pattern code lines for this frame
	Lines []string

	// Optional custom frame name (if not using default naming)
	CustomName string
}

// StateAddResult contains the results of adding a state
type StateAddResult struct {
	// Characters that were successfully updated
	UpdatedCharacters []string

	// Characters that failed to update
	FailedCharacters map[string]error

	// Total frames added
	FramesAdded int

	// Documentation files updated
	DocsUpdated []string

	// Whether any height changes were made
	HeightChanged bool
}

// ValidationError represents a validation error with context
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (e *ValidationError) Error() string {
	if e.Value != nil {
		return e.Field + ": " + e.Message + " (value: " + toString(e.Value) + ")"
	}
	return e.Field + ": " + e.Message
}

func toString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// FrameValidationError represents an error in frame validation
type FrameValidationError struct {
	FrameIndex int
	LineIndex  int
	Message    string
}

func (e *FrameValidationError) Error() string {
	if e.LineIndex >= 0 {
		return "frame " + string(rune(e.FrameIndex+1)) + ", line " + string(rune(e.LineIndex+1)) + ": " + e.Message
	}
	return "frame " + string(rune(e.FrameIndex+1)) + ": " + e.Message
}
