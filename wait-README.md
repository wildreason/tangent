# wait Character

## Character Information

- **Name:** wait
- **Dimensions:** 11x4
- **States:** 1

## States

- **wait** (standard): 3 frames

## Preview

```
```

## Usage

```go
agent, _ := characters.LibraryAgent("wait")
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
