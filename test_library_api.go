package main

import (
	"fmt"
	"os"
	"github.com/wildreason/tangent/pkg/characters"
)

func main() {
	fmt.Println("🎯 Testing Library-First API")
	fmt.Println()

	// Test 1: Get library character
	fmt.Println("1. Getting library character...")
	alien, err := characters.Library("alien")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✅ Got character: %s (%dx%d, %d frames)\n", 
		alien.Name, alien.Width, alien.Height, len(alien.Frames))

	// Test 2: Show static character
	fmt.Println("\n2. Showing static character...")
	characters.ShowIdle(os.Stdout, alien)
	fmt.Println("✅ Static display complete")

	// Test 3: Create custom character
	fmt.Println("\n3. Creating custom character...")
	robot, err := characters.NewBuilder("my-robot", 6, 3).
		Pattern("FRF").
		Block(6).
		NewFrame().
		Pattern("LRL").
		Block(6).
		NewFrame().
		Pattern("FRF").
		Block(6).
		Build()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	fmt.Printf("✅ Created custom character: %s (%dx%d, %d frames)\n", 
		robot.Name, robot.Width, robot.Height, len(robot.Frames))

	// Test 4: Show custom character
	fmt.Println("\n4. Showing custom character...")
	characters.ShowIdle(os.Stdout, robot)
	fmt.Println("✅ Custom character display complete")

	fmt.Println("\n🎉 All tests passed! Library-first API is working perfectly.")
}
