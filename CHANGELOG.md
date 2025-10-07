# Changelog

## [0.1.0-beta.1] - 2025-10-07

**Strategic Pivot: Character Design System**

- **BREAKING REPOSITIONING**: Tangent is now a Character Design System, not just an animation library
- Add frame extraction API (GetFrames, Normalize, ToSpinnerFrames)
- Add Bubble Tea adapter package (pkg/adapters/bubbletea)
- Add complete Bubble Tea integration example
- Add Bubble Tea integration guide (docs/BUBBLETEA_INTEGRATION.md)
- Reposition README: framework-agnostic, "Two Ways to Use" pattern
- Built-in animation now "optional convenience" for simple CLIs
- Frame extraction is now first-class API

**New Dependencies:**
- github.com/charmbracelet/bubbletea (adapter only)
- github.com/charmbracelet/bubbles (adapter only)
- Core package remains dependency-free

## [0.1.0-alpha.5] - 2025-10-07

- **CRITICAL FIX**: Correct module path from `local/characters` to `github.com/wildreason/tangent`
- Fix `go install` command now works correctly
- Fix all import paths in examples and documentation
- Fix generated code uses correct module path

## [0.1.0-alpha.4] - 2025-10-07

- Add non-interactive CLI mode for AI agents (create, animate, export commands)
- Add comprehensive AI agents guide (AGENTS-README.md)
- Add 3 new library characters: pulse, wave, rocket
- Add `tangent gallery` command to browse all characters
- Shell/Python integration examples for agent workflows

## [0.1.0-alpha.3] - 2025-10-07

- Add one-line installer script for easy setup
- Simplify README (290 → 150 lines) with clear install → use workflow
- Remove confusing CLI vs Package sections
- Single install command provides both tangent CLI and Go package

## [0.1.0-alpha.2] - 2025-10-07

- Add Tangent CLI builder with interactive character designer
- Add live animation preview and multi-frame session management
- Add export options: terminal code or save to .go file
- Add duplicate frame feature and library system (alien character)
- Centralize pattern definitions with improved recall (1-8 quadrants)
- Add MIT License, contributing guidelines, and GoReleaser automation

## [0.0.1] - 2025-10-07

- Initial release with hex-style pattern system
- Character builder API and terminal animation engine
- Unicode Block Elements support (U+2580–U+259F)
