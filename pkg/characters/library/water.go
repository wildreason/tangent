package library

func init() {
	register(waterCharacter)
}

var waterCharacter = LibraryCharacter{
	Name:        "water",
	Description: "water - efficient AI Agent Character",
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
				"L",
			},
		},
		{
			Name: "plan_2",
			Lines: []string{
				"R",
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
				"T",
			},
		},
		{
			Name: "think_2",
			Lines: []string{
				"B",
			},
		},
		{
			Name: "think_3",
			Lines: []string{
				".",
			},
		},
		{
			Name: "execute_1",
			Lines: []string{
				"6",
			},
		},
		{
			Name: "execute_2",
			Lines: []string{
				"5",
			},
		},
		{
			Name: "execute_3",
			Lines: []string{
				"7",
			},
		},
	},
}
