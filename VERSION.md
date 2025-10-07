# Version 0.0.1

**Release Date:** October 7, 2025  
**Commit:** 36faf8d  
**Tag:** v0.0.1

## Package Structure

```
characters/
├── README.md              # Simple, focused documentation
├── CHANGELOG.md           # Version history
├── PATTERN_GUIDE.md       # Detailed pattern syntax guide
├── .gitignore            # Excludes archive/
├── go.mod                # Module definition
├── block_elements.go     # Block element constants
├── pkg/characters/       # Core package
│   ├── pattern.go        # Hex-style pattern compiler
│   ├── spec.go           # Character specification
│   ├── builder.go        # Builder API
│   ├── animator.go       # Terminal animation
│   ├── registry.go       # Character registry
│   └── characters.go     # Package interface
├── examples/             # 7 working examples
│   ├── hex_style/        # Full hex-style demo
│   ├── compact/          # Minimal hex-style
│   ├── basic/            # Builder API
│   ├── simple/           # 2-frame animation
│   ├── one_line/         # Single line character
│   ├── alien_spec/       # Alien character demo
│   └── advanced/         # Multi-character
└── archive/              # Experimental files (gitignored)
```

## Core Features

1. **Hex-Style Patterns** - `"00R9FFF9L00"` (like hex colors)
2. **Pattern Compiler** - Single-character codes (F, T, B, L, R, etc.)
3. **Character Builder** - Fluent API with validation
4. **Animation Engine** - ANSI escape sequences for terminal
5. **Character Registry** - Global thread-safe storage
6. **Block Elements** - All 18 Unicode block elements (U+2580-U+259F)

## Usage

```go
alien := characters.NewCharacterSpec("alien", 11, 3).
    AddFrame("idle", []string{
        "00R9FFF9L00",
        "0T9FFFFF7T0",
        "00011000220",
    })

char, _ := alien.Build()
characters.Animate(os.Stdout, char, 4, 2)
```

## Pattern Codes

```
F=█  T=▀  B=▄  L=▌  R=▐  (Basic blocks)
7=▛  9=▜  6=▙  8=▟     (Three-quads)
1=▘  2=▝  3=▖  4=▗     (Single quads)
.=░  :=▒  #=▓          (Shading)
0=Space (or _ for readability)
```

## What's Archived

- Old cmd/ and internal/ directories (superseded by pkg/)
- Experimental examples (alien_easy, pattern_based, single_line)
- Test files (alien-*.txt, launch-viewer scripts)
- Planning documents (challenge.md, IMPLEMENTATION_SUMMARY.md)
- Unicode palette JSON (out of scope)

## Next Steps

This is a stable v0.0.1 release ready for builders to use. The package is focused, documented, and examples work correctly.

