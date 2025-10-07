package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Session represents a character building session
type Session struct {
	Name      string    `json:"name"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Frames    []Frame   `json:"frames"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Frame represents a single animation frame
type Frame struct {
	Name  string   `json:"name"`
	Lines []string `json:"lines"`
}

const sessionDir = ".characters"

// SaveSession saves the session to disk
func (s *Session) Save() error {
	// Create session directory if it doesn't exist
	if err := os.MkdirAll(sessionDir, 0755); err != nil {
		return fmt.Errorf("failed to create session directory: %w", err)
	}

	s.UpdatedAt = time.Now()

	// Marshal to JSON
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	// Write to file
	filename := filepath.Join(sessionDir, s.Name+".json")
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write session file: %w", err)
	}

	return nil
}

// LoadSession loads a session from disk
func LoadSession(name string) (*Session, error) {
	filename := filepath.Join(sessionDir, name+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read session file: %w", err)
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

// ListSessions returns all saved session names
func ListSessions() ([]string, error) {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(sessionDir, 0755); err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(sessionDir)
	if err != nil {
		return nil, err
	}

	var sessions []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			name := entry.Name()[:len(entry.Name())-5] // Remove .json
			sessions = append(sessions, name)
		}
	}

	return sessions, nil
}

// DeleteSession deletes a session file
func DeleteSession(name string) error {
	filename := filepath.Join(sessionDir, name+".json")
	return os.Remove(filename)
}

// SessionExists checks if a session file exists
func SessionExists(name string) bool {
	filename := filepath.Join(sessionDir, name+".json")
	_, err := os.Stat(filename)
	return err == nil
}

// NewSession creates a new session
func NewSession(name string, width, height int) *Session {
	return &Session{
		Name:      name,
		Width:     width,
		Height:    height,
		Frames:    []Frame{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// AddFrame adds a frame to the session
func (s *Session) AddFrame(frame Frame) {
	s.Frames = append(s.Frames, frame)
	s.UpdatedAt = time.Now()
}

// UpdateFrame updates a frame at the given index
func (s *Session) UpdateFrame(index int, frame Frame) error {
	if index < 0 || index >= len(s.Frames) {
		return fmt.Errorf("frame index out of range")
	}
	s.Frames[index] = frame
	s.UpdatedAt = time.Now()
	return nil
}

// DeleteFrame deletes a frame at the given index
func (s *Session) DeleteFrame(index int) error {
	if index < 0 || index >= len(s.Frames) {
		return fmt.Errorf("frame index out of range")
	}
	s.Frames = append(s.Frames[:index], s.Frames[index+1:]...)
	s.UpdatedAt = time.Time{}
	return nil
}

// GetFrame returns a frame at the given index
func (s *Session) GetFrame(index int) (*Frame, error) {
	if index < 0 || index >= len(s.Frames) {
		return nil, fmt.Errorf("frame index out of range")
	}
	return &s.Frames[index], nil
}
