package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wildreason/tangent/pkg/characters"
)

// handleList shows a simple list of available agents
func handleList() {
	// Get all library characters
	names := characters.ListLibrary()

	if len(names) == 0 {
		fmt.Println("No agents available.")
		return
	}

	fmt.Println("Available Agents:")
	fmt.Println()

	for _, name := range names {
		fmt.Printf("  • %s\n", name)
	}

	fmt.Println()
	fmt.Printf("Total: %d agent%s available\n", len(names), pluralize(len(names)))
	fmt.Println()
	fmt.Println("View agent: tangent browse <name>")
}

// handleListAgent shows and animates a specific agent
func handleListAgent(name string) {
	// Parse optional flags (reuse logic from handleDemo)
	var targetState string
	var overrideFPS int
	var overrideLoops int

	// Parse flags from os.Args starting from index 3 (after "tangent browse <name>")
	for i := 3; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--state":
			if i+1 < len(os.Args) {
				targetState = os.Args[i+1]
				i++
			}
		case "--fps":
			if i+1 < len(os.Args) {
				if fps, err := strconv.Atoi(os.Args[i+1]); err == nil {
					overrideFPS = fps
				}
				i++
			}
		case "--loops":
			if i+1 < len(os.Args) {
				if loops, err := strconv.Atoi(os.Args[i+1]); err == nil {
					overrideLoops = loops
				}
				i++
			}
		}
	}

	// Load character
	agent, err := characters.LibraryAgent(name)
	if err != nil {
		fmt.Printf("Error: agent '%s' not found\n", name)
		fmt.Println("Available agents:")
		names := characters.ListLibrary()
		for _, name := range names {
			fmt.Printf("  • %s\n", name)
		}
		os.Exit(1)
	}

	char := agent.GetCharacter()

	// Print agent header
	fmt.Printf("Agent: %s (%dx%d)\n", char.Name, char.Width, char.Height)
	if char.Personality != "" {
		fmt.Printf("Personality: %s\n", char.Personality)
	}
	fmt.Println()

	// Show base character
	fmt.Println("Base Character:")
	agent.ShowBase(os.Stdout)
	fmt.Println()

	// Animate states
	if targetState != "" {
		// Animate specific state
		state, exists := char.States[targetState]
		if !exists {
			fmt.Printf("Error: state '%s' not found\n", targetState)
			fmt.Println("Available states:")
			for name := range char.States {
				fmt.Printf("  • %s\n", name)
			}
			os.Exit(1)
		}

		fps := state.AnimationFPS
		loops := state.AnimationLoops
		if overrideFPS > 0 {
			fps = overrideFPS
		}
		if overrideLoops > 0 {
			loops = overrideLoops
		}

		fmt.Printf("Animating '%s' (%d frames) at %d FPS for %d loops\n", targetState, len(state.Frames), fps, loops)
		agent.AnimateState(os.Stdout, targetState, fps, loops)
		fmt.Println()
	} else {
		// List available states when no specific state requested
		if len(char.States) > 0 {
			fmt.Println("Available States:")
			stateNames := make([]string, 0, len(char.States))
			for stateName := range char.States {
				stateNames = append(stateNames, stateName)
			}

			// Sort for consistent order
			for i := 0; i < len(stateNames); i++ {
				for j := i + 1; j < len(stateNames); j++ {
					if stateNames[i] > stateNames[j] {
						stateNames[i], stateNames[j] = stateNames[j], stateNames[i]
					}
				}
			}

			for _, stateName := range stateNames {
				state := char.States[stateName]
				fmt.Printf("  • %s (%d frames, %d FPS, %d loops)\n",
					stateName, len(state.Frames), state.AnimationFPS, state.AnimationLoops)
			}
			fmt.Println()
		}
		// Animate all states in stable order
		stateNames := make([]string, 0, len(char.States))
		for name := range char.States {
			stateNames = append(stateNames, name)
		}

		// Sort for consistent order
		for i := 0; i < len(stateNames); i++ {
			for j := i + 1; j < len(stateNames); j++ {
				if stateNames[i] > stateNames[j] {
					stateNames[i], stateNames[j] = stateNames[j], stateNames[i]
				}
			}
		}

		for _, stateName := range stateNames {
			state := char.States[stateName]
			fps := state.AnimationFPS
			loops := state.AnimationLoops
			if overrideFPS > 0 {
				fps = overrideFPS
			}
			if overrideLoops > 0 {
				loops = overrideLoops
			}

			fmt.Printf("Animating '%s' (%d frames) at %d FPS for %d loops\n", stateName, len(state.Frames), fps, loops)
			agent.AnimateState(os.Stdout, stateName, fps, loops)
			fmt.Println()
		}
	}

	fmt.Println("✅ View complete!")
}
