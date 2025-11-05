# State Registry Workflow

How to add new states to Tangent.

## Architecture

**Source of truth**: JSON files in `pkg/characters/stateregistry/states/`

**Loader**: Automatic via `go:embed` on package init

**No registration needed**: Just commit the JSON file

## Adding a New State

### 1. Create JSON file

File: `pkg/characters/stateregistry/states/{state_name}.json`

Format:
```json
{
  "name": "state_name",
  "frames": [
    {
      "lines": [
        "line1 (11 chars)",
        "line2 (11 chars)",
        "line3 (11 chars)",
        "line4 (11 chars)"
      ]
    },
    {
      "lines": [
        "next frame..."
      ]
    }
  ]
}
```

### 2. Pattern codes

Use these ASCII codes (compiled to Unicode blocks):

* `F` - Full block (face center)
* `R` - Right half
* `L` - Left half
* `T` - Top half
* `B` - Bottom half
* `1-8` - Eighth blocks (gradients)
* `_` - Empty space
* `r/l` - Small corner blocks
* `f` - Shading

### 3. Commit

```bash
git add pkg/characters/stateregistry/states/{state_name}.json
git commit -m "feat: Add {state_name} state"
```

Done. State registry automatically loads it on next import.

## Example: approval state

File: `pkg/characters/stateregistry/states/approval.json`

```json
{
  "name": "approval",
  "frames": [
    {
      "lines": [
        "l_r5fff6l__",
        "rbffffff51_",
        "___11_22___",
        "           "
      ]
    },
    {
      "lines": [
        "___________",
        "l_r5fff6l__",
        "rbffffff51_",
        "___11_22___"
      ]
    }
  ]
}
```

This creates a nodding animation (8 frames total).

## Using States

### Go library

```go
import "github.com/wildreason/tangent/pkg/characters/stateregistry"

// Get state
state, ok := stateregistry.Get("approval")

// List all states
names := stateregistry.List() // ["arise", "wait", "think", ...]

// Get all states
all := stateregistry.All()
```

### Internal tool

```bash
tangent-cli browse sa --state approval
```

## Current States

16 states available:

* arise, wait, think, plan, execute
* error
* read, search, write, bash, build, communicate
* block, blocked
* resting_face, approval

## Adding State to All Characters

**Before**: Used `tangent-cli admin batch-register` with template + colors

**After**: States are global (not character-specific)

All 7 characters automatically get new states when registry loads.

## Implementation Details

**Loader**: `pkg/characters/stateregistry/loader.go`

```go
//go:embed states
var statesFS embed.FS

func init() {
    DefaultRegistry, _ = LoadEmbedded()
}

func LoadEmbedded() (*Registry, error) {
    registry := NewRegistry()
    entries, _ := statesFS.ReadDir("states")

    for _, entry := range entries {
        if !strings.HasSuffix(entry.Name(), ".json") {
            continue
        }

        data, _ := statesFS.ReadFile(filepath.Join("states", entry.Name()))
        var state StateDefinition
        json.Unmarshal(data, &state)
        registry.Register(state)
    }

    return registry, nil
}
```

**Tests**: `pkg/characters/stateregistry/loader_test.go`

Run: `go test ./pkg/characters/stateregistry/...`
