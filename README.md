# Tangent

Terminal avatars for AI agents. Go library.

## Install

```bash
go get github.com/wildreason/tangent/pkg/characters
```

## Usage

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("sa")
agent.Plan(os.Stdout)
agent.Think(os.Stdout)
agent.Success(os.Stdout)
```

## API

```go
// Get agent
agent, err := characters.LibraryAgent(name)

// State methods
agent.Plan(writer)
agent.Think(writer)
agent.Execute(writer)
agent.Wait(writer)
agent.Error(writer)
agent.Success(writer)

// Custom states
agent.ShowState(writer, "arise")
agent.ShowState(writer, "approval")

// Introspection
states := agent.ListStates()
hasState := agent.HasState("think")

// Themes - set color theme for all agents
characters.SetTheme("latte")        // Switch to latte theme
themes := characters.ListThemes()   // ["bright", "cozy", "garden", "latte"]
current := characters.GetCurrentTheme()  // "latte"
```

## Avatars

7 characters × 17 states × 4 themes

### Characters

- **sa** - Shadja (musical note)
- **ri** - Rishabha
- **ga** - Gandhara
- **ma** - Madhyama
- **pa** - Panchama
- **da** - Dhaivata
- **ni** - Nishada

### Themes

**bright** (default) - Original bright colors, 100% saturation
- sa: #FF0000 | ri: #FF8800 | ga: #FFD700 | ma: #00FF00 | pa: #0088FF | da: #8800FF | ni: #FF0088

**latte** - Catppuccin-inspired warm pastels, GUI-friendly
- sa: #E78284 | ri: #EF9F76 | ga: #E5C890 | ma: #A6D189 | pa: #85C1DC | da: #CA9EE6 | ni: #F4B8E4

**garden** - Earthy natural colors, reduces terminal intimidation
- sa: #D4787D | ri: #D89C6A | ga: #C9B68C | ma: #8FB378 | pa: #7CA8B8 | da: #A888BA | ni: #C895A8

**cozy** - Modern GUI hybrid with professional warmth
- sa: #E18B8B | ri: #E5A679 | ga: #E6CC94 | ma: #99C794 | pa: #78AED4 | da: #B592D4 | ni: #DE99B8

## States

arise, wait, think, plan, execute, error, success, read, search, write, bash, build, communicate, block, blocked, resting, approval

## License

MIT © 2025 Wildreason, Inc
