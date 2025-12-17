package client

import "time"

// DefaultIdleExpressions - 25 filler phrases for agent personality
// Used when no active events exist (idle state)
var DefaultIdleExpressions = []string{
	"hmm...",
	"uhh...",
	"well...",
	"let me think...",
	"pondering...",
	"...",
	"ah...",
	"okay...",
	"right...",
	"so...",
	"um...",
	"erm...",
	"hm.",
	"thinking...",
	"processing...",
	"considering...",
	"wait...",
	"one moment...",
	"let's see...",
	"now then...",
	"alright...",
	"mhm...",
	"indeed...",
	"interesting...",
	"curious...",
}

// ExpressionChangeInterval controls how often idle expression changes
const ExpressionChangeInterval = 2 * time.Second
