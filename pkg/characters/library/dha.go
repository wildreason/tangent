package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "dha",
		Description: "dha - Musical note avatar (Purple)",
		Author:      "Wildreason, Inc",
		Color:       "#8800FF",
		Width:       11,
		Height:      4,
	}))
}
