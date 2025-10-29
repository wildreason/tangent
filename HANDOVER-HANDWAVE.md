# Handover: Tangent alpha.17

**To**: handwave team
**From**: Tangent maintainers
**Date**: 2025-10-29
**Version**: alpha.17 (latest)

## What Changed

Tangent has transitioned to a library-first architecture. The CLI you may have used in alpha.16 is now an internal tool.

## Quick Migration

### Before (alpha.16)
```bash
tangent browse sa
tangent create
```

### After (alpha.17)
```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("sa")
agent.Think(os.Stdout)
agent.Plan(os.Stdout)
agent.Success(os.Stdout)
```

## Installation

**Old way** (alpha.16):
```bash
go install github.com/wildreason/tangent/cmd/tangent@latest
```

**New way** (alpha.17):
```bash
go get github.com/wildreason/tangent/pkg/characters
```

## What Still Works

* All Go library APIs (LibraryAgent, state methods)
* All 7 characters (sa, ri, ga, ma, pa, dha, ni)
* All 17 states (arise, wait, think, plan, execute, error, success, read, search, write, bash, build, communicate, block, blocked, resting_face, approval)
* Character colors and personalities
* Animation functionality

## What's Different

### No More Distributed CLI

The `tangent` binary is not distributed anymore. It's now `tangent-cli` (internal tool for Tangent contributors).

**Impact**: If you were using `tangent browse` or `tangent create`, switch to the Go library API.

### Library-First

Tangent is now purely a Go package. All functionality is available through `pkg/characters` APIs.

### State Registry System (New)

States are now JSON-first with automatic loading:

```go
import "github.com/wildreason/tangent/pkg/characters/stateregistry"

// Get specific state
state, ok := stateregistry.Get("arise")

// List all states
names := stateregistry.List()

// Get all states
all := stateregistry.All()
```

## Documentation

See these files for details:

* `MIGRATION-alpha.16-to-alpha.17.md` - Complete migration guide
* `README.md` - Updated library documentation
* `docs/STATE-REGISTRY.md` - State registry workflow (if contributing states)

## Support

For issues or questions:
* Open an issue at the repository
* Check README.md for usage examples
* Review MIGRATION guide for specific scenarios

## Testing Your Integration

Recommended test:

```go
package main

import (
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    agent, err := characters.LibraryAgent("sa")
    if err != nil {
        panic(err)
    }

    agent.Think(os.Stdout)
    agent.Plan(os.Stdout)
    agent.Execute(os.Stdout)
    agent.Success(os.Stdout)
}
```

Run:
```bash
go mod tidy
go run main.go
```

You should see the sa character animating through states.

## Questions

Any questions about this handover? Open an issue or reach out to the maintainers.

**Note**: alpha.16 is still available if you need time to migrate. Use `@v0.1.0-alpha.16` tag.
