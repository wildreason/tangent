package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/domain"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
	"github.com/wildreason/tangent/pkg/characters/library"
	"github.com/wildreason/tangent/pkg/characters/patterns"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// handleError provides user-friendly error handling with helpful suggestions
func handleError(message string, err error) {
	fmt.Printf("✗ %s: %v\n", message, err)
}

func main() {
	// Handle subcommands
	if len(os.Args) > 1 {
		handleCLI()
		return
	}

	// No arguments - show usage
	printUsage()
}

func showBanner() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  TANGENT - Terminal Avatars for AI      ║")
	fmt.Println("║  Give your AI agent a face              ║")
	fmt.Printf("║  %-40s ║\n", version)
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
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
		handleError("Character creation failed", domain.NewValidationError("name", name, "character name cannot be empty"))
		return
	}

	// Validate character name for Go identifier compatibility
	if !isValidGoIdentifier(name) {
		handleError("Character creation failed", domain.ErrCharacterNameInvalid)
		return
	}

	// Check if character already exists in library
	availableChars := characters.ListLibrary()
	for _, existingName := range availableChars {
		if existingName == name {
			fmt.Printf("✗ Character '%s' already exists in library. Use a different name.\n\n", name)
			return
		}
	}

	// Get dimensions
	width := getIntInput("◢ Enter width (e.g., 11): ", 1, 100)
	height := getIntInput("◢ Enter height (e.g., 3): ", 1, 50)

	fmt.Println()
	fmt.Printf("✓ Creating character '%s' (%dx%d)\n", name, width, height)
	fmt.Printf("✓ Character '%s' is starting\n\n", name)

	// Create session
	session := NewSession(name, width, height)
	session.Save()

	// Start Bubbletea TUI
	if err := StartCreationTUI(session); err != nil {
		handleError("TUI failed", err)
		return
	}
}

// convertAgentToSession converts an AgentCharacter to a Session for UI compatibility
func convertAgentToSession(agent *characters.AgentCharacter) *Session {
	// Get the underlying domain character
	domainChar := agent.GetCharacter()

	session := NewSession(domainChar.Name, domainChar.Width, domainChar.Height)

	// Convert base frame
	if len(domainChar.BaseFrame.Lines) > 0 {
		session.BaseFrame = Frame{
			Name:  domainChar.BaseFrame.Name,
			Lines: domainChar.BaseFrame.Lines,
		}
	}

	// Convert states
	for _, state := range domainChar.States {
		stateSession := StateSession{
			Name:           state.Name,
			Description:    state.Description,
			StateType:      state.StateType,
			AnimationFPS:   state.AnimationFPS,
			AnimationLoops: state.AnimationLoops,
		}

		// Convert state frames
		for _, frame := range state.Frames {
			stateFrame := Frame{
				Name:  frame.Name,
				Lines: frame.Lines,
			}
			stateSession.Frames = append(stateSession.Frames, stateFrame)
		}

		session.States = append(session.States, stateSession)
	}

	// Convert legacy frames for backward compatibility
	for _, frame := range domainChar.Frames {
		sessionFrame := Frame{
			Name:  frame.Name,
			Lines: frame.Lines,
		}
		session.Frames = append(session.Frames, sessionFrame)
	}

	return session
}

// convertSessionToCharacterSpec converts a Session to a characters.CharacterSpec
func convertSessionToCharacterSpec(session *Session) *characters.CharacterSpec {
	spec := characters.NewCharacterSpec(session.Name, session.Width, session.Height)

	for _, frame := range session.Frames {
		spec.AddFrame(frame.Name, frame.Lines)
	}

	return spec
}

// saveSessionAsCharacter saves a session as a character
func saveSessionAsCharacter(session *Session) error {
	// For now, just return success since we're not persisting characters
	// In the future, this could save to a file or database
	return nil
}

func characterBuilder(session *Session) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("▢ CHARACTER: " + session.Name)

		// Show status
		baseStatus := "✗ Not created"
		if len(session.BaseFrame.Lines) > 0 {
			baseStatus = "✓ Created"
		}
		stateNames := []string{}
		for _, state := range session.States {
			stateNames = append(stateNames, state.Name)
		}
		stateList := "none"
		if len(stateNames) > 0 {
			stateList = strings.Join(stateNames, ", ")
		}

		fmt.Printf("  Base: %s | States: %d (%s)\n", baseStatus, len(session.States), stateList)
		fmt.Println()

		// Show appropriate tip
		if len(session.BaseFrame.Lines) == 0 {
			fmt.Println("  ◢ Tip: Start by creating the base character (idle state)!")
			fmt.Println()
		} else if len(session.States) == 0 {
			fmt.Println("  ◢ Tip: Now add agent states (think, plan, search)!")
			fmt.Println()
		}

		fmt.Println("  1. Create base character")
		fmt.Println("  2. Add agent state")
		fmt.Println("  3. Animate all states")
		fmt.Println("  4. Export for contribution (JSON)")
		fmt.Println("  5. Exit")
		fmt.Println()
		fmt.Print("◢ Choose option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			createBaseCharacter(session)
		case "2":
			addAgentStateWithBase(session)
		case "3":
			previewAllStates(session)
		case "4":
			exportForContribution(session)
		case "5":
			// Save session
			if err := session.Save(); err != nil {
				handleError("Failed to save session", err)
			} else {
				fmt.Println("✓ Progress saved. Goodbye!\n")
			}
			os.Exit(0)
		default:
			fmt.Println("✗ Invalid option\n")
		}
	}
}

func addFrame(session *Session) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("◢ Adding agent state to character: " + session.Name)

	// Show which states are already added
	existingStates := make(map[string]bool)
	for _, frame := range session.Frames {
		existingStates[frame.Name] = true
	}

	// Show required states
	requiredStates := []string{"plan", "think", "execute"}
	missingRequired := []string{}
	for _, req := range requiredStates {
		if !existingStates[req] {
			missingRequired = append(missingRequired, req)
		}
	}

	fmt.Println()
	if len(missingRequired) > 0 {
		fmt.Println("  ◢ Required states (choose one):")
		if !existingStates["plan"] {
			fmt.Println("    • plan     - Agent analyzing and planning")
		}
		if !existingStates["think"] {
			fmt.Println("    • think    - Agent processing information")
		}
		if !existingStates["execute"] {
			fmt.Println("    • execute  - Agent performing actions")
		}
		fmt.Println()
	}

	// Show optional states
	fmt.Println("  ◢ Optional states:")
	if !existingStates["wait"] {
		fmt.Println("    • wait     - Agent waiting for input")
	}
	if !existingStates["error"] {
		fmt.Println("    • error    - Agent handling errors")
	}
	if !existingStates["success"] {
		fmt.Println("    • success  - Agent celebrating success")
	}
	fmt.Println()
	fmt.Println("  ◢ Or enter custom state name")
	fmt.Println()

	fmt.Print("◢ Agent state name: ")
	frameName, _ := reader.ReadString('\n')
	frameName = strings.TrimSpace(frameName)

	if frameName == "" {
		fmt.Println("✗ State name cannot be empty\n")
		return
	}

	// Check if frame exists
	for _, frame := range session.Frames {
		if frame.Name == frameName {
			fmt.Printf("✗ State '%s' already exists\n\n", frameName)
			return
		}
	}

	// Determine state type
	stateType := "custom"
	standardStates := []string{"plan", "think", "execute", "wait", "error", "success"}
	for _, std := range standardStates {
		if frameName == std {
			stateType = "standard"
			break
		}
	}

	fmt.Println()
	fmt.Println(patterns.GetPatternHelp())
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

	// Auto-save the frame with state type
	session.Frames = append(session.Frames, Frame{
		Name:      frameName,
		Lines:     lines,
		StateType: stateType,
	})
	session.Save()
	fmt.Printf("\n✓ %s state '%s' added and saved!\n\n", strings.Title(stateType), frameName)

	// Show progress on required states
	existingStates[frameName] = true
	missingCount := 0
	for _, req := range requiredStates {
		if !existingStates[req] {
			missingCount++
		}
	}
	if missingCount > 0 {
		fmt.Printf("  ◢ Tip: %d required state(s) remaining\n\n", missingCount)
	} else {
		fmt.Println("  ✓ All required states added! You can now export for contribution.\n")
	}
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

	// Convert session to domain character
	spec := convertSessionToCharacterSpec(session)
	character, err := spec.Build()
	if err != nil {
		handleError("Failed to create character for animation", err)
		return
	}

	fmt.Printf("◢ Animating '%s' with %d frames at 5 FPS for 3 cycles\n", session.Name, len(session.Frames))
	fmt.Println("◢ Press Ctrl+C to stop\n")

	// Animate using AgentCharacter
	agent := characters.NewAgentCharacter(character)
	if err := agent.AnimateState(os.Stdout, "plan", 5, 3); err != nil {
		handleError("Animation failed", err)
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
	fmt.Println("   agent := characters.NewAgentCharacter(char)")
	fmt.Println("   agent.AnimateState(os.Stdout, \"plan\", 5, 3)")
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
	code.WriteString(fmt.Sprintf("// Generated by Tangent %s\n", version))
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
	compiler := infrastructure.NewPatternCompiler()
	return compiler.Compile(pattern)
}

// convertFramesToDomain converts local Frame slice to domain.Frame slice
func convertFramesToDomain(frames []Frame) []domain.Frame {
	domainFrames := make([]domain.Frame, len(frames))
	for i, frame := range frames {
		domainFrames[i] = domain.Frame{
			Name:  frame.Name,
			Lines: frame.Lines,
		}
	}
	return domainFrames
}

// isValidGoIdentifier checks if a string is a valid Go identifier
func isValidGoIdentifier(name string) bool {
	if name == "" {
		return false
	}

	// First character must be letter or underscore
	first := rune(name[0])
	if !unicode.IsLetter(first) && first != '_' {
		return false
	}

	// All other characters must be letter, digit, or underscore
	for _, r := range name[1:] {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}

	return true
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
		showBanner()
		createCharacter()
	case "browse":
		if len(os.Args) > 2 {
			handleListAgent(os.Args[2])
		} else {
			handleList()
		}
	case "view":
		handleView(os.Args[2:])
	case "admin":
		handleAdminCLI()
	case "version", "--version", "-v":
		fmt.Printf("tangent %s (commit: %s, built: %s)\n", version, commit, date)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n\n", command)
		fmt.Fprintln(os.Stderr, "Try 'tangent help' for usage information")
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Tangent - Terminal Avatars for AI Agents")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  tangent browse              List all avatars")
	fmt.Println("  tangent browse <name>       Preview avatar with states")
	fmt.Println("  tangent browse <name> [--state plan|think|execute] [--fps N] [--loops N]")
	fmt.Println("  tangent version             Show version information")
	fmt.Println("  tangent help                Show this help message")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  tangent browse              # Discover available avatars")
	fmt.Println("  tangent browse mercury      # Preview mercury avatar")
	fmt.Println("  tangent browse water --state plan --fps 8")
	fmt.Println()
	fmt.Println("For API integration: https://github.com/wildreason/tangent")
}

func handleAdminCLI() {
	if len(os.Args) < 3 {
		printAdminUsage()
		os.Exit(1)
	}

	subcommand := os.Args[2]

	switch subcommand {
	case "register":
		if len(os.Args) < 4 {
			fmt.Println("Error: missing JSON file path")
			printAdminUsage()
			os.Exit(1)
		}
		// Check for --force flag
		forceUpdate := false
		if len(os.Args) > 4 && os.Args[4] == "--force" {
			forceUpdate = true
		}
		adminRegister(os.Args[3], forceUpdate)
	case "validate":
		if len(os.Args) < 4 {
			fmt.Println("Error: missing JSON file path")
			printAdminUsage()
			os.Exit(1)
		}
		adminValidate(os.Args[3])
	case "export":
		if len(os.Args) < 4 {
			fmt.Println("Error: missing character name")
			printAdminUsage()
			os.Exit(1)
		}
		adminExport(os.Args[3])
	case "batch-register":
		if len(os.Args) < 4 {
			fmt.Println("Error: missing template JSON file path")
			printAdminUsage()
			os.Exit(1)
		}
		if len(os.Args) < 5 {
			fmt.Println("Error: missing colors JSON file path")
			printAdminUsage()
			os.Exit(1)
		}
		adminBatchRegister(os.Args[3], os.Args[4])
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown admin command '%s'\n\n", subcommand)
		printAdminUsage()
		os.Exit(1)
	}
}

func printAdminUsage() {
	fmt.Println("Admin Commands:")
	fmt.Println("  tangent admin export <character>              Export library character to JSON")
	fmt.Println("  tangent admin register <json>                 Register new character to library")
	fmt.Println("  tangent admin register <json> --force         Update existing character")
	fmt.Println("  tangent admin batch-register <template> <colors>  Register multiple characters from template")
	fmt.Println("  tangent admin validate <json>                 Validate character JSON")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  tangent admin export mercury                  # Export to mercury.json")
	fmt.Println("  tangent admin register egon.json              # Add new character")
	fmt.Println("  tangent admin register mercury.json --force   # Update existing character")
	fmt.Println("  tangent admin batch-register template.json colors.json  # Create all characters")
	fmt.Println("  tangent admin validate egon.json              # Validate before registering")
}

func adminRegister(jsonPath string, forceUpdate bool) {
	if forceUpdate {
		fmt.Printf("Registering/Updating character from %s...\n", jsonPath)
	} else {
		fmt.Printf("Registering character from %s...\n", jsonPath)
	}

	// Load and parse JSON
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	var charData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Color       string `json:"color"`
		Personality string `json:"personality"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
		BaseFrame   struct {
			Name  string   `json:"name"`
			Lines []string `json:"lines"`
		} `json:"base_frame"`
		States []struct {
			Name   string `json:"name"`
			Frames []struct {
				Lines []string `json:"lines"`
			} `json:"frames"`
		} `json:"states"`
	}

	if err := json.Unmarshal(data, &charData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Validate required fields
	if charData.Name == "" {
		fmt.Println("Error: missing 'name' field")
		os.Exit(1)
	}
	if !isValidGoIdentifier(charData.Name) {
		fmt.Printf("Error: %s\n", domain.ErrCharacterNameInvalid.Error())
		os.Exit(1)
	}
	if charData.Width == 0 || charData.Height == 0 {
		fmt.Println("Error: missing or invalid 'width'/'height' fields")
		os.Exit(1)
	}
	if len(charData.BaseFrame.Lines) == 0 {
		fmt.Println("Error: missing 'base_frame' field")
		os.Exit(1)
	}
	if len(charData.States) == 0 {
		fmt.Println("Error: missing 'states' field")
		os.Exit(1)
	}

	// Check if character already exists (unless --force is used)
	if !forceUpdate {
		availableChars := characters.ListLibrary()
		for _, charName := range availableChars {
			if charName == charData.Name {
				fmt.Printf("Error: character '%s' already exists in library\n", charData.Name)
				fmt.Println("Use --force to update existing character:")
				fmt.Printf("  tangent admin register %s --force\n", jsonPath)
				os.Exit(1)
			}
		}
	} else {
		// Backup existing file before update
		libraryFile := fmt.Sprintf("pkg/characters/library/%s.go", charData.Name)
		if _, err := os.Stat(libraryFile); err == nil {
			backupFile := libraryFile + ".backup"
			data, _ := os.ReadFile(libraryFile)
			os.WriteFile(backupFile, data, 0644)
			fmt.Printf("📦 Backed up existing file to %s\n", backupFile)
		}
	}

	// Generate library file
	libraryFile := fmt.Sprintf("pkg/characters/library/%s.go", charData.Name)

	// Use default personality if not provided
	personality := charData.Personality
	if personality == "" {
		personality = "efficient"
	}

	// Create patterns array
	patterns := []struct {
		Name  string
		Lines []string
	}{
		{
			Name:  "base",
			Lines: charData.BaseFrame.Lines,
		},
	}

	// Add states - preserve individual frames for proper animation
	for _, state := range charData.States {
		for i, frame := range state.Frames {
			// Create individual pattern for each frame
			patternName := state.Name
			if len(state.Frames) > 1 {
				// Add frame number if multiple frames exist
				patternName = fmt.Sprintf("%s_%d", state.Name, i+1)
			}

			patterns = append(patterns, struct {
				Name  string
				Lines []string
			}{
				Name:  patternName,
				Lines: frame.Lines,
			})
		}
	}

	// Use provided values or defaults
	description := charData.Description
	if description == "" {
		description = fmt.Sprintf("%s - efficient AI Agent Character", charData.Name)
	}
	author := charData.Author
	if author == "" {
		author = "Wildreason, Inc"
	}
	color := charData.Color
	if color == "" {
		color = "#FFFFFF" // Default white
	}

	// Generate Go code
	code := generateLibraryCode(charData.Name, description, author, color, personality, charData.Width, charData.Height, patterns)

	// Write file
	if err := os.WriteFile(libraryFile, []byte(code), 0644); err != nil {
		fmt.Printf("Error writing library file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Character '%s' registered successfully!\n", charData.Name)
	fmt.Printf("📁 Library file: %s\n", libraryFile)
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("1. Run: make build")
	fmt.Println("2. Test: tangent gallery")
	fmt.Println("3. Commit the changes")
}

func adminValidate(jsonPath string) {
	fmt.Printf("Validating character JSON: %s\n", jsonPath)

	// Load and parse JSON
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	var charData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Author      string `json:"author"`
		Color       string `json:"color"`
		Personality string `json:"personality"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
		BaseFrame   struct {
			Name  string   `json:"name"`
			Lines []string `json:"lines"`
		} `json:"base_frame"`
		States []struct {
			Name   string `json:"name"`
			Frames []struct {
				Lines []string `json:"lines"`
			} `json:"frames"`
		} `json:"states"`
	}

	if err := json.Unmarshal(data, &charData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Validate required fields
	valid := true

	if charData.Name == "" {
		fmt.Println("❌ Missing 'name' field")
		valid = false
	}
	if charData.Width == 0 || charData.Height == 0 {
		fmt.Println("❌ Missing or invalid 'width'/'height' fields")
		valid = false
	}
	if len(charData.BaseFrame.Lines) == 0 {
		fmt.Println("❌ Missing 'base_frame' field")
		valid = false
	}
	if len(charData.States) == 0 {
		fmt.Println("❌ Missing 'states' field")
		valid = false
	}

	// Check required states and minimum frames
	requiredStates := []string{"plan", "think", "execute"}
	stateNames := make(map[string]bool)
	minFrames := 3

	for _, state := range charData.States {
		stateNames[state.Name] = true

		// Check minimum frames per state
		if len(state.Frames) < minFrames {
			fmt.Printf("❌ State '%s' has %d frames; minimum is %d\n", state.Name, len(state.Frames), minFrames)
			valid = false
		}
	}

	for _, required := range requiredStates {
		if !stateNames[required] {
			fmt.Printf("❌ Missing required state: %s\n", required)
			valid = false
		}
	}

	if valid {
		fmt.Println("✅ Character JSON is valid!")
		fmt.Printf("   Name: %s\n", charData.Name)
		fmt.Printf("   Size: %dx%d\n", charData.Width, charData.Height)
		fmt.Printf("   States: %d\n", len(charData.States))
	} else {
		fmt.Println("❌ Character JSON has validation errors")
		os.Exit(1)
	}
}

func adminExport(characterName string) {
	fmt.Printf("Exporting character '%s' to JSON...\n", characterName)

	// Load character from library
	char, err := library.Get(characterName)
	if err != nil {
		fmt.Printf("Error: character '%s' not found in library\n", characterName)
		fmt.Println("Available characters:", strings.Join(characters.ListLibrary(), ", "))
		os.Exit(1)
	}

	// Prepare JSON structure
	type Frame struct {
		Lines []string `json:"lines"`
	}

	type State struct {
		Name   string  `json:"name"`
		Frames []Frame `json:"frames"`
	}

	type CharacterExport struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Author      string `json:"author,omitempty"`
		Color       string `json:"color,omitempty"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
		BaseFrame   struct {
			Name  string   `json:"name"`
			Lines []string `json:"lines"`
		} `json:"base_frame"`
		States []State `json:"states"`
	}

	export := CharacterExport{
		Name:        char.Name,
		Description: char.Description,
		Author:      char.Author,
		Color:       char.Color,
		Width:       char.Width,
		Height:      char.Height,
	}

	// Extract base frame (first frame or frame named "base")
	baseFound := false
	for _, pattern := range char.Patterns {
		if pattern.Name == "base" {
			export.BaseFrame.Name = "base"
			export.BaseFrame.Lines = pattern.Lines
			baseFound = true
			break
		}
	}
	if !baseFound && len(char.Patterns) > 0 {
		// Use first frame as base
		export.BaseFrame.Name = "base"
		export.BaseFrame.Lines = char.Patterns[0].Lines
	}

	// Group frames by state name
	stateMap := make(map[string][]Frame)
	stateOrder := []string{} // Preserve order

	for _, pattern := range char.Patterns {
		if pattern.Name == "base" {
			continue // Skip base frame
		}

		// Extract state name from pattern name
		// Examples: "wait_1" -> "wait", "plan_2" -> "plan", "arise" -> "arise"
		stateName := pattern.Name
		if idx := strings.LastIndex(pattern.Name, "_"); idx != -1 {
			// Check if suffix is a number
			suffix := pattern.Name[idx+1:]
			if _, err := strconv.Atoi(suffix); err == nil {
				stateName = pattern.Name[:idx]
			}
		}

		// Add frame to state
		if _, exists := stateMap[stateName]; !exists {
			stateOrder = append(stateOrder, stateName)
			stateMap[stateName] = []Frame{}
		}

		stateMap[stateName] = append(stateMap[stateName], Frame{
			Lines: pattern.Lines,
		})
	}

	// Build states array in order
	for _, stateName := range stateOrder {
		export.States = append(export.States, State{
			Name:   stateName,
			Frames: stateMap[stateName],
		})
	}

	// Write JSON to file
	outputFile := characterName + ".json"
	data, err := json.MarshalIndent(export, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputFile, data, 0644); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Character exported successfully!\n")
	fmt.Printf("📁 Output file: %s\n", outputFile)
	fmt.Println()
	fmt.Println("Stats:")
	fmt.Printf("  - Dimensions: %dx%d\n", char.Width, char.Height)
	fmt.Printf("  - States: %d\n", len(export.States))
	fmt.Printf("  - Total frames: %d\n", len(char.Patterns))
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Edit the JSON file to add/modify states")
	fmt.Println("  2. Run: tangent admin register " + outputFile + " --force")
}

func adminBatchRegister(templatePath, colorsPath string) {
	fmt.Printf("Batch registering characters from template: %s\n", templatePath)
	fmt.Printf("Using color config: %s\n\n", colorsPath)

	// Read template JSON
	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("Error reading template file: %v\n", err)
		os.Exit(1)
	}

	var template struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		States []struct {
			Name   string `json:"name"`
			Frames []struct {
				Lines []string `json:"lines"`
			} `json:"frames"`
		} `json:"states"`
	}

	if err := json.Unmarshal(templateData, &template); err != nil {
		fmt.Printf("Error parsing template JSON: %v\n", err)
		os.Exit(1)
	}

	// Read colors config
	colorsData, err := os.ReadFile(colorsPath)
	if err != nil {
		fmt.Printf("Error reading colors file: %v\n", err)
		os.Exit(1)
	}

	var colors map[string]struct {
		Color       string `json:"color"`
		Description string `json:"description"`
	}

	if err := json.Unmarshal(colorsData, &colors); err != nil {
		fmt.Printf("Error parsing colors JSON: %v\n", err)
		os.Exit(1)
	}

	// Process each character
	successCount := 0
	for charName, charConfig := range colors {
		fmt.Printf("📝 Registering character: %s (color: %s)\n", charName, charConfig.Color)

		// Prepare patterns from template states
		var patterns []struct {
			Name  string
			Lines []string
		}

		for _, state := range template.States {
			for i, frame := range state.Frames {
				frameNum := i + 1
				frameName := fmt.Sprintf("%s_%d", state.Name, frameNum)
				patterns = append(patterns, struct {
					Name  string
					Lines []string
				}{
					Name:  frameName,
					Lines: frame.Lines,
				})
			}
		}

		// Generate library code with character-specific color
		code := generateLibraryCode(
			charName,
			charConfig.Description,
			"Wildreason, Inc",
			charConfig.Color,
			"",
			template.Width,
			template.Height,
			patterns,
		)

		// Write to library file (with backup if exists)
		libraryFile := fmt.Sprintf("pkg/characters/library/%s.go", charName)
		if _, err := os.Stat(libraryFile); err == nil {
			backupFile := libraryFile + ".backup"
			existingData, _ := os.ReadFile(libraryFile)
			os.WriteFile(backupFile, existingData, 0644)
			fmt.Printf("  📦 Backed up existing file to %s\n", backupFile)
		}

		if err := os.WriteFile(libraryFile, []byte(code), 0644); err != nil {
			fmt.Printf("  ❌ Error writing library file: %v\n", err)
			continue
		}

		fmt.Printf("  ✅ Successfully registered %s\n\n", charName)
		successCount++
	}

	fmt.Printf("🎉 Batch registration complete!\n")
	fmt.Printf("   Created/Updated: %d characters\n", successCount)
	fmt.Printf("   Total in config: %d characters\n\n", len(colors))
	fmt.Println("Next steps:")
	fmt.Println("  1. Run: make build")
	fmt.Println("  2. Test: tangent browse")
}

func generateLibraryCode(name, description, author, color, personality string, width, height int, patterns []struct {
	Name  string
	Lines []string
}) string {
	var sb strings.Builder

	sb.WriteString("package library\n\n")
	sb.WriteString("func init() {\n")
	sb.WriteString(fmt.Sprintf("\tregister(%sCharacter)\n", name))
	sb.WriteString("}\n\n")
	sb.WriteString(fmt.Sprintf("var %sCharacter = LibraryCharacter{\n", name))
	sb.WriteString(fmt.Sprintf("\tName:        \"%s\",\n", name))
	sb.WriteString(fmt.Sprintf("\tDescription: \"%s\",\n", description))
	sb.WriteString(fmt.Sprintf("\tAuthor:      \"%s\",\n", author))
	sb.WriteString(fmt.Sprintf("\tColor:       \"%s\",\n", color))
	sb.WriteString(fmt.Sprintf("\tWidth:       %d,\n", width))
	sb.WriteString(fmt.Sprintf("\tHeight:      %d,\n", height))
	sb.WriteString("\tPatterns: []Frame{\n")

	for _, pattern := range patterns {
		sb.WriteString("\t\t{\n")
		sb.WriteString(fmt.Sprintf("\t\t\tName: \"%s\",\n", pattern.Name))
		sb.WriteString("\t\t\tLines: []string{\n")
		for _, line := range pattern.Lines {
			sb.WriteString(fmt.Sprintf("\t\t\t\t\"%s\",\n", line))
		}
		sb.WriteString("\t\t\t},\n")
		sb.WriteString("\t\t},\n")
	}

	sb.WriteString("\t},\n")
	sb.WriteString("}\n")

	return sb.String()
}

func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

// hasRequiredStates checks if a session has the minimum required agent states
func hasRequiredStates(session *Session) bool {
	required := []string{"plan", "think", "execute"}
	found := make(map[string]bool)

	// Check in States (new structure)
	for _, state := range session.States {
		for _, req := range required {
			if state.Name == req {
				found[req] = true
			}
		}
	}

	// Also check in Frames for backward compatibility
	for _, frame := range session.Frames {
		for _, req := range required {
			if frame.Name == req {
				found[req] = true
			}
		}
	}

	return len(found) >= 3
}

// getMissingRequiredStates returns a list of missing required states
func getMissingRequiredStates(session *Session) []string {
	required := []string{"plan", "think", "execute"}
	found := make(map[string]bool)

	// Check in States (new structure)
	for _, state := range session.States {
		found[state.Name] = true
	}

	// Also check in Frames for backward compatibility
	for _, frame := range session.Frames {
		found[frame.Name] = true
	}

	missing := []string{}
	for _, req := range required {
		if !found[req] {
			missing = append(missing, req)
		}
	}

	return missing
}

// exportForContribution exports a character as JSON for GitHub contribution
func exportForContribution(session *Session) {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  EXPORT FOR CONTRIBUTION                 ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Show character info
	fmt.Printf("◢ Character: %s\n", session.Name)
	fmt.Printf("◢ Dimensions: %dx%d\n", session.Width, session.Height)
	fmt.Printf("◢ States: %d\n", len(session.States))
	fmt.Println()

	// List states
	fmt.Println("  States included:")
	for _, state := range session.States {
		stateIcon := "●"
		if state.StateType == "standard" {
			stateIcon = "✓"
		}
		fmt.Printf("    %s %s (%s)\n", stateIcon, state.Name, state.StateType)
	}
	fmt.Println()

	// Export to JSON
	filename := session.Name + ".json"
	data, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		fmt.Printf("✗ Error marshaling JSON: %v\n\n", err)
		return
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Printf("✗ Error writing file: %v\n\n", err)
		return
	}

	fmt.Printf("✓ Exported to %s\n\n", filename)

	// Generate contribution README
	readmeFilename := session.Name + "-README.md"
	readme := generateContributionReadme(session)
	if err := os.WriteFile(readmeFilename, []byte(readme), 0644); err != nil {
		fmt.Printf("✗ Error writing README: %v\n\n", err)
		return
	}

	fmt.Printf("✓ Generated %s\n\n", readmeFilename)

	// Show next steps
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  NEXT STEPS                              ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("  1. Review the exported JSON file")
	fmt.Println("  2. Read the contribution README")
	fmt.Println("  3. Fork the Tangent repository on GitHub")
	fmt.Println("  4. Create a new branch for your character")
	fmt.Println("  5. Add your JSON file to characters/ directory")
	fmt.Println("  6. Submit a Pull Request")
	fmt.Println()
	fmt.Println("  ◢ See .github/CONTRIBUTING_CHARACTERS.md for details")
	fmt.Println()
}

// generateContributionReadme generates a README for character contribution
func generateContributionReadme(session *Session) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s Character Contribution\n\n", strings.Title(session.Name)))
	sb.WriteString("## Character Information\n\n")
	sb.WriteString(fmt.Sprintf("- **Name:** %s\n", session.Name))
	sb.WriteString(fmt.Sprintf("- **Dimensions:** %dx%d\n", session.Width, session.Height))
	sb.WriteString(fmt.Sprintf("- **States:** %d\n\n", len(session.States)))

	sb.WriteString("## States Included\n\n")
	for _, state := range session.States {
		sb.WriteString(fmt.Sprintf("- **%s** (%s)\n", state.Name, state.StateType))
	}
	sb.WriteString("\n")

	sb.WriteString("## Preview\n\n")
	sb.WriteString("```\n")
	if len(session.BaseFrame.Lines) > 0 {
		for _, line := range session.BaseFrame.Lines {
			compiler := infrastructure.NewPatternCompiler()
			sb.WriteString(compiler.Compile(line) + "\n")
		}
	}
	sb.WriteString("```\n\n")

	sb.WriteString("## Contribution Checklist\n\n")
	sb.WriteString("- [x] Minimum 3 required states (plan, think, execute)\n")
	sb.WriteString("- [x] Valid pattern codes\n")
	sb.WriteString("- [x] Tested in Tangent CLI\n")
	sb.WriteString("- [ ] JSON file added to characters/ directory\n")
	sb.WriteString("- [ ] Pull Request submitted\n\n")

	sb.WriteString("## How to Contribute\n\n")
	sb.WriteString("1. Fork the Tangent repository\n")
	sb.WriteString("2. Create a new branch: `git checkout -b add-" + session.Name + "-character`\n")
	sb.WriteString("3. Add " + session.Name + ".json to characters/ directory\n")
	sb.WriteString("4. Commit: `git commit -m 'Add " + session.Name + " character'`\n")
	sb.WriteString("5. Push: `git push origin add-" + session.Name + "-character`\n")
	sb.WriteString("6. Submit Pull Request on GitHub\n\n")

	sb.WriteString("See `.github/CONTRIBUTING_CHARACTERS.md` for full contribution guidelines.\n")

	return sb.String()
}

// createBaseCharacter creates the base (idle) character
func createBaseCharacter(session *Session) {
	reader := bufio.NewReader(os.Stdin)

	// Check if base already exists
	if len(session.BaseFrame.Lines) > 0 {
		fmt.Println("\n✗ Base character already exists!")
		fmt.Print("  Overwrite? (y/n): ")
		confirm, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
			fmt.Println()
			return
		}
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  CREATE BASE CHARACTER                   ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("◢ Designing %s's base (idle) state\n", session.Name)
	fmt.Println("  This is the immutable foundation for all states")
	fmt.Println()
	fmt.Println("  ◢ Live Preview: Press 'p' to toggle split-pane preview")
	fmt.Println()
	fmt.Println(patterns.GetPatternHelp())
	fmt.Println()

	lines := make([]string, session.Height)
	livePreview := false

	for i := 0; i < session.Height; i++ {
		for {
			// Show live preview if enabled
			if livePreview {
				showLivePreview(session, lines, i, "base")
			}

			fmt.Printf("◢ Line %d/%d: ", i+1, session.Height)
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			// Toggle live preview
			if line == "p" {
				livePreview = !livePreview
				if livePreview {
					fmt.Println("  ✓ Live preview enabled")
				} else {
					fmt.Println("  ✓ Live preview disabled")
				}
				continue
			}

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

	// Final preview
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  BASE CHARACTER PREVIEW                  ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	for _, line := range lines {
		fmt.Println(compilePattern(line))
	}

	// Save
	session.BaseFrame = Frame{
		Name:  "base",
		Lines: lines,
	}
	session.Save()
	fmt.Println("\n✓ Base character created! Now add animated states.\n")
}

// previewBaseCharacter shows the base character
func previewBaseCharacter(session *Session) {
	if len(session.BaseFrame.Lines) == 0 {
		fmt.Println("\n✗ No base character created yet\n")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  BASE CHARACTER                          ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	for _, line := range session.BaseFrame.Lines {
		fmt.Println(compilePattern(line))
	}
	fmt.Println()
}

// addAgentStateWithBase adds an agent state with reference to base
func addAgentStateWithBase(session *Session) {
	reader := bufio.NewReader(os.Stdin)

	// Check if base exists
	if len(session.BaseFrame.Lines) == 0 {
		fmt.Println("\n✗ Create base character first!\n")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  ADD AGENT STATE                         ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Show existing states
	existingStates := make(map[string]bool)
	for _, state := range session.States {
		existingStates[state.Name] = true
	}

	// Show required states
	requiredStates := []string{"plan", "think", "execute"}
	missingRequired := []string{}
	for _, req := range requiredStates {
		if !existingStates[req] {
			missingRequired = append(missingRequired, req)
		}
	}

	if len(missingRequired) > 0 {
		fmt.Println("  ◢ Required states (choose one):")
		for _, req := range missingRequired {
			descriptions := map[string]string{
				"plan":    "Agent analyzing and planning",
				"think":   "Agent processing information",
				"execute": "Agent performing actions",
			}
			fmt.Printf("    • %-8s - %s\n", req, descriptions[req])
		}
		fmt.Println()
	}

	// Show optional states
	fmt.Println("  ◢ Optional states:")
	if !existingStates["wait"] {
		fmt.Println("    • wait     - Agent waiting for input")
	}
	if !existingStates["error"] {
		fmt.Println("    • error    - Agent handling errors")
	}
	if !existingStates["success"] {
		fmt.Println("    • success  - Agent celebrating success")
	}
	fmt.Println()
	fmt.Println("  ◢ Or enter custom state name")
	fmt.Println()

	fmt.Print("◢ Agent state name: ")
	stateName, _ := reader.ReadString('\n')
	stateName = strings.TrimSpace(stateName)

	if stateName == "" {
		fmt.Println("✗ State name cannot be empty\n")
		return
	}

	// Check if state exists
	if existingStates[stateName] {
		fmt.Printf("✗ State '%s' already exists\n\n", stateName)
		return
	}

	// Determine state type
	stateType := "custom"
	standardStates := []string{"plan", "think", "execute", "wait", "error", "success"}
	for _, std := range standardStates {
		if stateName == std {
			stateType = "standard"
			break
		}
	}

	// Ask for number of frames
	fmt.Print("◢ How many animation frames? (default: 3): ")
	frameCountInput, _ := reader.ReadString('\n')
	frameCountInput = strings.TrimSpace(frameCountInput)
	frameCount := 3
	if frameCountInput != "" {
		if count, err := strconv.Atoi(frameCountInput); err == nil && count > 0 {
			frameCount = count
		}
	}

	fmt.Printf("\n✓ Creating '%s' state with %d animation frames\n\n", stateName, frameCount)

	// Create frames for this state
	stateFrames := []Frame{}

	for frameIdx := 0; frameIdx < frameCount; frameIdx++ {
		fmt.Printf("╔══════════════════════════════════════════╗\n")
		fmt.Printf("║  FRAME %d/%d                               ║\n", frameIdx+1, frameCount)
		fmt.Printf("╚══════════════════════════════════════════╝\n")
		fmt.Println()

		// Show base as reference
		fmt.Println("  ◢ Base character (reference):")
		for _, line := range session.BaseFrame.Lines {
			fmt.Printf("    %s\n", compilePattern(line))
		}
		fmt.Println()

		// Ask if starting from base
		fmt.Print("  Start from base? (y/n): ")
		startFromBase, _ := reader.ReadString('\n')
		startFromBase = strings.TrimSpace(strings.ToLower(startFromBase))

		lines := make([]string, session.Height)

		if startFromBase == "y" {
			// Copy base lines
			copy(lines, session.BaseFrame.Lines)
			fmt.Println("  ✓ Copied base. Edit lines as needed (press Enter to keep):\n")
		} else {
			fmt.Println("  Creating from scratch:\n")
		}

		fmt.Println(patterns.GetPatternHelp())
		fmt.Println()
		fmt.Println("  ◢ Live Preview: Press 'p' to toggle split-pane preview")
		fmt.Println()

		// Input lines
		livePreview := false
		for i := 0; i < session.Height; i++ {
			for {
				// Show live preview if enabled
				if livePreview {
					showLivePreview(session, lines, i, stateName)
				}

				currentLine := ""
				if startFromBase == "y" && i < len(lines) {
					currentLine = lines[i]
					fmt.Printf("◢ Line %d/%d (current: %s): ", i+1, session.Height, compilePattern(currentLine))
				} else {
					fmt.Printf("◢ Line %d/%d: ", i+1, session.Height)
				}

				line, _ := reader.ReadString('\n')
				line = strings.TrimSpace(line)

				// Toggle live preview
				if line == "p" {
					livePreview = !livePreview
					if livePreview {
						fmt.Println("  ✓ Live preview enabled")
					} else {
						fmt.Println("  ✓ Live preview disabled")
					}
					continue
				}

				// If empty and we have a current line, keep it
				if line == "" && currentLine != "" {
					lines[i] = currentLine
					fmt.Printf("  ✓ Kept: %s\n", compilePattern(currentLine))
					break
				}

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
		}

		// Preview this frame
		fmt.Println("\n  ◢ Frame preview:")
		for _, line := range lines {
			fmt.Printf("    %s\n", compilePattern(line))
		}
		fmt.Println()

		// Save frame
		stateFrames = append(stateFrames, Frame{
			Name:  fmt.Sprintf("%s_frame_%d", stateName, frameIdx+1),
			Lines: lines,
		})
	}

	// Save state
	newState := StateSession{
		Name:           stateName,
		Description:    fmt.Sprintf("Agent %s state", stateName),
		StateType:      stateType,
		Frames:         stateFrames,
		AnimationFPS:   5,
		AnimationLoops: 1,
	}
	session.States = append(session.States, newState)
	session.Save()

	fmt.Printf("\n✓ %s state '%s' added with %d frames!\n\n", strings.Title(stateType), stateName, frameCount)

	// Show progress on required states
	existingStates[stateName] = true
	missingCount := 0
	for _, req := range requiredStates {
		if !existingStates[req] {
			missingCount++
		}
	}
	if missingCount > 0 {
		fmt.Printf("  ◢ Tip: %d required state(s) remaining\n\n", missingCount)
	} else {
		fmt.Println("  ✓ All required states added! You can now export for contribution.\n")
	}
}

// previewDualPane renders a side-by-side formation vs end-state preview for a session.
// If stateName is empty, shows base-only on both sides.
func previewDualPane(session *Session, stateName string) {
	// Determine preview height and width
	height := session.Height
	width := session.Width

	// Assemble left (formation) frames
	leftFrames := [][]string{}
	if stateName == "" {
		// Base static
		lf := make([]string, height)
		copy(lf, session.BaseFrame.Lines)
		leftFrames = append(leftFrames, lf)
	} else {
		// Find the state and use its frames as formation
		for _, st := range session.States {
			if st.Name == stateName {
				for _, f := range st.Frames {
					lf := make([]string, height)
					copy(lf, f.Lines)
					leftFrames = append(leftFrames, lf)
				}
				break
			}
		}
		if len(leftFrames) == 0 {
			// Fallback to base static
			lf := make([]string, height)
			copy(lf, session.BaseFrame.Lines)
			leftFrames = append(leftFrames, lf)
		}
	}

	// Assemble right (end-state) frames using domain + AgentCharacter grouping
	// Build temporary domain character
	dStates := map[string]domain.State{}
	for _, st := range session.States {
		dStates[st.Name] = domain.State{
			Name:           st.Name,
			Description:    st.Description,
			Frames:         convertFramesToDomain(st.Frames),
			StateType:      st.StateType,
			AnimationFPS:   st.AnimationFPS,
			AnimationLoops: st.AnimationLoops,
		}
	}
	tempChar := &domain.Character{
		Name:      session.Name,
		Width:     session.Width,
		Height:    session.Height,
		BaseFrame: domain.Frame{Name: session.BaseFrame.Name, Lines: session.BaseFrame.Lines},
		States:    dStates,
	}

	// Extract right frames for the selected state
	rightFrames := [][]string{}
	if stateName == "" {
		rf := make([]string, height)
		copy(rf, session.BaseFrame.Lines)
		rightFrames = append(rightFrames, rf)
	} else if st, ok := tempChar.States[stateName]; ok && len(st.Frames) > 0 {
		for _, f := range st.Frames {
			rf := make([]string, height)
			copy(rf, f.Lines)
			rightFrames = append(rightFrames, rf)
		}
	} else {
		rf := make([]string, height)
		copy(rf, session.BaseFrame.Lines)
		rightFrames = append(rightFrames, rf)
	}

	// Prepare compiler
	compiler := infrastructure.NewPatternCompiler()

	// Animation parameters
	fps := 5
	loops := 2
	frameDur := time.Second / time.Duration(fps)

	// Ensure both sides have at least one frame
	if len(leftFrames) == 0 {
		leftFrames = append(leftFrames, make([]string, height))
	}
	if len(rightFrames) == 0 {
		rightFrames = append(rightFrames, make([]string, height))
	}

	// Header
	fmt.Println("\n╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║  PREVIEW (FORMATION | END-STATE)                             ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("  Left: Formation (what you're building now)")
	fmt.Println("  Right: End-state (as it will animate)\n")

	// Hide cursor
	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[?25h")

	// Animate
	for loop := 0; loop < loops; loop++ {
		for fi := 0; fi < max(len(leftFrames), len(rightFrames)); fi++ {
			l := leftFrames[fi%len(leftFrames)]
			r := rightFrames[fi%len(rightFrames)]

			// Render combined rows
			for row := 0; row < height; row++ {
				lc := ""
				rc := ""
				if row < len(l) {
					lc = compiler.Compile(l[row])
				}
				if row < len(r) {
					rc = compiler.Compile(r[row])
				}
				fmt.Printf("\r\x1b[2K  %-*s    %s\n", width, lc, rc)
			}
			// Move cursor back up
			fmt.Printf("\x1b[%dA", height)
			time.Sleep(frameDur)
		}
	}

	// Print final combined frame cleanly
	l := leftFrames[len(leftFrames)-1]
	r := rightFrames[len(rightFrames)-1]
	for row := 0; row < height; row++ {
		lc := ""
		rc := ""
		if row < len(l) {
			lc = compiler.Compile(l[row])
		}
		if row < len(r) {
			rc = compiler.Compile(r[row])
		}
		fmt.Printf("\r\x1b[2K  %-*s    %s\n", width, lc, rc)
	}
	fmt.Println("\n✓ Preview complete. Press Enter to return.\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// showLivePreview displays a persistent vertical split-pane with live updates
func showLivePreview(session *Session, lines []string, currentLine int, context string) {
	// Clear screen and set up split-pane layout
	fmt.Print("\x1b[2J\x1b[H") // Clear screen and move cursor to top-left

	// Get terminal width (assume 80 columns as fallback)
	terminalWidth := 80
	leftWidth := terminalWidth / 2
	rightWidth := terminalWidth - leftWidth - 1 // -1 for separator

	// Header
	fmt.Println("╔══════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║  LIVE PREVIEW MODE - Press 'p' to toggle, 'q' to quit preview              ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Split-pane layout
	fmt.Printf("┌%s┬%s┐\n", strings.Repeat("─", leftWidth-2), strings.Repeat("─", rightWidth-2))

	// Left pane: Current creation interface
	fmt.Printf("│ %-*s │ %-*s │\n", leftWidth-4, "CREATION", rightWidth-4, "LIVE PREVIEW")
	fmt.Printf("├%s┼%s┤\n", strings.Repeat("─", leftWidth-2), strings.Repeat("─", rightWidth-2))

	// Show current line being edited
	fmt.Printf("│ %-*s │ %-*s │\n", leftWidth-4, fmt.Sprintf("Line %d/%d", currentLine+1, session.Height), rightWidth-4, "Real-time render")
	fmt.Printf("├%s┼%s┤\n", strings.Repeat("─", leftWidth-2), strings.Repeat("─", rightWidth-2))

	// Show the character being built
	compiler := infrastructure.NewPatternCompiler()
	for row := 0; row < session.Height; row++ {
		leftContent := ""
		rightContent := ""

		if row < len(lines) && lines[row] != "" {
			// Show pattern codes on left
			leftContent = lines[row]
			// Show compiled result on right
			rightContent = compiler.Compile(lines[row])
		} else if row == currentLine {
			// Show current line being edited
			leftContent = repeatString("█", session.Width) // Cursor indicator
			rightContent = repeatString("█", session.Width)
		} else {
			// Show empty lines
			leftContent = strings.Repeat("_", session.Width)
			rightContent = strings.Repeat(" ", session.Width)
		}

		// Pad content to fit pane width
		if len(leftContent) > leftWidth-4 {
			leftContent = leftContent[:leftWidth-4]
		}
		if len(rightContent) > rightWidth-4 {
			rightContent = rightContent[:rightWidth-4]
		}

		fmt.Printf("│ %-*s │ %-*s │\n", leftWidth-4, leftContent, rightWidth-4, rightContent)
	}

	// Bottom border
	fmt.Printf("└%s┴%s┘\n", strings.Repeat("─", leftWidth-2), strings.Repeat("─", rightWidth-2))
	fmt.Println()
	fmt.Println("  ◢ Commands: 'p' = toggle preview, 'q' = quit preview, or continue editing")
	fmt.Print("  ◢ Line input: ")
}

// Helper function for string repetition
func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

// editAgentState edits an existing agent state
func editAgentState(session *Session) {
	if len(session.States) == 0 {
		fmt.Println("\n✗ No states to edit\n")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  EDIT AGENT STATE                        ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("  States:")
	for i, state := range session.States {
		fmt.Printf("    %d. %s (%d frames)\n", i+1, state.Name, len(state.Frames))
	}
	fmt.Println()
	fmt.Print("◢ Choose state number: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	stateIdx := -1
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(session.States) {
		stateIdx = num - 1
	}

	if stateIdx == -1 {
		fmt.Println("✗ Invalid state\n")
		return
	}

	state := &session.States[stateIdx]

	fmt.Printf("\n  Editing state: %s\n", state.Name)
	fmt.Println("  1. Add more frames")
	fmt.Println("  2. Edit existing frame")
	fmt.Println("  3. Remove frame")
	fmt.Println("  4. Change animation speed (FPS)")
	fmt.Println("  5. Cancel")
	fmt.Println()
	fmt.Print("◢ Choose option: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("\n  ◢ Feature coming soon: Add more frames")
	case "2":
		fmt.Println("\n  ◢ Feature coming soon: Edit existing frame")
	case "3":
		fmt.Println("\n  ◢ Feature coming soon: Remove frame")
	case "4":
		fmt.Print("\n  ◢ Current FPS: ", state.AnimationFPS)
		fmt.Print("\n  ◢ New FPS (1-30): ")
		fpsInput, _ := reader.ReadString('\n')
		fpsInput = strings.TrimSpace(fpsInput)
		if fps, err := strconv.Atoi(fpsInput); err == nil && fps > 0 && fps <= 30 {
			state.AnimationFPS = fps
			session.Save()
			fmt.Printf("\n  ✓ Animation speed updated to %d FPS\n\n", fps)
		} else {
			fmt.Println("\n  ✗ Invalid FPS\n")
		}
	case "5":
		return
	default:
		fmt.Println("\n✗ Invalid option\n")
	}
}

// previewStateAnimation previews a single state's animation
func previewStateAnimation(session *Session) {
	if len(session.States) == 0 {
		fmt.Println("\n✗ No states to preview\n")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  PREVIEW STATE ANIMATION                 ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("  States:")
	for i, state := range session.States {
		fmt.Printf("    %d. %s (%d frames, %d FPS)\n", i+1, state.Name, len(state.Frames), state.AnimationFPS)
	}
	fmt.Println()
	fmt.Print("◢ Choose state number: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	stateIdx := -1
	if num, err := strconv.Atoi(input); err == nil && num > 0 && num <= len(session.States) {
		stateIdx = num - 1
	}

	if stateIdx == -1 {
		fmt.Println("✗ Invalid state\n")
		return
	}

	state := session.States[stateIdx]

	fmt.Printf("\n◢ Animating '%s' state with %d frames at %d FPS for 2 cycles\n", state.Name, len(state.Frames), state.AnimationFPS)
	fmt.Println("◢ Press Ctrl+C to stop\n")

	// Create temporary character for animation
	tempChar := &domain.Character{
		Name:   session.Name + "-" + state.Name,
		Width:  session.Width,
		Height: session.Height,
		Frames: []domain.Frame{},
	}

	for _, frame := range state.Frames {
		tempChar.Frames = append(tempChar.Frames, domain.Frame{
			Name:  frame.Name,
			Lines: frame.Lines,
		})
	}

	// Animate using AgentCharacter
	agent := characters.NewAgentCharacter(tempChar)
	if err := agent.AnimateState(os.Stdout, state.Name, state.AnimationFPS, 2); err != nil {
		handleError("Animation failed", err)
		return
	}

	fmt.Println("\n✓ Animation complete!\n")
}

// previewAllStates previews all states in sequence
func previewAllStates(session *Session) {
	if len(session.States) == 0 {
		fmt.Println("\n✗ No states to preview\n")
		return
	}

	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║  PREVIEW ALL STATES                      ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Show base first
	if len(session.BaseFrame.Lines) > 0 {
		fmt.Println("  Base character:")
		for _, line := range session.BaseFrame.Lines {
			fmt.Printf("    %s\n", compilePattern(line))
		}
		fmt.Println()
	}

	// Show each state
	for _, state := range session.States {
		fmt.Printf("  State: %s (%d frames)\n", state.Name, len(state.Frames))
		fmt.Printf("  Animating at %d FPS...\n\n", state.AnimationFPS)

		// Create temporary character for animation using new state-based approach
		tempChar := &domain.Character{
			Name:   session.Name + "-" + state.Name,
			Width:  session.Width,
			Height: session.Height,
			BaseFrame: domain.Frame{
				Name:  session.BaseFrame.Name,
				Lines: session.BaseFrame.Lines,
			},
			States: map[string]domain.State{
				state.Name: {
					Name:           state.Name,
					Description:    state.Description,
					Frames:         convertFramesToDomain(state.Frames),
					StateType:      state.StateType,
					AnimationFPS:   state.AnimationFPS,
					AnimationLoops: state.AnimationLoops,
				},
			},
		}

		// Create AgentCharacter and use AnimateState method
		agent := characters.NewAgentCharacter(tempChar)
		if err := agent.AnimateState(os.Stdout, state.Name, state.AnimationFPS, 1); err != nil {
			handleError("Animation failed", err)
			continue
		}

		fmt.Println()
	}

	fmt.Println("✓ All states previewed!\n")
}
