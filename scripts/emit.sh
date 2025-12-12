#!/bin/bash
# Interactive state emitter for testing tangent avatar animations
# Sends state events to murmur socket to test jump/tview integration

# Default socket - override with MURMUR_SOCKET env var
SOCKET="${MURMUR_SOCKET:-$HOME/.murmur/projects/-Users-btsznh-wild-characters/agents/sam/room-1/socket.sock}"
SCRIPTS_DIR="${EMIT_SCRIPTS:-$HOME/.murmur/emit-scripts}"
mkdir -p "$SCRIPTS_DIR"

# Available states (maps to tangent animations)
# resting, arise, wait, read, write, search, approval

emit() {
    local state="$1"
    local msg="${2:-$state}"
    local event="{\"type\":\"state\",\"state\":\"$state\",\"message\":\"$msg\",\"room\":1}"
    echo "$event" | nc -U "$SOCKET" 2>/dev/null
    echo "-> $state"
}

expand_shortcut() {
    case "$1" in
        # Direct tangent states
        x) echo "resting" ;;
        a) echo "arise" ;;
        W) echo "wait" ;;
        R) echo "read" ;;
        w) echo "write" ;;
        S) echo "search" ;;
        A) echo "approval" ;;
        # Murmur states (test alias mapping)
        e) echo "edit" ;;
        b) echo "bash" ;;
        g) echo "grep" ;;
        t) echo "think" ;;
        s) echo "success" ;;
        f) echo "find" ;;
        F) echo "failed" ;;
        *) echo "$1" ;;
    esac
}

run_script() {
    local script_file="$SCRIPTS_DIR/$1"
    if [ ! -f "$script_file" ]; then
        echo "Script not found: $script_file"
        return 1
    fi

    echo "Running: $1"
    local content=$(cat "$script_file")
    IFS=',' read -ra STEPS <<< "$content"

    for step in "${STEPS[@]}"; do
        local state=$(echo "$step" | cut -d: -f1)
        local duration=$(echo "$step" | cut -d: -f2)
        state=$(expand_shortcut "$state")
        emit "$state"
        sleep "$duration"
    done
    echo "Done."
}

list_scripts() {
    echo "Saved scripts:"
    ls -1 "$SCRIPTS_DIR" 2>/dev/null | while read f; do
        echo "  $f: $(cat "$SCRIPTS_DIR/$f")"
    done
}

# Interactive script builder
build_script() {
    echo ""
    echo "Building script - add steps one at a time"
    echo "Format: <state> <seconds>  (e.g., 'e 2' for edit:2)"
    echo "Type 'done' when finished, 'cancel' to abort"
    echo ""

    local steps=""
    local count=0

    while true; do
        read -p "[$count] state seconds: " state_input sec_input

        if [ "$state_input" = "done" ]; then
            if [ -z "$steps" ]; then
                echo "No steps added. Cancelled."
                return
            fi
            # Remove leading comma
            steps="${steps:1}"

            read -p "Script name: " script_name
            if [ -n "$script_name" ]; then
                echo "$steps" > "$SCRIPTS_DIR/$script_name"
                echo "Saved: $script_name = $steps"
            fi
            return
        fi

        if [ "$state_input" = "cancel" ]; then
            echo "Cancelled."
            return
        fi

        if [ -z "$state_input" ] || [ -z "$sec_input" ]; then
            echo "  Need both state and seconds"
            continue
        fi

        local full_state=$(expand_shortcut "$state_input")
        steps="$steps,$full_state:$sec_input"
        count=$((count + 1))
        echo "  Added: $full_state:$sec_input"
    done
}

show_help() {
    echo "Tangent Avatar State Emitter"
    echo ""
    echo "Usage: emit.sh [command] [args]"
    echo ""
    echo "Commands:"
    echo "  <state>       Emit single state (e.g., 'emit.sh bash')"
    echo "  run <name>    Run saved script"
    echo "  list          List saved scripts"
    echo "  (none)        Interactive mode"
    echo ""
    echo "Shortcuts (interactive mode):"
    echo "  Tangent states: x=resting a=arise W=wait R=read w=write S=search A=approval"
    echo "  Murmur states:  e=edit b=bash g=grep t=think s=success f=find F=failed"
    echo ""
    echo "Environment:"
    echo "  MURMUR_SOCKET  Socket path (default: ~/.murmur/projects/.../socket.sock)"
    echo "  EMIT_SCRIPTS   Scripts dir (default: ~/.murmur/emit-scripts)"
    echo ""
    echo "Current socket: $SOCKET"
}

# Command line mode
if [ "$1" = "run" ]; then
    run_script "$2"
    exit 0
fi

if [ "$1" = "list" ]; then
    list_scripts
    exit 0
fi

if [ "$1" = "-h" ] || [ "$1" = "--help" ]; then
    show_help
    exit 0
fi

if [ -n "$1" ]; then
    emit "$1" "$2"
    exit 0
fi

# Interactive mode
echo "Tangent State Emitter"
echo "Socket: $SOCKET"
echo ""
echo "Shortcuts: x=resting a=arise W=wait R=read w=write S=search A=approval"
echo "Murmur:    e=edit b=bash g=grep t=think s=success f=find F=failed"
echo "Commands:  new, run <name>, list, help, q"
echo ""

while true; do
    read -p "> " cmd arg

    case "$cmd" in
        q|quit) exit 0 ;;
        help) show_help ;;
        new) build_script ;;
        run)
            if [ -n "$arg" ]; then
                run_script "$arg"
            else
                echo "Usage: run <name>"
            fi
            ;;
        list) list_scripts ;;
        # Tangent states
        x) emit "resting" ;;
        a) emit "arise" ;;
        W) emit "wait" ;;
        R) emit "read" ;;
        w) emit "write" ;;
        S) emit "search" ;;
        A) emit "approval" ;;
        # Murmur states (test alias mapping)
        e) emit "edit" ;;
        b) emit "bash" ;;
        g) emit "grep" ;;
        t) emit "think" ;;
        s) emit "success" ;;
        f) emit "find" ;;
        F) emit "failed" ;;
        "") ;;
        *) emit "$cmd" ;;
    esac
done
