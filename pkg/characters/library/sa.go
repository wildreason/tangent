package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "sa",
		Description: "sa - Musical note avatar (Pure Red)",
		Author:      "Wildreason, Inc",
		Color:       "#FF0000",
		Width:       11,
		Height:      4,
	}))
}
