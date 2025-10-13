package characters

import (
	"bytes"
	"testing"

	"github.com/wildreason/tangent/pkg/characters/domain"
)

func TestAgentCharacter_Plan(t *testing.T) {
	// Create a test character with states
	char := &domain.Character{
		Name:        "test",
		Personality: "efficient",
		Width:       5,
		Height:      3,
		States: map[string]domain.State{
			"plan": {
				Name:        "Planning",
				Description: "Test planning state",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name:  "plan_frame",
						Lines: []string{"?????", ".....", "_____"},
					},
				},
			},
		},
	}

	agent := NewAgentCharacter(char)
	buf := &bytes.Buffer{}

	err := agent.Plan(buf)
	if err != nil {
		t.Errorf("Plan() error = %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Error("Plan() produced no output")
	}
}

func TestAgentCharacter_Think(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "analytical",
		Width:       5,
		Height:      3,
		States: map[string]domain.State{
			"think": {
				Name:        "Thinking",
				Description: "Test thinking state",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name:  "think_frame",
						Lines: []string{".....", ".....", "....."},
					},
				},
			},
		},
	}

	agent := NewAgentCharacter(char)
	buf := &bytes.Buffer{}

	err := agent.Think(buf)
	if err != nil {
		t.Errorf("Think() error = %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Error("Think() produced no output")
	}
}

func TestAgentCharacter_Execute(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "efficient",
		Width:       5,
		Height:      3,
		States: map[string]domain.State{
			"execute": {
				Name:        "Executing",
				Description: "Test executing state",
				StateType:   "standard",
				Frames: []domain.Frame{
					{
						Name:  "execute_frame",
						Lines: []string{">>>>>", ">>>>>", ">>>>>"},
					},
				},
			},
		},
	}

	agent := NewAgentCharacter(char)
	buf := &bytes.Buffer{}

	err := agent.Execute(buf)
	if err != nil {
		t.Errorf("Execute() error = %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Error("Execute() produced no output")
	}
}

func TestAgentCharacter_ShowState(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "friendly",
		Width:       5,
		Height:      3,
		States: map[string]domain.State{
			"custom": {
				Name:        "Custom",
				Description: "Test custom state",
				StateType:   "custom",
				Frames: []domain.Frame{
					{
						Name:  "custom_frame",
						Lines: []string{"*****", "*****", "*****"},
					},
				},
			},
		},
	}

	agent := NewAgentCharacter(char)
	buf := &bytes.Buffer{}

	err := agent.ShowState(buf, "custom")
	if err != nil {
		t.Errorf("ShowState() error = %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Error("ShowState() produced no output")
	}
}

func TestAgentCharacter_ShowState_NotFound(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "efficient",
		Width:       5,
		Height:      3,
		States:      map[string]domain.State{},
	}

	agent := NewAgentCharacter(char)
	buf := &bytes.Buffer{}

	err := agent.ShowState(buf, "nonexistent")
	if err == nil {
		t.Error("ShowState() should return error for nonexistent state")
	}
}

func TestAgentCharacter_ListStates(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "efficient",
		Width:       5,
		Height:      3,
		States: map[string]domain.State{
			"plan":    {Name: "Planning", StateType: "standard"},
			"think":   {Name: "Thinking", StateType: "standard"},
			"execute": {Name: "Executing", StateType: "standard"},
			"custom":  {Name: "Custom", StateType: "custom"},
		},
	}

	agent := NewAgentCharacter(char)
	states := agent.ListStates()

	if len(states) != 4 {
		t.Errorf("ListStates() returned %d states, expected 4", len(states))
	}

	// Check that states are sorted
	expectedStates := []string{"custom", "execute", "plan", "think"}
	for i, state := range states {
		if state != expectedStates[i] {
			t.Errorf("ListStates()[%d] = %s, expected %s", i, state, expectedStates[i])
		}
	}
}

func TestAgentCharacter_HasState(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "efficient",
		Width:       5,
		Height:      3,
		States: map[string]domain.State{
			"plan": {Name: "Planning", StateType: "standard"},
		},
	}

	agent := NewAgentCharacter(char)

	if !agent.HasState("plan") {
		t.Error("HasState('plan') should return true")
	}

	if agent.HasState("nonexistent") {
		t.Error("HasState('nonexistent') should return false")
	}
}

func TestAgentCharacter_GetStateDescription(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "efficient",
		Width:       5,
		Height:      3,
		States: map[string]domain.State{
			"plan": {
				Name:        "Planning",
				Description: "Test description",
				StateType:   "standard",
			},
		},
	}

	agent := NewAgentCharacter(char)

	desc, err := agent.GetStateDescription("plan")
	if err != nil {
		t.Errorf("GetStateDescription() error = %v", err)
	}

	if desc != "Test description" {
		t.Errorf("GetStateDescription() = %s, expected 'Test description'", desc)
	}

	_, err = agent.GetStateDescription("nonexistent")
	if err == nil {
		t.Error("GetStateDescription() should return error for nonexistent state")
	}
}

func TestAgentCharacter_Name(t *testing.T) {
	char := &domain.Character{
		Name:        "test-character",
		Personality: "efficient",
		Width:       5,
		Height:      3,
		States:      map[string]domain.State{},
	}

	agent := NewAgentCharacter(char)

	if agent.Name() != "test-character" {
		t.Errorf("Name() = %s, expected 'test-character'", agent.Name())
	}
}

func TestAgentCharacter_Personality(t *testing.T) {
	char := &domain.Character{
		Name:        "test",
		Personality: "friendly",
		Width:       5,
		Height:      3,
		States:      map[string]domain.State{},
	}

	agent := NewAgentCharacter(char)

	if agent.Personality() != "friendly" {
		t.Errorf("Personality() = %s, expected 'friendly'", agent.Personality())
	}
}

func TestAgentCharacter_NilCharacter(t *testing.T) {
	agent := NewAgentCharacter(nil)

	if agent.Name() != "" {
		t.Error("Name() should return empty string for nil character")
	}

	if agent.Personality() != "" {
		t.Error("Personality() should return empty string for nil character")
	}

	if len(agent.ListStates()) != 0 {
		t.Error("ListStates() should return empty slice for nil character")
	}

	if agent.HasState("any") {
		t.Error("HasState() should return false for nil character")
	}
}

