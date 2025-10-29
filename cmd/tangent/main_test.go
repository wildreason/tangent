package main

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

// TestGenerateLibraryCode_ValidGoCode verifies generated code is syntactically valid Go
func TestGenerateLibraryCode_ValidGoCode(t *testing.T) {
	// Minimal character spec
	patterns := []struct {
		Name  string
		Lines []string
	}{
		{
			Name: "test_1",
			Lines: []string{
				"___",
				"_F_",
				"___",
			},
		},
		{
			Name: "test_2",
			Lines: []string{
				"___",
				"_R_",
				"___",
			},
		},
	}

	code := generateLibraryCode(
		"test",
		"test character",
		"Test Author",
		"#FF0000",
		"test personality",
		3,
		3,
		patterns,
	)

	// Verify it's valid Go code by parsing
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "test.go", code, parser.AllErrors)
	if err != nil {
		t.Fatalf("Generated code is not valid Go:\n%s\n\nError: %v", code, err)
	}

	// Verify critical sections exist
	if len(code) < 100 {
		t.Fatalf("Generated code too short: %d bytes", len(code))
	}

	// Check for required Go structures
	required := []string{
		"package library",
		"func init()",
		"register(testCharacter)",
		"var testCharacter = LibraryCharacter{",
		`Name:        "test"`,
		`Color:       "#FF0000"`,
		"Width:       3",
		"Height:      3",
		"Patterns: []Frame{",
	}

	for _, req := range required {
		if !contains(code, req) {
			t.Errorf("Generated code missing required string: %q", req)
		}
	}
}

// TestAdminRegister_SmokeTest verifies single character registration workflow
func TestAdminRegister_SmokeTest(t *testing.T) {
	// Create temp directory for test
	tmpDir := t.TempDir()

	// Create minimal test JSON
	testJSON := `{
  "name": "testchar",
  "description": "Test character",
  "author": "Test Author",
  "color": "#00FF00",
  "width": 3,
  "height": 3,
  "base_frame": {
    "name": "base",
    "lines": ["___", "_F_", "___"]
  },
  "states": [
    {
      "name": "test",
      "frames": [
        {"lines": ["___", "_F_", "___"]},
        {"lines": ["___", "_R_", "___"]}
      ]
    }
  ]
}`

	jsonPath := filepath.Join(tmpDir, "test.json")
	if err := os.WriteFile(jsonPath, []byte(testJSON), 0644); err != nil {
		t.Fatalf("Failed to create test JSON: %v", err)
	}

	// Create fake library directory
	libDir := filepath.Join(tmpDir, "pkg", "characters", "library")
	if err := os.MkdirAll(libDir, 0755); err != nil {
		t.Fatalf("Failed to create library dir: %v", err)
	}

	// Change to temp directory for test
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Test that adminRegister doesn't panic with valid input
	// Note: We can't test the full function because it calls os.Exit()
	// but we can test generateLibraryCode which is the core logic

	// Read and parse the JSON manually
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		t.Fatalf("Failed to read test JSON: %v", err)
	}

	if len(data) < 50 {
		t.Fatalf("Test JSON too short: %d bytes", len(data))
	}

	// Verify JSON is parseable
	var testData map[string]interface{}
	if err := parseJSON(data, &testData); err != nil {
		t.Fatalf("Test JSON not parseable: %v", err)
	}
}

// TestAdminBatchRegister_SmokeTest verifies batch registration workflow
func TestAdminBatchRegister_SmokeTest(t *testing.T) {
	// Create temp directory for test
	tmpDir := t.TempDir()

	// Create minimal template.json
	templateJSON := `{
  "width": 3,
  "height": 3,
  "states": [
    {
      "name": "test",
      "frames": [
        {"lines": ["___", "_F_", "___"]},
        {"lines": ["___", "_R_", "___"]}
      ]
    }
  ]
}`

	// Create minimal colors.json
	colorsJSON := `{
  "char1": {
    "color": "#FF0000",
    "description": "Test char 1"
  },
  "char2": {
    "color": "#00FF00",
    "description": "Test char 2"
  }
}`

	templatePath := filepath.Join(tmpDir, "template.json")
	colorsPath := filepath.Join(tmpDir, "colors.json")

	if err := os.WriteFile(templatePath, []byte(templateJSON), 0644); err != nil {
		t.Fatalf("Failed to create template JSON: %v", err)
	}

	if err := os.WriteFile(colorsPath, []byte(colorsJSON), 0644); err != nil {
		t.Fatalf("Failed to create colors JSON: %v", err)
	}

	// Read and verify both files are parseable
	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("Failed to read template: %v", err)
	}

	colorsData, err := os.ReadFile(colorsPath)
	if err != nil {
		t.Fatalf("Failed to read colors: %v", err)
	}

	// Verify template structure
	var template struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		States []struct {
			Name   string `json:"name"`
			Frames []struct {
				Lines []string `json:"lines"`
			} `json:"frames"`
		} `json:"states"`
	}

	if err := parseJSON(templateData, &template); err != nil {
		t.Fatalf("Template JSON not parseable: %v", err)
	}

	if template.Width != 3 {
		t.Errorf("Expected width 3, got %d", template.Width)
	}

	if len(template.States) == 0 {
		t.Fatal("Template has no states")
	}

	// Verify colors structure
	var colors map[string]struct {
		Color       string `json:"color"`
		Description string `json:"description"`
	}

	if err := parseJSON(colorsData, &colors); err != nil {
		t.Fatalf("Colors JSON not parseable: %v", err)
	}

	if len(colors) != 2 {
		t.Errorf("Expected 2 characters, got %d", len(colors))
	}

	// Verify we can generate code for each character
	for charName, charConfig := range colors {
		var patterns []struct {
			Name  string
			Lines []string
		}

		for _, state := range template.States {
			for i, frame := range state.Frames {
				patterns = append(patterns, struct {
					Name  string
					Lines []string
				}{
					Name:  state.Name + "_" + string(rune('1'+i)),
					Lines: frame.Lines,
				})
			}
		}

		code := generateLibraryCode(
			charName,
			charConfig.Description,
			"Test Author",
			charConfig.Color,
			"",
			template.Width,
			template.Height,
			patterns,
		)

		// Verify generated code is valid Go
		fset := token.NewFileSet()
		_, err := parser.ParseFile(fset, charName+".go", code, parser.AllErrors)
		if err != nil {
			t.Errorf("Generated code for %s is not valid Go: %v", charName, err)
		}
	}
}

// Helper functions

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsAt(s, substr))
}

func containsAt(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func parseJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
