package library

// Character name constants
// The seven musical notes (sargam) representing the standard character set
const (
	CharacterSa = "sa" // Shadja (Red)
	CharacterRi = "ri" // Rishabha (Orange)
	CharacterGa = "ga" // Gandhara (Gold)
	CharacterMa = "ma" // Madhyama (Green)
	CharacterPa = "pa" // Panchama (Blue)
	CharacterDa = "da" // Dhaivata (Purple)
	CharacterNi = "ni" // Nishada (Pink)
)

// AllCharacterNames returns all standard character names in order
func AllCharacterNames() []string {
	return []string{
		CharacterSa,
		CharacterRi,
		CharacterGa,
		CharacterMa,
		CharacterPa,
		CharacterDa,
		CharacterNi,
	}
}

// Character color constants
const (
	ColorSa = "#FF0000" // Red
	ColorRi = "#FF8800" // Orange
	ColorGa = "#FFD700" // Gold
	ColorMa = "#00FF00" // Green
	ColorPa = "#0088FF" // Blue
	ColorDa = "#8800FF" // Purple
	ColorNi = "#FF0088" // Pink
)
