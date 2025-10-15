package library

func init() {
	register(mercuryCharacter)
}

var mercuryCharacter = LibraryCharacter{
	Name:        "mercury",
	Description: "mercury - efficient AI Agent Character",
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
				":",
			},
		},
		{
			Name: "plan_2",
			Lines: []string{
				".",
			},
		},
		{
			Name: "plan_3",
			Lines: []string{
				"#",
			},
		},
		{
			Name: "think_1",
			Lines: []string{
				"F",
			},
		},
		{
			Name: "think_2",
			Lines: []string{
				"T",
			},
		},
		{
			Name: "think_3",
			Lines: []string{
				"B",
			},
		},
		{
			Name: "execute_1",
			Lines: []string{
				"F",
			},
		},
		{
			Name: "execute_2",
			Lines: []string{
				"R",
			},
		},
		{
			Name: "execute_3",
			Lines: []string{
				"L",
			},
		},
	},
}
