# Design Discovery System

**Rapid animation design exploration through automated mutation and human selection**

## Overview

Instead of manually designing animations one at a time, this system:
1. **Mutates** existing designs randomly
2. **Builds** the changes
3. **Previews** the animation
4. **Asks** for approval/decline
5. **Repeats** continuously

This enables "theory of discovery" - exploring the design space rapidly to find great variations.

## Quick Start

```bash
# From project root
cd /Users/btsznh/.cursor/worktrees/characters/VjDfZ

# Basic usage - explore 'search' state
python scripts/discover.py --state search --iterations 10

# Different agent and state
python scripts/discover.py --agent ga --state wait --iterations 20

# Quick exploration
python scripts/discover.py --state search --iterations 5
```

## How It Works

### Mutation Types

The system randomly applies 1-2 mutations per iteration:

1. **Frame Removal** - Reduces frame count (keeps minimum 2)
2. **Line Modification** - Changes individual lines to random patterns
3. **Frame Duplication** - Duplicates existing frames
4. **Frame Shuffling** - Reorders frames
5. **Multi-Line Changes** - Modifies 2-3 lines at once

### Workflow

```
┌─────────────────────────────────────────┐
│  Load Current State                     │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│  Apply Random Mutations                 │
│  (1-2 mutations per iteration)          │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│  Build CLI (make build-cli)             │
│  If fails → revert, continue            │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│  Preview Animation                      │
│  (3 FPS, 2 loops)                       │
└──────────┬──────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────────┐
│  Human Decision:                        │
│  [A]pprove - Keep as base for next      │
│  [D]ecline - Revert to previous         │
│  [Q]uit    - Exit and restore original  │
└──────────┬──────────────────────────────┘
           │
           ▼
        Repeat
```

## Command Options

```bash
python scripts/discover.py [OPTIONS]

Options:
  --agent AGENT           Agent character (default: ga)
  --state STATE          State to explore (default: search)
  --iterations N         Number of iterations (default: 10)
  -h, --help            Show help message
```

## Examples

### Explore Different States

```bash
# Search animation
python scripts/discover.py --state search --iterations 15

# Wait animation  
python scripts/discover.py --state wait --iterations 10

# Write animation
python scripts/discover.py --state write --iterations 20

# Read animation
python scripts/discover.py --state read --iterations 10
```

### Session Examples

```bash
# Quick 5-iteration exploration
python scripts/discover.py --state search --iterations 5

# Deep 50-iteration exploration
python scripts/discover.py --state search --iterations 50

# Different agent
python scripts/discover.py --agent ma --state search --iterations 10
```

## Output

### During Session

```
╔════════════════════════════════════════════════════════════╗
║  🔍 DESIGN DISCOVERY SYSTEM                                ║
╚════════════════════════════════════════════════════════════╝
Agent: ga
State: search
Iterations: 10

💾 Backup created: search.backup

────────────────────────────────────────────────────────────
🎨 ITERATION 1/10
────────────────────────────────────────────────────────────
🔧 Generating variation...
   Applying 2 mutation(s):
   - remove random frames
      Frames: 6 → 3
   - modify line
      Line 1 in frame 0: modified

🔨 Building...
✅ Build successful

📺 Preview:
────────────────────────────────────────────────────────────
[Animation plays here]
────────────────────────────────────────────────────────────

[A]pprove  [D]ecline  [Q]uit: A
✅ APPROVED! Keeping this design as base for next iteration
```

### After Session

```
════════════════════════════════════════════════════════════
📊 DISCOVERY SESSION COMPLETE
════════════════════════════════════════════════════════════
Total iterations: 10
✅ Approved designs: 3
❌ Declined designs: 7

💡 Your current state contains the last approved design.
📁 Original backup: search.backup
💾 Saved approved designs to: testing/discovery/search_20251114_143022_approved.json
```

## Safety Features

- **Automatic Backup** - Original state saved as `.backup`
- **Build Validation** - Failed builds are automatically reverted
- **Keyboard Interrupt** - Ctrl+C safely restores original
- **Error Handling** - Any errors trigger restoration
- **Approved History** - All approved designs saved to `testing/discovery/`

## Tips

1. **Start Small** - Begin with 5-10 iterations to understand the system
2. **Approve Freely** - Don't overthink; if it looks interesting, approve it
3. **Use as Base** - Each approval becomes the foundation for the next mutation
4. **Compound Evolution** - Great designs emerge from multiple approved iterations
5. **Explore Boldly** - The system is safe; you can always quit and restore

## Restoring Original

If you want to revert to your original design:

```bash
# The backup file is created automatically
cp pkg/characters/stateregistry/states/search.json.backup \
   pkg/characters/stateregistry/states/search.json

# Then rebuild
make build-cli
```

## Review Approved Designs

All approved designs are saved to `testing/discovery/` with timestamps:

```bash
# View approved designs
cat testing/discovery/search_20251114_143022_approved.json

# Extract a specific iteration
jq '.[] | select(.iteration == 3) | .data' \
   testing/discovery/search_20251114_143022_approved.json \
   > search_v3.json
```

## Philosophy: Discovery over Design

Traditional approach:
- Think → Design → Implement → Test
- Linear, slow, limited by imagination

Discovery approach:
- Generate → Build → Preview → Select → Repeat
- Exponential exploration of possibility space
- Discover designs you wouldn't have imagined

## Troubleshooting

### Build Failures

If builds repeatedly fail:
```bash
# Check if base state is valid
make build-cli

# If successful, run discovery with fewer iterations
python scripts/discover.py --state search --iterations 3
```

### State Not Found

```bash
# List available states
ls pkg/characters/stateregistry/states/*.json

# Use correct state name (without .json)
python scripts/discover.py --state [state-name]
```

### Python Not Found

```bash
# Use Python 3 explicitly
python3 scripts/discover.py --state search
```

## Future Enhancements

Ideas for extending the system:

- [ ] Constraint-based mutations (e.g., "only reduce frames")
- [ ] Side-by-side comparison view
- [ ] Undo/redo stack
- [ ] Export to GIF
- [ ] ML-guided mutations based on approvals
- [ ] Batch approval mode
- [ ] Design similarity detection

## License

Same as parent project.

