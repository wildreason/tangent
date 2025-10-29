# Tangent

Terminal avatars for AI agents

## Install

```bash
curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash
```

## Usage

```bash
tangent browse              # List all avatars
tangent browse sa         # Preview fire avatar
tangent browse sa --state resting --fps 5 --loop 5
```

## Integrate

```go
import "github.com/wildreason/tangent/pkg/characters"

agent, _ := characters.LibraryAgent("fire")
agent.Plan(os.Stdout)       // Agent is planning
agent.Think(os.Stdout)      // Agent is thinking
agent.Success(os.Stdout)    // Agent succeeded
```

## States

arise, wait, think, plan, execute, error, success, read, search, write, bash, build, communicate, block, blocked, resting

## Avatars

7 characters (sa, ri, ga, ma, pa, dha, ni) × 16 states × 11x4 dimensions

- **sa** - Pure Red (#FF0000)
- **ri** - Orange (#FF8800)
- **ga** - Gold (#FFD700)
- **ma** - Green (#00FF00)
- **pa** - Blue (#0088FF)
- **dha** - Purple (#8800FF)
- **ni** - Pink (#FF0088)

## For Contributors

### Add New States

```bash
# 1. Create state interactively
tangent create

# 2. Export template
tangent admin export sa -o template.json

# 3. Add your state to template.json

# 4. Batch register to all characters
tangent admin batch-register template.json colors.json --force
```

## License

MIT © 2025 Wildreason, Inc
