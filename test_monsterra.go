package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/domain"
)

func main() {
	fmt.Println("🧪 Testing Monsterra Character")
	fmt.Println("==============================")

	// Load the monsterra.json file
	data, err := os.ReadFile("monsterra.json")
	if err != nil {
		fmt.Printf("❌ Error reading monsterra.json: %v\n", err)
		return
	}

	// Parse the JSON into a Session structure
	var session struct {
		Name        string `json:"name"`
		Personality string `json:"personality"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
		BaseFrame   struct {
			Name  string   `json:"name"`
			Lines []string `json:"lines"`
		} `json:"base_frame"`
		States []struct {
			Name           string `json:"name"`
			Description    string `json:"description"`
			StateType      string `json:"state_type"`
			AnimationFPS   int    `json:"animation_fps"`
			AnimationLoops int    `json:"animation_loops"`
			Frames         []struct {
				Name  string   `json:"name"`
				Lines []string `json:"lines"`
			} `json:"frames"`
		} `json:"states"`
	}

	if err := json.Unmarshal(data, &session); err != nil {
		fmt.Printf("❌ Error parsing JSON: %v\n", err)
		return
	}

	fmt.Printf("✅ Loaded character: %s (%s)\n", session.Name, session.Personality)
	fmt.Printf("   Dimensions: %dx%d\n", session.Width, session.Height)
	fmt.Printf("   States: %d\n", len(session.States))

	// Convert to domain.Character
	character := &domain.Character{
		Name:        session.Name,
		Personality: session.Personality,
		Width:       session.Width,
		Height:      session.Height,
		BaseFrame: domain.Frame{
			Name:  session.BaseFrame.Name,
			Lines: session.BaseFrame.Lines,
		},
		States: make(map[string]domain.State),
	}

	// Convert states
	for _, stateData := range session.States {
		frames := make([]domain.Frame, len(stateData.Frames))
		for i, frameData := range stateData.Frames {
			frames[i] = domain.Frame{
				Name:  frameData.Name,
				Lines: frameData.Lines,
			}
		}

		character.States[stateData.Name] = domain.State{
			Name:           stateData.Name,
			Description:    stateData.Description,
			StateType:      stateData.StateType,
			Frames:         frames,
			AnimationFPS:   stateData.AnimationFPS,
			AnimationLoops: stateData.AnimationLoops,
		}
	}

	// Create AgentCharacter wrapper
	agent := characters.NewAgentCharacter(character)

	fmt.Println("\n🎭 Testing Agent States")
	fmt.Println("=======================")

	// Test base character
	fmt.Println("\n📋 Base Character (Idle):")
	agent.ShowBase(os.Stdout)

	// Test each state
	states := []string{"plan", "think", "execute"}
	for _, stateName := range states {
		fmt.Printf("\n🔄 %s State:\n", stateName)
		if err := agent.ShowState(os.Stdout, stateName); err != nil {
			fmt.Printf("❌ Error showing %s state: %v\n", stateName, err)
		}
	}

	// Test state inspection
	fmt.Println("\n🔍 State Inspection")
	fmt.Println("==================")
	fmt.Printf("Available states: %v\n", agent.ListStates())

	for _, stateName := range states {
		if agent.HasState(stateName) {
			desc, _ := agent.GetStateDescription(stateName)
			fmt.Printf("  %s: %s\n", stateName, desc)
		}
	}

	fmt.Println("\n✅ Monsterra character test completed!")
}
