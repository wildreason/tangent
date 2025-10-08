package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wildreason/tangent/pkg/characters"
)

var (
	version = "v0.1.0-beta.2"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Non-interactive CLI mode
	if len(os.Args) > 1 {
		handleCLI()
		return
	}

	// Interactive mode
	showBanner()
	mainMenu()
}

func showBanner() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  TANGENT - Terminal Agent Designer      ║")
	fmt.Println("║  Design characters for your CLI agents  ║")
	fmt.Printf("║  %-40s ║\n", version)
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
}

func mainMenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("▢ MAIN MENU")
		fmt.Println("  1. Create new character")
		fmt.Println("  2. Load character project")
		fmt.Println("  3. Browse library characters")
		fmt.Println("  4. Preview library character")
		fmt.Println("  5. View palette")
		fmt.Println("  6. Exit")
		fmt.Println()
		fmt.Print("◢ Choose option (1-6): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\n✓ Goodbye!")
				return
			}
			fmt.Printf("\n✗ Error reading input: %v\n", err)
			return
		}
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			createCharacter()
		case "2":
			loadCharacter()
		case "3":
			browseLibrary()
		case "4":
			previewLibrary()
		case "5":
			showPalette()
		case "6":
			fmt.Println("\n✓ Thanks for using Tangent!")
			os.Exit(0)
		default:
			fmt.Println("✗ Invalid option. Please choose 1-6.\n")
		}
	}
}

func createCharacter() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  CREATE NEW CHARACTER                    ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	// Get character name
	fmt.Print("◢ Character name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("✗ Name cannot be empty\n")
		return
	}

	// Check if session exists
	if sessionExists(name) {
		fmt.Printf("✗ Character '%s' already exists. Use 'Load character project' to continue working on it.\n\n", name)
		return
	}

	// Get dimensions
	width := getIntInput("◢ Enter width (e.g., 11): ", 1, 100)
	height := getIntInput("◢ Enter height (e.g., 3): ", 1, 50)

	fmt.Println()
	fmt.Printf("✓ Creating character '%s' (%dx%d)\n\n", name, width, height)

	// Create session
	session := NewSession(name, width, height)
	session.Save()

	// Enter character builder
	characterBuilder(session)
}

func loadCharacter() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  LOAD CHARACTER PROJECT                  ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	sessions, err := ListSessions()
	if err != nil || len(sessions) == 0 {
		fmt.Println("✗ No saved character projects found\n")
		return
	}

	fmt.Println("▢ Saved Projects:")
	for i, name := range sessions {
		fmt.Printf("  %d. %s\n", i+1, name)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("◢ Choose project (number or name): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var sessionName string
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(sessions) {
		sessionName = sessions[num-1]
	} else {
		sessionName = input
	}

	session, err := LoadSession(sessionName)
	if err != nil {
		fmt.Printf("✗ Failed to load project: %v\n\n", err)
		return
	}

	fmt.Printf("✓ Loaded '%s' (%dx%d) with %d frame(s)\n\n", session.Name, session.Width, session.Height, len(session.Frames))

	characterBuilder(session)
}

func characterBuilder(session *Session) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("▢ CHARACTER: " + session.Name)
		fmt.Printf("  Dimensions: %dx%d | Frames: %d\n", session.Width, session.Height, len(session.Frames))
		fmt.Println()

		if len(session.Frames) == 0 {
			fmt.Println("  ◢ Tip: Start by adding your first frame!")
			fmt.Println()
		}

		fmt.Println("  1. Add new frame")
		fmt.Println("  2. Duplicate frame")
		fmt.Println("  3. Edit frame")
		fmt.Println("  4. Preview character")
		fmt.Println("  5. Animate character")
		fmt.Println("  6. Export code (terminal)")
		fmt.Println("  7. Save to file")
		fmt.Println("  8. Delete frame")
		fmt.Println("  9. Back to main menu")
		fmt.Println("  10. Exit")
		fmt.Println()
		fmt.Print("◢ Choose option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addFrame(session)
		case "2":
			duplicateFrame(session)
		case "3":
			editFrame(session)
		case "4":
			previewCharacter(session)
		case "5":
			animateCharacter(session)
		case "6":
			exportCode(session)
		case "7":
			saveToFile(session)
		case "8":
			deleteFrame(session)
		case "9":
			session.Save()
			fmt.Println("✓ Progress saved\n")
			return
		case "10":
			session.Save()
			fmt.Println("✓ Progress saved. Goodbye!\n")
			os.Exit(0)
		default:
			fmt.Println("✗ Invalid option\n")
		}
	}
}

func addFrame(session *Session) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("◢ Adding frame to character: " + session.Name)
	fmt.Print("◢ Frame name (e.g., 'idle', 'wave', 'jump'): ")
	frameName, _ := reader.ReadString('\n')
	frameName = strings.TrimSpace(frameName)

	if frameName == "" {
		fmt.Println("✗ Frame name cannot be empty\n")
		return
	}

	// Check if frame exists
	for _, frame := range session.Frames {
		if frame.Name == frameName {
			fmt.Printf("✗ Frame '%s' already exists\n\n", frameName)
			return
		}
	}

	fmt.Println()
	showPalette()
	fmt.Println()

	lines := make([]string, session.Height)

	for i := 0; i < session.Height; i++ {
		for {
			fmt.Printf("◢ Line %d/%d: ", i+1, session.Height)
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			// Apply mirroring
			line = applyMirroring(line)

			if len(line) != session.Width {
				fmt.Printf("  ✗ Error: Expected %d characters, got %d. Try again.\n", session.Width, len(line))
				continue
			}

			// Show preview
			compiled := compilePattern(line)
			fmt.Printf("  Preview: %s\n", compiled)

			// Confirm
			fmt.Print("  ✓ Keep this line? (y/n): ")
			confirm, _ := reader.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
				lines[i] = line
				break
			}
		}

		// Show progressive preview
		if i < session.Height-1 {
			fmt.Println("\n  ◢ Building up...")
			for j := 0; j <= i; j++ {
				fmt.Printf("  %s\n", compilePattern(lines[j]))
			}
			fmt.Println()
		}
	}

	// Final preview - auto-save
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  FINAL PREVIEW                           ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	for _, line := range lines {
		fmt.Println(compilePattern(line))
	}

	// Auto-save the frame
	session.Frames = append(session.Frames, Frame{
		Name:  frameName,
		Lines: lines,
	})
	session.Save()
	fmt.Println("\n✓ Frame added and saved!\n")
}

func duplicateFrame(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("\n✗ No frames to duplicate\n")
		return
	}

	fmt.Println()
	fmt.Println("▢ Frames:")
	for i, frame := range session.Frames {
		fmt.Printf("  %d. %s\n", i+1, frame.Name)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("◢ Choose frame to duplicate: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	frameIdx := -1
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(session.Frames) {
		frameIdx = num - 1
	}

	if frameIdx == -1 {
		fmt.Println("✗ Invalid frame\n")
		return
	}

	sourceFrame := session.Frames[frameIdx]

	fmt.Printf("\n◢ Duplicating frame: %s\n", sourceFrame.Name)
	fmt.Print("◢ New frame name: ")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)

	if newName == "" {
		fmt.Println("✗ Frame name cannot be empty\n")
		return
	}

	// Check if frame name already exists
	for _, frame := range session.Frames {
		if frame.Name == newName {
			fmt.Printf("✗ Frame '%s' already exists\n\n", newName)
			return
		}
	}

	// Create duplicate with copied lines
	duplicateLines := make([]string, len(sourceFrame.Lines))
	copy(duplicateLines, sourceFrame.Lines)

	newFrame := Frame{
		Name:  newName,
		Lines: duplicateLines,
	}

	session.Frames = append(session.Frames, newFrame)
	session.Save()

	// Show preview
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  DUPLICATED FRAME                        ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	for _, line := range newFrame.Lines {
		fmt.Println(compilePattern(line))
	}
	fmt.Printf("\n✓ Frame '%s' duplicated as '%s'!\n", sourceFrame.Name, newName)
	fmt.Println("◢ Tip: Use 'Edit frame' to modify it\n")
}

func editFrame(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("\n✗ No frames to edit\n")
		return
	}

	fmt.Println()
	fmt.Println("▢ Frames:")
	for i, frame := range session.Frames {
		fmt.Printf("  %d. %s\n", i+1, frame.Name)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("◢ Choose frame to edit: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	frameIdx := -1
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(session.Frames) {
		frameIdx = num - 1
	}

	if frameIdx == -1 {
		fmt.Println("✗ Invalid frame\n")
		return
	}

	frame := &session.Frames[frameIdx]

	// Show current frame
	fmt.Printf("\n▢ Editing Frame: %s\n", frame.Name)
	for i, line := range frame.Lines {
		fmt.Printf("  %d. %s → %s\n", i+1, line, compilePattern(line))
	}
	fmt.Println()

	fmt.Print("◢ Line number to edit (or 'done'): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "done" {
		return
	}

	lineIdx := -1
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(frame.Lines) {
		lineIdx = num - 1
	}

	if lineIdx == -1 {
		fmt.Println("✗ Invalid line\n")
		return
	}

	// Edit the line
	fmt.Printf("◢ Current: %s\n", frame.Lines[lineIdx])
	fmt.Print("◢ New pattern: ")
	newLine, _ := reader.ReadString('\n')
	newLine = strings.TrimSpace(newLine)
	newLine = applyMirroring(newLine)

	if len(newLine) != session.Width {
		fmt.Printf("✗ Error: Expected %d characters, got %d\n\n", session.Width, len(newLine))
		return
	}

	fmt.Printf("  Preview: %s\n", compilePattern(newLine))
	fmt.Print("✓ Update this line? (y/n): ")
	confirm, _ := reader.ReadString('\n')

	if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
		frame.Lines[lineIdx] = newLine
		session.Save()
		fmt.Println("✓ Line updated!\n")
	}
}

func previewCharacter(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("\n✗ No frames to preview\n")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  CHARACTER PREVIEW                       ║")
	fmt.Println("╚══════════════════════════════════════════╝")

	for _, frame := range session.Frames {
		fmt.Printf("\n▢ Frame: %s\n", frame.Name)
		for _, line := range frame.Lines {
			fmt.Println(compilePattern(line))
		}
	}
	fmt.Println()
}

func animateCharacter(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("\n✗ No frames to animate\n")
		return
	}

	if len(session.Frames) == 1 {
		fmt.Println("\n✗ Need at least 2 frames to animate (you have 1)\n")
		fmt.Println("◢ Tip: Use 'Duplicate frame' to create variations for animation\n")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  ANIMATE CHARACTER                       ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Build the character from session
	spec := characters.NewCharacterSpec(session.Name, session.Width, session.Height)
	for _, frame := range session.Frames {
		spec = spec.AddFrame(frame.Name, frame.Lines)
	}

	char, err := spec.Build()
	if err != nil {
		fmt.Printf("✗ Error building character: %v\n\n", err)
		return
	}

	fmt.Printf("◢ Animating '%s' with %d frames at 5 FPS for 3 cycles\n", session.Name, len(session.Frames))
	fmt.Println("◢ Press Ctrl+C to stop\n")

	// Animate at 5 FPS for 3 cycles
	if err := characters.Animate(os.Stdout, char, 5, 3); err != nil {
		fmt.Printf("\n✗ Animation error: %v\n\n", err)
		return
	}

	fmt.Println("\n✓ Animation complete!\n")
}

func exportCode(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("\n✗ No frames to export\n")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  EXPORT GO CODE                          ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	fmt.Println("// Pattern codes:")
	for _, frame := range session.Frames {
		fmt.Printf("// Frame: %s\n", frame.Name)
		for i, line := range frame.Lines {
			fmt.Printf("// Line %d: %s\n", i+1, line)
		}
		fmt.Println()
	}

	fmt.Println("// Go code:")
	fmt.Printf("spec := characters.NewCharacterSpec(\"%s\", %d, %d)\n", session.Name, session.Width, session.Height)
	for _, frame := range session.Frames {
		fmt.Printf(".AddFrame(\"%s\", []string{\n", frame.Name)
		for _, line := range frame.Lines {
			fmt.Printf("    \"%s\",\n", line)
		}
		fmt.Println("})")
	}
	fmt.Println()
	fmt.Println("char, _ := spec.Build()")
	fmt.Println("characters.ShowIdle(os.Stdout, char)")
	fmt.Println()
}

func saveToFile(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("\n✗ No frames to export\n")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  SAVE TO FILE                            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	// Get package name
	fmt.Print("◢ Package name (default: characters): ")
	pkgName, _ := reader.ReadString('\n')
	pkgName = strings.TrimSpace(pkgName)
	if pkgName == "" {
		pkgName = "characters"
	}

	// Get directory
	fmt.Print("◢ Save to directory (default: .): ")
	dir, _ := reader.ReadString('\n')
	dir = strings.TrimSpace(dir)
	if dir == "" {
		dir = "."
	}

	// Generate filename
	filename := filepath.Join(dir, session.Name+".go")

	// Show preview
	fmt.Println()
	fmt.Printf("◢ Will create: %s\n", filename)
	fmt.Printf("◢ Package: %s\n", pkgName)
	fmt.Printf("◢ Function: %s()\n", capitalize(session.Name))
	fmt.Println()

	// Confirm
	fmt.Print("◢ Confirm? (y/n): ")
	confirm, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
		fmt.Println("✗ Cancelled\n")
		return
	}

	// Generate code
	code := generateGoFile(session, pkgName)

	// Create directory if needed
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("✗ Error creating directory: %v\n\n", err)
		return
	}

	// Write file
	if err := os.WriteFile(filename, []byte(code), 0644); err != nil {
		fmt.Printf("✗ Error writing file: %v\n\n", err)
		return
	}

	fmt.Printf("✓ Saved to %s\n", filename)
	fmt.Println()
	fmt.Println("◢ Usage:")
	fmt.Printf("   import \"%s\"\n", pkgName)
	fmt.Printf("   char, _ := %s.%s()\n", pkgName, capitalize(session.Name))
	fmt.Println("   characters.Animate(os.Stdout, char, 5, 3)")
	fmt.Println()
}

func generateGoFile(session *Session, pkgName string) string {
	var code strings.Builder

	// Package and imports
	code.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	code.WriteString("import \"github.com/wildreason/tangent/pkg/characters\"\n\n")

	// Comment
	code.WriteString(fmt.Sprintf("// %s returns a %dx%d character with %d frame(s)\n",
		capitalize(session.Name), session.Width, session.Height, len(session.Frames)))
	code.WriteString("// Generated by Tangent v0.1.0-beta.2\n")
	code.WriteString("//\n")
	code.WriteString("// Frames:\n")
	for _, frame := range session.Frames {
		code.WriteString(fmt.Sprintf("//   - %s\n", frame.Name))
	}
	code.WriteString("func " + capitalize(session.Name) + "() (*characters.Character, error) {\n")

	// Character spec
	code.WriteString(fmt.Sprintf("\tspec := characters.NewCharacterSpec(\"%s\", %d, %d)",
		session.Name, session.Width, session.Height))

	for _, frame := range session.Frames {
		code.WriteString(".\n\t\tAddFrame(\"" + frame.Name + "\", []string{\n")
		for _, line := range frame.Lines {
			code.WriteString(fmt.Sprintf("\t\t\t\"%s\",\n", line))
		}
		code.WriteString("\t\t})")
	}

	code.WriteString("\n\n\treturn spec.Build()\n")
	code.WriteString("}\n")

	return code.String()
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func deleteFrame(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("\n✗ No frames to delete\n")
		return
	}

	fmt.Println()
	fmt.Println("▢ Frames:")
	for i, frame := range session.Frames {
		fmt.Printf("  %d. %s\n", i+1, frame.Name)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("◢ Choose frame to delete: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	frameIdx := -1
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(session.Frames) {
		frameIdx = num - 1
	}

	if frameIdx == -1 {
		fmt.Println("✗ Invalid frame\n")
		return
	}

	frameName := session.Frames[frameIdx].Name
	fmt.Printf("✗ Delete frame '%s'? (y/n): ", frameName)
	confirm, _ := reader.ReadString('\n')

	if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
		session.Frames = append(session.Frames[:frameIdx], session.Frames[frameIdx+1:]...)
		session.Save()
		fmt.Println("✓ Frame deleted\n")
	}
}

func browseLibrary() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  LIBRARY CHARACTERS                      ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	names := characters.ListLibrary()
	if len(names) == 0 {
		fmt.Println("✗ No library characters available\n")
		return
	}

	for _, name := range names {
		description, _ := characters.LibraryInfo(name)
		fmt.Printf("◆ %s\n", name)
		fmt.Printf("  %s\n", description)
		fmt.Println()
	}
}

func previewLibrary() {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  PREVIEW LIBRARY CHARACTER               ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	names := characters.ListLibrary()
	if len(names) == 0 {
		fmt.Println("✗ No library characters available\n")
		return
	}

	fmt.Println("▢ Available:")
	for i, name := range names {
		fmt.Printf("  %d. %s\n", i+1, name)
	}
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("◢ Choose character: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	charName := input
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(names) {
		charName = names[num-1]
	}

	char, err := characters.Library(charName)
	if err != nil {
		fmt.Printf("✗ Failed to load: %v\n\n", err)
		return
	}

	fmt.Printf("\n▢ Character: %s\n", charName)
	characters.ShowIdle(os.Stdout, char)
	fmt.Println()

	fmt.Print("◢ Animate it? (y/n): ")
	animate, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(animate)) == "y" {
		characters.Animate(os.Stdout, char, 4, 2)
	}
	fmt.Println()
}

func showPalette() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  BLOCK PALETTE                           ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("  Basic Blocks:")
	fmt.Println("    F = █  (Full)")
	fmt.Println("    T = ▀  (Top half)")
	fmt.Println("    B = ▄  (Bottom half)")
	fmt.Println("    L = ▌  (Left half)")
	fmt.Println("    R = ▐  (Right half)")
	fmt.Println()
	fmt.Println("  Shading:")
	fmt.Println("    . = ░  (Light)")
	fmt.Println("    : = ▒  (Medium)")
	fmt.Println("    # = ▓  (Dark)")
	fmt.Println()
	fmt.Println("  Quadrants (1-4):")
	fmt.Println("    1 = ▘  (Upper Left)     ↔ reverse of 8")
	fmt.Println("    2 = ▝  (Upper Right)    ↔ reverse of 7")
	fmt.Println("    3 = ▖  (Lower Left)     ↔ reverse of 6")
	fmt.Println("    4 = ▗  (Lower Right)    ↔ reverse of 5")
	fmt.Println()
	fmt.Println("  Three-Quadrants (5-8):")
	fmt.Println("    5 = ▛  (UL+UR+LL)       ↔ reverse of 4")
	fmt.Println("    6 = ▜  (UL+UR+LR)       ↔ reverse of 3")
	fmt.Println("    7 = ▙  (UL+LL+LR)       ↔ reverse of 2")
	fmt.Println("    8 = ▟  (UR+LL+LR)       ↔ reverse of 1")
	fmt.Println()
	fmt.Println("  Diagonals:")
	fmt.Println("    \\ = ▚  (Backward diagonal)")
	fmt.Println("    / = ▞  (Forward diagonal)")
	fmt.Println()
	fmt.Println("  Special:")
	fmt.Println("    _ = Space")
	fmt.Println("    X = Mirror marker")
	fmt.Println()
	fmt.Println("  ◢ Tip: Use X to auto-mirror")
	fmt.Println("         Example: __R5FX → __R5F5R__")
}

func getIntInput(prompt string, min, max int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < min || num > max {
			fmt.Printf("  ✗ Error: Please enter a valid number between %d-%d\n", min, max)
			continue
		}
		return num
	}
}

func applyMirroring(pattern string) string {
	idx := strings.Index(pattern, "X")
	if idx == -1 {
		return pattern
	}

	left := pattern[:idx]
	reversed := reverseString(left)
	return left + reversed
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func compilePattern(pattern string) string {
	compiler := characters.NewPatternCompiler()
	return compiler.CompilePattern(pattern)
}

func sessionExists(name string) bool {
	sessions, _ := ListSessions()
	for _, s := range sessions {
		if s == name {
			return true
		}
	}
	return false
}

// ============================================================================
// NON-INTERACTIVE CLI MODE
// ============================================================================

func handleCLI() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "create":
		handleCreate()
	case "animate":
		handleAnimate()
	case "export":
		handleExport()
	case "gallery":
		handleGallery()
	case "version", "--version", "-v":
		fmt.Printf("tangent %s (commit: %s, built: %s)\n", version, commit, date)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Tangent - Terminal Agent Designer")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  tangent                           Start interactive builder")
	fmt.Println("  tangent create [options]          Create character from CLI")
	fmt.Println("  tangent animate [options]         Show character animation")
	fmt.Println("  tangent export [options]          Export character to Go code")
	fmt.Println("  tangent gallery                   Browse library characters")
	fmt.Println("  tangent version                   Show version information")
	fmt.Println("  tangent help                      Show this help message")
	fmt.Println()
	fmt.Println("CREATE OPTIONS:")
	fmt.Println("  --name <string>                   Character name (required)")
	fmt.Println("  --width <int>                     Character width (required)")
	fmt.Println("  --height <int>                    Character height (required)")
	fmt.Println("  --frame <name> <pattern>          Add frame (can repeat)")
	fmt.Println("  --output <file>                   Save to .go file")
	fmt.Println("  --package <name>                  Go package name (default: main)")
	fmt.Println()
	fmt.Println("ANIMATE OPTIONS:")
	fmt.Println("  --name <string>                   Character name (from library)")
	fmt.Println("  --session <string>                Load from session file")
	fmt.Println("  --fps <int>                       Frames per second (default: 5)")
	fmt.Println("  --loops <int>                     Number of loops (default: 3)")
	fmt.Println()
	fmt.Println("EXPORT OPTIONS:")
	fmt.Println("  --session <string>                Load from session file (required)")
	fmt.Println("  --output <file>                   Output file (default: stdout)")
	fmt.Println("  --package <name>                  Go package name (default: main)")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  # Create character")
	fmt.Println("  tangent create --name robot --width 11 --height 3 \\")
	fmt.Println("    --frame idle '__R6FFF6L__,_T5FFFFF6T_,___11_22___'")
	fmt.Println()
	fmt.Println("  # Create multi-frame character")
	fmt.Println("  tangent create --name alien --width 7 --height 3 \\")
	fmt.Println("    --frame idle 'L6FFF6R,T5FFF6T,_1_2_' \\")
	fmt.Println("    --frame wave 'R6FFF6L,T5FFF6T,_1_2_' \\")
	fmt.Println("    --output alien.go --package myagent")
	fmt.Println()
	fmt.Println("  # Animate library character")
	fmt.Println("  tangent animate --name alien --fps 10 --loops 5")
	fmt.Println()
	fmt.Println("  # Animate from session")
	fmt.Println("  tangent animate --session robot --fps 5 --loops 3")
	fmt.Println()
	fmt.Println("  # Export session to file")
	fmt.Println("  tangent export --session robot --output robot.go --package agent")
	fmt.Println()
	fmt.Println("PATTERN CODES:")
	fmt.Println("  F=█ T=▀ B=▄ L=▌ R=▐ 1-8=quads .=#:=shades _=space X=mirror")
	fmt.Println()
	fmt.Println("For full documentation: https://github.com/wildreason/tangent")
}

func handleCreate() {
	var name, output, pkg string
	var width, height int
	var frames []struct {
		name    string
		pattern string
	}

	// Parse args manually to handle multiple --frame
	args := os.Args[2:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--name":
			if i+1 < len(args) {
				name = args[i+1]
				i++
			}
		case "--width":
			if i+1 < len(args) {
				fmt.Sscanf(args[i+1], "%d", &width)
				i++
			}
		case "--height":
			if i+1 < len(args) {
				fmt.Sscanf(args[i+1], "%d", &height)
				i++
			}
		case "--output":
			if i+1 < len(args) {
				output = args[i+1]
				i++
			}
		case "--package":
			if i+1 < len(args) {
				pkg = args[i+1]
				i++
			}
		case "--frame":
			if i+2 < len(args) {
				frames = append(frames, struct {
					name    string
					pattern string
				}{
					name:    args[i+1],
					pattern: args[i+2],
				})
				i += 2
			}
		}
	}

	// Set defaults
	if pkg == "" {
		pkg = "main"
	}

	// Validate required fields
	if name == "" {
		fmt.Fprintln(os.Stderr, "Error: --name is required")
		os.Exit(1)
	}
	if width == 0 {
		fmt.Fprintln(os.Stderr, "Error: --width is required")
		os.Exit(1)
	}
	if height == 0 {
		fmt.Fprintln(os.Stderr, "Error: --height is required")
		os.Exit(1)
	}
	if len(frames) == 0 {
		fmt.Fprintln(os.Stderr, "Error: at least one --frame is required")
		os.Exit(1)
	}

	// Build character spec
	spec := characters.NewCharacterSpec(name, width, height)

	for _, frame := range frames {
		// Split pattern by comma for multi-line
		lines := strings.Split(frame.pattern, ",")

		// Validate line count
		if len(lines) != height {
			fmt.Fprintf(os.Stderr, "Error: frame '%s' has %d lines, expected %d\n", frame.name, len(lines), height)
			os.Exit(1)
		}

		spec = spec.AddFrame(frame.name, lines)
	}

	// Build character
	char, err := spec.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building character: %v\n", err)
		os.Exit(1)
	}

	// Output handling
	if output != "" {
		// Save to file
		code := generateExportCode(name, pkg, spec)
		err := os.WriteFile(output, []byte(code), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Character saved to %s\n", output)
	} else {
		// Show preview
		fmt.Printf("✓ Character '%s' created\n\n", name)
		characters.ShowIdle(os.Stdout, char)
		fmt.Println()
		fmt.Println("To export: tangent create ... --output file.go")
	}
}

func handleAnimate() {
	fs := flag.NewFlagSet("animate", flag.ExitOnError)
	name := fs.String("name", "", "Character name (from library)")
	session := fs.String("session", "", "Session name to load")
	fps := fs.Int("fps", 5, "Frames per second")
	loops := fs.Int("loops", 3, "Number of loops")

	fs.Parse(os.Args[2:])

	var char *characters.Character
	var err error

	if *name != "" {
		// Load from library
		char, err = characters.Library(*name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: library character '%s' not found\n", *name)
			fmt.Fprintln(os.Stderr, "Available: alien")
			os.Exit(1)
		}
	} else if *session != "" {
		// Load from session
		sessData, err := LoadSession(*session)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading session '%s': %v\n", *session, err)
			os.Exit(1)
		}

		// Build character from session
		spec := characters.NewCharacterSpec(sessData.Name, sessData.Width, sessData.Height)
		for _, frame := range sessData.Frames {
			spec = spec.AddFrame(frame.Name, frame.Lines)
		}

		char, err = spec.Build()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error building character: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Error: either --name or --session is required")
		os.Exit(1)
	}

	// Animate
	fmt.Printf("Animating '%s' at %d FPS for %d loops...\n\n", char.Name, *fps, *loops)
	characters.Animate(os.Stdout, char, *fps, *loops)
	fmt.Println()
}

func handleExport() {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	session := fs.String("session", "", "Session name to load")
	output := fs.String("output", "", "Output file (default: stdout)")
	pkg := fs.String("package", "main", "Go package name")

	fs.Parse(os.Args[2:])

	if *session == "" {
		fmt.Fprintln(os.Stderr, "Error: --session is required")
		os.Exit(1)
	}

	// Load session
	sessData, err := LoadSession(*session)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading session '%s': %v\n", *session, err)
		os.Exit(1)
	}

	// Build spec
	spec := characters.NewCharacterSpec(sessData.Name, sessData.Width, sessData.Height)
	for _, frame := range sessData.Frames {
		spec = spec.AddFrame(frame.Name, frame.Lines)
	}

	// Generate code
	code := generateExportCode(sessData.Name, *pkg, spec)

	// Output
	if *output != "" {
		err := os.WriteFile(*output, []byte(code), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Exported to %s\n", *output)
	} else {
		fmt.Println(code)
	}
}

func generateExportCode(name, pkg string, spec *characters.CharacterSpec) string {
	// Get spec data via JSON (hacky but works)
	data, _ := json.Marshal(spec)
	var specData struct {
		Name   string `json:"name"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Frames []struct {
			Name     string   `json:"name"`
			Patterns []string `json:"patterns"`
		} `json:"frames"`
	}
	json.Unmarshal(data, &specData)

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("package %s\n\n", pkg))
	sb.WriteString("import (\n")
	sb.WriteString("\t\"github.com/wildreason/tangent/pkg/characters\"\n")
	sb.WriteString(")\n\n")
	sb.WriteString(fmt.Sprintf("// Get%s returns the %s character\n", strings.Title(name), name))
	sb.WriteString(fmt.Sprintf("func Get%s() (*characters.Character, error) {\n", strings.Title(name)))
	sb.WriteString(fmt.Sprintf("\tspec := characters.NewCharacterSpec(\"%s\", %d, %d)\n",
		specData.Name, specData.Width, specData.Height))

	for _, frame := range specData.Frames {
		sb.WriteString(fmt.Sprintf("\tspec = spec.AddFrame(\"%s\", []string{\n", frame.Name))
		for _, pattern := range frame.Patterns {
			sb.WriteString(fmt.Sprintf("\t\t\"%s\",\n", pattern))
		}
		sb.WriteString("\t})\n")
	}

	sb.WriteString("\treturn spec.Build()\n")
	sb.WriteString("}\n")

	return sb.String()
}

func handleGallery() {
	// Get all library characters
	names := characters.ListLibrary()

	if len(names) == 0 {
		fmt.Println("No library characters available.")
		return
	}

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║             TANGENT CHARACTER GALLERY                        ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	for i, name := range names {
		if i > 0 {
			fmt.Println()
			fmt.Println("────────────────────────────────────────────────────────────")
			fmt.Println()
		}

		// Load character
		char, err := characters.Library(name)
		if err != nil {
			fmt.Printf("Error loading %s: %v\n", name, err)
			continue
		}

		// Get description
		info, _ := characters.LibraryInfo(name)

		// Display header
		fmt.Printf("◆ %s\n", strings.ToUpper(name))
		fmt.Println()

		// Show first frame
		characters.ShowIdle(os.Stdout, char)
		fmt.Println()

		// Show metadata
		fmt.Printf("  Size: %dx%d | Frames: %d | Author: Wildreason, Inc\n",
			char.Width, char.Height, len(char.Frames))
		fmt.Println()

		// Show description (first line only)
		descLines := strings.Split(info, "\n")
		if len(descLines) > 0 {
			fmt.Printf("  %s\n", strings.TrimSpace(descLines[0]))
		}
		fmt.Println()

		// Show usage
		fmt.Printf("  Try it:\n")
		fmt.Printf("    tangent animate --name %s --fps 5 --loops 3\n", name)
		fmt.Println()
		fmt.Printf("  Use in code:\n")
		fmt.Printf("    %s, _ := characters.Library(\"%s\")\n", name, name)
		fmt.Printf("    characters.Animate(os.Stdout, %s, 5, 3)\n", name)
	}

	fmt.Println()
	fmt.Println("────────────────────────────────────────────────────────────")
	fmt.Println()
	fmt.Printf("Total: %d character%s available\n", len(names), pluralize(len(names)))
	fmt.Println()
	fmt.Println("Create your own: tangent (interactive mode)")
	fmt.Println("Full docs: https://github.com/wildreason/tangent")
}

func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}
