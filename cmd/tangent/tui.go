package main

import (
	"fmt"
	"strings"

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
	currentStateName  string
	currentStateType  string
	stateFrameCount   int
	stateFrames       [][]string
	currentFrame      int
	currentFrameLine  int
	frameLines        []string

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
	ti.CharLimit = session.Width
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

// Init initializes the model
func (m CreationModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages
func (m CreationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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
			m.textInput.Focus()
			m.statusMsg = "Creating base character - Press Enter to confirm each line"
		case 1: // Add agent state
			if len(m.session.BaseFrame.Lines) == 0 {
				m.statusMsg = "Create base character first!"
				m.err = fmt.Errorf("base character required")
			} else {
				m.screen = ScreenStateNameInput
				m.textInput.SetValue("")
				m.textInput.Placeholder = "Enter state name (plan, think, execute)"
				m.textInput.Focus()
				m.statusMsg = "Enter agent state name"
			}
		case 2: // Animate all states
			m.statusMsg = "Animation preview coming soon"
		case 3: // Export
			m.screen = ScreenExport
			m.statusMsg = "Export functionality"
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
			// Save state to session
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
			
			m.statusMsg = fmt.Sprintf("✓ State '%s' created with %d frames!", m.currentStateName, m.stateFrameCount)
			m.screen = ScreenMenu
			m.currentStateName = ""
			m.stateFrames = nil
			m.currentFrame = 0
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
	case ScreenExport:
		content = "Export screen"
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
		
	case ScreenExport:
		preview.WriteString("Export preview")
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

// StartCreationTUI starts the Bubbletea TUI for character creation
func StartCreationTUI(session *Session) error {
	p := tea.NewProgram(
		NewCreationModel(session),
		tea.WithAltScreen(),
	)

	_, err := p.Run()
	return err
}
