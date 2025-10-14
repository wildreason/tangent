# Contributing Characters to Tangent

Thank you for your interest in contributing characters to the Tangent character library! This guide will help you create and submit your character.

## Quick Start

1. **Create your character** using the Tangent CLI
2. **Export for contribution** (JSON format)
3. **Submit a Pull Request** with your character file
4. **Await review** from maintainers

## Requirements

### Minimum Requirements

All character contributions must meet these requirements:

- **Minimum 3 agent states**: `plan`, `think`, `execute` (required)
- **Valid pattern codes**: Use only supported pattern characters
- **Dimensions**: Between 5x3 and 20x10 characters
- **Tested**: Character must be tested in Tangent CLI
- **JSON format**: Properly formatted JSON file

### Optional States

You can add these optional standard states:

- `wait` - Agent waiting for input
- `error` - Agent handling errors
- `success` - Agent celebrating success

### Custom States

You may add custom states for unique character behaviors, but the 3 required states must be present.

## Step-by-Step Guide

### Step 1: Create Your Character

Use the interactive Tangent CLI:

```bash
tangent
```

Choose "Create new character" and follow the prompts:

1. Enter character name (e.g., "dragon", "robot", "wizard")
2. Choose personality (efficient, friendly, analytical, creative)
3. Set dimensions (width x height)
4. Add agent states (minimum: plan, think, execute)

### Step 2: Design Agent States

For each state, design the visual representation:

**Plan State** - Show the character analyzing/planning
- Use question marks (?), dots (.), or thinking symbols
- Example: Character with thought bubbles or analysis indicators

**Think State** - Show the character processing information
- Use processing dots, spinning indicators, or brain symbols
- Example: Character with processing animation

**Execute State** - Show the character performing actions
- Use action symbols, movement indicators, or progress bars
- Example: Character in motion or working

**Optional States:**
- **Wait**: Idle, patient pose
- **Error**: Confused, error indicators (X marks)
- **Success**: Celebrating, checkmarks (✓)

### Step 3: Export for Contribution

In the character builder menu, choose:

```
6. Export for contribution (JSON)
```

This will generate:
- `your-character.json` - Character definition
- `your-character-README.md` - Contribution guide

### Step 4: Submit Pull Request

1. **Fork** the Tangent repository on GitHub
2. **Clone** your fork locally
3. **Create a branch**: `git checkout -b add-your-character`
4. **Add your JSON** to the `characters/` directory
5. **Commit**: `git commit -m "Add your-character"`
6. **Push**: `git push origin add-your-character`
7. **Create Pull Request** on GitHub

Use the character contribution PR template when submitting.

## Character Design Guidelines

### Visual Quality

- **Clear and readable**: Characters should be recognizable
- **Consistent style**: Use similar visual language across states
- **Appropriate size**: Not too large (max 20x10) or too small (min 5x3)
- **Good contrast**: Use block characters effectively

### State Design

- **Distinct states**: Each state should be visually different
- **Meaningful animations**: States should represent the behavior
- **Smooth transitions**: Consider how states flow together
- **Personality match**: Design should match character personality

### Pattern Codes

Use these pattern characters:

```
Basic Blocks:
  F = █  (Full)
  T = ▀  (Top half)
  B = ▄  (Bottom half)
  L = ▌  (Left half)
  R = ▐  (Right half)

Shading:
  . = ░  (Light)
  : = ▒  (Medium)
  # = ▓  (Dark)

Quadrants (1-4):
  1 = ▘  (Upper Left)
  2 = ▝  (Upper Right)
  3 = ▖  (Lower Left)
  4 = ▗  (Lower Right)

Three-Quadrants (5-8):
  5 = ▛  (UL+UR+LL)
  6 = ▜  (UL+UR+LR)
  7 = ▙  (UL+LL+LR)
  8 = ▟  (UR+LL+LR)

Special:
  _ = Space
  X = Mirror marker
```

## Review Process

### Automated Validation

When you submit a PR, automated checks will:

- Validate JSON structure
- Check for required states (plan, think, execute)
- Verify dimensions are within limits
- Validate pattern codes

### Manual Review

Maintainers will review:

- Visual quality and clarity
- State appropriateness
- Overall design consistency
- Uniqueness (not duplicate of existing character)

### Approval Process

1. **Automated checks pass** - PR is ready for review
2. **Maintainer review** - Visual and quality check
3. **Feedback/changes** - If needed, make requested changes
4. **Approval** - Maintainer approves PR
5. **Merge** - Character is added to library
6. **Compilation** - Character is compiled into next release

## Example Character

Here's a simple example of a character JSON structure:

```json
{
  "name": "robot",
  "personality": "efficient",
  "width": 9,
  "height": 4,
  "frames": [
    {
      "name": "plan",
      "lines": [
        "_L5FFF5R_",
        "_6FFFFF6_",
        "__?F_F?__",
        "__FF_FF__"
      ],
      "state_type": "standard"
    },
    {
      "name": "think",
      "lines": [
        "_L5FFF5R_",
        "_6FFFFF6_",
        "__._._.__",
        "__FF_FF__"
      ],
      "state_type": "standard"
    },
    {
      "name": "execute",
      "lines": [
        "_L5FFF5R_",
        "_6FFFFF6_",
        "__>F_F>__",
        "__FF_FF__"
      ],
      "state_type": "standard"
    }
  ]
}
```

## Tips for Success

1. **Start simple** - Don't overcomplicate your first character
2. **Test thoroughly** - Preview and animate before submitting
3. **Be creative** - Unique characters are more likely to be accepted
4. **Follow guidelines** - Ensure all requirements are met
5. **Be patient** - Review process may take a few days

## Questions?

If you have questions about contributing characters:

1. Check existing characters for inspiration
2. Review the pattern guide in the CLI
3. Open an issue on GitHub for clarification
4. Join the community discussions

## License

By contributing a character, you agree that your contribution will be licensed under the same license as the Tangent project (MIT License).

---

Thank you for contributing to Tangent! Your characters help make AI agents more expressive and engaging.



