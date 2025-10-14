package library

func init() {
	register(egonCharacter)
}

var egonCharacter = LibraryCharacter{
	Name:        "egon",
	Description: "Egon - Efficient AI Agent Character",
	Author:      "Wildreason, Inc",
	Width:       11,
	Height:      3,
	Patterns: []Frame{
		// Base frame
		{
			Name: "base",
			Lines: []string{
				"____FF_____",
				"___FFF_____",
				"____FFF____",
			},
		},
		// plan state
		{
			Name: "plan",
			Lines: []string{
				"___FF______",
				"___FFF_____",
				"___FFFF____",
				"___FF______",
				"___FFFF____",
				"____FFF____",
			},
		},
		// think state
		{
			Name: "think",
			Lines: []string{
				"___FFF____F",
				"___FFFF___F",
				"___FFFFF__F",
				"F___FFF____",
				"___FFF_____",
				"___FFFFF___",
			},
		},
		// execute state
		{
			Name: "execute",
			Lines: []string{
				"___FFFFFF__",
				"___TTTTTT__",
				"___________",
				"___________",
				"____FFFF___",
				"____FF_____",
			},
		},
	},
}
