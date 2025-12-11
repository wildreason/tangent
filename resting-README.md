# resting Character

## Character Information

- **Name:** resting
- **Dimensions:** 10x2
- **States:** 7

## States

- **resting** (custom): 3 frames
- **arise** (custom): 3 frames
- **wait** (standard): 2 frames
- **read** (custom): 4 frames
- **write** (custom): 5 frames
- **search** (custom): 4 frames
- **approval** (custom): 3 frames

## Preview

```
 ▐██████▌ 
 ▐█▙█ █▙▌ 
```

## Usage

```go
agent, _ := characters.LibraryAgent("resting")
agent.Plan(os.Stdout)   // Show plan state
agent.Think(os.Stdout)  // Show think state
agent.Execute(os.Stdout) // Show execute state
```

## Contribution

This character was created using Tangent character builder.

### Next Steps

1. Review the exported JSON file
2. Fork the Tangent repository on GitHub
3. Add your character JSON to the repository
4. Submit a Pull Request
