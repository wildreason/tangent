# Tangent

Terminal avatars for AI agents. Go library only.

## Install

```bash
go get github.com/wildreason/tangent/pkg/characters
```

## Usage

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("sa")
agent.Plan(os.Stdout)       // Agent is planning
agent.Think(os.Stdout)      // Agent is thinking
agent.Success(os.Stdout)    // Agent succeeded
```

## States

arise, wait, think, plan, execute, error, success, read, search, write, bash, build, communicate, block, blocked, resting, approval

## Avatars

7 characters (sa, ri, ga, ma, pa, dha, ni) × 17 states × 11x4 dimensions

- **sa** - Pure Red (#FF0000)
- **ri** - Orange (#FF8800)
- **ga** - Gold (#FFD700)
- **ma** - Green (#00FF00)
- **pa** - Blue (#0088FF)
- **dha** - Purple (#8800FF)
- **ni** - Pink (#FF0088)

## For Contributors

States are defined in `pkg/characters/stateregistry/states/*.json`

To add a new state:
1. Create JSON file in `pkg/characters/stateregistry/states/`
2. Commit the JSON file
3. State registry automatically loads it

Internal tools: `tangent-cli` (not distributed)

## License

MIT © 2025 Wildreason, Inc
