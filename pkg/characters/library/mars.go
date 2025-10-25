package library

func init() {
	register(marsCharacter)
}

var marsCharacter = LibraryCharacter{
	Name:        "mars",
	Description: "mars - War energy avatar (14 states)",
	Author:      "Wildreason, Inc",
	Width:       11,
	Height:      3,
	ColorPalette: map[string]string{
		"R": "#FF0000", // Red
		"L": "#DC143C", // Crimson
		"F": "#FF6B6B", // Light red
		"6": "#E63946", // Light crimson
		"5": "#D62828", // Medium red
		"4": "#C61818", // Dark red
		"2": "#8B0000", // Dark red
		"1": "#6B0000", // Darkest red
	},
	Patterns: []Frame{
		{
			Name: "base",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "wait_1",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "wait_2",
			Lines: []string{
				"__R5FFFFL__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "wait_3",
			Lines: []string{
				"__RFFFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "plan_1",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "plan_2",
			Lines: []string{
				"__R5FFF6L__",
				"TT6FFFFF5__",
				"___11_22___",
			},
		},
		{
			Name: "plan_3",
			Lines: []string{
				"__R5FFF6L__",
				"__6FFFFF5TT",
				"___11_22___",
			},
		},
		{
			Name: "think_1",
			Lines: []string{
				"__RF5F6FL__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "think_2",
			Lines: []string{
				"__RFFFFFL__",
				"_26F7F8F51_",
				"___11_22___",
			},
		},
		{
			Name: "think_3",
			Lines: []string{
				"__RFFFFFL__",
				"_26F8F7F51_",
				"___11_22___",
			},
		},
		{
			Name: "bash_1",
			Lines: []string{
				"__R5FFF6L__",
				"_26:::::51_",
				"__22___11__",
			},
		},
		{
			Name: "bash_2",
			Lines: []string{
				"__R5FFF6L__",
				"_26:...:51_",
				"__22___11__",
			},
		},
		{
			Name: "bash_3",
			Lines: []string{
				"__R5FFF6L  ",
				" 26#####51 ",
				"  22   11  ",
			},
		},
		{
			Name: "error_1",
			Lines: []string{
				"__RTTFTTL__",
				"_26BBFBB51_",
				"__2TTTTT1__",
			},
		},
		{
			Name: "error_2",
			Lines: []string{
				"__RTTFTTL__",
				"_26BBTBB51_",
				"__2TTTTT1__",
			},
		},
		{
			Name: "error_3",
			Lines: []string{
				"__RTTFTTL__",
				"_26BBFBB51_",
				"__2TTTTT1__",
			},
		},
		{
			Name: "success_1",
			Lines: []string{
				"__R5FFF6L__",
				"7B8FFFFF7B8",
				"___11_22___",
			},
		},
		{
			Name: "success_2",
			Lines: []string{
				"__R5FFF6L__",
				" 78FFFFF78 ",
				"___11_22___",
			},
		},
		{
			Name: "success_3",
			Lines: []string{
				"__R5FFF6L__",
				"7B8FFFFF7B8",
				"___11_22___",
			},
		},
		{
			Name: "read_1",
			Lines: []string{
				"__RFFFFFL__",
				"_26FRFLF51_",
				"___11_22___",
			},
		},
		{
			Name: "read_2",
			Lines: []string{
				"__RFRFLFL__",
				"_26FRFLF51_",
				"___11_22___",
			},
		},
		{
			Name: "read_3",
			Lines: []string{
				"__RFRFLFL__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "search_1",
			Lines: []string{
				"__RF5FF5L__",
				"__6FFFFF5T_",
				"____11_22__",
			},
		},
		{
			Name: "search_2",
			Lines: []string{
				"__RF5FF5L__",
				"__6FFFFF5T_",
				"___22__11__",
			},
		},
		{
			Name: "search_3",
			Lines: []string{
				"__RF5FF5L__",
				"__6FFFFF5T_",
				"___2_2_1_1_",
			},
		},
		{
			Name: "write_1",
			Lines: []string{
				"__RFFFFFL__",
				"_266F6FF51_",
				"__2TTTTT1__",
			},
		},
		{
			Name: "write_2",
			Lines: []string{
				"__RFFFF##__",
				"_266F65T___",
				"__2TTT1____",
			},
		},
		{
			Name: "write_3",
			Lines: []string{
				"__RFFF:::__",
				"_265F51____",
				"__2TTT_____",
			},
		},
		{
			Name: "build_1",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF511",
				"___11_22___",
			},
		},
		{
			Name: "build_2",
			Lines: []string{
				"__R5FFF6L__",
				"226FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "build_3",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "communicate_1",
			Lines: []string{
				"__R5FFF6L__",
				"_26FF#FF51_",
				"___11_22___",
			},
		},
		{
			Name: "communicate_2",
			Lines: []string{
				"__R5FFF6L__",
				"_26FF.FF51_",
				"___11_22___",
			},
		},
		{
			Name: "communicate_3",
			Lines: []string{
				"__R5FFF6L__",
				"_26FF:FF51_",
				"___11_22___",
			},
		},
		{
			Name: "blocked_1",
			Lines: []string{
				"__R57B86L__",
				"_48FFFFF73_",
				"____6T5____",
			},
		},
		{
			Name: "blocked_2",
			Lines: []string{
				"__R5L_R6L__",
				"_48FFFFF73_",
				"____6F5____",
			},
		},
		{
			Name: "blocked_3",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "block_1",
			Lines: []string{
				"__RFFFFFL__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "block_2",
			Lines: []string{
				"__RFFFFFL__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "block_3",
			Lines: []string{
				"__RFFBFFL__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "execute_1",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "execute_2",
			Lines: []string{
				"__R5FFF6L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
		{
			Name: "execute_3",
			Lines: []string{
				"__R6FFF5L__",
				"_26FFFFF51_",
				"___11_22___",
			},
		},
	},
}
