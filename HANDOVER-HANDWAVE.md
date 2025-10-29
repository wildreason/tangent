# Handover: Tangent v0.1.0-beta.0

**To**: handwave team
**From**: Tangent maintainers
**Date**: 2025-10-29

## TL;DR

**No code changes required.** Just update your dependency version.

```bash
go get github.com/wildreason/tangent@v0.1.0-beta.0
go mod tidy
```

## What's Unchanged

Your existing code continues to work exactly the same:

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("sa")
agent.Think(os.Stdout)
agent.Plan(os.Stdout)
agent.Success(os.Stdout)
```

All APIs are **100% backward compatible**:
* `characters.LibraryAgent(name)` - Same
* `agent.Think()`, `agent.Plan()`, `agent.Execute()` - Same
* All 7 characters (sa, ri, ga, ma, pa, dha, ni) - Same
* All 17 states - Same
* Character colors - Same

## What Changed Internally

These changes don't affect your code, but FYI:

1. **Character generation**: Characters now auto-generate from a state registry (JSON files)
2. **New states auto-included**: When we add new states to the registry, all 7 characters automatically get them
3. **CLI moved**: The `tangent` CLI is now `tangent-cli` (internal tool, not distributed)

## Testing Your Integration

Run your existing tests. Everything should pass without changes.

If you want to verify manually:

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
    agent.Success(os.Stdout)
}
```

```bash
go mod tidy
go run main.go
```

You should see the sa character animating through states (same as alpha.16).

## Questions?

Open an issue at the repository if you encounter any problems.

## Rollback Option

If needed, you can stay on alpha.16:
```bash
go get github.com/wildreason/tangent@v0.1.0-alpha.16
```
