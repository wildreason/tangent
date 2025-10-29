# tangent-cli

Internal development tool for Tangent contributors. Not distributed.

## Build

```bash
make build-cli
# or
./scripts/build.sh
```

Output: `./tangent-cli` binary

## Install Locally

```bash
make install
```

Installs to: `~/.local/bin/tangent-cli`

## Commands

### browse

Preview characters and states:

```bash
tangent-cli browse              # List all characters
tangent-cli browse sa           # Browse sa character states
tangent-cli browse sa --state think
tangent-cli browse sa --state think --fps 10 --loops 3
```

### create

Interactive TUI for creating characters:

```bash
tangent-cli create
```

**Workflow**:
1. Name character
2. Set metadata (description, author, color, personality)
3. Set dimensions (width × height)
4. Create states with frames
5. Export to Go library code

**TUI Controls**:
- Enter: Submit line/advance
- Ctrl+D: Duplicate previous frame
- Ctrl+R: Paste line from previous frame
- Ctrl+F: Finish state (manual)
- Ctrl+C: Cancel

**Pattern Codes**: F, R, L, T, B, 1-8, _ (see library documentation)

### admin

Internal admin commands:

```bash
# Export character to JSON
tangent-cli admin export sa > sa.json

# Register single character from JSON
tangent-cli admin register character.json
tangent-cli admin register character.json --force

# Batch register multiple characters
tangent-cli admin batch-register template.json colors.json
```

**Batch register workflow**:
1. Create template.json (single character with color placeholder)
2. Create colors.json (array of color mappings)
3. Run batch-register to generate all characters

### version

```bash
tangent-cli version
```

Shows build info (version, commit, date).

## Development

**Build with version injection**:
```bash
VERSION=$(git describe --tags)
COMMIT=$(git rev-parse --short HEAD)
DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

go build -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE" \
  -o tangent-cli ./cmd/tangent-cli
```

**Tests**:
```bash
go test ./cmd/tangent-cli/...
```

## Files

- `cmd/tangent-cli/main.go` - Entry point, admin commands
- `cmd/tangent-cli/tui.go` - Bubbletea creation UI
- `cmd/tangent-cli/session.go` - Session management
- `cmd/tangent-cli/list.go` - Browse functionality
- `cmd/tangent-cli/view.go` - Preview helpers
- `cmd/tangent-cli/main_test.go` - Smoke tests
