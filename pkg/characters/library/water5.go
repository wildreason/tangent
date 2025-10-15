package library

func init() {
	register(water5Character)
}

var water5Character = LibraryCharacter{
	Name:        "water5",
	Description: "water5 - efficient AI Agent Character",
	Author:      "Wildreason, Inc",
	Width:       2,
	Height:      2,
	Patterns: []Frame{
		{
			Name: "base",
			Lines: []string{
				"FF",
				"FF",
			},
		},
		{
			Name: "plan_1",
			Lines: []string{
				"R6",
				"FF",
			},
		},
		{
			Name: "plan_2",
			Lines: []string{
				"RF",
				"F8",
			},
		},
		{
			Name: "plan_3",
			Lines: []string{
				"RF",
				"5F",
			},
		},
		{
			Name: "think_1",
			Lines: []string{
				"TF",
				"FF",
			},
		},
		{
			Name: "think_2",
			Lines: []string{
				"FT",
				"FF",
			},
		},
		{
			Name: "think_3",
			Lines: []string{
				"FF",
				"FB",
			},
		},
		{
			Name: "execute_1",
			Lines: []string{
				"FL",
				"FF",
			},
		},
		{
			Name: "execute_2",
			Lines: []string{
				"FF",
				"FL",
			},
		},
		{
			Name: "execute_3",
			Lines: []string{
				"FF",
				"RF",
			},
		},
	},
}
