package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wildreason/tangent/pkg/characters/infrastructure"
	"github.com/wildreason/tangent/pkg/characters/patterns"
)

// Screen represents different screens in the TUI
type Screen int

const (
	ScreenMenu Screen = iota
	ScreenCreateBase
	ScreenCreateState
	ScreenStateNameInput
	ScreenStateFrameInput
	ScreenStatePreview
	ScreenAnimateAll
	ScreenExport
)

// CreationModel is the Bubbletea model for character creation
type CreationModel struct {
	session     *Session
	screen      Screen
	menuCursor  int
	menuOptions []string

	// Base creation state
	baseLines   []string
	currentLine int
	textInput   textinput.Model

	// State creation state
	currentStateName string
	currentStateType string
	stateFrameCount  int
	stateFrames      [][]string
	currentFrame     int
	currentFrameLine int
	frameLines       []string

	// Animation state for preview
	previewFrameIndex int
	animating         bool
	currentStateIndex int // For animating all states

	// UI dimensions
	width  int
	height int

	// Styles
	styles Styles

	// Status message
	statusMsg string
	err       error
}

// Styles holds all lipgloss styles
type Styles struct {
	border       lipgloss.Style
	title        lipgloss.Style
	menuItem     lipgloss.Style
	selectedItem lipgloss.Style
	preview      lipgloss.Style
	statusBar    lipgloss.Style
	errorMsg     lipgloss.Style
	helpText     lipgloss.Style
}

// NewCreationModel creates a new Bubbletea model
func NewCreationModel(session *Session) CreationModel {
	ti := textinput.New()
	ti.Placeholder = "Enter pattern (e.g., FFFFFFFF)"
	ti.Focus()
	ti.CharLimit = 50 // Allow longer input for state names
	ti.Width = 40

	return CreationModel{
		session:    session,
		screen:     ScreenMenu,
		menuCursor: 0,
		menuOptions: []string{
			"Create base character",
			"Add agent state",
			"Animate all states",
			"Export for contribution",
			"Save and exit",
		},
		baseLines: make([]string, session.Height),
		textInput: ti,
		styles:    NewStyles(),
	}
}

// NewStyles creates the lipgloss styles
func NewStyles() Styles {
	return Styles{
		border: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1),
		title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			Padding(0, 1),
		menuItem: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 2),
		selectedItem: lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Padding(0, 1),
		preview: lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2),
		statusBar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Padding(0, 1),
		errorMsg: lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true).
			Padding(0, 1),
		helpText: lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true).
			Padding(0, 1),
	}
}

// tickMsg is sent on animation ticks
type tickMsg struct{}

// Init initializes the model
func (m CreationModel) Init() tea.Cmd {
	return textinput.Blink
}

// tick returns a command that sends a tickMsg after a delay
func tick() tea.Cmd {
	return tea.Tick(time.Second/5, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

// Update handles messages
func (m CreationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tickMsg:
		// Handle animation frame updates
		if m.animating && m.screen == ScreenStatePreview {
			if len(m.session.States) > 0 {
				state := m.session.States[len(m.session.States)-1]
				m.previewFrameIndex = (m.previewFrameIndex + 1) % len(state.Frames)
			}
			return m, tick() // Continue animation
		}
		if m.animating && m.screen == ScreenAnimateAll {
			if len(m.session.States) > 0 {
				state := m.session.States[m.currentStateIndex]
				m.previewFrameIndex = (m.previewFrameIndex + 1) % len(state.Frames)
			}
			return m, tick() // Continue animation
		}
		return m, nil

	case tea.KeyMsg:
		// Handle screen-specific keys first (to allow text input to work)
		switch m.screen {
		case ScreenMenu:
			// Menu-specific keys
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			}
			return m.updateMenu(msg)

		case ScreenCreateBase:
			// Only handle special keys, let text input handle regular keys
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.screen = ScreenMenu
				m.statusMsg = "Cancelled"
				return m, nil
			case "enter":
				return m.handleBaseLineSubmit()
			case "ctrl+d":
				return m.handleBaseLineDelete()
			default:
				// Let text input handle all other keys
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}

		case ScreenStateNameInput:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.screen = ScreenMenu
				m.statusMsg = "Cancelled"
				return m, nil
			case "enter":
				return m.handleStateNameSubmit()
			default:
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}

		case ScreenStateFrameInput:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				m.screen = ScreenMenu
				m.statusMsg = "Cancelled"
				return m, nil
			case "enter":
				return m.handleStateFrameLineSubmit()
			case "ctrl+d":
				return m.handleStateFrameLineDelete()
			default:
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}

		case ScreenStatePreview:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "enter", "space":
				// Save and go back to menu
				m.animating = false
				m.screen = ScreenMenu
				m.statusMsg = fmt.Sprintf("✓ State '%s' saved!", m.currentStateName)
				m.currentStateName = ""
				m.stateFrames = nil
				return m, nil
			case "esc":
				// Discard and go back to menu
				m.animating = false
				// Remove the last state
				if len(m.session.States) > 0 {
					m.session.States = m.session.States[:len(m.session.States)-1]
					m.session.Save()
				}
				m.screen = ScreenMenu
				m.statusMsg = "State discarded"
				m.currentStateName = ""
				m.stateFrames = nil
				return m, nil
			}

		case ScreenAnimateAll:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "esc", "q":
				m.animating = false
				m.screen = ScreenMenu
				m.statusMsg = "Animation stopped"
				return m, nil
			case "left", "h":
				// Previous state
				if m.currentStateIndex > 0 {
					m.currentStateIndex--
					m.previewFrameIndex = 0
					m.statusMsg = fmt.Sprintf("State %d/%d: %s", m.currentStateIndex+1, len(m.session.States), m.session.States[m.currentStateIndex].Name)
				}
			case "right", "l":
				// Next state
				if m.currentStateIndex < len(m.session.States)-1 {
					m.currentStateIndex++
					m.previewFrameIndex = 0
					m.statusMsg = fmt.Sprintf("State %d/%d: %s", m.currentStateIndex+1, len(m.session.States), m.session.States[m.currentStateIndex].Name)
				}
			}

		case ScreenExport:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			case "enter":
				// Export the character
				err := m.exportCharacter()
				if err != nil {
					m.err = err
					m.statusMsg = fmt.Sprintf("Export failed: %v", err)
				} else {
					m.statusMsg = fmt.Sprintf("✓ Exported %s.json and %s-README.md!", m.session.Name, m.session.Name)
					m.screen = ScreenMenu
				}
				return m, nil
			case "esc":
				m.screen = ScreenMenu
				m.statusMsg = "Export cancelled"
				return m, nil
			}
		}
	}

	return m, cmd
}

// updateMenu handles menu navigation
func (m CreationModel) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.menuCursor > 0 {
			m.menuCursor--
		}
	case "down", "j":
		if m.menuCursor < len(m.menuOptions)-1 {
			m.menuCursor++
		}
	case "enter":
		switch m.menuCursor {
		case 0: // Create base character
			m.screen = ScreenCreateBase
			m.currentLine = 0
			m.textInput.SetValue("")
			m.textInput.CharLimit = m.session.Width // Limit to character width for patterns
			m.textInput.Placeholder = "Enter pattern (e.g., FFFFFFFF)"
			m.textInput.Focus()
			m.statusMsg = "Creating base character - Press Enter to confirm each line"
		case 1: // Add agent state
			if len(m.session.BaseFrame.Lines) == 0 {
				m.statusMsg = "Create base character first!"
				m.err = fmt.Errorf("base character required")
			} else {
				m.screen = ScreenStateNameInput
				m.textInput.SetValue("")
				m.textInput.CharLimit = 50 // Allow full state names
				m.textInput.Placeholder = "Enter state name (plan, think, execute)"
				m.textInput.Focus()
				m.statusMsg = "Enter agent state name"
			}
		case 2: // Animate all states
			if len(m.session.States) == 0 {
				m.statusMsg = "No states to animate. Create some states first!"
				m.err = fmt.Errorf("no states available")
			} else {
				m.screen = ScreenAnimateAll
				m.currentStateIndex = 0
				m.previewFrameIndex = 0
				m.animating = true
				m.statusMsg = "Animating all states - Use ←/→ to switch states, Esc to exit"
				return m, tick()
			}
		case 3: // Export
			m.screen = ScreenExport
			m.statusMsg = "Export for contribution"
		case 4: // Save and exit
			m.session.Save()
			return m, tea.Quit
		}
	}
	return m, nil
}

// handleBaseLineSubmit handles Enter key in base creation
func (m CreationModel) handleBaseLineSubmit() (tea.Model, tea.Cmd) {
	value := m.textInput.Value()

	// Validate length
	if len(value) != m.session.Width {
		m.err = fmt.Errorf("expected %d characters, got %d", m.session.Width, len(value))
		return m, nil
	}

	// Apply mirroring
	value = applyMirroring(value)

	// Store the line
	m.baseLines[m.currentLine] = value
	m.currentLine++

	// Clear input for next line
	m.textInput.SetValue("")
	m.err = nil

	// Check if we're done
	if m.currentLine >= m.session.Height {
		// Save base to session
		m.session.BaseFrame = Frame{
			Name:  "base",
			Lines: m.baseLines,
		}
		m.session.Save()

		m.statusMsg = "✓ Base character created!"
		m.screen = ScreenMenu
		m.currentLine = 0
	} else {
		m.statusMsg = fmt.Sprintf("Line %d/%d completed", m.currentLine, m.session.Height)
	}

	return m, nil
}

// handleBaseLineDelete handles Ctrl+D in base creation
func (m CreationModel) handleBaseLineDelete() (tea.Model, tea.Cmd) {
	// Delete last line
	if m.currentLine > 0 {
		m.currentLine--
		m.baseLines[m.currentLine] = ""
		m.statusMsg = fmt.Sprintf("Deleted line %d", m.currentLine+1)
	}
	return m, nil
}

// updateCreateBase is now deprecated - logic moved to handlers
func (m CreationModel) updateCreateBase(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

// handleStateNameSubmit handles state name input
func (m CreationModel) handleStateNameSubmit() (tea.Model, tea.Cmd) {
	stateName := strings.TrimSpace(m.textInput.Value())

	if stateName == "" {
		m.err = fmt.Errorf("state name cannot be empty")
		return m, nil
	}

	// Check if state already exists
	for _, state := range m.session.States {
		if state.Name == stateName {
			m.err = fmt.Errorf("state '%s' already exists", stateName)
			return m, nil
		}
	}

	// Determine state type
	standardStates := []string{"plan", "think", "execute", "wait", "error", "success"}
	m.currentStateType = "custom"
	for _, std := range standardStates {
		if stateName == std {
			m.currentStateType = "standard"
			break
		}
	}

	m.currentStateName = stateName
	m.stateFrameCount = 3 // Default to 3 frames (minimum)
	m.stateFrames = make([][]string, 0, m.stateFrameCount)
	m.currentFrame = 0
	m.currentFrameLine = 0
	m.frameLines = make([]string, m.session.Height)

	// Move to frame input
	m.screen = ScreenStateFrameInput
	m.textInput.SetValue("")
	m.textInput.CharLimit = m.session.Width // Limit to character width for patterns
	m.textInput.Placeholder = "Enter pattern (e.g., FFFFFFFF)"
	m.textInput.Focus()
	m.err = nil
	m.statusMsg = fmt.Sprintf("Creating '%s' state - Frame 1/%d, Line 1/%d", stateName, m.stateFrameCount, m.session.Height)

	return m, nil
}

// handleStateFrameLineSubmit handles frame line input for state
func (m CreationModel) handleStateFrameLineSubmit() (tea.Model, tea.Cmd) {
	value := m.textInput.Value()

	// Validate length
	if len(value) != m.session.Width {
		m.err = fmt.Errorf("expected %d characters, got %d", m.session.Width, len(value))
		return m, nil
	}

	// Apply mirroring
	value = applyMirroring(value)

	// Store the line
	m.frameLines[m.currentFrameLine] = value
	m.currentFrameLine++

	// Clear input for next line
	m.textInput.SetValue("")
	m.err = nil

	// Check if current frame is complete
	if m.currentFrameLine >= m.session.Height {
		// Save current frame
		frameCopy := make([]string, m.session.Height)
		copy(frameCopy, m.frameLines)
		m.stateFrames = append(m.stateFrames, frameCopy)

		m.currentFrame++
		m.currentFrameLine = 0
		m.frameLines = make([]string, m.session.Height)

		// Check if all frames are complete
		if m.currentFrame >= m.stateFrameCount {
			// Save state to session temporarily
			newState := StateSession{
				Name:           m.currentStateName,
				Description:    fmt.Sprintf("Agent %s state", m.currentStateName),
				StateType:      m.currentStateType,
				AnimationFPS:   5,
				AnimationLoops: 1,
				Frames:         make([]Frame, 0, len(m.stateFrames)),
			}

			for i, frameLines := range m.stateFrames {
				newState.Frames = append(newState.Frames, Frame{
					Name:  fmt.Sprintf("%s_frame_%d", m.currentStateName, i+1),
					Lines: frameLines,
				})
			}

			m.session.States = append(m.session.States, newState)
			m.session.Save()

			// Go to preview screen instead of menu
			m.screen = ScreenStatePreview
			m.statusMsg = fmt.Sprintf("Preview '%s' state - Press Enter to confirm, Esc to discard", m.currentStateName)
			m.currentFrame = 0
			m.previewFrameIndex = 0
			m.animating = true
			return m, tick() // Start animation
		} else {
			m.statusMsg = fmt.Sprintf("Frame %d/%d completed - Starting frame %d", m.currentFrame, m.stateFrameCount, m.currentFrame+1)
		}
	} else {
		m.statusMsg = fmt.Sprintf("Frame %d/%d, Line %d/%d completed", m.currentFrame+1, m.stateFrameCount, m.currentFrameLine, m.session.Height)
	}

	return m, nil
}

// handleStateFrameLineDelete handles Ctrl+D in state frame input
func (m CreationModel) handleStateFrameLineDelete() (tea.Model, tea.Cmd) {
	if m.currentFrameLine > 0 {
		m.currentFrameLine--
		m.frameLines[m.currentFrameLine] = ""
		m.statusMsg = fmt.Sprintf("Deleted line %d", m.currentFrameLine+1)
	} else if m.currentFrame > 0 {
		// Go back to previous frame
		m.currentFrame--
		if len(m.stateFrames) > m.currentFrame {
			m.stateFrames = m.stateFrames[:m.currentFrame]
		}
		m.currentFrameLine = m.session.Height - 1
		m.statusMsg = fmt.Sprintf("Back to frame %d", m.currentFrame+1)
	}
	return m, nil
}

// View renders the UI
func (m CreationModel) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	// Calculate split-pane widths
	leftWidth := m.width / 2
	rightWidth := m.width - leftWidth

	// Render left pane (menu/creation)
	leftPane := m.renderLeftPane(leftWidth)

	// Render right pane (live preview)
	rightPane := m.renderRightPane(rightWidth)

	// Combine panes side by side
	splitView := lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftPane,
		rightPane,
	)

	// Add status bar at bottom
	statusBar := m.renderStatusBar()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		splitView,
		statusBar,
	)
}

// renderLeftPane renders the left pane (menu or creation interface)
func (m CreationModel) renderLeftPane(width int) string {
	var content string

	title := m.styles.title.Render(fmt.Sprintf("◢ CHARACTER: %s", m.session.Name))

	switch m.screen {
	case ScreenMenu:
		content = m.renderMenu()
	case ScreenCreateBase:
		content = m.renderCreateBase()
	case ScreenStateNameInput:
		content = m.renderStateNameInput()
	case ScreenStateFrameInput:
		content = m.renderStateFrameInput()
	case ScreenStatePreview:
		content = m.renderStatePreview()
	case ScreenAnimateAll:
		content = m.renderAnimateAll()
	case ScreenExport:
		content = m.renderExport()
	}

	pane := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		content,
	)

	return m.styles.border.
		Width(width - 4).
		Height(m.height - 5).
		Render(pane)
}

// renderMenu renders the main menu
func (m CreationModel) renderMenu() string {
	var menu strings.Builder

	menu.WriteString(m.styles.helpText.Render("Use ↑/↓ to navigate, Enter to select, q to quit"))
	menu.WriteString("\n\n")

	// Show base status
	baseStatus := "✗ Not created"
	if len(m.session.BaseFrame.Lines) > 0 {
		baseStatus = "✓ Created"
	}
	menu.WriteString(fmt.Sprintf("Base: %s | States: %d\n\n", baseStatus, len(m.session.States)))

	// Show menu options
	for i, option := range m.menuOptions {
		cursor := "  "
		if i == m.menuCursor {
			cursor = "▶ "
			menu.WriteString(m.styles.selectedItem.Render(cursor + option))
		} else {
			menu.WriteString(m.styles.menuItem.Render(cursor + option))
		}
		menu.WriteString("\n")
	}

	return menu.String()
}

// renderCreateBase renders the base creation interface
func (m CreationModel) renderCreateBase() string {
	var content strings.Builder

	content.WriteString(m.styles.title.Render("CREATE BASE CHARACTER"))
	content.WriteString("\n\n")
	content.WriteString(m.styles.helpText.Render(fmt.Sprintf("Line %d/%d", m.currentLine+1, m.session.Height)))
	content.WriteString("\n\n")
	content.WriteString(patterns.GetPatternHelp())
	content.WriteString("\n\n")

	// Show text input
	content.WriteString(m.textInput.View())
	content.WriteString("\n\n")

	// Show error if any
	if m.err != nil {
		content.WriteString(m.styles.errorMsg.Render(fmt.Sprintf("✗ %s", m.err.Error())))
		content.WriteString("\n")
	}

	// Show completed lines
	content.WriteString("\nCompleted lines:\n")
	compiler := infrastructure.NewPatternCompiler()
	for i := 0; i < m.currentLine; i++ {
		if m.baseLines[i] != "" {
			compiled := compiler.Compile(m.baseLines[i])
			content.WriteString(fmt.Sprintf("  %d: %s\n", i+1, compiled))
		}
	}

	content.WriteString("\n")
	content.WriteString(m.styles.helpText.Render("Enter: confirm line | Ctrl+D: delete last | Esc: cancel"))

	return content.String()
}

// renderStateNameInput renders the state name input
func (m CreationModel) renderStateNameInput() string {
	var content strings.Builder

	content.WriteString(m.styles.title.Render("ADD AGENT STATE"))
	content.WriteString("\n\n")
	content.WriteString("Enter a state name for your character:\n\n")
	content.WriteString("Standard states:\n")
	content.WriteString("  • plan     - Agent analyzing and planning\n")
	content.WriteString("  • think    - Agent processing information\n")
	content.WriteString("  • execute  - Agent performing actions\n")
	content.WriteString("  • wait     - Agent waiting for input\n")
	content.WriteString("  • error    - Agent handling errors\n")
	content.WriteString("  • success  - Agent celebrating success\n\n")
	content.WriteString("Or enter a custom state name.\n\n")

	// Show text input
	content.WriteString(m.textInput.View())
	content.WriteString("\n\n")

	// Show error if any
	if m.err != nil {
		content.WriteString(m.styles.errorMsg.Render(fmt.Sprintf("✗ %s", m.err.Error())))
		content.WriteString("\n")
	}

	// Show existing states
	if len(m.session.States) > 0 {
		content.WriteString("\nExisting states:\n")
		for _, state := range m.session.States {
			content.WriteString(fmt.Sprintf("  • %s (%d frames)\n", state.Name, len(state.Frames)))
		}
	}

	content.WriteString("\n")
	content.WriteString(m.styles.helpText.Render("Enter: confirm | Esc: cancel"))

	return content.String()
}

// renderStateFrameInput renders the state frame input
func (m CreationModel) renderStateFrameInput() string {
	var content strings.Builder

	content.WriteString(m.styles.title.Render(fmt.Sprintf("CREATE STATE: %s", m.currentStateName)))
	content.WriteString("\n\n")
	content.WriteString(m.styles.helpText.Render(fmt.Sprintf("Frame %d/%d | Line %d/%d",
		m.currentFrame+1, m.stateFrameCount, m.currentFrameLine+1, m.session.Height)))
	content.WriteString("\n\n")
	content.WriteString(patterns.GetPatternHelp())
	content.WriteString("\n\n")

	// Show text input
	content.WriteString(m.textInput.View())
	content.WriteString("\n\n")

	// Show error if any
	if m.err != nil {
		content.WriteString(m.styles.errorMsg.Render(fmt.Sprintf("✗ %s", m.err.Error())))
		content.WriteString("\n")
	}

	// Show current frame progress
	content.WriteString(fmt.Sprintf("\nCurrent frame (%d/%d):\n", m.currentFrame+1, m.stateFrameCount))
	compiler := infrastructure.NewPatternCompiler()
	for i := 0; i < m.currentFrameLine; i++ {
		if m.frameLines[i] != "" {
			compiled := compiler.Compile(m.frameLines[i])
			content.WriteString(fmt.Sprintf("  %d: %s\n", i+1, compiled))
		}
	}

	// Show completed frames
	if len(m.stateFrames) > 0 {
		content.WriteString(fmt.Sprintf("\nCompleted frames: %d/%d\n", len(m.stateFrames), m.stateFrameCount))
	}

	content.WriteString("\n")
	content.WriteString(m.styles.helpText.Render("Enter: confirm line | Ctrl+D: delete last | Esc: cancel"))

	return content.String()
}

// renderStatePreview renders the final preview of completed state
func (m CreationModel) renderStatePreview() string {
	var content strings.Builder

	content.WriteString(m.styles.title.Render(fmt.Sprintf("PREVIEW: %s", m.currentStateName)))
	content.WriteString("\n\n")
	content.WriteString(m.styles.helpText.Render("Review your animated state"))
	content.WriteString("\n\n")

	// Get the last state (the one we just created)
	if len(m.session.States) > 0 {
		state := m.session.States[len(m.session.States)-1]

		content.WriteString(fmt.Sprintf("State: %s (%s)\n", state.Name, state.StateType))
		content.WriteString(fmt.Sprintf("Frames: %d\n", len(state.Frames)))
		content.WriteString(fmt.Sprintf("Animation: %d FPS, %d loops\n\n", state.AnimationFPS, state.AnimationLoops))

		content.WriteString("All frames:\n\n")
		compiler := infrastructure.NewPatternCompiler()
		for i, frame := range state.Frames {
			content.WriteString(fmt.Sprintf("Frame %d:\n", i+1))
			for _, line := range frame.Lines {
				compiled := compiler.Compile(line)
				content.WriteString("  " + compiled + "\n")
			}
			content.WriteString("\n")
		}
	}

	content.WriteString(m.styles.helpText.Render("Enter: confirm and save | Esc: discard and cancel"))

	return content.String()
}

// renderAnimateAll renders the animate all states screen
func (m CreationModel) renderAnimateAll() string {
	var content strings.Builder

	content.WriteString(m.styles.title.Render("ANIMATE ALL STATES"))
	content.WriteString("\n\n")

	if len(m.session.States) == 0 {
		content.WriteString("No states to animate.\n")
		content.WriteString("\nCreate some states first!\n")
		return content.String()
	}

	state := m.session.States[m.currentStateIndex]

	content.WriteString(fmt.Sprintf("State %d/%d: %s (%s)\n\n",
		m.currentStateIndex+1, len(m.session.States), state.Name, state.StateType))

	content.WriteString(fmt.Sprintf("Frames: %d | FPS: %d | Loops: %d\n\n",
		len(state.Frames), state.AnimationFPS, state.AnimationLoops))

	// Show base character for reference
	if len(m.session.BaseFrame.Lines) > 0 {
		content.WriteString("Base Character:\n")
		compiler := infrastructure.NewPatternCompiler()
		for _, line := range m.session.BaseFrame.Lines {
			compiled := compiler.Compile(line)
			content.WriteString("  " + compiled + "\n")
		}
		content.WriteString("\n")
	}

	// Show all states
	content.WriteString("All States:\n")
	for i, st := range m.session.States {
		prefix := "  "
		if i == m.currentStateIndex {
			prefix = "▶ "
		}
		content.WriteString(fmt.Sprintf("%s%s (%d frames)\n", prefix, st.Name, len(st.Frames)))
	}

	content.WriteString("\n")
	content.WriteString(m.styles.helpText.Render("←/→: switch states | Esc: back to menu"))

	return content.String()
}

// renderExport renders the export screen
func (m CreationModel) renderExport() string {
	var content strings.Builder

	content.WriteString(m.styles.title.Render("EXPORT FOR CONTRIBUTION"))
	content.WriteString("\n\n")

	// Validate character is complete
	if len(m.session.BaseFrame.Lines) == 0 {
		content.WriteString(m.styles.errorMsg.Render("✗ Cannot export: No base character"))
		content.WriteString("\n\nCreate a base character first.")
		return content.String()
	}

	if len(m.session.States) < 3 {
		content.WriteString(m.styles.errorMsg.Render(fmt.Sprintf("✗ Cannot export: Only %d states (minimum 3 required)", len(m.session.States))))
		content.WriteString("\n\nMinimum required states: plan, think, execute")
		content.WriteString("\n\nCreate more states to export.")
		return content.String()
	}

	// Check for required states
	hasRequired := map[string]bool{"plan": false, "think": false, "execute": false}
	for _, state := range m.session.States {
		if _, ok := hasRequired[state.Name]; ok {
			hasRequired[state.Name] = true
		}
	}

	missingStates := []string{}
	for state, has := range hasRequired {
		if !has {
			missingStates = append(missingStates, state)
		}
	}

	if len(missingStates) > 0 {
		content.WriteString(m.styles.errorMsg.Render(fmt.Sprintf("✗ Missing required states: %s", strings.Join(missingStates, ", "))))
		content.WriteString("\n\nCreate these required states before exporting.")
		return content.String()
	}

	// Show export summary
	content.WriteString(fmt.Sprintf("Character: %s (%dx%d)\n\n", m.session.Name, m.session.Width, m.session.Height))
	content.WriteString(fmt.Sprintf("Base: ✓ Created\n"))
	content.WriteString(fmt.Sprintf("States: %d\n\n", len(m.session.States)))

	content.WriteString("States to export:\n")
	for _, state := range m.session.States {
		content.WriteString(fmt.Sprintf("  • %s (%s, %d frames)\n", state.Name, state.StateType, len(state.Frames)))
	}

	content.WriteString("\n")
	content.WriteString(m.styles.title.Render("Export Files:"))
	content.WriteString("\n\n")
	content.WriteString(fmt.Sprintf("  📄 %s.json\n", m.session.Name))
	content.WriteString(fmt.Sprintf("  📄 %s-README.md\n\n", m.session.Name))

	content.WriteString("✓ Ready to export!\n\n")
	content.WriteString(m.styles.helpText.Render("Enter: export files | Esc: cancel"))

	return content.String()
}

// renderRightPane renders the right pane (live preview)
func (m CreationModel) renderRightPane(width int) string {
	title := m.styles.title.Render("◢ LIVE PREVIEW")

	var preview strings.Builder
	compiler := infrastructure.NewPatternCompiler()

	switch m.screen {
	case ScreenMenu:
		// Show base if available
		if len(m.session.BaseFrame.Lines) > 0 {
			preview.WriteString("Base Character:\n\n")
			for _, line := range m.session.BaseFrame.Lines {
				compiled := compiler.Compile(line)
				preview.WriteString("  " + compiled + "\n")
			}
		} else {
			preview.WriteString("No character created yet.\n\n")
			preview.WriteString("Select 'Create base character' to begin.")
		}

	case ScreenCreateBase:
		// Show current progress
		preview.WriteString("Building base character:\n\n")
		for i := 0; i < m.session.Height; i++ {
			if i < m.currentLine && m.baseLines[i] != "" {
				compiled := compiler.Compile(m.baseLines[i])
				preview.WriteString("  " + compiled + "\n")
			} else if i == m.currentLine {
				// Show current line being edited
				currentValue := m.textInput.Value()
				if currentValue != "" {
					compiled := compiler.Compile(currentValue)
					preview.WriteString("  " + compiled + " ◀\n")
				} else {
					preview.WriteString("  " + strings.Repeat("_", m.session.Width) + " ◀\n")
				}
			} else {
				preview.WriteString("  " + strings.Repeat("_", m.session.Width) + "\n")
			}
		}

	case ScreenStateNameInput:
		// Show base character while choosing state name
		if len(m.session.BaseFrame.Lines) > 0 {
			preview.WriteString("Base Character:\n\n")
			for _, line := range m.session.BaseFrame.Lines {
				compiled := compiler.Compile(line)
				preview.WriteString("  " + compiled + "\n")
			}
		}

		// Show existing states
		if len(m.session.States) > 0 {
			preview.WriteString("\nExisting States:\n")
			for _, state := range m.session.States {
				preview.WriteString(fmt.Sprintf("\n%s (%d frames):\n", state.Name, len(state.Frames)))
				if len(state.Frames) > 0 {
					for _, line := range state.Frames[0].Lines {
						compiled := compiler.Compile(line)
						preview.WriteString("  " + compiled + "\n")
					}
				}
			}
		}

	case ScreenStateFrameInput:
		// Show current frame being built
		preview.WriteString(fmt.Sprintf("Building '%s' state:\n", m.currentStateName))
		preview.WriteString(fmt.Sprintf("Frame %d/%d\n\n", m.currentFrame+1, m.stateFrameCount))

		for i := 0; i < m.session.Height; i++ {
			if i < m.currentFrameLine && m.frameLines[i] != "" {
				compiled := compiler.Compile(m.frameLines[i])
				preview.WriteString("  " + compiled + "\n")
			} else if i == m.currentFrameLine {
				// Show current line being edited
				currentValue := m.textInput.Value()
				if currentValue != "" {
					compiled := compiler.Compile(currentValue)
					preview.WriteString("  " + compiled + " ◀\n")
				} else {
					preview.WriteString("  " + strings.Repeat("_", m.session.Width) + " ◀\n")
				}
			} else {
				preview.WriteString("  " + strings.Repeat("_", m.session.Width) + "\n")
			}
		}

		// Show completed frames preview
		if len(m.stateFrames) > 0 {
			preview.WriteString(fmt.Sprintf("\n\nCompleted frames: %d\n", len(m.stateFrames)))
			for fi, frame := range m.stateFrames {
				preview.WriteString(fmt.Sprintf("\nFrame %d:\n", fi+1))
				for _, line := range frame {
					compiled := compiler.Compile(line)
					preview.WriteString("  " + compiled + "\n")
				}
			}
		}

	case ScreenStatePreview:
		// Show animated preview of the completed state
		if len(m.session.States) > 0 {
			state := m.session.States[len(m.session.States)-1]
			preview.WriteString(fmt.Sprintf("◢ ANIMATING: %s\n\n", state.Name))
			preview.WriteString(fmt.Sprintf("Frame %d/%d @ 5 FPS\n\n", m.previewFrameIndex+1, len(state.Frames)))

			// Show only the current frame (animated)
			if m.previewFrameIndex < len(state.Frames) {
				frame := state.Frames[m.previewFrameIndex]
				// Center the animation
				preview.WriteString("\n")
				for _, line := range frame.Lines {
					compiled := compiler.Compile(line)
					preview.WriteString("    " + compiled + "\n")
				}
				preview.WriteString("\n")
			}

			preview.WriteString("\n━━━━━━━━━━━━━━━━━━━━━━\n\n")
			preview.WriteString("✓ Animation cycling!\n")
			preview.WriteString(fmt.Sprintf("  %d frames @ 5 FPS\n", len(state.Frames)))
			preview.WriteString("\n")

			// Show frame indicators
			preview.WriteString("Frames: ")
			for i := range state.Frames {
				if i == m.previewFrameIndex {
					preview.WriteString("● ")
				} else {
					preview.WriteString("○ ")
				}
			}
		}

	case ScreenAnimateAll:
		// Show current state animating
		if len(m.session.States) > 0 && m.currentStateIndex < len(m.session.States) {
			state := m.session.States[m.currentStateIndex]
			preview.WriteString(fmt.Sprintf("◢ %s\n\n", strings.ToUpper(state.Name)))
			preview.WriteString(fmt.Sprintf("Frame %d/%d @ %d FPS\n\n", m.previewFrameIndex+1, len(state.Frames), state.AnimationFPS))

			// Show current animated frame
			if m.previewFrameIndex < len(state.Frames) {
				frame := state.Frames[m.previewFrameIndex]
				preview.WriteString("\n")
				for _, line := range frame.Lines {
					compiled := compiler.Compile(line)
					preview.WriteString("    " + compiled + "\n")
				}
				preview.WriteString("\n")
			}

			preview.WriteString("\n━━━━━━━━━━━━━━━━━━━━━━\n\n")

			// Frame indicators
			preview.WriteString("Frames: ")
			for i := range state.Frames {
				if i == m.previewFrameIndex {
					preview.WriteString("● ")
				} else {
					preview.WriteString("○ ")
				}
			}
		}

	case ScreenExport:
		// Show export preview
		preview.WriteString("◢ EXPORT PREVIEW\n\n")

		if len(m.session.BaseFrame.Lines) > 0 {
			preview.WriteString("Base Character:\n\n")
			for _, line := range m.session.BaseFrame.Lines {
				compiled := compiler.Compile(line)
				preview.WriteString("  " + compiled + "\n")
			}
			preview.WriteString("\n")
		}

		if len(m.session.States) > 0 {
			preview.WriteString("States:\n\n")
			for _, state := range m.session.States {
				preview.WriteString(fmt.Sprintf("%s (%d frames):\n", state.Name, len(state.Frames)))
				if len(state.Frames) > 0 {
					for _, line := range state.Frames[0].Lines {
						compiled := compiler.Compile(line)
						preview.WriteString("  " + compiled + "\n")
					}
				}
				preview.WriteString("\n")
			}
		}
	}

	pane := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		preview.String(),
	)

	return m.styles.preview.
		Width(width - 4).
		Height(m.height - 5).
		Render(pane)
}

// renderStatusBar renders the bottom status bar
func (m CreationModel) renderStatusBar() string {
	status := m.statusMsg
	if status == "" {
		status = "Ready"
	}

	return m.styles.statusBar.
		Width(m.width).
		Render(fmt.Sprintf("◢ %s", status))
}

// exportCharacter exports the character to JSON and README files
func (m *CreationModel) exportCharacter() error {
	// Generate JSON content
	jsonData := struct {
		Name      string `json:"name"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		BaseFrame struct {
			Name  string   `json:"name"`
			Lines []string `json:"lines"`
		} `json:"base_frame"`
		States []struct {
			Name   string `json:"name"`
			Frames []struct {
				Lines []string `json:"lines"`
			} `json:"frames"`
		} `json:"states"`
	}{
		Name:   m.session.Name,
		Width:  m.session.Width,
		Height: m.session.Height,
	}

	// Add base frame
	jsonData.BaseFrame.Name = m.session.BaseFrame.Name
	jsonData.BaseFrame.Lines = m.session.BaseFrame.Lines

	// Add states
	for _, state := range m.session.States {
		stateData := struct {
			Name   string `json:"name"`
			Frames []struct {
				Lines []string `json:"lines"`
			} `json:"frames"`
		}{
			Name: state.Name,
		}

		for _, frame := range state.Frames {
			frameData := struct {
				Lines []string `json:"lines"`
			}{
				Lines: frame.Lines,
			}
			stateData.Frames = append(stateData.Frames, frameData)
		}

		jsonData.States = append(jsonData.States, stateData)
	}

	// Marshal to JSON
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write JSON file
	jsonFilename := m.session.Name + ".json"
	if err := os.WriteFile(jsonFilename, jsonBytes, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	// Generate README
	readme := m.generateReadme()
	readmeFilename := m.session.Name + "-README.md"
	if err := os.WriteFile(readmeFilename, []byte(readme), 0644); err != nil {
		return fmt.Errorf("failed to write README file: %w", err)
	}

	return nil
}

// generateReadme generates the contribution README content
func (m *CreationModel) generateReadme() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s Character\n\n", m.session.Name))
	sb.WriteString("## Character Information\n\n")
	sb.WriteString(fmt.Sprintf("- **Name:** %s\n", m.session.Name))
	sb.WriteString(fmt.Sprintf("- **Dimensions:** %dx%d\n", m.session.Width, m.session.Height))
	sb.WriteString(fmt.Sprintf("- **States:** %d\n\n", len(m.session.States)))

	sb.WriteString("## States\n\n")
	for _, state := range m.session.States {
		sb.WriteString(fmt.Sprintf("- **%s** (%s): %d frames\n", state.Name, state.StateType, len(state.Frames)))
	}
	sb.WriteString("\n")

	sb.WriteString("## Preview\n\n")
	sb.WriteString("```\n")
	if len(m.session.BaseFrame.Lines) > 0 {
		compiler := infrastructure.NewPatternCompiler()
		for _, line := range m.session.BaseFrame.Lines {
			sb.WriteString(compiler.Compile(line) + "\n")
		}
	}
	sb.WriteString("```\n\n")

	sb.WriteString("## Usage\n\n")
	sb.WriteString("```go\n")
	sb.WriteString(fmt.Sprintf("agent, _ := characters.LibraryAgent(\"%s\")\n", m.session.Name))
	sb.WriteString("agent.Plan(os.Stdout)   // Show plan state\n")
	sb.WriteString("agent.Think(os.Stdout)  // Show think state\n")
	sb.WriteString("agent.Execute(os.Stdout) // Show execute state\n")
	sb.WriteString("```\n\n")

	sb.WriteString("## Contribution\n\n")
	sb.WriteString("This character was created using Tangent character builder.\n\n")
	sb.WriteString("### Next Steps\n\n")
	sb.WriteString("1. Review the exported JSON file\n")
	sb.WriteString("2. Fork the Tangent repository on GitHub\n")
	sb.WriteString("3. Add your character JSON to the repository\n")
	sb.WriteString("4. Submit a Pull Request\n")

	return sb.String()
}

// StartCreationTUI starts the Bubbletea TUI for character creation
func StartCreationTUI(session *Session) error {
	p := tea.NewProgram(
		NewCreationModel(session),
		tea.WithAltScreen(),
	)

	_, err := p.Run()
	return err
}
