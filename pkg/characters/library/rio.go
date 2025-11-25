package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "rio",
		Description: "rio - Musical note avatar (Orange)",
		Author:      "Wildreason, Inc",
		Color:       "#FF8800",
		Width:       11,
		Height:      4,
	}))
}
