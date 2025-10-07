# Library Characters

Pre-built characters you can use immediately without creating patterns.

## Usage

### Quick Start

```go
package main

import (
    "os"
    "local/characters/pkg/characters"
)

func main() {
    // Load a library character
    alien, _ := characters.Library("alien")
    
    // Animate it
    characters.Animate(os.Stdout, alien, 4, 2)
}
```

### Discovery

```go
// List all available library characters
for _, name := range characters.ListLibrary() {
    info, _ := characters.LibraryInfo(name)
    fmt.Printf("%s: %s\n", name, info)
}
```

### Registry Integration

```go
// Load and register a library character
characters.UseLibrary("alien")

// Now retrieve it from registry
alien, _ := characters.Get("alien")
```

## Available Characters

### alien

**Description:** Animated alien with waving hands - three-frame idle animation

**Dimensions:** 7x3  
**Frames:** 3  
**Recommended FPS:** 5

**Animation:** Friendly alien mascot with alternating hand waves

**Try it:**
```bash
tangent animate --name alien --fps 5 --loops 3
```

---

### pulse

**Description:** Heartbeat/thinking indicator with expanding pulse effect

**Dimensions:** 9x5  
**Frames:** 3  
**Recommended FPS:** 8

**Animation:** Expanding and contracting heartbeat, perfect for AI processing indicators

**Try it:**
```bash
tangent animate --name pulse --fps 8 --loops 3
```

**Use cases:**
- AI thinking indicators
- Loading animations
- Background processing
- Heartbeat monitors

---

### wave

**Description:** Friendly greeting bot with waving hands

**Dimensions:** 11x5  
**Frames:** 5  
**Recommended FPS:** 6

**Animation:** Bot character waving hello with alternating left and right hands

**Try it:**
```bash
tangent animate --name wave --fps 6 --loops 3
```

**Use cases:**
- Welcome screens
- Greeting animations
- Onboarding sequences
- Friendly AI interactions

---

### rocket

**Description:** Launch sequence animation for deployments

**Dimensions:** 7x7  
**Frames:** 4  
**Recommended FPS:** 5

**Animation:** Rocket ignition, liftoff, and flight sequence

**Try it:**
```bash
tangent animate --name rocket --fps 5 --loops 3
```

**Use cases:**
- CI/CD pipelines
- Deployment scripts
- Build celebrations
- Launch sequences

## API

### Library Functions

```go
// Library(name) - Load and build a library character
char, err := characters.Library("alien")

// ListLibrary() - Get all library character names
names := characters.ListLibrary()  // ["alien"]

// LibraryInfo(name) - Get character description
info, err := characters.LibraryInfo("alien")

// UseLibrary(name) - Load and register in one step
char, err := characters.UseLibrary("alien")
```

## How It Works

Library characters are defined using the same hex-style pattern system:

```go
// pkg/characters/library/alien.go
var alienCharacter = LibraryCharacter{
    Name:        "alien",
    Description: "Animated alien with waving hands",
    Author:      "Wildreason, Inc",
    Width:       11,
    Height:      3,
    Patterns: []Frame{
        {
            Name: "idle",
            Lines: []string{
                "00R9FFF9L00",
                "0T9FFFFF7T0",
                "00011000220",
            },
        },
        // ... more frames
    },
}
```

This means library characters:
- ✓ Use the same pattern format as custom characters
- ✓ Are easy to maintain and extend
- ✓ Serve as examples for creating custom characters
- ✓ Work with all the same API functions (Animate, ShowIdle, etc.)

## Benefits

**For Beginners:**
- Start immediately without learning pattern syntax
- See working examples of animated characters
- Learn by examining library character patterns

**For Everyone:**
- Consistent, tested characters
- No setup required
- Same API as custom characters
- Mix library and custom characters freely

## Example

Complete example showing library usage:

```go
package main

import (
    "fmt"
    "os"
    "local/characters/pkg/characters"
)

func main() {
    // Discover what's available
    fmt.Println("Library Characters:")
    for _, name := range characters.ListLibrary() {
        info, _ := characters.LibraryInfo(name)
        fmt.Printf("  • %s - %s\n", name, info)
    }

    // Load and use
    alien, _ := characters.Library("alien")
    
    // Same API as custom characters
    characters.ShowIdle(os.Stdout, alien)
    characters.Animate(os.Stdout, alien, 4, 2)
}
```

## Future Characters

The library can be extended with more characters:
- Loading spinners
- Progress indicators
- Emoji-style faces
- Simple icons
- Animated symbols

All using the same pattern-based system!

