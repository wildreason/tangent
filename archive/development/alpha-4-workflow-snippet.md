(venv) ~/wild/characters $ ./tangent create
╔══════════════════════════════════════════╗
║  TANGENT - Terminal Agent Designer      ║
║  Design characters for your CLI agents  ║
║  v0.1.0-41-gc5ce809-dirty                 ║
╚══════════════════════════════════════════╝


╔══════════════════════════════════════════╗
║  CREATE NEW CHARACTER                    ║
╚══════════════════════════════════════════╝

◢ Character name: water5
◢ Enter width (e.g., 11): 2
◢ Enter height (e.g., 3): 2

✓ Creating character 'water5' (2x2)
✓ Character 'water5' is starting

(venv) ~/wild/characters $ ./tangent admin register water5.json
Registering character from water5.json...
✅ Character 'water5' registered successfully!
📁 Library file: pkg/characters/library/water5.go

Next steps:
1. Run: make build
2. Test: tangent gallery
3. Commit the changes
(venv) ~/wild/characters $ make build
Building Tangent...
Building Tangent...
  Version: v0.1.0-41-gc5ce809-dirty
  Commit:  c5ce809
  Date:    2025-10-15T13:40:20Z
✓ Built: tangent
  Run: ./tangent version
(venv) ~/wild/characters $ ./tangent browse
Available Agents:

  • demo4
  • mercury
  • water
  • water5

Total: 4 agents available

View agent: tangent browse <name>
(venv) ~/wild/characters $ ./tangent list
Error: unknown command 'list'

Tangent - Terminal Agent Designer

USAGE:
  tangent create                    Start interactive character builder
  tangent browse [name] [options]   List agents or view specific agent
  tangent demo <name> [options]     Animate character for testing
  tangent view [--session|--json]   View WIP character without register
  tangent admin <command>           Admin commands
  tangent version                   Show version information
  tangent help                      Show this help message

BROWSE OPTIONS:
  --state <name>                    Animate specific state (plan|think|execute)
  --fps <N>                         Override animation FPS
  --loops <N>                       Override animation loops

DEMO OPTIONS:
  --state <name>                    Animate specific state (plan|think|execute)
  --fps <N>                         Override animation FPS
  --loops <N>                       Override animation loops

VIEW OPTIONS:
  --session <name>                  Load saved session and preview it
  --json <file>                     Load a contribution JSON and preview it
  --state <name>                    Animate specific state (plan|think|execute)
  --fps <N>                         Override animation FPS
  --loops <N>                       Override animation loops

ADMIN COMMANDS:
  tangent admin register <json>     Register character to library
  tangent admin validate <json>     Validate character JSON

EXAMPLES:
  # Create character (interactive)
  tangent create

  # List all agents
  tangent browse

  # View specific agent
  tangent browse alex
  tangent browse alex --state plan
  tangent browse alex --fps 10 --loops 2

  # Test character animations (alternative)
  tangent demo alex
  tangent demo alex --state plan

  # View WIP character (no admin)
  tangent view --session mercury --state plan --fps 8 --loops 2
  tangent view --json mercury.json --state think --fps 6 --loops 3

  # Admin: Register character
  tangent admin register alex.json

For full documentation: https://github.com/wildreason/tangent
(venv) ~/wild/characters $ ./tangent view
Available Sessions:

  • demo-2
  • demo-3
  • demo-4
  • demo3
  • dusk
  • egon
  • egon2
  • egon3
  • er
  • mercury
  • mercuy
  • mira
  • monsterra
  • neptune
  • new
  • paper
  • paris
  • pods
  • rocket
  • ski
  • test-bot-2
  • test-bot
  • test-robot
  • test-test
  • test
  • test2
  • tokyo
  • tui
  • venus
  • water
  • water2
  • water3
  • water5

Preview a session: tangent view --session <name> --state plan --fps 8 --loops 2
(venv) ~/wild/characters $ ./tangent view water5
Available Sessions:

  • demo-2
  • demo-3
  • demo-4
  • demo3
  • dusk
  • egon
  • egon2
  • egon3
  • er
  • mercury
  • mercuy
  • mira
  • monsterra
  • neptune
  • new
  • paper
  • paris
  • pods
  • rocket
  • ski
  • test-bot-2
  • test-bot
  • test-robot
  • test-test
  • test
  • test2
  • tokyo
  • tui
  • venus
  • water
  • water2
  • water3
  • water5

Preview a session: tangent view --session <name> --state plan --fps 8 --loops 2
(venv) ~/wild/characters $ ./tanget view --session water5 --state plan --fps 8 --loops 2
zsh: no such file or directory: ./tanget
(venv) ~/wild/characters $ ./tangent view --session water5 --state plan --fps 8 --loops 2

Previewing 'plan' state for water5 (2x2) at 8 FPS for 2 loops

▐█
▛█

✓ View complete!

(venv) ~/wild/characters $
