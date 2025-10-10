#!/bin/bash
# Test script for Phase 5 CLI updates

echo "Testing Phase 5 CLI updates..."

# Test 1: Create a character with proper input
echo "Test 1: Creating a character"
{
    echo "1"           # Create new character
    echo "test-char"   # Character name
    echo "8"           # Width
    echo "3"           # Height
    echo "9"           # Back to main menu
    echo "6"           # Exit
} | timeout 10 ./tangent

echo -e "\nTest 1 completed\n"

# Test 2: Load character (should show empty list)
echo "Test 2: Loading character (empty list)"
{
    echo "2"           # Load character
    echo "6"           # Exit
} | timeout 10 ./tangent

echo -e "\nTest 2 completed\n"

# Test 3: Library preview
echo "Test 3: Library preview"
{
    echo "4"           # Preview library character
    echo "alien"       # Character name
    echo "6"           # Exit
} | timeout 10 ./tangent

echo -e "\nTest 3 completed\n"

echo "All tests completed!"
