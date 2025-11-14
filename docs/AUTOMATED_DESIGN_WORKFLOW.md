# Automated Design Iteration Workflow

## Overview

A system that automatically generates design variations, builds, previews, and presents them for user approval/decline, with save/discard and rebuild capabilities.

## Architecture Components

### 1. Design Generator (`pkg/characters/designer/`)
```go
// Core generator interface
type DesignGenerator interface {
    GenerateVariation(original StateDefinition, strategy VariationStrategy) (StateDefinition, error)
    GenerateMultiple(original StateDefinition, count int, strategies []VariationStrategy) ([]StateDefinition, error)
}

// Variation strategies
type VariationStrategy struct {
    Type        string  // "frame_reduction", "line_modification", "pattern_addition"
    Parameters  map[string]interface{}
    Constraints Constraints
}
```

**Responsibilities:**
- Generate design variations based on strategies
- Apply constraints (frame count limits, pattern rules)
- Maintain design consistency
- Track generation history

### 2. Build Manager (`pkg/characters/build/`)
```go
type BuildManager struct {
    ProjectRoot string
    TempDir     string
    BuildCmd    string  // "make build-cli"
}

func (bm *BuildManager) BuildWithChanges(changes []StateChange) (BuildResult, error)
func (bm *BuildManager) Revert() error
func (bm *BuildManager) Save() error
```

**Responsibilities:**
- Apply changes to JSON files
- Execute build process
- Manage temporary/staging directories
- Handle build failures gracefully

### 3. Preview System (`pkg/characters/preview/`)
```go
type PreviewRenderer struct {
    OutputFormat string  // "terminal", "html", "gif", "video"
    FPS          int
    Loops        int
}

func (pr *PreviewRenderer) Render(stateName string, frames []Frame) (Preview, error)
func (pr *PreviewRenderer) Compare(original, modified Preview) (Diff, error)
```

**Responsibilities:**
- Render animations for preview
- Generate side-by-side comparisons
- Create visual diffs
- Export previews in multiple formats

### 4. Approval Workflow (`pkg/characters/workflow/`)
```go
type WorkflowManager struct {
    CurrentIteration *Iteration
    History          []Iteration
    State           WorkflowState
}

type Iteration struct {
    ID          string
    Timestamp   time.Time
    Changes     []StateChange
    Preview     Preview
    BuildResult BuildResult
    Status      IterationStatus  // "pending", "approved", "declined", "saved"
    Metadata    map[string]interface{}
}

type WorkflowState struct {
    CurrentState string  // "generating", "building", "previewing", "awaiting_approval"
    ActiveChanges map[string]StateChange
}
```

**Responsibilities:**
- Manage iteration lifecycle
- Track approval history
- Handle save/discard operations
- Maintain state across sessions

### 5. User Interface (`cmd/tangent-cli/workflow.go`)
```go
type WorkflowUI struct {
    Interactive bool
    AutoMode    bool
    BatchSize   int
}

func (ui *WorkflowUI) RunWorkflow(generator DesignGenerator) error
func (ui *WorkflowUI) ShowPreview(preview Preview) error
func (ui *WorkflowUI) PromptApproval(iteration *Iteration) (bool, error)
```

**Responsibilities:**
- Interactive approval prompts
- Batch processing mode
- Progress indicators
- Error handling and recovery

## Workflow Steps

### Phase 1: Generation
```
1. Load original state (e.g., search.json)
2. Analyze current design (frame count, patterns, structure)
3. Generate variation based on strategy:
   - Frame reduction: Remove frames 3,4,5
   - Line modification: Add _rfffffffl_ to first line
   - Pattern changes: Modify specific patterns
4. Validate generated design (constraints, rules)
5. Create StateChange object
```

### Phase 2: Build
```
1. Create temporary copy of state file
2. Apply changes to temporary file
3. Validate JSON structure
4. Execute: make build-cli
5. Capture build output and errors
6. If build fails: mark iteration as failed, return to generation
```

### Phase 3: Preview
```
1. Load modified state from temporary build
2. Render animation preview:
   - Terminal: ./tangent-cli browse ga --state search
   - Or programmatic: agent.AnimateState()
3. Generate comparison (before/after)
4. Create preview artifacts (screenshots, gifs, etc.)
5. Store preview in iteration
```

### Phase 4: Approval
```
1. Display preview to user
2. Show change summary:
   - Frames: 6 → 3
   - Lines modified: [list]
   - Build status: ✓ success
3. Present options:
   [A]pprove  [D]ecline  [S]ave  [Q]uit  [N]ext variation
4. Handle user choice
```

### Phase 5: Save/Discard
```
If Approved:
  1. Copy temporary file to actual location
  2. Commit changes (optional)
  3. Update iteration status
  4. Rebuild final version
  5. Verify final build

If Declined:
  1. Discard temporary files
  2. Mark iteration as declined
  3. Optionally generate new variation
  4. Return to Phase 1

If Saved (for later):
  1. Store iteration in pending queue
  2. Keep temporary files
  3. Allow batch approval later
```

### Phase 6: Rebuild
```
1. If approved, rebuild with final changes
2. Run tests
3. Update binary
4. Verify all states still work
5. Generate final preview
```

## Data Structures

### StateChange
```go
type StateChange struct {
    StateName    string
    ChangeType   string  // "frame_removal", "line_modification", "pattern_addition"
    Original     StateDefinition
    Modified     StateDefinition
    Diff         ChangeDiff
    Metadata     map[string]interface{}
}
```

### ChangeDiff
```go
type ChangeDiff struct {
    FramesRemoved    []int
    FramesAdded      []Frame
    LinesModified    []LineChange
    PatternsChanged  []PatternChange
}
```

### Preview
```go
type Preview struct {
    Format      string  // "terminal", "gif", "html"
    Content     []byte
    Metadata    PreviewMetadata
    Comparison  *ComparisonPreview
}

type PreviewMetadata struct {
    FrameCount    int
    AnimationFPS   int
    Duration      time.Duration
    FileSize      int64
}
```

## Implementation Approach

### Phase 1: Core Generator
```go
// pkg/characters/designer/generator.go
package designer

type Generator struct {
    strategies []VariationStrategy
    constraints Constraints
}

func (g *Generator) GenerateFrameReduction(state StateDefinition, framesToRemove []int) (StateDefinition, error) {
    // Implementation
}

func (g *Generator) GenerateLineModification(state StateDefinition, lineIndex int, newLine string) (StateDefinition, error) {
    // Implementation
}
```

### Phase 2: Build Integration
```go
// pkg/characters/build/manager.go
package build

func (bm *BuildManager) ApplyChangesToFile(filePath string, changes []StateChange) error {
    // Read original
    // Apply changes
    // Write to temp location
    // Validate JSON
}

func (bm *BuildManager) ExecuteBuild() (BuildResult, error) {
    // Run: make build-cli
    // Capture output
    // Check for errors
}
```

### Phase 3: Preview Integration
```go
// pkg/characters/preview/renderer.go
package preview

func (pr *PreviewRenderer) RenderFromState(stateName string) (Preview, error) {
    // Load state
    // Create agent character
    // Animate and capture
    // Generate preview
}
```

### Phase 4: Workflow Orchestration
```go
// cmd/tangent-cli/workflow.go
func runAutomatedWorkflow(stateName string, strategy VariationStrategy) error {
    // 1. Generate
    generator := designer.NewGenerator()
    variation, err := generator.GenerateVariation(original, strategy)
    
    // 2. Build
    buildMgr := build.NewManager()
    result, err := buildMgr.BuildWithChanges([]StateChange{change})
    
    // 3. Preview
    renderer := preview.NewRenderer()
    preview, err := renderer.RenderFromState(stateName)
    
    // 4. Approval
    ui := workflow.NewUI()
    approved, err := ui.PromptApproval(iteration)
    
    // 5. Save/Discard
    if approved {
        buildMgr.Save()
        buildMgr.Rebuild()
    } else {
        buildMgr.Discard()
    }
}
```

## CLI Command Structure

```bash
# Interactive workflow
tangent-cli workflow --state search --strategy frame_reduction

# Batch mode (generate multiple, approve all)
tangent-cli workflow --state search --batch 5 --auto-approve

# Preview only (no changes)
tangent-cli workflow --preview --state search

# Apply saved changes
tangent-cli workflow --apply-saved

# List pending iterations
tangent-cli workflow --list-pending
```

## Configuration

### Workflow Config (`workflow.yaml`)
```yaml
generator:
  strategies:
    - type: frame_reduction
      max_removals: 3
    - type: line_modification
      allowed_patterns: ["_rfffffffl_", "_fffffffff_"]
  
build:
  command: "make build-cli"
  timeout: 30s
  retries: 2
  
preview:
  format: "terminal"
  fps: 5
  loops: 2
  comparison: true
  
approval:
  interactive: true
  auto_approve_on_success: false
  batch_size: 5
```

## Benefits

1. **Speed**: Automated generation and testing
2. **Consistency**: All changes follow same validation rules
3. **Safety**: Build verification before approval
4. **History**: Track all iterations and decisions
5. **Batch Processing**: Handle multiple variations
6. **Rollback**: Easy to revert declined changes

## Challenges & Solutions

### Challenge 1: Design Quality
**Problem**: Generated designs might not be aesthetically pleasing
**Solution**: 
- Constraint-based generation
- Pattern library for valid modifications
- User feedback learning

### Challenge 2: Build Failures
**Problem**: Changes might break compilation
**Solution**:
- Pre-validation before build
- JSON schema validation
- Graceful error handling

### Challenge 3: Preview Accuracy
**Problem**: Preview might not match final result
**Solution**:
- Use same rendering pipeline as production
- Side-by-side comparison
- Multiple preview formats

### Challenge 4: State Management
**Problem**: Managing temporary files and iterations
**Solution**:
- Dedicated temp directory per iteration
- Cleanup on approval/decline
- Persistent iteration history

## Future Enhancements

1. **AI-Powered Generation**: Use ML models to suggest better variations
2. **A/B Testing**: Automatically test multiple variations
3. **Performance Metrics**: Measure animation smoothness, file size
4. **Collaborative Approval**: Multiple reviewers
5. **Version Control Integration**: Auto-commit approved changes
6. **Web UI**: Browser-based preview and approval

## Example Usage

```bash
$ tangent-cli workflow --state search --strategy frame_reduction

🔧 Generating variation...
   Strategy: frame_reduction
   Removing frames: [3, 4, 5]
   Adding first line: _rfffffffl_

🔨 Building...
   ✓ Build successful

📺 Previewing...
   [Animation plays]

📊 Changes:
   - Frames: 6 → 3
   - Lines modified: 3
   - Build: ✓ success

[A]pprove  [D]ecline  [S]ave  [N]ext  [Q]uit: A

✅ Approved! Saving changes...
🔨 Rebuilding final version...
✓ Complete! Changes saved to search.json
```

