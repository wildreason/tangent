package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/domain"
)

func main() {
	fmt.Println("🎬 Testing Monsterra Character Animation")
	fmt.Println("========================================")

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

	fmt.Println("\n🎬 Testing Animated States")
	fmt.Println("==========================")

	// Test animated states
	states := []string{"plan", "think", "execute"}
	for _, stateName := range states {
		fmt.Printf("\n🎭 Animating %s state (3 seconds):\n", stateName)
		fmt.Println("Press Ctrl+C to skip...")

		// Start animation in background
		go func(name string) {
			agent.AnimateState(os.Stdout, name, 5, 3) // 5 FPS, 3 loops
		}(stateName)

		// Wait 3 seconds
		time.Sleep(3 * time.Second)
		fmt.Println("\n--- Animation complete ---\n")
	}

	fmt.Println("✅ Monsterra animation test completed!")
}
