package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/wildreason/tangent/pkg/adapters/bubbletea"
)

// Model represents the Bubble Tea application state
type model struct {
	agent    spinner.Model
	loading  spinner.Model
	rocket   spinner.Model
	quitting bool
}

func initialModel() model {
	// Load Tangent library characters and create spinners
	agentSpinner, _ := bubbletea.LibrarySpinner("wave", 6)
	loadingSpinner, _ := bubbletea.LibrarySpinner("pulse", 8)
	rocketSpinner, _ := bubbletea.LibrarySpinner("rocket", 5)

	return model{
		agent:   agentSpinner,
		loading: loadingSpinner,
		rocket:  rocketSpinner,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.agent.Tick,
		m.loading.Tick,
		m.rocket.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	var agentCmd, loadingCmd, rocketCmd tea.Cmd
	m.agent, agentCmd = m.agent.Update(msg)
	m.loading, loadingCmd = m.loading.Update(msg)
	m.rocket, rocketCmd = m.rocket.Update(msg)

	return m, tea.Batch(agentCmd, loadingCmd, rocketCmd)
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye! 👋\n"
	}

	// Define styles with Lip Gloss
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00FFFF")).
		Padding(1, 0)

	agentStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00FF00")).
		Padding(0, 2)

	loadingStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF00FF")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF00FF")).
		Padding(0, 2)

	rocketStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFF00")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFFF00")).
		Padding(0, 2)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Padding(1, 0)

	// Build view with styled characters
	title := titleStyle.Render("Tangent + Bubble Tea Integration Demo")

	agentView := lipgloss.JoinVertical(
		lipgloss.Center,
		agentStyle.Render(m.agent.View()),
		"Agent (wave)",
	)

	loadingView := lipgloss.JoinVertical(
		lipgloss.Center,
		loadingStyle.Render(m.loading.View()),
		"Loading (pulse)",
	)

	rocketView := lipgloss.JoinVertical(
		lipgloss.Center,
		rocketStyle.Render(m.rocket.View()),
		"Status (rocket)",
	)

	// Horizontal layout of characters
	charactersRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		agentView,
		"   ",
		loadingView,
		"   ",
		rocketView,
	)

	help := helpStyle.Render("Press 'q' to quit")

	// Vertical layout
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		charactersRow,
		"",
		help,
	)
}

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Tangent + Bubble Tea Example           ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("This example demonstrates:")
	fmt.Println("• Loading Tangent library characters")
	fmt.Println("• Creating Bubble Tea spinners")
	fmt.Println("• Styling with Lip Gloss")
	fmt.Println("• Multiple animated characters")
	fmt.Println()
	fmt.Println("Starting...")
	fmt.Println()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
