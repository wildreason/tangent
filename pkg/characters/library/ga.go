package library

func init() {
	register(GenerateFromRegistry(CharacterMetadata{
		Name:        "ga",
		Description: "ga - Musical note avatar (Gold)",
		Author:      "Wildreason, Inc",
		Color:       "#FFD700",
		Width:       11,
		Height:      4,
	}))
}
