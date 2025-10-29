# Tangent Creator Guide

**For Advanced Users**: Avatar creation and contribution workflow.

This guide is for contributors who want to create custom avatars for AI agents. Most users should focus on browsing and integrating existing avatars (see main [README](../README.md)).

## Overview

Tangent provides a creation tool for building terminal avatars for AI agents. All avatars go through a curated review process to maintain quality and AI-native focus.

## Commands

### tangent create

Start the interactive character builder with live preview.

```bash
tangent create
```

**What it does:**
- Prompts for character name and dimensions
- Launches Bubbletea TUI with split-pane interface
- Left pane: Creation interface
- Right pane: Live preview
- Guides you through creating base and states

**Workflow:**
1. Enter character name (valid Go identifier)
2. Enter width and height
3. Create base character (idle state)
4. Add agent states (plan, think, execute minimum)
5. Preview animations
6. Export JSON for contribution

**Example:**
```bash
$ tangent create
╔══════════════════════════════════════════╗
║  CREATE NEW CHARACTER                    ║
╚══════════════════════════════════════════╝

◢ Character name: mercury
◢ Enter width (e.g., 11): 2
◢ Enter height (e.g., 3): 2

✓ Creating character 'mercury' (2x2)
✓ Character 'mercury' is starting

[TUI launches with live preview]
```

### tangent browse

List all available agents or view/animate a specific agent.

#### List All Agents

```bash
tangent browse
```

**Output:**
```
Available Agents:

  • demo4
  • mercury
  • water
  • water5

Total: 4 agents available

View agent: tangent browse <name>
```

#### View Specific Agent

```bash
tangent browse <name> [options]
```

**Options:**
- `--state <name>` - Animate specific state (plan|think|execute)
- `--fps <N>` - Override animation FPS
- `--loops <N>` - Override animation loops

**Examples:**
```bash
# View agent (shows base + all states)
tangent browse mercury

# View specific state
tangent browse mercury --state plan

# Custom animation settings
tangent browse mercury --state think --fps 10 --loops 3
```

**Output:**
```
Agent: mercury (2x2)

Base Character:
██
██

Animating 'plan' (3 frames) at 5 FPS for 1 loops
[Animation plays]

Animating 'think' (3 frames) at 5 FPS for 1 loops
[Animation plays]

Animating 'execute' (3 frames) at 5 FPS for 1 loops
[Animation plays]

✅ View complete!
```

### tangent view

Preview work-in-progress characters without admin registration.

#### View Session

```bash
tangent view --session <name> [options]
```

**Options:**
- `--state <name>` - Animate specific state (default: plan)
- `--fps <N>` - Frames per second (default: 5)
- `--loops <N>` - Number of loops (default: 1)

**Example:**
```bash
tangent view --session mercury --state plan --fps 8 --loops 2
```

#### View JSON File

```bash
tangent view --json <file> [options]
```

**Options:**
- Same as `--session` above

**Example:**
```bash
tangent view --json mercury.json --state think --fps 6 --loops 3
```

#### List Sessions

```bash
tangent view
```

Shows all saved sessions (work-in-progress characters).

**Output:**
```
Available Sessions:

  • demo-2
  • mercury
  • water
  • test-robot

Preview a session: tangent view --session <name> --state plan --fps 8 --loops 2
```

### tangent version

Show version information.

```bash
tangent version
```

**Output:**
```
tangent v0.1.0-alpha.4 (commit: abc1234, built: 2025-10-15T14:30:00Z)
```

### tangent help

Show usage information.

```bash
tangent help
```

Shows concise usage with commands, options, and examples.

## Common Workflows

### Creating a New Character

```bash
# Start creation
tangent create

# In TUI:
# 1. Create base character
# 2. Add agent state (plan)
# 3. Add agent state (think)
# 4. Add agent state (execute)
# 5. Animate all states (preview)
# 6. Export for contribution
#    → Creates yourcharacter.json
#    → Creates yourcharacter-README.md

# After export, submit via GitHub PR
```

### Testing a Character

```bash
# Quick preview
tangent browse yourcharacter

# Test specific state
tangent browse yourcharacter --state plan

# Test with custom settings
tangent browse yourcharacter --state execute --fps 10 --loops 5
```

### Previewing Work-in-Progress

```bash
# List your sessions
tangent view

# Preview a session
tangent view --session yourcharacter --state plan

# Preview exported JSON before submission
tangent view --json yourcharacter.json --state think
```

## Pattern Codes Reference

When creating characters in the TUI:

```
F = █  (Full Block)
T = ▀  (Top Half)
B = ▄  (Bottom Half)
L = ▌  (Left Half)
R = ▐  (Right Half)

1-8 = Quadrants (▘▝▖▗ ▛▜▙▟)
. = ░  (Light Shade)
: = ▒  (Medium Shade)
# = ▓  (Dark Shade)
_ =    (Space)
X = ◐  (Mirror)
```

See [Pattern Guide](PATTERNS.md) for complete reference.

## Tips and Tricks

### Efficient Creation

1. **Start Small**: Begin with 2x2 or 3x3, expand if needed
2. **Plan First**: Sketch your character on paper first
3. **Use Base**: Start states from base for consistency
4. **Preview Often**: Check animations frequently during creation

### Testing

1. **Test Both Themes**: Preview on light and dark terminal backgrounds
2. **Try Different FPS**: Test at various speeds (3, 5, 8, 10 FPS)
3. **Check Loops**: Ensure animations loop smoothly

### Debugging

1. **Session Recovery**: Sessions auto-save to `.tangent-sessions/`
2. **Pattern Errors**: Double-check pattern codes against reference
3. **Dimension Mismatch**: Ensure all lines match character width

## Output

All commands output to stdout, suitable for:
- Terminal display
- Piping to files
- Integration with other tools

**Example:**
```bash
# Save output to file
tangent browse mercury > mercury-output.txt

# Use in shell scripts
if tangent browse myagent; then
    echo "Agent exists"
fi
```

## Exit Codes

- `0` - Success
- `1` - Error (invalid command, character not found, etc.)

## Environment

### Sessions Directory

Work-in-progress characters are saved to:
```
$HOME/.tangent-sessions/
```

### Binary Location

After `make build`, the binary is at:
```
./tangent
```

Install globally:
```bash
sudo cp tangent /usr/local/bin/
```

## Integration Examples

### Shell Script

```bash
#!/bin/bash
# Show agent status in script

agent, _ := characters.LibraryAgent("mercury")

echo "Planning..."
tangent browse mercury --state plan --loops 1

echo "Executing..."
tangent browse mercury --state execute --loops 1

echo "Complete!"
```

### Makefile

```makefile
.PHONY: status
status:
	@tangent browse mercury --state think
	@echo "Build in progress..."
```

## See Also

- [API Reference](API.md) - Using characters in Go code
- [States Guide](STATES.md) - Designing agent states
- [Pattern Guide](PATTERNS.md) - Unicode block patterns
- [Contributing](CONTRIBUTING_CHARACTERS.md) - Submission workflow

## Support

For issues or questions:
- GitHub Issues: https://github.com/wildreason/tangent/issues
- Documentation: `docs/`
- Examples: `examples/`

---

**Tangent** - Terminal Avatars for AI Agents

