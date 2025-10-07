package library

func init() {
	register(alienCharacter)
}

var alienCharacter = LibraryCharacter{
	Name:        "alien",
	Description: "Animated alien with waving hands - three-frame idle animation",
	Author:      "Wildreason, Inc",
	Width:       11,
	Height:      3,
	Patterns: []Frame{
		{
			Name: "idle",
			Lines: []string{
				"__R6FFF6L__",
				"_T6FFFFF5T_",
				"___11__22__",
			},
		},
		{
			Name: "left",
			Lines: []string{
				"__R6FFF6L__",
				"5T6FFFFF5T_",
				"__11__22___",
			},
		},
		{
			Name: "right",
			Lines: []string{
				"__R6FFF6L__",
				"_T6FFFFF5T6",
				"___11__22__",
			},
		},
	},
}
