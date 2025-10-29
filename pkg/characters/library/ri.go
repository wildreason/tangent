package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "ri",
		Description: "ri - Musical note avatar (Orange)",
		Author:      "Wildreason, Inc",
		Color:       "#FF8800",
		Width:       11,
		Height:      4,
	}))
}
