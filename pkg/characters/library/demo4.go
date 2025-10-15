package library

func init() {
	register(demo4Character)
}

var demo4Character = LibraryCharacter{
	Name:        "demo4",
	Description: "demo4 - efficient AI Agent Character",
	Author:      "Wildreason, Inc",
	Width:       1,
	Height:      1,
	Patterns: []Frame{
		{
			Name: "base",
			Lines: []string{
				"F",
			},
		},
		{
			Name: "plan_1",
			Lines: []string{
				"F",
			},
		},
		{
			Name: "plan_2",
			Lines: []string{
				":",
			},
		},
		{
			Name: "plan_3",
			Lines: []string{
				".",
			},
		},
	},
}
