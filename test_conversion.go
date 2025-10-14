package main

import (
	"fmt"

	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	agent, err := characters.LibraryAgent("testchar")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	char := agent.GetCharacter()
	fmt.Printf("Success: %s\n", char.Name)
	fmt.Printf("BaseFrame: %+v\n", char.BaseFrame)
	fmt.Printf("States: %d\n", len(char.States))
	for name, state := range char.States {
		fmt.Printf("  State %s: %d frames\n", name, len(state.Frames))
		for i, frame := range state.Frames {
			fmt.Printf("    Frame %d: %s\n", i+1, frame.Name)
		}
	}
}
