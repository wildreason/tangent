package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/bubbletea"
)

// model wraps the AnimatedCharacter with additional UI state
type model struct {
	character *bubbletea.AnimatedCharacter
	states    []string
	stateIdx  int
	err       error
}

func initialModel() (model, error) {
	// Load a character from the library
	agent, err := characters.LibraryAgent("sam")
	if err != nil {
		return model{}, fmt.Errorf("failed to load character: %w", err)
	}

	// Create animated character (10 FPS = 100ms per frame)
	char := bubbletea.NewAnimatedCharacter(agent, 100*time.Millisecond)

	// Get available states
	states := char.ListStates()
	if len(states) > 0 {
		char.SetState(states[0])
	}

	return model{
		character: char,
		states:    states,
		stateIdx:  0,
	}, nil
}

func (m model) Init() tea.Cmd {
	return m.character.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "left":
			// Previous state
			if len(m.states) > 0 {
				m.stateIdx = (m.stateIdx - 1 + len(m.states)) % len(m.states)
				m.character.SetState(m.states[m.stateIdx])
			}

		case "right":
			// Next state
			if len(m.states) > 0 {
				m.stateIdx = (m.stateIdx + 1) % len(m.states)
				m.character.SetState(m.states[m.stateIdx])
			}

		case " ":
			// Toggle play/pause
			if m.character.IsPlaying() {
				m.character.Pause()
			} else {
				m.character.Play()
			}
		}
	}

	// Forward to character component
	charModel, cmd := m.character.Update(msg)
	m.character = charModel.(*bubbletea.AnimatedCharacter)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	s := "\n"
	s += "  Tangent Bubble Tea Integration Demo\n"
	s += "  ====================================\n\n"

	// Show character animation
	s += m.character.View() + "\n\n"

	// Show current state
	currentState := m.character.GetState()
	playStatus := "Playing"
	if !m.character.IsPlaying() {
		playStatus = "Paused"
	}
	s += fmt.Sprintf("  State: %s (%s)\n", currentState, playStatus)
	s += fmt.Sprintf("  Frame Rate: %.0f FPS\n\n", float64(time.Second)/float64(m.character.GetTickInterval()))

	// Show controls
	s += "  Controls:\n"
	s += "  ← →     Switch state\n"
	s += "  Space   Play/Pause\n"
	s += "  q       Quit\n"

	return s
}

func main() {
	m, err := initialModel()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
