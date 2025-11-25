package library

//go:generate go run generator_codegen.go

// Character name constants
// The seven musical notes (sargam) representing the standard character set
const (
	CharacterSa = "sam" // Shadja (Red)
	CharacterRi = "rio" // Rishabha (Orange)
	CharacterGa = "ga"  // Gandhara (Gold)
	CharacterMa = "ma"  // Madhyama (Green)
	CharacterPa = "pa"  // Panchama (Blue)
	CharacterDa = "da"  // Dhaivata (Purple)
	CharacterNi = "ni"  // Nishada (Pink)
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

// Character color constants (Theme 1 - Original Bright)
const (
	ColorSa = "#FF0000" // Red
	ColorRi = "#FF8800" // Orange
	ColorGa = "#FFD700" // Gold
	ColorMa = "#00FF00" // Green
	ColorPa = "#0088FF" // Blue
	ColorDa = "#8800FF" // Purple
	ColorNi = "#FF0088" // Pink
)

// Theme 1: Original Bright (100% saturation - high impact, maximum distinction)
const (
	Theme1ColorSa = "#FF0000" // Pure red
	Theme1ColorRi = "#FF8800" // Bright orange
	Theme1ColorGa = "#FFD700" // Gold
	Theme1ColorMa = "#00FF00" // Pure green
	Theme1ColorPa = "#0088FF" // Bright blue
	Theme1ColorDa = "#8800FF" // Bright purple
	Theme1ColorNi = "#FF0088" // Bright pink
)

// Theme 2: Latte Window (Catppuccin-inspired - GUI user friendly, warm pastels)
const (
	Theme2ColorSa = "#E78284" // Soft coral - approachable warmth
	Theme2ColorRi = "#EF9F76" // Peach - friendly energy
	Theme2ColorGa = "#E5C890" // Warm sand - comfort
	Theme2ColorMa = "#A6D189" // Sage - growth/trust
	Theme2ColorPa = "#85C1DC" // Sky blue - calm clarity
	Theme2ColorDa = "#CA9EE6" // Lavender - creative
	Theme2ColorNi = "#F4B8E4" // Rose - welcoming
)

// Theme 3: Garden Terminal (Earthy natural - reduces terminal intimidation)
const (
	Theme3ColorSa = "#D4787D" // Terracotta rose - earthy warmth
	Theme3ColorRi = "#D89C6A" // Clay orange - natural
	Theme3ColorGa = "#C9B68C" // Wheat - harvest calm
	Theme3ColorMa = "#8FB378" // Moss green - forest trust
	Theme3ColorPa = "#7CA8B8" // River blue - flowing ease
	Theme3ColorDa = "#A888BA" // Dusty iris - twilight
	Theme3ColorNi = "#C895A8" // Dusty mauve - soft dusk
)

// Theme 4: Cozy Workspace (Modern GUI hybrid - professional warmth)
const (
	Theme4ColorSa = "#E18B8B" // Rose quartz - gentle strength
	Theme4ColorRi = "#E5A679" // Amber glow - warm energy
	Theme4ColorGa = "#E6CC94" // Honey - productive sweetness
	Theme4ColorMa = "#99C794" // Balanced green - growth
	Theme4ColorPa = "#78AED4" // Clear sky - open possibility
	Theme4ColorDa = "#B592D4" // Soft violet - creative depth
	Theme4ColorNi = "#DE99B8" // Blush - approachable
)
