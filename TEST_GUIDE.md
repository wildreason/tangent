# Testing the Agent State Design

## Quick Test Options

### Option 1: Run the Demo Application ⭐ (Recommended)

See all features in action:

```bash
cd /Users/btsznh/wild/characters
go run examples/agent_states.go
```

This demonstrates:
- Creating characters with agent states
- Using state methods (Plan, Think, Execute, etc.)
- State inspection and validation
- Custom states
- Practical AI agent workflows

---

### Option 2: Test the Interactive CLI

Create a character with the new agent state workflow:

```bash
cd /Users/btsznh/wild/characters
./tangent
```

**What to test:**

1. **Create new character**
   - Choose option 1
   - Enter a name (e.g., "test-bot")
   - Choose dimensions (e.g., 7x5)
   - **NEW**: Choose personality (efficient/friendly/analytical/creative)

2. **Add agent states**
   - Choose option 1 (Add new frame)
   - **NEW**: See required states listed (plan, think, execute)
   - **NEW**: See optional states (wait, error, success)
   - Add "plan" state first
   - Add "think" state
   - Add "execute" state
   - **NEW**: See progress message "All required states added!"

3. **Export for contribution**
   - Choose option 6 (Export for contribution - JSON)
   - **NEW**: Validates you have required states
   - **NEW**: Exports JSON + README
   - Check the generated files

---

### Option 3: Test the API Programmatically

Create a simple test file:

```bash
cat > test_agent_api.go << 'EOF'
package main

import (
    "fmt"
    "os"
    "github.com/wildreason/tangent/pkg/characters"
    "github.com/wildreason/tangent/pkg/characters/domain"
)

func main() {
    // Create a test character with states
    char := &domain.Character{
        Name:        "test-agent",
        Personality: "efficient",
        Width:       7,
        Height:      3,
        States: map[string]domain.State{
            "plan": {
                Name:        "Planning",
                Description: "Test planning",
                StateType:   "standard",
                Frames: []domain.Frame{
                    {Lines: []string{"_?????_", "_?????_", "_?????_"}},
                },
            },
            "think": {
                Name:        "Thinking",
                Description: "Test thinking",
                StateType:   "standard",
                Frames: []domain.Frame{
                    {Lines: []string{"_....._", "_....._", "_....._"}},
                },
            },
            "execute": {
                Name:        "Executing",
                Description: "Test executing",
                StateType:   "standard",
                Frames: []domain.Frame{
                    {Lines: []string{"_>>>>>_", "_>>>>>_", "_>>>>>_"}},
                },
            },
        },
    }

    // Wrap in AgentCharacter
    agent := characters.NewAgentCharacter(char)

    // Test state methods
    fmt.Println("Testing agent states:")
    fmt.Println()

    fmt.Println("1. Planning:")
    agent.Plan(os.Stdout)
    fmt.Println()

    fmt.Println("2. Thinking:")
    agent.Think(os.Stdout)
    fmt.Println()

    fmt.Println("3. Executing:")
    agent.Execute(os.Stdout)
    fmt.Println()

    // Test state inspection
    fmt.Println("Available states:", agent.ListStates())
    fmt.Println("Has 'plan' state:", agent.HasState("plan"))
    fmt.Println("Character name:", agent.Name())
    fmt.Println("Personality:", agent.Personality())
}
EOF

go run test_agent_api.go
```

---

### Option 4: Run Unit Tests

Run the comprehensive test suite:

```bash
cd /Users/btsznh/wild/characters
go test ./pkg/characters/agent_test.go ./pkg/characters/agent.go -v
```

Expected output: All 11 tests should pass ✅

---

## What to Look For

### ✅ CLI Enhancements

1. **Personality Selection**
   - Should see 4 personality options when creating character
   - Default is "efficient"

2. **Agent State Guidance**
   - When adding frames, should see "Adding agent state" (not "Adding frame")
   - Should see required states listed with descriptions
   - Should see optional states listed
   - Should see progress on required states

3. **State Validation**
   - Should show "X required state(s) remaining"
   - Should show "All required states added!" when complete
   - Export should fail if missing required states

4. **JSON Export**
   - Should create `character-name.json`
   - Should create `character-name-README.md`
   - Should show next steps for GitHub PR

### ✅ API Features

1. **State Methods**
   - `agent.Plan(os.Stdout)` - Should display planning state
   - `agent.Think(os.Stdout)` - Should display thinking state
   - `agent.Execute(os.Stdout)` - Should display executing state
   - `agent.Wait(os.Stdout)` - Should display waiting state
   - `agent.Error(os.Stdout)` - Should display error state
   - `agent.Success(os.Stdout)` - Should display success state

2. **State Inspection**
   - `agent.ListStates()` - Should return array of state names
   - `agent.HasState("plan")` - Should return true/false
   - `agent.GetStateDescription("plan")` - Should return description
   - `agent.Name()` - Should return character name
   - `agent.Personality()` - Should return personality

3. **Custom States**
   - `agent.ShowState(os.Stdout, "custom")` - Should display custom state
   - Should work with any state name

### ✅ Backward Compatibility

Old API should still work:

```bash
cat > test_backward_compat.go << 'EOF'
package main

import (
    "os"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    // Old API should still work
    alien, _ := characters.Library("alien")
    characters.ShowIdle(os.Stdout, alien)
}
EOF

go run test_backward_compat.go
```

---

## Complete Test Workflow

### Step 1: Test Demo App
```bash
go run examples/agent_states.go
```
**Expected**: See 6 demos with agent states working

### Step 2: Test CLI Creation
```bash
./tangent
# Create character with states
# Export for contribution
```
**Expected**: JSON + README files created

### Step 3: Test Unit Tests
```bash
go test ./pkg/characters/agent_test.go ./pkg/characters/agent.go -v
```
**Expected**: 11/11 tests passing

### Step 4: Test API Integration
```bash
go run test_agent_api.go
```
**Expected**: States display correctly

### Step 5: Test Backward Compatibility
```bash
go run test_backward_compat.go
```
**Expected**: Old API still works

---

## Troubleshooting

### Issue: "undefined: characters.LibraryAgent"

**Solution**: Make sure you're using the updated code:
```bash
go mod tidy
go build ./...
```

### Issue: CLI doesn't show personality option

**Solution**: Rebuild the tangent binary:
```bash
cd cmd/tangent
go build -o ../../tangent
```

### Issue: Tests fail

**Solution**: Check you have all dependencies:
```bash
go mod download
go test ./...
```

---

## Quick Verification Checklist

- [ ] Demo app runs successfully
- [ ] CLI shows personality selection
- [ ] CLI shows agent state guidance
- [ ] CLI validates required states
- [ ] JSON export works
- [ ] Unit tests pass (11/11)
- [ ] State methods work (Plan, Think, Execute, etc.)
- [ ] State inspection works (ListStates, HasState, etc.)
- [ ] Custom states work
- [ ] Old API still works (backward compatibility)

---

## Next Steps After Testing

1. **Try creating your own character** with the CLI
2. **Export it for contribution** to see the workflow
3. **Integrate agent states** into your own project
4. **Read the documentation** in `docs/AGENT_STATES.md`
5. **Check the quick start** in `QUICK_START_AGENT_STATES.md`

---

## Questions?

- Check `IMPLEMENTATION_SUMMARY.md` for complete details
- See `docs/AGENT_STATES.md` for full documentation
- Run `tangent help` for CLI usage
- Look at `examples/agent_states.go` for code examples

