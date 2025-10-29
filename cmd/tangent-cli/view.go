package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wildreason/tangent/pkg/characters"
	"github.com/wildreason/tangent/pkg/characters/domain"
)

// handleView implements `tangent view` to preview sessions or JSON without admin register
func handleView(args []string) {
	fs := flag.NewFlagSet("view", flag.ContinueOnError)
	sessionName := fs.String("session", "", "Session name to load and preview")
	jsonPath := fs.String("json", "", "JSON file to load and preview")
	state := fs.String("state", "plan", "State to animate")
	fps := fs.Int("fps", 5, "Frames per second")
	loops := fs.Int("loops", 1, "Number of loops")

	if err := fs.Parse(args); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if *sessionName == "" && *jsonPath == "" {
		// List sessions for convenience
		names, _ := ListSessions()
		if len(names) == 0 {
			fmt.Println("No sessions found. Create one with: tangent create")
			return
		}
		fmt.Println("Available Sessions:\n")
		for _, n := range names {
			fmt.Println("  •", n)
		}
		fmt.Println("\nPreview a session: tangent view --session <name> --state plan --fps 8 --loops 2")
		return
	}

	var sess *Session
	var err error
	if *sessionName != "" {
		sess, err = LoadSession(*sessionName)
		if err != nil {
			fmt.Println("✗ Failed to load session:", err)
			return
		}
	} else {
		// Load from JSON
		data, rerr := os.ReadFile(*jsonPath)
		if rerr != nil {
			fmt.Println("✗ Failed to read JSON:", rerr)
			return
		}
		// Expect same schema as exported contribution JSON
		var tmp struct {
			Name      string `json:"name"`
			Width     int    `json:"width"`
			Height    int    `json:"height"`
			BaseFrame struct {
				Name  string   `json:"name"`
				Lines []string `json:"lines"`
			} `json:"base_frame"`
			States []struct {
				Name   string `json:"name"`
				Frames []struct {
					Lines []string `json:"lines"`
				} `json:"frames"`
			} `json:"states"`
		}
		if uerr := json.Unmarshal(data, &tmp); uerr != nil {
			fmt.Println("✗ Failed to parse JSON:", uerr)
			return
		}
		sess = &Session{
			Name: tmp.Name, Width: tmp.Width, Height: tmp.Height,
			BaseFrame: Frame{Name: tmp.BaseFrame.Name, Lines: tmp.BaseFrame.Lines},
		}
		for _, st := range tmp.States {
			s := StateSession{Name: st.Name, Description: st.Name + " state", StateType: "standard", AnimationFPS: 5, AnimationLoops: 1}
			for i, f := range st.Frames {
				s.Frames = append(s.Frames, Frame{Name: fmt.Sprintf("%s_frame_%d", st.Name, i+1), Lines: f.Lines})
			}
			sess.States = append(sess.States, s)
		}
	}

	// Build temporary domain character and animate
	dStates := map[string]domain.State{}
	for _, st := range sess.States {
		dStates[st.Name] = domain.State{
			Name:           st.Name,
			Description:    st.Description,
			Frames:         convertFramesToDomain(st.Frames),
			StateType:      st.StateType,
			AnimationFPS:   st.AnimationFPS,
			AnimationLoops: st.AnimationLoops,
		}
	}
	tempChar := &domain.Character{
		Name:      sess.Name,
		Width:     sess.Width,
		Height:    sess.Height,
		BaseFrame: domain.Frame{Name: sess.BaseFrame.Name, Lines: sess.BaseFrame.Lines},
		States:    dStates,
	}

	agent := characters.NewAgentCharacter(tempChar)
	fmt.Printf("\nPreviewing '%s' state for %s (%dx%d) at %d FPS for %d loops\n\n", *state, sess.Name, sess.Width, sess.Height, *fps, *loops)
	if err := agent.AnimateState(os.Stdout, *state, *fps, *loops); err != nil {
		handleError("Animation failed", err)
		return
	}
	fmt.Println("\n✓ View complete!\n")
}
