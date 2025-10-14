package library

func init() {
	register(testcharCharacter)
}

var testcharCharacter = LibraryCharacter{
	Name:        "testchar",
	Description: "testchar - friendly AI Agent Character",
	Author:      "Wildreason, Inc",
	Width:       7,
	Height:      3,
	Patterns: []Frame{
		{
			Name: "base",
			Lines: []string{
				"___F___",
				"__FFF__",
				"___F___",
			},
		},
		{
			Name: "plan_1",
			Lines: []string{
				"___F___",
				"__FFF__",
				"___F___",
			},
		},
		{
			Name: "plan_2",
			Lines: []string{
				"__F_F__",
				"_F___F_",
				"__F_F__",
			},
		},
		{
			Name: "think",
			Lines: []string{
				"___F___",
				"__FFF__",
				"___F___",
			},
		},
		{
			Name: "execute_1",
			Lines: []string{
				"___F___",
				"__FFF__",
				"___F___",
			},
		},
		{
			Name: "execute_2",
			Lines: []string{
				"___F___",
				"__FFF__",
				"___F___",
			},
		},
		{
			Name: "execute_3",
			Lines: []string{
				"___F___",
				"__FFF__",
				"___F___",
			},
		},
	},
}
