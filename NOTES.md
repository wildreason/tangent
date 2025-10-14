# Tangent Workflow - Planet Series

## New Workflow (v0.1.0-alpha.3)

### 1. Create Character
- Character name (e.g., mercury, venus, earth)
- Enter width (e.g., 11)
- Enter height (e.g., 3)
- Choose personality (efficient, friendly, analytical, creative)

### 2. Create Base Character
- Design the base (idle) state
- This is the immutable foundation
- Line-by-line input with pattern codes
- Preview and confirm

### 3. Add Agent States
For each state (minimum 3 required):
- **Required states**: think, plan, search
- **Optional states**: wait, error, success
- **Custom states**: any name you choose

For each state:
- Choose number of animation frames (default: 3)
- For each frame:
  - Option to start from base (copy base lines)
  - Or create from scratch
  - Line-by-line input with live preview
  - Progressive building with reference to base

### 4. Preview & Animate
- Preview base character
- Preview individual state animations
- Animate all states in sequence
- Adjust animation speed (FPS) per state

### 5. Export
- Export for contribution (JSON format)
- Validates minimum 3 required states
- Generates contribution README
- Ready for GitHub PR submission

## Menu Structure

```
▢ CHARACTER: mercury
  Base: ✓ Created | States: 3 (think, plan, search)

  1. Create base character
  2. Add agent state
  3. Edit agent state
  4. Preview base character
  5. Preview state animation
  6. Animate all states
  7. Export for contribution (JSON)
  8. Back to main menu
  9. Exit
```

## Key Features

### Base Character
- Immutable foundation for all states
- Always shown when agent is idle
- Reference point for creating state animations

### Agent States
- Multiple animation frames per state
- Configurable FPS and loops per state
- Start from base or create from scratch
- Live preview during creation

### State Types
- **Standard**: think, plan, search, wait, error, success
- **Custom**: any name for unique behaviors

### Animation
- Smooth frame transitions (fixed animation engine)
- Per-state FPS control
- Configurable loop count
- Preview before exporting

## Pattern Codes

```
Basic Blocks:
  F = █  T = ▀  B = ▄  L = ▌  R = ▐

Shading:
  . = ░  : = ▒  # = ▓

Quadrants (1-4):
  1 = ▘  2 = ▝  3 = ▖  4 = ▗

Three-Quadrants (5-8):
  5 = ▛  6 = ▜  7 = ▙  8 = ▟

Special:
  _ = Space
  X = Mirror marker
```

## Planet Series Characters

Target: 8 characters
- Mercury
- Venus
- Earth
- Mars
- Jupiter
- Saturn
- Uranus
- Neptune

Each character:
- Unique base (idle) design
- Same 3 required states (think, plan, search)
- Optional additional states
- Distinct personality and visual style

## Example: Creating Mercury

1. **Create Character**
   ```
   Name: mercury
   Width: 11
   Height: 3
   Personality: efficient
   ```

2. **Create Base**
   ```
   Design Mercury's idle face
   (immutable foundation)
   ```

3. **Add Think State**
   ```
   3 animation frames
   Start from base: yes
   Modify lines to show thinking
   ```

4. **Add Plan State**
   ```
   3 animation frames
   Start from base: yes
   Modify lines to show planning
   ```

5. **Add Search State**
   ```
   3 animation frames
   Start from base: yes
   Modify lines to show searching
   ```

6. **Preview & Export**
   ```
   Preview all states
   Export for contribution
   Submit GitHub PR
   ```

## Technical Details

### Session Structure
```json
{
  "name": "mercury",
  "personality": "efficient",
  "width": 11,
  "height": 3,
  "base_frame": {
    "name": "base",
    "lines": ["...", "...", "..."]
  },
  "states": [
    {
      "name": "think",
      "description": "Agent thinking state",
      "state_type": "standard",
      "frames": [...],
      "animation_fps": 5,
      "animation_loops": 1
    }
  ]
}
```

### Domain Model
- `Character.BaseFrame` - Idle/immutable base
- `Character.States` - Map of agent states
- `State.Frames` - Multiple animation frames
- `State.AnimationFPS` - Per-state FPS
- `State.AnimationLoops` - Per-state loop count

### API
- `agent.ShowBase()` - Display base character
- `agent.AnimateState(name, fps, loops)` - Animate specific state
- `agent.Think()`, `agent.Plan()`, `agent.Search()` - Standard states

## Migration from Old Workflow

**Old (alpha.2)**:
- Generic frame creation
- No base character concept
- Frames treated equally
- Manual duplication for variations

**New (alpha.3)**:
- Base character first
- States built FROM base
- Structured state creation
- Copy-from-base option
- Reduced friction

## Backward Compatibility

- Old session files still work
- Legacy `Frames []Frame` preserved
- New workflow only for new characters
- Export includes both structures

## Building

For development builds with proper version tracking:

```bash
make build
# or
./scripts/build.sh
```

Version will show: `v0.1.0-alpha.2-3-gabc1234-dirty`

- `v0.1.0-alpha.2` - Last git tag
- `3` - Commits ahead of tag
- `gabc1234` - Commit hash
- `dirty` - Uncommitted changes present

This provides complete traceability:
- Which release we're building towards
- How many commits ahead
- Exact commit hash
- Whether there are uncommitted changes

### Quick Build & Test

```bash
# Clean build
make clean
make build

# Check version
./tangent version

# Should show: tangent v0.1.0-alpha.2-N-g<hash>[-dirty] (commit: <hash>, built: <date>)
```

## Release Process

### Development Builds
- Use `make build` for local testing
- Version shows commit count and dirty status
- Example: `v0.1.1-11-gc80a59a-dirty`

### Production Releases
- Use `make release` to create tagged release
- GitHub Actions automatically builds and publishes
- Version shows clean tag only
- Example: `v0.1.2`

**See [RELEASE.md](RELEASE.md) for complete release guide.**
