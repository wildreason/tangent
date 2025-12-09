package library

func init() {
	// Register all micro (10x2) character variants
	// All use the same micro patterns from microstateregistry, just with different colors

	registerMicro(GenerateMicroFromRegistry(CharacterMetadata{
		Name:        CharacterSa + "-micro",
		Description: "sam micro - Musical note avatar (Red)",
		Author:      "Wildreason, Inc",
		Color:       ColorSa,
	}))

	registerMicro(GenerateMicroFromRegistry(CharacterMetadata{
		Name:        CharacterRi + "-micro",
		Description: "rio micro - Musical note avatar (Orange)",
		Author:      "Wildreason, Inc",
		Color:       ColorRi,
	}))

	registerMicro(GenerateMicroFromRegistry(CharacterMetadata{
		Name:        CharacterGa + "-micro",
		Description: "ga micro - Musical note avatar (Gold)",
		Author:      "Wildreason, Inc",
		Color:       ColorGa,
	}))

	registerMicro(GenerateMicroFromRegistry(CharacterMetadata{
		Name:        CharacterMa + "-micro",
		Description: "ma micro - Musical note avatar (Green)",
		Author:      "Wildreason, Inc",
		Color:       ColorMa,
	}))

	registerMicro(GenerateMicroFromRegistry(CharacterMetadata{
		Name:        CharacterPa + "-micro",
		Description: "pa micro - Musical note avatar (Blue)",
		Author:      "Wildreason, Inc",
		Color:       ColorPa,
	}))

	registerMicro(GenerateMicroFromRegistry(CharacterMetadata{
		Name:        CharacterDa + "-micro",
		Description: "da micro - Musical note avatar (Purple)",
		Author:      "Wildreason, Inc",
		Color:       ColorDa,
	}))

	registerMicro(GenerateMicroFromRegistry(CharacterMetadata{
		Name:        CharacterNi + "-micro",
		Description: "ni micro - Musical note avatar (Pink)",
		Author:      "Wildreason, Inc",
		Color:       ColorNi,
	}))
}
