# Character Contributions

This directory contains community-contributed characters in JSON format.

Each character file should follow the format specified in `.github/CONTRIBUTING_CHARACTERS.md`.

## Submitting a Character

1. Create your character using the Tangent CLI
2. Export for contribution (JSON format)
3. Place your JSON file in this directory
4. Submit a Pull Request

See `.github/CONTRIBUTING_CHARACTERS.md` for full guidelines.

## Character Format

Each JSON file should contain:

```json
{
  "name": "character-name",
  "personality": "efficient|friendly|analytical|creative",
  "width": 7,
  "height": 5,
  "frames": [
    {
      "name": "plan",
      "lines": ["...", "..."],
      "state_type": "standard"
    },
    {
      "name": "think",
      "lines": ["...", "..."],
      "state_type": "standard"
    },
    {
      "name": "execute",
      "lines": ["...", "..."],
      "state_type": "standard"
    }
  ]
}
```

## Requirements

- Minimum 3 states: plan, think, execute
- Valid pattern codes only
- Dimensions: 5x3 to 20x10
- Tested in Tangent CLI
