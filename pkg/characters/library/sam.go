package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "sam",
		Description: "sam - Musical note avatar (Pure Red)",
		Author:      "Wildreason, Inc",
		Color:       "#FF0000",
		Width:       11,
		Height:      4,
	}))
}
