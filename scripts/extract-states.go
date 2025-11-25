package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Frame struct {
	Lines []string `json:"lines"`
}

type State struct {
	Name   string  `json:"name"`
	Frames []Frame `json:"frames"`
}

func main() {
	// Read sam.go file
	content, err := os.ReadFile("pkg/characters/library/sam.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	text := string(content)

	// Extract all pattern frames
	framePattern := regexp.MustCompile(`\{\s*Name:\s*"([^"]+)",\s*Lines:\s*\[\]string\{([^}]+)\}`)
	matches := framePattern.FindAllStringSubmatch(text, -1)

	// Group frames by state
	states := make(map[string][]Frame)

	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		frameName := match[1]
		linesText := match[2]

		// Extract state name (everything before last underscore)
		parts := strings.Split(frameName, "_")
		if len(parts) < 2 {
			continue
		}
		stateName := strings.Join(parts[:len(parts)-1], "_")

		// Parse lines
		linePattern := regexp.MustCompile(`"([^"]*)"`)
		lineMatches := linePattern.FindAllStringSubmatch(linesText, -1)

		var lines []string
		for _, lineMatch := range lineMatches {
			if len(lineMatch) >= 2 {
				lines = append(lines, lineMatch[1])
			}
		}

		if len(lines) > 0 {
			states[stateName] = append(states[stateName], Frame{Lines: lines})
		}
	}

	// Create output directory
	os.MkdirAll("pkg/characters/states", 0755)

	// Write each state to JSON file
	for stateName, frames := range states {
		state := State{
			Name:   stateName,
			Frames: frames,
		}

		data, err := json.MarshalIndent(state, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling %s: %v\n", stateName, err)
			continue
		}

		filename := fmt.Sprintf("pkg/characters/states/%s.json", stateName)
		err = os.WriteFile(filename, data, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✓ Extracted %s (%d frames)\n", stateName, len(frames))
	}

	fmt.Printf("\n✅ Extracted %d states to pkg/characters/states/\n", len(states))
}
