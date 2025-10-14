package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const sessionDir = ".tangent"

// Frame represents a single animation frame.
type Frame struct {
	Name      string   `json:"name"`
	Lines     []string `json:"lines"`
	StateType string   `json:"state_type,omitempty"` // For backward compatibility
}

// StateSession represents an agent state with animation frames
type StateSession struct {
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	StateType      string  `json:"state_type"`
	Frames         []Frame `json:"frames"`
	AnimationFPS   int     `json:"animation_fps"`
	AnimationLoops int     `json:"animation_loops"`
}

// Session represents a character project.
type Session struct {
	Name        string         `json:"name"`
	Personality string         `json:"personality,omitempty"`
	Width       int            `json:"width"`
	Height      int            `json:"height"`
	BaseFrame   Frame          `json:"base_frame"`
	States      []StateSession `json:"states"`
	Frames      []Frame        `json:"frames"` // Deprecated, keep for backward compatibility
}

// NewSession creates a new session.
func NewSession(name string, width, height int) *Session {
	return &Session{
		Name:   name,
		Width:  width,
		Height: height,
		States: []StateSession{},
		Frames: []Frame{}, // Keep for backward compatibility
	}
}

// Save persists the session to disk.
func (s *Session) Save() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	dir := filepath.Join(homeDir, sessionDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create session directory: %w", err)
	}

	filePath := filepath.Join(dir, s.Name+".json")

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write session file: %w", err)
	}

	return nil
}

// LoadSession loads a session from disk.
func LoadSession(name string) (*Session, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	filePath := filepath.Join(homeDir, sessionDir, name+".json")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read session file: %w", err)
	}

	var s Session
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &s, nil
}

// ListSessions lists all saved sessions.
func ListSessions() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	dir := filepath.Join(homeDir, sessionDir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("failed to read session directory: %w", err)
	}

	var names []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			names = append(names, strings.TrimSuffix(file.Name(), ".json"))
		}
	}

	return names, nil
}

// DeleteSession deletes a saved session.
func DeleteSession(name string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	filePath := filepath.Join(homeDir, sessionDir, name+".json")
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete session file: %w", err)
	}

	return nil
}
