package domain

// CharacterRepository defines the interface for character persistence
type CharacterRepository interface {
	Save(character *Character) error
	Load(id string) (*Character, error)
	List() ([]string, error)
	Delete(id string) error
}

// PatternCompiler defines the interface for pattern compilation
type PatternCompiler interface {
	Compile(pattern string) string
	Validate(pattern string) error
}

// AnimationEngine defines the interface for character animation
type AnimationEngine interface {
	Animate(character *Character, fps int, loops int) error
}
