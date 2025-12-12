package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// EditSession extends Session with source file tracking
type EditSession struct {
	*Session
	SourcePath string // Path to save back to
	IsMicro    bool
}

// getMicroJSONPath returns the path to micro.json source file
func getMicroJSONPath() (string, error) {
	// Get the directory of this source file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get caller info")
	}

	// Navigate from cmd/tangent-cli/edit.go to pkg/characters/microstateregistry/states/micro.json
	dir := filepath.Dir(filename)
	microPath := filepath.Join(dir, "..", "..", "pkg", "characters", "microstateregistry", "states", "micro.json")

	// Clean and resolve the path
	microPath = filepath.Clean(microPath)

	// Verify it exists
	if _, err := os.Stat(microPath); err != nil {
		return "", fmt.Errorf("micro.json not found at %s: %w", microPath, err)
	}

	return microPath, nil
}

// MicroJSON represents the micro.json structure
type MicroJSON struct {
	Name      string           `json:"name"`
	Width     int              `json:"width"`
	Height    int              `json:"height"`
	BaseFrame MicroJSONFrame   `json:"base_frame"`
	States    []MicroJSONState `json:"states"`
}

type MicroJSONFrame struct {
	Name  string   `json:"name,omitempty"`
	Lines []string `json:"lines"`
}

type MicroJSONState struct {
	Name   string           `json:"name"`
	Frames []MicroJSONFrame `json:"frames"`
}

// loadMicroJSON loads micro.json from filesystem
func loadMicroJSON() (*MicroJSON, string, error) {
	path, err := getMicroJSONPath()
	if err != nil {
		return nil, "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read micro.json: %w", err)
	}

	var micro MicroJSON
	if err := json.Unmarshal(data, &micro); err != nil {
		return nil, "", fmt.Errorf("failed to parse micro.json: %w", err)
	}

	return &micro, path, nil
}

// saveMicroJSON saves micro.json back to filesystem
func saveMicroJSON(micro *MicroJSON, path string) error {
	data, err := json.MarshalIndent(micro, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal micro.json: %w", err)
	}

	// Add trailing newline
	data = append(data, '\n')

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write micro.json: %w", err)
	}

	return nil
}

// convertMicroToSession converts MicroJSON to Session for TUI editing
func convertMicroToSession(micro *MicroJSON) *Session {
	session := NewSession("micro", micro.Width, micro.Height)

	// Convert base frame
	session.BaseFrame = Frame{
		Name:  micro.BaseFrame.Name,
		Lines: micro.BaseFrame.Lines,
	}

	// Convert states - need to fix the JSON structure issue
	for _, state := range micro.States {
		stateSession := StateSession{
			Name:           state.Name,
			AnimationFPS:   5, // Default
			AnimationLoops: 1,
		}

		for _, frame := range state.Frames {
			stateSession.Frames = append(stateSession.Frames, Frame{
				Lines: frame.Lines,
			})
		}

		session.States = append(session.States, stateSession)
	}

	return session
}

// convertSessionToMicro converts Session back to MicroJSON
func convertSessionToMicro(session *Session) *MicroJSON {
	micro := &MicroJSON{
		Name:   session.Name,
		Width:  session.Width,
		Height: session.Height,
		BaseFrame: MicroJSONFrame{
			Name:  session.BaseFrame.Name,
			Lines: session.BaseFrame.Lines,
		},
	}

	for _, state := range session.States {
		microState := MicroJSONState{
			Name: state.Name,
		}

		for _, frame := range state.Frames {
			microState.Frames = append(microState.Frames, MicroJSONFrame{
				Lines: frame.Lines,
			})
		}

		micro.States = append(micro.States, microState)
	}

	return micro
}

// handleEdit handles the edit subcommand
func handleEdit() {
	// Parse flags
	var stateName string
	var useMicro bool

	args := os.Args[2:] // Skip "tangent-cli edit"

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--micro", "-m":
			useMicro = true
		case "--help", "-h":
			printEditUsage()
			return
		default:
			if stateName == "" && args[i][0] != '-' {
				stateName = args[i]
			}
		}
	}

	// For now, only micro editing is supported
	if !useMicro {
		fmt.Println("Currently only --micro editing is supported")
		fmt.Println("Usage: tangent-cli edit [state] --micro")
		os.Exit(1)
	}

	// Load micro.json
	micro, path, err := loadMicroJSON()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded: %s\n", path)
	fmt.Printf("Character: %s (%dx%d)\n", micro.Name, micro.Width, micro.Height)
	fmt.Printf("States: %d\n", len(micro.States))

	// Convert to session
	session := convertMicroToSession(micro)

	// Find target state index if specified
	targetStateIdx := -1
	if stateName != "" {
		for i, state := range session.States {
			if state.Name == stateName {
				targetStateIdx = i
				break
			}
		}
		if targetStateIdx == -1 {
			fmt.Printf("\nError: state '%s' not found\n", stateName)
			fmt.Println("Available states:")
			for _, state := range session.States {
				fmt.Printf("  - %s (%d frames)\n", state.Name, len(state.Frames))
			}
			os.Exit(1)
		}
		fmt.Printf("\nEditing state: %s\n", stateName)
	}

	fmt.Println()

	// Start TUI with edit mode
	if err := StartEditTUI(session, path, targetStateIdx); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func printEditUsage() {
	fmt.Println("tangent-cli edit - Edit existing character states")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  tangent-cli edit [state] --micro    Edit micro avatar state")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  tangent-cli edit --micro            Open micro editor (menu)")
	fmt.Println("  tangent-cli edit approval --micro   Jump to approval state")
	fmt.Println("  tangent-cli edit resting --micro    Jump to resting state")
	fmt.Println()
	fmt.Println("The editor saves changes back to the source micro.json file.")
}
