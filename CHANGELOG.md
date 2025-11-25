# Changelog

## [0.1.1] - 2025-11-24

**Patch Release: Character Name Updates**

### Changed
- Renamed character "sa" to "sam" for consistency with Handwave integration
- Renamed character "ri" to "rio" for consistency with Handwave integration
- Updated all documentation to reflect new character names
- Updated all tests to use new character names

### Migration Guide
If you're using the old character names in your code:
```go
// Old
agent, _ := characters.LibraryAgent("sa")
agent, _ := characters.LibraryAgent("ri")

// New
agent, _ := characters.LibraryAgent("sam")
agent, _ := characters.LibraryAgent("rio")
```

## [0.1.0] - 2025-11-18

**First stable release**

Terminal avatars for AI agents. Go library providing 7 characters with 16 animated states and 4 color themes.

### Features

**Character Library**
- 7 musical note characters (sam, rio, ga, ma, pa, da, ni)
- 16 animated states per character
- 4 color themes (latte, bright, garden, cozy)
- 28 total combinations (7 × 4 themes)

**States**
- Core workflow: arise, wait, think, plan, execute, error, resting
- File operations: read, write, search
- Development: bash, build, communicate
- Blocking: block, blocked
- Approval: approval (confirmation/nodding animation)

**Theme System**
- `SetTheme(name)` - Switch color themes globally
- `GetCurrentTheme()` - Get active theme
- `ListThemes()` - List available themes
- Default: latte (Catppuccin-inspired pastels)

**API**
- `LibraryAgent(name)` - Load character by name
- State methods: `Plan()`, `Think()`, `Execute()`, `Wait()`, `Error()`
- Custom states: `ShowState(writer, stateName)`
- Introspection: `ListStates()`, `HasState(name)`

**Performance**
- Per-state FPS control
- Optimized frame animations
- Zero-allocation rendering where possible

### Installation

```bash
go get github.com/wildreason/tangent/pkg/characters
```

### License

MIT © 2025 Wildreason, Inc
