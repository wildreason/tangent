package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	fmt.Println("Testing Alex character...")

	// Load the character
	agent, err := characters.LibraryAgent("alex")
	if err != nil {
		fmt.Printf("Error loading alex: %v\n", err)
		return
	}

	char := agent.GetCharacter()
	fmt.Printf("✓ Loaded: %s (%dx%d)\n", char.Name, char.Width, char.Height)
	fmt.Printf("✓ BaseFrame: %d lines\n", len(char.BaseFrame.Lines))
	fmt.Printf("✓ States: %d\n", len(char.States))

	// Show base character
	fmt.Println("\n🔹 Base Character:")
	agent.ShowBase(os.Stdout)
	fmt.Println()

	// Test each state
	for stateName, state := range char.States {
		fmt.Printf("🔹 %s State (%d frames):\n", stateName, len(state.Frames))
		agent.AnimateState(os.Stdout, stateName, state.AnimationFPS, state.AnimationLoops)
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("✅ Alex character test completed!")
}
