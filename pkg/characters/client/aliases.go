// Package client provides a framework-agnostic animation controller for tangent characters.
package client

// DefaultAliases maps common state names to tangent character states.
// These are applied automatically when SetState is called with an unrecognized name.
// Custom aliases set via SetAlias take precedence over these defaults.
var DefaultAliases = map[string]string{
	// Thinking states -> resting (calm idle)
	"think":    "resting",
	"thinking": "resting",

	// Tool/search states -> read (scanning)
	"bash": "read",
	"grep": "read",
	"glob": "read",
	"find": "read",

	// Writing/editing states -> write
	"edit":    "write",
	"editing": "write",
	"writing": "write",

	// Reading states -> read
	"reading": "read",

	// Success/approval states -> approval
	"success":  "approval",
	"complete": "approval",
	"done":     "approval",

	// Error/failure states -> wait (uncertain)
	"fail":   "wait",
	"failed": "wait",

	// Arise/startup states -> arise
	"start":   "arise",
	"startup": "arise",
}

// DefaultStateFPS provides recommended FPS for states.
// These can be applied via client.SetStateFPS() or used as reference.
var DefaultStateFPS = map[string]int{
	"resting":  2, // Slow, calm
	"wait":     3, // Moderate
	"read":     4, // Active scanning (2x for bash/grep)
	"write":    4, // Normal write speed
	"search":   6, // Fast scanning
	"approval": 3, // Moderate celebration
	"arise":    5, // Quick startup
}
