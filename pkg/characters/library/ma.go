package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "ma",
		Description: "ma - Musical note avatar (Green)",
		Author:      "Wildreason, Inc",
		Color:       "#00FF00",
		Width:       11,
		Height:      4,
	}))
}
