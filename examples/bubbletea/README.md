# Bubble Tea Integration Example

This example demonstrates how to use Tangent characters with Bubble Tea.

## What This Shows

- Loading Tangent library characters (`wave`, `pulse`, `rocket`)
- Creating Bubble Tea spinners using the `bubbletea` adapter
- Styling with Lip Gloss
- Running multiple animated characters simultaneously
- Full event loop control with Bubble Tea

## Architecture

```
┌─────────────────────────────────────────┐
│ Bubble Tea (Event Loop & Rendering)     │
├─────────────────────────────────────────┤
│ Lip Gloss (Styling & Layout)            │
├─────────────────────────────────────────┤
│ Tangent Adapter (Frame Extraction)      │
├─────────────────────────────────────────┤
│ Tangent Characters (Frame Design)       │
└─────────────────────────────────────────┘
```

## Running

```bash
cd examples/bubbletea
go run main.go
```

## Key Integration Points

### 1. Load Library Character and Create Spinner

```go
spinner, err := bubbletea.LibrarySpinner("wave", 6)
```

This single line:
- Loads the "wave" character from Tangent's library
- Extracts frames
- Creates a Bubble Tea spinner at 6 FPS

### 2. Custom Character Integration

```go
// Load custom character
char, _ := characters.Library("mychar")

// Create spinner with adapter
s := bubbletea.SpinnerFromCharacter(char, 5)
```

### 3. Style with Lip Gloss

```go
style := lipgloss.NewStyle().
    Foreground(lipgloss.Color("#00FF00")).
    Border(lipgloss.RoundedBorder()).
    Padding(0, 2)

view := style.Render(spinner.View())
```

## Why This Works

- **Tangent** provides character frames (shapes)
- **Bubble Tea** handles event loop and rendering
- **Lip Gloss** handles styling and layout
- **No conflicts** - each layer has clear responsibilities

## What Not To Do

❌ Don't use `characters.Animate()` in Bubble Tea apps:

```go
// DON'T DO THIS
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    char, _ := characters.Library("alien")
    characters.Animate(os.Stdout, char, 5, 3)  // ❌ Conflicts!
    return m, nil
}
```

✅ Instead, extract frames and use Bubble Tea's render loop:

```go
// DO THIS
func initialModel() model {
    s, _ := bubbletea.LibrarySpinner("alien", 5)
    return model{spinner: s}
}

func (m model) View() string {
    return m.spinner.View()  // ✅ Bubble Tea controls rendering
}
```

## See Also

- [Bubble Tea Integration Guide](../../docs/BUBBLETEA_INTEGRATION.md)
- [Tangent Library Characters](../../docs/LIBRARY.md)
- [Pattern Guide](../../docs/PATTERN_GUIDE.md)

