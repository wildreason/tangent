# Contributing Characters

Thank you for your interest in contributing characters to Tangent! This guide walks you through the character creation and submission process.

## Overview

The contribution workflow:
1. Create your character using `tangent create`
2. Export JSON and README files
3. Fork the repository on GitHub
4. Submit a Pull Request

## Character Requirements

All contributed characters must meet these requirements:

### Required Components
- ✓ **Base character**: Idle/immutable foundation
- ✓ **Plan state**: 3+ frames showing planning/analysis
- ✓ **Think state**: 3+ frames showing processing/reasoning
- ✓ **Execute state**: 3+ frames showing action/implementation

### Technical Requirements
- ✓ Minimum 3 frames per state (for animation)
- ✓ Valid Unicode block patterns only
- ✓ Consistent dimensions across all frames
- ✓ Valid Go identifier name (letters, numbers, underscore only)
- ✓ Dimensions between 1x1 and 100x50

### Optional Enhancements
- Additional states (wait, error, success)
- Custom FPS settings per state
- Custom loop counts

## Step-by-Step Guide

### Step 1: Create Your Character

Start the interactive character builder:

```bash
tangent create
```

#### 1.1 Enter Basic Information
- **Character name**: Use descriptive name (e.g., "mercury", "rocket")
  - Must be valid Go identifier (no hyphens or special characters)
  - Should be unique and memorable
- **Width**: Character width in terminal columns (e.g., 2, 11)
- **Height**: Character height in terminal rows (e.g., 2, 5)

#### 1.2 Create Base Character
The base character is your foundation - the idle/default appearance:

```
Select: Create base character
Enter each line using pattern codes:
F=█ T=▀ B=▄ L=▌ R=▐ 1-8=quads .=░ :=▒ #=▓ _=space X=◐

Example (2x2):
Line 1: FF
Line 2: FF
```

#### 1.3 Add Agent States
Create at least 3 required states (plan, think, execute):

```
Select: Add agent state
State name: plan
Number of frames: 3

Frame 1:
Line 1: TF
Line 2: FF

Frame 2:
Line 1: FT
Line 2: FF

Frame 3:
Line 1: FF
Line 2: BF
```

**Tips:**
- Each state needs 3+ frames for smooth animation
- Use pattern codes consistently
- Preview frequently to check animation
- You can start from base or create from scratch

#### 1.4 Preview Your Work
Use the built-in preview to see animations:

```
Select: Animate all states
- Navigate between states with arrow keys
- See how each state animates
- Verify smoothness and clarity
```

### Step 2: Export Your Character

When satisfied with your character:

```
Select: Export for contribution (JSON)
```

This creates two files:
- `yourcharacter.json` - Character definition
- `yourcharacter-README.md` - Documentation

**Review these files** before submitting to ensure correctness.

### Step 3: Submit Your Contribution

#### 3.1 Fork the Repository
1. Go to https://github.com/wildreason/tangent
2. Click "Fork" button (top right)
3. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR-USERNAME/tangent
   cd tangent
   ```

#### 3.2 Add Your Character
1. Create a new branch:
   ```bash
   git checkout -b add-yourcharacter
   ```

2. Copy your JSON file to the repository:
   ```bash
   cp /path/to/yourcharacter.json characters/
   ```

3. Commit your changes:
   ```bash
   git add characters/yourcharacter.json
   git commit -m "Add yourcharacter character"
   ```

4. Push to your fork:
   ```bash
   git push origin add-yourcharacter
   ```

#### 3.3 Create Pull Request
1. Go to your fork on GitHub
2. Click "Pull Request" button
3. Title: "Add [character-name] character"
4. Description: Include character README content
5. Submit the PR

### Step 4: Review Process

Once submitted, maintainers will:

1. **Validate** your character:
   - Check requirements (base + 3 states, 3+ frames)
   - Verify pattern codes are valid
   - Test animation quality

2. **Register** if approved:
   - Run `tangent admin register yourcharacter.json`
   - Generate Go library file
   - Rebuild and test

3. **Merge** your contribution:
   - Your character becomes part of the library
   - Available to all users via `characters.LibraryAgent("yourcharacter")`

## Character Design Guidelines

### Visual Design

**Do:**
- Create distinct, recognizable characters
- Use clear, readable patterns
- Make states visually different
- Test on both light and dark terminals
- Keep animations smooth and purposeful

**Don't:**
- Make states too similar (hard to distinguish)
- Use overly complex patterns (hard to read)
- Create jerky animations (too few frames)
- Exceed practical dimensions (keep under 20x10 for most use cases)

### State Design

Each state should clearly communicate its purpose:

**Plan:**
- Show analysis, questioning, strategizing
- Use symbols: ?, ., thinking postures
- Should feel contemplative

**Think:**
- Show processing, reasoning, computation
- Use symbols: ..., brain indicators
- Should feel active

**Execute:**
- Show action, movement, implementation
- Use symbols: →, !, progress indicators
- Should feel decisive

### Animation Tips

1. **Frame Count**: Start with 3 frames, add more if needed
2. **Timing**: 5 FPS is good default for terminal animations
3. **Transitions**: Make frame changes smooth and natural
4. **Purpose**: Every frame should serve the animation
5. **Testing**: Preview frequently during creation

## Examples

### Minimal Character (2x2)

```json
{
  "name": "dot",
  "width": 2,
  "height": 2,
  "base_frame": {
    "name": "base",
    "lines": ["FF", "FF"]
  },
  "states": [
    {
      "name": "plan",
      "frames": [
        {"lines": ["TF", "FF"]},
        {"lines": ["FF", "TF"]},
        {"lines": ["FF", "FT"]}
      ]
    }
  ]
}
```

### Medium Character (5x3)

See existing characters in `pkg/characters/library/` for examples like `mercury`, `water`, etc.

## Troubleshooting

### "Character name must be valid Go identifier"
- Use only letters, numbers, and underscores
- No hyphens, spaces, or special characters
- Examples: `my_char`, `robot2`, `agent_alpha`

### "Expected N characters, got M"
- Each line must match the character width exactly
- Count your pattern codes carefully
- Use `_` for spaces

### "Invalid pattern code"
- Use only valid codes: F T B L R 1-8 . : # _ X
- Check for typos in your patterns
- See [Pattern Guide](PATTERNS.md) for reference

### "Minimum 3 frames required"
- Each state needs at least 3 frames for animation
- Add more frames to create smooth transitions

## FAQ

**Q: Can I contribute multiple characters?**  
A: Yes! Submit each character as a separate PR for easier review.

**Q: Can I update an existing character?**  
A: Yes, submit a PR with the updated JSON file and explain changes.

**Q: What if my character is rejected?**  
A: We'll provide feedback. You can revise and resubmit.

**Q: How long does review take?**  
A: Usually within a few days, depending on maintainer availability.

**Q: Can I use colors?**  
A: No, characters use Unicode block patterns only (terminal-agnostic).

**Q: What dimensions should I use?**  
A: Start small (2x2 to 5x3), increase if needed. Most characters are under 11x5.

## Need Help?

- Check the [Pattern Guide](PATTERNS.md) for pattern codes
- See [States Guide](STATES.md) for state design tips
- Look at existing characters in `pkg/characters/library/`
- Open an issue on GitHub for questions

## Thank You!

Your contributions make Tangent better for everyone. We appreciate your creativity and effort in designing agent characters!

---

**Tangent** - Terminal Agent Designer

