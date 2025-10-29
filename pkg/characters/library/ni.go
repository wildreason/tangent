package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "ni",
		Description: "ni - Musical note avatar (Pink)",
		Author:      "Wildreason, Inc",
		Color:       "#FF0088",
		Width:       11,
		Height:      4,
	}))
}
