package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Block element palette with codes
var blocks = map[string]rune{
	"F": '█', "T": '▀', "B": '▄', "L": '▌', "R": '▐',
	".": '░', ":": '▒', "#": '▓',
	"1": '▘', "2": '▝', "3": '▖', "4": '▗',
	"5": '▛', "6": '▜', "7": '▙', "8": '▟',
	"\\": '▚', "/": '▞',
	"_": ' ',
	"X": ' ', // Mirror marker (will be replaced)
}

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Block Character Builder v0.1.0         ║")
	fmt.Println("║  Multi-Frame Session Manager            ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Show saved sessions
	sessions, _ := ListSessions()
	if len(sessions) > 0 {
		fmt.Println("📁 Saved characters:")
		for _, name := range sessions {
			fmt.Printf("   • %s\n", name)
		}
		fmt.Println()
	}

	// Ask to load or create new
	reader := bufio.NewReader(os.Stdin)
	var session *Session

	if len(sessions) > 0 {
		fmt.Print("◢ (N)ew character or (L)oad existing? (N/l): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(strings.ToLower(choice))

		if choice == "l" || choice == "load" {
			fmt.Print("◢ Enter character name to load: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			loaded, err := LoadSession(name)
			if err != nil {
				fmt.Printf("  ✗ Error loading: %v\n", err)
				fmt.Println("  Starting new character instead...")
			} else {
				session = loaded
				fmt.Printf("  ✓ Loaded '%s' with %d frame(s)\n", session.Name, len(session.Frames))
			}
		}
	}

	// Create new session if not loaded
	if session == nil {
		fmt.Print("◢ Character name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if SessionExists(name) {
			fmt.Print("  ◢ Character exists. Overwrite? (y/N): ")
			confirm, _ := reader.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
				fmt.Println("  Cancelled.")
				return
			}
		}

		width, height := getDimensions()
		session = NewSession(name, width, height)
		fmt.Printf("  ✓ Created '%s' (%dx%d)\n", session.Name, session.Width, session.Height)
	}

	fmt.Println()
	showPalette()
	fmt.Println()

	// Main session loop
	for {
		if err := session.Save(); err != nil {
			fmt.Printf("  ◢ Warning: Auto-save failed: %v\n", err)
		}

		showSessionMenu(session)
		action := getMenuChoice(reader)

		switch action {
		case "a": // Add frame
			addFrame(session, reader)
		case "e": // Edit frame
			editFrame(session, reader)
		case "d": // Delete frame
			deleteFrame(session, reader)
		case "v": // View all frames
			viewAllFrames(session)
		case "x": // Export
			exportSession(session)
			return
		case "q": // Quit
			fmt.Print("  ◢ Save before quitting? (Y/n): ")
			save, _ := reader.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(save)) != "n" {
				session.Save()
				fmt.Println("  ✓ Saved!")
			}
			fmt.Println("  Goodbye!")
			return
		default:
			fmt.Println("  ✗ Invalid choice")
		}
	}
}

func showSessionMenu(session *Session) {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Printf("║  CHARACTER: %-28s ║\n", session.Name)
	fmt.Printf("║  Size: %dx%-2d  Frames: %-2d               ║\n", session.Width, session.Height, len(session.Frames))
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("  (A)dd frame")
	fmt.Println("  (E)dit frame")
	fmt.Println("  (D)elete frame")
	fmt.Println("  (V)iew all frames")
	fmt.Println("  (X)port code")
	fmt.Println("  (Q)uit")
	fmt.Println()
}

func getMenuChoice(reader *bufio.Reader) string {
	fmt.Print("◢ Choice: ")
	choice, _ := reader.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(choice))
}

func addFrame(session *Session, reader *bufio.Reader) {
	fmt.Print("\n◢ Frame name (e.g., idle, walk, jump): ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		name = fmt.Sprintf("frame%d", len(session.Frames)+1)
	}

	fmt.Printf("\n▢ Building frame '%s' (%d lines)\n", name, session.Height)

	var lines []string
	for i := 0; i < session.Height; i++ {
		line := buildLineWithRetry(i+1, session.Width, session.Height)
		lines = append(lines, line)

		if i < session.Height-1 {
			showProgressivePreview(lines, session.Height)
		}
	}

	frame := Frame{Name: name, Lines: lines}

	// Show frame preview
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Printf("║  FRAME PREVIEW: %-24s ║\n", name)
	fmt.Println("╚══════════════════════════════════════════╝")
	for _, line := range lines {
		fmt.Println("  " + compilePattern(line))
	}
	fmt.Println()

	fmt.Print("◢ Add this frame? (Y/n): ")
	confirm, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(confirm)) == "n" {
		fmt.Println("  ✗ Frame discarded")
		return
	}

	session.AddFrame(frame)
	session.Save()
	fmt.Printf("  ✓ Frame '%s' added! Total frames: %d\n", name, len(session.Frames))
}

func editFrame(session *Session, reader *bufio.Reader) {
	if len(session.Frames) == 0 {
		fmt.Println("  ◢ No frames to edit. Add a frame first.")
		return
	}

	// List frames
	fmt.Println("\n📋 Frames:")
	for i, frame := range session.Frames {
		fmt.Printf("  %d. %s\n", i+1, frame.Name)
	}
	fmt.Println()

	fmt.Print("◢ Frame number to edit (1-" + fmt.Sprint(len(session.Frames)) + "): ")
	var frameNum int
	fmt.Scanf("%d\n", &frameNum)
	reader.ReadString('\n') // Clear buffer

	if frameNum < 1 || frameNum > len(session.Frames) {
		fmt.Println("  ✗ Invalid frame number")
		return
	}

	frameIdx := frameNum - 1
	frame := session.Frames[frameIdx]

	// Show current frame
	fmt.Println()
	fmt.Printf("▢ Editing frame '%s'\n", frame.Name)
	for i, line := range frame.Lines {
		fmt.Printf("  %d: %s → %s\n", i+1, line, compilePattern(line))
	}
	fmt.Println()

	fmt.Print("◢ Edit (L)ine, (R)ename, or (C)ancel? (L/r/c): ")
	action, _ := reader.ReadString('\n')
	action = strings.ToLower(strings.TrimSpace(action))

	switch action {
	case "l", "":
		// Edit line
		fmt.Print("◢ Line number to edit (1-" + fmt.Sprint(len(frame.Lines)) + "): ")
		var lineNum int
		fmt.Scanf("%d\n", &lineNum)
		reader.ReadString('\n') // Clear buffer

		if lineNum < 1 || lineNum > len(frame.Lines) {
			fmt.Println("  ✗ Invalid line number")
			return
		}

		lineIdx := lineNum - 1
		fmt.Printf("\n▢ Current line %d: %s\n", lineNum, frame.Lines[lineIdx])
		fmt.Printf("  Preview: %s\n", compilePattern(frame.Lines[lineIdx]))

		newLine := buildLineWithRetry(lineNum, session.Width, session.Height)
		frame.Lines[lineIdx] = newLine

		session.UpdateFrame(frameIdx, frame)
		session.Save()
		fmt.Println("  ✓ Line updated!")

	case "r":
		// Rename frame
		fmt.Print("◢ New name: ")
		newName, _ := reader.ReadString('\n')
		newName = strings.TrimSpace(newName)
		if newName != "" {
			frame.Name = newName
			session.UpdateFrame(frameIdx, frame)
			session.Save()
			fmt.Println("  ✓ Frame renamed!")
		}

	default:
		fmt.Println("  Cancelled")
	}
}

func deleteFrame(session *Session, reader *bufio.Reader) {
	if len(session.Frames) == 0 {
		fmt.Println("  ◢ No frames to delete.")
		return
	}

	// List frames
	fmt.Println("\n📋 Frames:")
	for i, frame := range session.Frames {
		fmt.Printf("  %d. %s\n", i+1, frame.Name)
	}
	fmt.Println()

	fmt.Print("◢ Frame number to delete (1-" + fmt.Sprint(len(session.Frames)) + "): ")
	var frameNum int
	fmt.Scanf("%d\n", &frameNum)
	reader.ReadString('\n') // Clear buffer

	if frameNum < 1 || frameNum > len(session.Frames) {
		fmt.Println("  ✗ Invalid frame number")
		return
	}

	frameIdx := frameNum - 1
	frameName := session.Frames[frameIdx].Name

	fmt.Printf("◢ Delete frame '%s'? (y/N): ", frameName)
	confirm, _ := reader.ReadString('\n')
	if strings.ToLower(strings.TrimSpace(confirm)) != "y" {
		fmt.Println("  Cancelled")
		return
	}

	session.DeleteFrame(frameIdx)
	session.Save()
	fmt.Printf("  ✓ Frame '%s' deleted\n", frameName)
}

func viewAllFrames(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("  ◢ No frames yet. Add a frame first.")
		return
	}

	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  ALL FRAMES PREVIEW                      ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	for i, frame := range session.Frames {
		fmt.Printf("Frame %d: %s\n", i+1, frame.Name)
		for _, line := range frame.Lines {
			fmt.Println("  " + compilePattern(line))
		}
		if i < len(session.Frames)-1 {
			fmt.Println()
		}
	}
	fmt.Println()

	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func exportSession(session *Session) {
	if len(session.Frames) == 0 {
		fmt.Println("  ◢ No frames to export. Add at least one frame first.")
		return
	}

	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  EXPORT CODE                             ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Show pattern code
	fmt.Println("▢ Pattern Code:")
	fmt.Printf("char := characters.NewCharacterSpec(\"%s\", %d, %d)\n", session.Name, session.Width, session.Height)
	for i, frame := range session.Frames {
		if i == 0 {
			fmt.Print("    .AddFrame(\"" + frame.Name + "\", []string{\n")
		} else {
			fmt.Print("    .AddFrame(\"" + frame.Name + "\", []string{\n")
		}

		for j, line := range frame.Lines {
			comma := ","
			if j == len(frame.Lines)-1 {
				comma = ""
			}
			fmt.Printf("        \"%s\"%s\n", line, comma)
		}

		if i < len(session.Frames)-1 {
			fmt.Print("    })\n")
		} else {
			fmt.Print("    })\n")
		}
	}
	fmt.Println()

	fmt.Println("✓ Code ready to copy!")
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

// Existing helper functions...
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
	fmt.Println("  Three-Quadrants (5-8 = reverse of 1-4):")
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
	fmt.Println("  ◢ Tip: Use X to auto-mirror left side to right")
	fmt.Println("         Example: __R5FX → __R5F5R__ (mirrored!)")
}

func getDimensions() (int, int) {
	reader := bufio.NewReader(os.Stdin)

	var width int
	for {
		fmt.Print("◢ Enter width (e.g., 11): ")
		widthStr, _ := reader.ReadString('\n')
		widthStr = strings.TrimSpace(widthStr)

		n, err := fmt.Sscanf(widthStr, "%d", &width)
		if err != nil || n != 1 || width <= 0 || width > 100 {
			fmt.Println("  ✗ Error: Please enter a valid number between 1-100")
			continue
		}
		break
	}

	var height int
	for {
		fmt.Print("◢ Enter height (e.g., 3): ")
		heightStr, _ := reader.ReadString('\n')
		heightStr = strings.TrimSpace(heightStr)

		n, err := fmt.Sscanf(heightStr, "%d", &height)
		if err != nil || n != 1 || height <= 0 || height > 50 {
			fmt.Println("  ✗ Error: Please enter a valid number between 1-50")
			continue
		}
		break
	}

	fmt.Printf("  ✓ Dimensions set: %dx%d\n", width, height)
	return width, height
}

func showProgressivePreview(lines []string, totalHeight int) {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Printf("║  PROGRESSIVE PREVIEW (%d/%d lines)       ║\n", len(lines), totalHeight)
	fmt.Println("╚══════════════════════════════════════════╝")

	for _, line := range lines {
		fmt.Println("  " + compilePattern(line))
	}

	remaining := totalHeight - len(lines)
	for i := 0; i < remaining; i++ {
		fmt.Println("  ...")
	}
}

func buildLineWithRetry(lineNum, width, totalHeight int) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n▢ Line %d/%d (enter %d codes):\n", lineNum, totalHeight, width)
		fmt.Printf("  Pattern: ")

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		line = processMirroring(line, width)

		if len(line) < width {
			line = line + strings.Repeat("_", width-len(line))
			fmt.Printf("  ◢ Padded with spaces to width %d\n", width)
		} else if len(line) > width {
			line = line[:width]
			fmt.Printf("  ◢ Truncated to width %d\n", width)
		}

		fmt.Printf("  Preview: %s\n", compilePattern(line))

		fmt.Print("  ◢ Accept this line? (Y/n): ")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(strings.ToLower(confirm))

		if confirm == "" || confirm == "y" || confirm == "yes" {
			fmt.Println("  ✓ Line accepted!")
			return line
		}

		fmt.Println("  ◢ Let's try again...")
	}
}

func processMirroring(pattern string, width int) string {
	xIndex := strings.Index(pattern, "X")
	if xIndex == -1 {
		return pattern
	}

	xCount := strings.Count(pattern, "X")
	if xCount > 1 {
		fmt.Println("  ◢ Mirror mode activated! Found", xCount, "X markers")
	}

	parts := strings.Split(pattern, "X")
	if len(parts) < 2 {
		return pattern
	}

	left := parts[0]
	reversed := reverseString(left)
	mirrored := left + reversed

	if len(mirrored) < width {
		mirrored = mirrored + strings.Repeat("_", width-len(mirrored))
	}

	fmt.Printf("  ◢ Mirrored: %s → %s\n", pattern, mirrored[:min(width, len(mirrored))])

	return mirrored
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func compilePattern(pattern string) string {
	var result strings.Builder
	for _, char := range pattern {
		if block, exists := blocks[string(char)]; exists {
			result.WriteRune(block)
		} else {
			result.WriteRune(' ')
		}
	}
	return result.String()
}
