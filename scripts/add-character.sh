#!/bin/bash
# Helper script for admins to add approved contributions to the library

CHARACTER_NAME=$1
JSON_FILE=$2

if [ -z "$CHARACTER_NAME" ] || [ -z "$JSON_FILE" ]; then
    echo "Usage: $0 <character_name> <json_file>"
    echo "Example: $0 mercury mercury.json"
    exit 1
fi

if [ ! -f "$JSON_FILE" ]; then
    echo "Error: JSON file '$JSON_FILE' not found"
    exit 1
fi

echo "Adding character: $CHARACTER_NAME"
echo "From JSON file: $JSON_FILE"

# Validate JSON has required fields
if ! jq -e '.name and .personality and .width and .height and .base_frame and .states' "$JSON_FILE" > /dev/null; then
    echo "Error: JSON file missing required fields"
    exit 1
fi

# Check if character already exists
LIBRARY_FILE="pkg/characters/library/${CHARACTER_NAME}.go"
if [ -f "$LIBRARY_FILE" ]; then
    echo "Error: Character '$CHARACTER_NAME' already exists in library"
    exit 1
fi

echo "Creating library file: $LIBRARY_FILE"

# Extract character data from JSON
NAME=$(jq -r '.name' "$JSON_FILE")
PERSONALITY=$(jq -r '.personality' "$JSON_FILE")
WIDTH=$(jq -r '.width' "$JSON_FILE")
HEIGHT=$(jq -r '.height' "$JSON_FILE")

# Create the library file
cat > "$LIBRARY_FILE" << EOF
package library

func init() {
	register(${CHARACTER_NAME}Character)
}

var ${CHARACTER_NAME}Character = LibraryCharacter{
	Name:        "$NAME",
	Personality: "$PERSONALITY",
	Width:       $WIDTH,
	Height:      $HEIGHT,
	Patterns: []Frame{
		// Base frame
		{
			Name: "base",
			Lines: [
EOF

# Add base frame
jq -r '.base_frame.lines[]' "$JSON_FILE" | while read -r line; do
    echo "				\"$line\"," >> "$LIBRARY_FILE"
done

echo "			],
		}," >> "$LIBRARY_FILE"

# Add states
jq -r '.states[] | .name' "$JSON_FILE" | while read -r state_name; do
    echo "		// $state_name state" >> "$LIBRARY_FILE"
    echo "		{" >> "$LIBRARY_FILE"
    echo "			Name: \"$state_name\"," >> "$LIBRARY_FILE"
    echo "			Lines: [" >> "$LIBRARY_FILE"
    
    # Get frames for this state
    jq -r --arg state "$state_name" '.states[] | select(.name == $state) | .frames[] | .lines[]' "$JSON_FILE" | while read -r line; do
        echo "				\"$line\"," >> "$LIBRARY_FILE"
    done
    
    echo "			]," >> "$LIBRARY_FILE"
    echo "		}," >> "$LIBRARY_FILE"
done

echo "	}," >> "$LIBRARY_FILE"
echo "}" >> "$LIBRARY_FILE"

echo "✅ Character '$CHARACTER_NAME' added to library"
echo "📁 File created: $LIBRARY_FILE"
echo ""
echo "Next steps:"
echo "1. Review the generated file"
echo "2. Test the character: go run examples/test_character.go"
echo "3. Commit and push changes"
echo "4. Character will be available in next release"
