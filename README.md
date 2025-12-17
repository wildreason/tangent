# Tangent

Animated terminal avatars for AI agents. Go library.

## Install

```bash
go get github.com/wildreason/tangent/pkg/characters
```

## Usage

```go
import "github.com/wildreason/tangent/pkg/characters/client"

// Create animation controller
tc, _ := client.NewMicro("sam")
tc.Start()

// Get current frame (thread-safe)
frame := tc.GetFrameRaw()

// Change state
tc.SetState("write")
```

## Features

- 7 characters, 16 states, 4 themes
- Thread-safe TangentClient API
- O(1) frame access via pre-rendered cache
- Bubble Tea component included

## Documentation

- [API Reference](docs/API.md)
- [Architecture Decisions](decisions/)
- [Release Notes](RELEASE-NOTES.md)

## Current

**v0.4.0** - Dynamic noise for extreme motion (Wall Street Rush)

---

Part of [WildReason](https://wildreason.com) agent infrastructure.

Not accepting contributions.
