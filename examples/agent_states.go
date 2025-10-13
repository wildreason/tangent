package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/domain"
)

func main() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Agent State API Demo                    ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()

	// Demo 1: Create a custom agent character with states
	fmt.Println("▢ Demo 1: Custom Agent Character with States")
	fmt.Println()

	// Create a robot character with agent states
	robot := &domain.Character{
		Name:        "demo-robot",
		Personality: "efficient",
		Width:       9,
		Height:      4,
		States: map[string]domain.State{
			"plan": {
				Name:        "Planning",
				Description: "Robot analyzing task",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name: "plan_frame",
						Lines: []string{
							"_L5FFF5R_",
							"_6FFFFF6_",
							"__?F_F?__",
							"__FF_FF__",
						},
					},
				},
			},
			"think": {
				Name:        "Thinking",
				Description: "Robot processing information",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name: "think_frame",
						Lines: []string{
							"_L5FFF5R_",
							"_6FFFFF6_",
							"__._._.__",
							"__FF_FF__",
						},
					},
				},
			},
			"execute": {
				Name:        "Executing",
				Description: "Robot performing action",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name: "execute_frame1",
						Lines: []string{
							"_L5FFF5R_",
							"_6FFFFF6_",
							"__>F_F>__",
							"__FF_FF__",
						},
					},
					{
						Name: "execute_frame2",
						Lines: []string{
							"_L5FFF5R_",
							"_6FFFFF6_",
							"__>>_>>__",
							"__FF_FF__",
						},
					},
				},
			},
			"wait": {
				Name:        "Waiting",
				Description: "Robot on standby",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name: "wait_frame",
						Lines: []string{
							"_L5FFF5R_",
							"_6FFFFF6_",
							"__FF_FF__",
							"__FF_FF__",
						},
					},
				},
			},
			"error": {
				Name:        "Error",
				Description: "Robot encountered error",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name: "error_frame",
						Lines: []string{
							"_L5FFF5R_",
							"_6FFFFF6_",
							"__XF_FX__",
							"__FF_FF__",
						},
					},
				},
			},
			"success": {
				Name:        "Success",
				Description: "Robot completed task",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name: "success_frame",
						Lines: []string{
							"_L5FFF5R_",
							"_6FFFFF6_",
							"__✓F_F✓__",
							"__FF_FF__",
						},
					},
				},
			},
		},
	}

	// Wrap in AgentCharacter for state-based API
	agent := characters.NewAgentCharacter(robot)

	fmt.Printf("Character: %s\n", agent.Name())
	fmt.Printf("Personality: %s\n", agent.Personality())
	fmt.Printf("Available states: %v\n", agent.ListStates())
	fmt.Println()

	// Demo 2: Use agent states
	fmt.Println("▢ Demo 2: Agent State Transitions")
	fmt.Println()

	fmt.Println("1. Agent is planning the task...")
	agent.Plan(os.Stdout)
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("2. Agent is thinking about the solution...")
	agent.Think(os.Stdout)
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("3. Agent is executing the task...")
	agent.Execute(os.Stdout)
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("4. Agent completed successfully!")
	agent.Success(os.Stdout)
	fmt.Println()

	// Demo 3: Error handling
	fmt.Println("▢ Demo 3: Error Handling")
	fmt.Println()

	fmt.Println("Simulating error condition...")
	agent.Error(os.Stdout)
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("Agent recovering...")
	agent.Think(os.Stdout)
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("Agent retrying...")
	agent.Execute(os.Stdout)
	fmt.Println()
	time.Sleep(1 * time.Second)

	fmt.Println("Success!")
	agent.Success(os.Stdout)
	fmt.Println()

	// Demo 4: State inspection
	fmt.Println("▢ Demo 4: State Inspection")
	fmt.Println()

	for _, stateName := range agent.ListStates() {
		desc, _ := agent.GetStateDescription(stateName)
		hasState := agent.HasState(stateName)
		fmt.Printf("State: %-10s | Has: %-5t | Description: %s\n", stateName, hasState, desc)
	}
	fmt.Println()

	// Demo 5: Custom state usage
	fmt.Println("▢ Demo 5: Custom State Access")
	fmt.Println()

	// Add a custom state
	robot.States["celebrate"] = domain.State{
		Name:        "Celebrating",
		Description: "Robot celebrating achievement",
		StateType:   "custom",
		Frames: []domain.Frame{
			{
				Name: "celebrate_frame",
				Lines: []string{
					"_L5FFF5R_",
					"56FFFFF65",
					"__*F_F*__",
					"__FF_FF__",
				},
			},
		},
	}

	fmt.Println("Agent celebrating with custom state:")
	agent.ShowState(os.Stdout, "celebrate")
	fmt.Println()

	// Demo 6: Practical workflow
	fmt.Println("▢ Demo 6: Practical AI Agent Workflow")
	fmt.Println()

	workflow := []struct {
		state   string
		message string
	}{
		{"wait", "Agent waiting for task..."},
		{"plan", "Analyzing requirements..."},
		{"think", "Designing solution..."},
		{"execute", "Implementing solution..."},
		{"success", "Task completed!"},
	}

	for _, step := range workflow {
		fmt.Println(step.message)
		agent.ShowState(os.Stdout, step.state)
		fmt.Println()
		time.Sleep(800 * time.Millisecond)
	}

	// Summary
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║  Summary                                 ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Agent State API provides:")
	fmt.Println("  ✓ State-based character access")
	fmt.Println("  ✓ Standard agent states (plan, think, execute, etc.)")
	fmt.Println("  ✓ Custom state support")
	fmt.Println("  ✓ State inspection and validation")
	fmt.Println("  ✓ Easy integration with AI agent workflows")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  agent, _ := characters.LibraryAgent(\"character-name\")")
	fmt.Println("  agent.Plan(os.Stdout)")
	fmt.Println("  agent.Think(os.Stdout)")
	fmt.Println("  agent.Execute(os.Stdout)")
	fmt.Println()
}

