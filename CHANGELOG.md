# Changelog

## [0.1.0-alpha.5] - 2025-10-16

**AI-Native Positioning: Terminal Avatars for AI Agents**

Alpha.5 repositions Tangent with laser focus on its true purpose: providing terminal avatars for AI agents. This is not a generic animation library - it's the single-source solution for giving AI agents a face, presence, and personality in terminal applications.

### Strategic Positioning Changes

- **NEW TAGLINE**: "Terminal Avatars for AI Agents" (was "Terminal Agent Designer")
- **NEW PHILOSOPHY**: Give your AI agent a face - not "add animations to CLIs"
- **NEW MESSAGING**: All documentation emphasizes AI-native, terminal-first design
- **NEW VISION**: By v1.0, become the standard library for AI agent terminal avatars

### Documentation Overhaul

- **REWRITTEN**: README.md with AI-native positioning and Stripe-like clarity
  - Clear value prop: "Give your AI agent a face"
  - Browse → pick → integrate workflow emphasized
  - Available avatars showcased with personality descriptions
  - Philosophy section explains terminal-first, AI-native approach
- **RENAMED**: docs/CLI.md → docs/CREATORS.md (for advanced users/contributors)
- **ENHANCED**: docs/API.md with discovery workflow and AI agent context
  - Added "Why Terminal Avatars?" section
  - Added discovery workflow using `tangent browse`
  - Emphasized guaranteed states and AI workflow mapping
- **UPDATED**: All taglines changed from "Terminal Agent Designer" to "Terminal Avatars for AI Agents"

### CLI Changes

- **HIDDEN**: `tangent create` removed from help output (still functional for contributors)
- **UPDATED**: Help text now emphasizes `tangent browse` as primary discovery tool
- **REMOVED**: References to creation workflow in main help
- **SIMPLIFIED**: Focus on avatar discovery and API integration
- **UPDATED**: Banner and usage text with new "Terminal Avatars for AI Agents" tagline

### Install Script Updates

- **UPDATED**: Installer messaging emphasizes avatars for AI agents
- **CHANGED**: Success message: "Next: tangent browse (discover avatars for your AI agent)"
- **REMOVED**: References to `tangent create` from installer output

### Positioning Strategy

**What This Changes:**
- From: "A tool to design and use CLI character animations"
- To: "The definitive terminal avatar system for AI agents"

**Target Audience:**
- Developers building AI-native CLI applications
- AI agent frameworks needing visible presence
- Terminal-first AI tools

**Not For:**
- Generic CLI beautification
- Web-based animations
- Non-AI terminal applications

### Roadmap Alignment

- **Alpha.5** (current): API-first positioning, curated avatar library
- **Beta** (next): Expanded library (10-20 avatars), richer state vocabulary
- **v1.0** (goal): The standard - every AI CLI tool uses Tangent avatars

### Files Modified

1. README.md - Complete rewrite with AI-native focus
2. cmd/tangent/main.go - Updated taglines, hidden `create` from help
3. docs/CLI.md → docs/CREATORS.md - Renamed and repositioned for advanced users
4. docs/API.md - Added AI-native context and discovery workflow
5. install.sh - Updated all messaging to emphasize avatars for AI agents
6. CHANGELOG.md - This entry

### Migration Notes

**For Users:**
- No API changes - all existing code continues to work
- `tangent browse` is now the primary discovery command
- `tangent create` still available but not promoted (see docs/CREATORS.md)

**For Contributors:**
- Character creation workflow unchanged
- Export → PR process maintained
- Focus on AI-agent-appropriate avatars

### Philosophy

Tangent is purpose-built for AI agents in terminal applications. Every design decision optimizes for giving agents visible presence, mapping states to AI workflows, building trust through personality, and maintaining terminal-native simplicity.

**The Vision**: By v1.0, become "the Stripe of terminal agent faces" - the single source every AI-native CLI reaches for when their agent needs a face.

---

## [0.1.0-alpha.4] - 2025-10-15

Alpha.4 focuses on a simpler, production-ready workflow for both creators and developers.

### Highlights
- New Bubbletea TUI with split-pane live preview and auto-animate
- Simplified CLI: `create`, `browse`, `view` (admin kept, hidden; demo removed)
- State engine: minimum 3 frames per state; CLI overrides for `--fps` and `--loops`
- Centralized pattern codes; improved pattern compiler
- Library workflow: `tangent admin register` preserves frames per state
- API ready: `LibraryAgent()` with `Plan`, `Think`, `Execute` required
- Documentation reorg to `docs/` (API, STATES, PATTERNS, CLI, CONTRIBUTING_CHARACTERS)
- Release guard: release requires CHANGELOG update

### Breaking/Notable Changes
- Removed `tangent demo`; help simplified
- Admin commands retained but removed from help output

### Migration Notes
- Use `tangent browse` and `tangent view` instead of `tangent demo`
- Characters must include `plan`, `think`, `execute` with ≥3 frames each

## [0.1.0-alpha.3] - 2025-10-14

**API Contract Unification & Planet Series MVP**

This release unifies the API contract across CLI, documentation, and code, establishing a clean, consistent system for agent state management. Introduces the Planet Series character library with 8 professional agent characters.

### Major Changes

#### API Contract Unification
- **FIXED**: Unified state names across CLI and API (`plan`, `think`, `execute`)
- **FIXED**: Removed inconsistent "search" state, replaced with "execute"
- **FIXED**: All documentation now shows consistent state names
- **FIXED**: CLI creates characters that work perfectly with API
- **IMPROVED**: Perfect alignment between CLI creation and API usage

#### Planet Series Character Library
- **NEW**: 8 Planet Series characters: Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune
- **NEW**: Each character has distinct personality (efficient, friendly, balanced, etc.)
- **NEW**: All characters include required states: `plan`, `think`, `execute`
- **NEW**: Optional states: `wait`, `error`, `success`
- **NEW**: Professional, cohesive character library

#### Documentation Overhaul
- **UPDATED**: README.md with Planet Series characters and unified API
- **UPDATED**: QUICK_START_AGENT_STATES.md with correct state names
- **UPDATED**: NOTES.md with unified workflow and Planet Series info
- **UPDATED**: All examples use Planet Series characters
- **UPDATED**: Architecture diagrams reflect new system

#### Breaking Changes (Intentional)
- **REMOVED**: Old character references (alien, robot, pulse, wave, rocket) from documentation
- **CHANGED**: Required states from mixed naming to strict `plan`, `think`, `execute`
- **FOCUSED**: Planet Series only for v0.1.0 release (clean MVP boundaries)

### Technical Improvements
- **FIXED**: CLI validation functions check correct required states
- **FIXED**: Export validation uses unified state names
- **IMPROVED**: State descriptions consistent across all interfaces
- **IMPROVED**: Binary rebuilt with all changes

### Migration Notes
- **No backward compatibility needed** (no users yet)
- **Clean slate** for Planet Series development
- **Unified contract** ensures CLI-created characters work with API
- **Professional foundation** for future character development

### Next Steps
1. Design Mercury character (base + 3 states)
2. Design remaining 7 Planet Series characters
3. Implement character library files
4. Test complete workflow for each character
5. Release v0.1.0 (stable)

---

## [0.1.0-alpha.2] - 2025-10-13

**Agent State Architecture Release**

This release transforms Tangent into an AI agent-centric character library with state-based behavioral animations. Characters now represent agent states (planning, thinking, executing) rather than generic animations.

### Major Changes

#### Agent State System
- **NEW**: State-based character API with `AgentCharacter` wrapper
- **NEW**: Standard agent states: `Plan()`, `Think()`, `Execute()`, `Wait()`, `Error()`, `Success()`
- **NEW**: State inspection methods: `ListStates()`, `HasState()`, `GetStateDescription()`
- **NEW**: Character personality system (efficient, friendly, analytical, creative)
- **NEW**: Hybrid state model: 3 required states + optional standard + unlimited custom states

#### CLI Enhancements
- **NEW**: Personality selection during character creation
- **NEW**: Agent state guidance (shows required vs optional states)
- **NEW**: State progress tracking ("X required state(s) remaining")
- **NEW**: State type classification (standard vs custom)
- **NEW**: "Export for contribution (JSON)" menu option
- **IMPROVED**: Renamed "frames" to "agent states" in UI for clarity

#### Contribution Workflow
- **NEW**: JSON export for character contributions
- **NEW**: Automated contribution README generation
- **NEW**: GitHub PR template for character submissions
- **NEW**: CI/CD validation workflow for character contributions
- **NEW**: Complete contribution guide (`.github/CONTRIBUTING_CHARACTERS.md`)
- **NEW**: Character contribution directory with validation

#### Documentation
- **NEW**: Complete agent states guide (`docs/AGENT_STATES.md`)
- **NEW**: Quick start guide (`QUICK_START_AGENT_STATES.md`)
- **NEW**: Implementation summary (`IMPLEMENTATION_SUMMARY.md`)
- **NEW**: Test guide (`TEST_GUIDE.md`)
- **UPDATED**: README with agent state examples and migration guide

#### Testing & Examples
- **NEW**: Comprehensive agent state tests (11 tests, all passing)
- **NEW**: Full demo application (`examples/agent_states.go`)
- **NEW**: API test script (`test_agent_api.go`)
- **NEW**: 6 complete demo scenarios

#### Backward Compatibility
- **MAINTAINED**: Legacy `Library()` API still works (marked deprecated)
- **MAINTAINED**: Existing character structure preserved
- **MAINTAINED**: All existing functionality intact
- **NEW**: `LibraryAgent()` for new state-based API

### API Changes

#### New API (Recommended)
```go
agent, _ := characters.LibraryAgent("rocket")
agent.Plan(os.Stdout)      // Planning state
agent.Think(os.Stdout)     // Thinking state
agent.Execute(os.Stdout)   // Executing state
agent.Success(os.Stdout)   // Success state
```

#### Legacy API (Still Works)
```go
char, _ := characters.Library("alien")
characters.Animate(os.Stdout, char, 5, 3)
```

### Files Added (15 new files)
- `pkg/characters/agent.go` - Agent character wrapper with state API
- `pkg/characters/agent_test.go` - Comprehensive agent tests
- `examples/agent_states.go` - Full demo application
- `test_agent_api.go` - Quick API test script
- `docs/AGENT_STATES.md` - Complete agent states documentation
- `QUICK_START_AGENT_STATES.md` - Quick reference guide
- `IMPLEMENTATION_SUMMARY.md` - Full implementation details
- `TEST_GUIDE.md` - Testing instructions
- `.github/CONTRIBUTING_CHARACTERS.md` - Contribution guide
- `.github/PULL_REQUEST_TEMPLATE/character_contribution.md` - PR template
- `.github/workflows/character-validation.yml` - CI validation
- `characters/README.md` - Contribution directory README

### Files Modified (5 files)
- `pkg/characters/domain/character.go` - Added States map and Personality
- `pkg/characters/characters.go` - Added LibraryAgent() function
- `cmd/tangent/session.go` - Added Personality and StateType fields
- `cmd/tangent/main.go` - Enhanced CLI with agent state workflow
- `README.md` - Updated with agent state examples

### Migration Guide

**For Users**: Replace `Library()` with `LibraryAgent()` and use state methods instead of `Animate()`.

**For Contributors**: Use interactive CLI to create characters, export as JSON, submit via GitHub PR.

### Technical Details
- **State Model**: Characters have `States map[string]State` with named states
- **Required States**: Minimum 3 (plan, think, execute)
- **Optional States**: 3 standard (wait, error, success) + unlimited custom
- **Validation**: Automated CI checks for required states and valid JSON
- **Distribution**: Characters compiled into library after admin approval

### Breaking Changes
None - Full backward compatibility maintained.

### Known Limitations
- Existing library characters (alien, pulse, rocket, wave) not yet converted to agent states
- Will be addressed in future release

---

## [0.1.0-alpha.1] - 2025-01-XX

**Architectural Refactoring Release**

This release represents a complete architectural refactoring of Tangent, implementing proper layered architecture with enhanced error handling, comprehensive testing, and improved maintainability.

### Major Changes

#### Architecture Improvements
- **NEW**: Implemented proper layered architecture (Domain, Service, Infrastructure layers)
- **NEW**: Added comprehensive error handling with contextual suggestions
- **NEW**: Implemented Builder Pattern v2 with fluent API and validation
- **NEW**: Added centralized error handling system with user-friendly messages
- **NEW**: Enhanced pattern compilation with detailed validation
- **NEW**: Improved file repository with pre-validation and error reporting

#### Code Quality
- **NEW**: Comprehensive test coverage for all layers
- **NEW**: Mock implementations for testing
- **NEW**: Enhanced validation throughout the system
- **NEW**: Proper dependency injection and interface-driven design
- **NEW**: Clean separation of concerns

#### Documentation
- **CLEANED**: Removed 30+ unnecessary documentation files
- **SIMPLIFIED**: Consolidated to 5 essential documentation files
- **IMPROVED**: Clear, focused user experience
- **MAINTAINED**: One source of truth for each type of information

#### Backward Compatibility
- **MAINTAINED**: All existing functionality preserved
- **MAINTAINED**: Zero external dependencies for core functionality
- **MAINTAINED**: Single binary distribution
- **MAINTAINED**: Existing API compatibility

### Technical Details
- **Domain Layer**: Core business logic and entities
- **Service Layer**: Application logic and orchestration
- **Infrastructure Layer**: File persistence, pattern compilation, animation
- **CLI Layer**: User interface with enhanced error handling
- **Testing**: Comprehensive test coverage with proper mocking
- **Error Handling**: Sophisticated error system with context and suggestions

### Files Structure
```
pkg/characters/
├── domain/           # Core business logic
├── service/          # Application services
├── infrastructure/   # Concrete implementations
├── error_handler.go  # Centralized error handling
└── builder_v2.go     # Enhanced builder pattern
```

## [0.0.1] - 2025-10-07

**Initial Development History**

This version represents the complete development history of Tangent from initial concept through architectural refactoring.

### Development Phases Completed
- **Phase 1**: Domain Layer implementation with proper interfaces
- **Phase 2**: Service Layer with application logic and validation
- **Phase 3**: Infrastructure Layer with concrete implementations
- **Phase 4**: Enhanced error handling with contextual suggestions
- **Phase 5**: Main function integration with new architecture
- **Phase 6**: Comprehensive testing implementation
- **Phase 7**: Documentation cleanup and simplification

### Key Features Developed
- **Character Design System**: Pattern-based character creation
- **Visual Builder**: Interactive CLI character designer
- **Library System**: Pre-built characters (alien, pulse, wave, rocket)
- **Animation Engine**: Built-in terminal animation
- **Frame Extraction API**: TUI framework integration
- **Bubble Tea Adapter**: Seamless Bubble Tea integration
- **CLI Tools**: Interactive and non-interactive modes
- **Session Management**: Save/load character projects
- **Export System**: Generate Go code from designs

### Technical Achievements
- **Zero Dependencies**: Core functionality uses only Go stdlib
- **Layered Architecture**: Clean separation of concerns
- **Comprehensive Testing**: Full test coverage with mocks
- **Error Handling**: User-friendly error messages with suggestions
- **Documentation**: Clean, focused documentation structure
- **Backward Compatibility**: All existing functionality preserved

### Installation & Distribution
- **One-command installer**: `curl -sSL ... | bash`
- **Go module support**: Standard `go mod tidy` workflow
- **Multi-platform builds**: Automated via GoReleaser
- **GitHub releases**: Automated version management

This represents the complete evolution of Tangent from initial concept to production-ready character design system.
