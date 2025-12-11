// Package client provides a framework-agnostic animation controller for tangent characters.
package client

// DefaultAliases maps common state names to tangent character states.
// These are applied automatically when SetState is called with an unrecognized name.
// Custom aliases set via SetAlias take precedence over these defaults.
var DefaultAliases = map[string]string{
	// Thinking/waiting states
	"think":    "wait",
	"thinking": "wait",

	// Writing/editing states
	"edit":    "write",
	"editing": "write",
	"writing": "write",

	// Reading states
	"reading": "read",

	// Search/tool states
	"bash":   "search",
	"grep":   "search",
	"glob":   "search",
	"find":   "search",
	"tool":   "search",
	"search": "search",

	// Success/approval states
	"success":  "approval",
	"complete": "approval",
	"done":     "approval",

	// Error states
	"fail":   "error",
	"failed": "error",

	// Arise/startup states
	"start":   "arise",
	"startup": "arise",
	"init":    "arise",
}
