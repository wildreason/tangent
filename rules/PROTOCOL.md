# Team Alignment Check ✅

## The Three-Team Architecture

```
┌─────────────────────────────────────────────────┐
│ UNDERSTORY (Orchestration Layer)                │
│ • Multi-agent coordination                      │
│ • Agent communication (call/answer/broadcast)   │
│ • Hook system & task management                 │
│ • Emits: JSON state updates via stdout         │
└──────────────┬──────────────────────────────────┘
               │ JSON over stdout
               ▼
┌─────────────────────────────────────────────────┐
│ HANDWAVE (Visualization Layer)                  │
│ • Receives JSON states via stdin                │
│ • Displays beautiful overlays                   │
│ • Toggle to agent terminals                     │
│ • Uses: Tangent characters                      │
└──────────────┬──────────────────────────────────┘
               │ Import
               ▼
┌─────────────────────────────────────────────────┐
│ TANGENT (Design System)                         │
│ • Unicode block character design                │
│ • Pattern system & visual builder               │
│ • Pre-built character library                   │
│ • Framework-agnostic (pure Go)                  │
└─────────────────────────────────────────────────┘
```

---

## Team Clarity Assessment

### ✅ **UNDERSTORY Team - WELL FOCUSED**

**Clear responsibilities:**
- Agent lifecycle management
- Inter-agent communication
- Task orchestration
- Hook system

**Interface with others:**
- **→ Handwave:** Emits JSON states to stdout
- **← Handwave:** None (one-way only)
- **← Tangent:** None (doesn't use characters)

**Action needed:**
- Add JSON state emission to agent output
- Document the state protocol for their agents

---

### ✅ **TANGENT Team - WELL FOCUSED**

**Clear responsibilities:**
- Character design system
- Pattern compiler
- Visual/CLI builders
- Character library

**Interface with others:**
- **→ Handwave:** Provides characters via Go import
- **← Understory:** None
- **← Handwave:** None (pure library)

**Action needed:**
- Continue building character library
- Maybe add "agent" themed characters (thinking, working, idle, etc.)

---

### ⚠️ **HANDWAVE Team - NEEDS CLARITY**

**Clear responsibilities:**
- Read JSON from stdin
- Display agent states using Tangent characters
- Toggle to agent terminals
- Overlay window management

**Interface with others:**
- **← Understory:** Receives JSON states via stdin
- **← Tangent:** Imports character library
- **→ Understory:** None (visualization only)

**Potential scope creep to avoid:**
- ❌ Don't build orchestration features
- ❌ Don't build agent communication
- ❌ Don't build character design tools
- ❌ Don't build multi-agent aggregation (not v0.1)

---

## The Flow (How It All Works Together)

```bash
# 1. Understory creates and runs agents
understory init
understory new  # tokyo

# 2. Understory agent pipes states to Handwave
understory agent tokyo | handwave overlay --session tokyo

# 3. Inside Understory's tokyo agent:
# - Agent does work
# - Emits: {"state":"plan","message":"Analyzing..."}
# - Handwave receives and displays

# 4. Handwave uses Tangent:
import "github.com/wildreason/tangent/pkg/characters"
planChar, _ := characters.Library("pulse")  // Tangent character
// Display planChar in overlay when state="plan"

# 5. User presses Ctrl+F:
# - Handwave toggles to tokyo's terminal
# - User sees full agent output
# - Press Ctrl+F again to return to overlay
```

---

## Dependency Graph

```
Understory (standalone)
    ↓ emits JSON
Handwave (depends on Tangent)
    ↓ imports
Tangent (standalone)
```

**Clean dependencies:**
- Understory: No dependencies on other two
- Tangent: No dependencies on other two  
- Handwave: Depends on Tangent (import), consumes Understory output (stdin)

---

## Scope Boundaries

### UNDERSTORY owns:
- ✅ `understory init` - Project setup
- ✅ `understory new` - Agent creation
- ✅ `understory call/answer/broadcast` - Communication
- ✅ `understory config hooks` - Hook system
- ✅ Agent process management
- ✅ Task coordination

### HANDWAVE owns:
- ✅ `handwave overlay` - Status window
- ✅ JSON protocol parser (stdin)
- ✅ State → Character mapping
- ✅ Toggle hotkey to terminal
- ✅ Tmux session integration
- ✅ Window spawning

### TANGENT owns:
- ✅ `tangent` - Visual character builder
- ✅ `tangent create` - CLI character creation
- ✅ `tangent gallery` - Library browser
- ✅ Character design system
- ✅ Pattern compiler
- ✅ Pre-built character library

---

## Potential Issues to Watch

### 🟡 Issue 1: State Protocol Coordination
**Problem:** Understory and Handwave must agree on JSON format.

**Solution:**
- Create shared `PROTOCOL.md` in both repos
- Version the protocol (v1)
- Understory team documents what they emit
- Handwave team documents what they expect

### 🟡 Issue 2: Character Requirements
**Problem:** Handwave needs specific characters for states.

**Solution:**
- Handwave documents character requirements (idle, plan, progress, done, error)
- Tangent team creates matching characters
- Use existing library + add new ones as needed

### 🟢 Issue 3: Independent Development
**Problem:** Teams might build in isolation.

**Solution:**
- Weekly sync: "Show what you built"
- Integration testing: Understory → Handwave → Tangent
- Example workflows in each repo

---

## Integration Points (v0.1)

### Understory → Handwave
**What Understory must provide:**
```bash
# Run agent with stdout that Handwave can consume
understory agent tokyo  # prints JSON states to stdout
```

**JSON format (agreed protocol):**
```json
{"state":"idle","message":"Waiting for tasks"}
{"state":"plan","message":"Analyzing request"}
{"state":"progress","message":"Processing files","progress":50}
{"state":"done","message":"Task completed"}
{"state":"error","message":"Failed: file not found"}
```

### Tangent → Handwave
**What Tangent must provide:**
```go
// Library characters Handwave can use
characters.Library("pulse")   // For "plan" state
characters.Library("wave")    // For "progress" state
characters.Library("rocket")  // For "done" state
// etc.
```

**Handwave imports:**
```go
import "github.com/wildreason/tangent/pkg/characters"
import "github.com/wildreason/tangent/pkg/adapters/bubbletea"
```

---

## v0.1 Success Criteria (Cross-Team)

**Integration works when:**

✅ Developer runs: `understory agent tokyo | handwave overlay --session tokyo`  
✅ Handwave displays Tokyo agent's state using Tangent characters  
✅ State changes (idle→plan→progress→done) animate correctly  
✅ Ctrl+F toggles to Tokyo's full terminal  
✅ All three tools work independently (can swap components)  

---

## Recommendations

### For UNDERSTORY Team:
1. ✅ Add JSON state emission to agent stdout
2. ✅ Document emitted states in README
3. ✅ Create example: "Using Understory with Handwave"
4. ⚠️ Don't build visualization - that's Handwave's job

### For HANDWAVE Team:
1. ✅ Focus only on visualization + toggle
2. ✅ Use Tangent for characters (don't build character system)
3. ✅ Document Understory integration prominently
4. ⚠️ Don't build orchestration - that's Understory's job
5. ⚠️ Don't build characters - that's Tangent's job

### For TANGENT Team:
1. ✅ Keep building character design system
2. ✅ Add "agent workflow" themed characters (thinking, working, celebrating, error)
3. ✅ Ensure Bubble Tea adapter is solid (Handwave uses this)
4. ⚠️ Don't build TUI framework - stay focused on characters

---

## The Big Picture

**You have clean separation of concerns:**

- **Understory** = The brain (orchestration)
- **Handwave** = The eyes (visualization)
- **Tangent** = The art department (design system)

**Each team can ship independently:**
- Understory works without Handwave (just CLI)
- Handwave works with any JSON source (not just Understory)
- Tangent works with any Go TUI app (not just Handwave)

**But together they're magical:**
- Beautiful multi-agent development experience
- Each tool does one thing perfectly
- Clean interfaces between them

---

## Final Assessment

### ✅ **YES, you are well-focused!**

**Strengths:**
- Clear boundaries between teams
- Each tool is useful independently
- Clean dependency graph (Tangent → Handwave → Understory output)
- No circular dependencies

**Watch out for:**
- Protocol drift (keep PROTOCOL.md in sync)
- Feature creep (each team stick to your lane)
- Integration testing (test the full stack regularly)

**Action items:**
1. Create shared `PROTOCOL.md` (Understory + Handwave)
2. Handwave documents required characters for Tangent team
3. Weekly integration demos across teams
4. Keep scope laser-focused for v0.1

You're building something really clean. Stay focused! 🎯
