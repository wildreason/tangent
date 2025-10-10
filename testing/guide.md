 ---
  Tangent v0.1.0 Beta Testing Guide

  Installation

  Install Command:
  curl -sSL https://raw.githubusercontent.com/wildreason/tangent/main/install.sh | bash

  Alternative (via Go):
  go install github.com/wildreason/tangent/cmd/tangent@v0.1.0

  Verify Installation:
  tangent --version
  # Should show: tangent 0.1.0

  ---
  Quick Start Testing

  1. Launch Interactive Builder
  tangent
  Test:
  - Navigation between frames
  - Pattern editing
  - Live preview
  - Duplicate frame feature
  - Save to file

  2. Test CLI Commands
  # Browse character library
  tangent gallery

  # Create character non-interactively
  tangent create my-character --width 11 --height 3

  # Animate existing character
  tangent animate examples/alien_spec/main.go

  # Export character frames
  tangent export my-character --format json

  3. Test Library Characters
  # Test built-in characters: alien, pulse, wave, rocket
  tangent animate --character alien
  tangent animate --character pulse
  tangent animate --character wave
  tangent animate --character rocket

  ---
  What to Test

  Core Functionality

  - Installation completes without errors
  - tangent --version shows correct version
  - Interactive builder launches and displays correctly
  - Pattern codes work (F, T, B, L, R, 1-9, etc.)
  - Animation preview plays smoothly
  - Save to .go file generates valid code

  Platform-Specific

  - Test on your OS (darwin/linux/windows)
  - Terminal colors render correctly
  - Keyboard navigation works
  - PATH configuration works

  Integration Testing

  - Generated Go code compiles successfully
  - Can import github.com/wildreason/tangent/pkg/characters in Go projects
  - Bubble Tea integration works (if testing framework integration)

  ---
  Known Issues to Watch For

  1. Installation: Should NOT show "Failed to download" error
  2. Frame Consistency: Wave character should animate without jitter
  3. PATH: Installer warns if ~/.local/bin not in PATH
  4. Colors: ANSI colors should display correctly in terminal

  ---
  Reporting Issues

  Include in Bug Reports:
  1. OS and architecture: uname -a
  2. Go version: go version
  3. Tangent version: tangent --version
  4. Full error output or screenshot
  5. Steps to reproduce

  Report Issues At:
  https://github.com/wildreason/tangent/issues

  Label Priority:
  - Critical: Installation fails, app crashes
  - High: Core features broken
  - Medium: UI issues, missing features
  - Low: Documentation, minor bugs

  ---
  Success Criteria

  Installation successful if:
  - ✓ Installer completes without errors
  - ✓ tangent --version returns v0.1.0
  - ✓ Interactive builder launches
  - ✓ Can create and animate characters
  - ✓ Generated code compiles and runs

  ---
  Additional Resources

  - README: https://github.com/wildreason/tangent
  - Bubble Tea Guide: https://github.com/wildreason/tangent/blob/main/docs/BUBBLETEA_INTEGRATION.md
  - AI Agents Guide: https://github.com/wildreason/tangent/blob/main/AGENTS-README.md
  - Changelog: https://github.com/wildreason/tangent/blob/main/CHANGELOG.md

  ---
  Expected Test Duration

  - Basic installation + quick test: 5-10 minutes
  - Full feature testing: 30-45 minutes
  - Integration testing: 1-2 hours

  ---
  Feedback Welcome

  Beyond bug reports, we'd love to hear:
  - First impressions and UX feedback
  - Missing features or use cases
  - Documentation clarity
  - Performance observations
  - Feature requests

  Note: This is the first stable release (v0.1.0), so thorough testing is appreciated!

