package domain

// PatternCompiler defines the interface for pattern compilation
type PatternCompiler interface {
	Compile(pattern string) string
	Validate(pattern string) error
}
