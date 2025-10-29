package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "pa",
		Description: "pa - Musical note avatar (Blue)",
		Author:      "Wildreason, Inc",
		Color:       "#0088FF",
		Width:       11,
		Height:      4,
	}))
}
