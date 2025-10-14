package library

func init() {
	register(alexCharacter)
}

var alexCharacter = LibraryCharacter{
	Name:        "alex",
	Description: "alex - analytical AI Agent Character",
	Author:      "Wildreason, Inc",
	Width:       9,
	Height:      4,
	Patterns: []Frame{
		{
			Name: "base",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "plan_1",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "plan_2",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "plan_3",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "think_1",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "think_2",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "execute_1",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "execute_2",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "execute_3",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
		{
			Name: "execute_4",
			Lines: []string{
				"  ███  ",
				" █████ ",
				"███████",
				"  ███  ",
			},
		},
	},
}
