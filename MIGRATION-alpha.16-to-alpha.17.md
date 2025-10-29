# Migration Guide: alpha.16 → alpha.17

**Target audience**: handwave team currently using alpha.16

## Breaking Changes

### 1. CLI Tool Moved (Not Distributed)

**alpha.16**: `tangent` binary available for character creation
**alpha.17**: CLI moved to `tangent-cli` (internal tool, not distributed)

**Impact**: If you were using `tangent browse` or `tangent create`, these are no longer available in the public package.

**Migration**:
- Use Go library API instead of CLI
- See "Usage" section below

### 2. Library-First Architecture

**alpha.16**: CLI-focused with library as secondary
**alpha.17**: Pure Go library, no distributed binaries

**What changed**:
- No more CLI installation instructions
- README is library-first
- Public API unchanged

## No Changes Required For

* **Go library usage** - All `pkg/characters` APIs remain identical
* **LibraryAgent()** function - Works exactly the same
* **State methods** - Plan(), Think(), Success(), etc. unchanged
* **Character names** - Still 7 characters (sa, ri, ga, ma, pa, dha, ni)
* **States** - All 17 states still available

## New in alpha.17

### State Registry System

States are now JSON-first with embedded loader:
- Source of truth: `pkg/characters/stateregistry/states/*.json`
- Automatic loading via `go:embed`
- No manual registration needed

**New package**: `pkg/characters/stateregistry`

```go
import "github.com/wildreason/tangent/pkg/characters/stateregistry"

// Get specific state
state, ok := stateregistry.Get("arise")

// List all states
names := stateregistry.List()

// Get all states
all := stateregistry.All()
```

## Usage Examples

### Before (alpha.16)

```bash
# CLI usage
tangent browse sa --state think
tangent create
```

### After (alpha.17)

```go
// Go library only
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("sa")
agent.Think(os.Stdout)
```

## Installation

**alpha.16**:
```bash
go install github.com/wildreason/tangent/cmd/tangent@latest
```

**alpha.17**:
```bash
go get github.com/wildreason/tangent/pkg/characters
```

## Questions

For issues or questions about this migration, open an issue at the repository.

## Internal Contributors

If you're contributing to Tangent (not just using it), see:
- `tangent-cli` documentation (internal tool)
- State registry workflow documentation
