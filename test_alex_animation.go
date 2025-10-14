package main

import (
    "os"
    "time"
    "github.com/wildreason/tangent/pkg/characters"
)

func main() {
    agent, _ := characters.LibraryAgent("alex")
    
    println("🔹 Base Character:")
    agent.ShowBase(os.Stdout)
    println()
    
    println("🔹 Plan State Animation:")
    agent.AnimateState(os.Stdout, "plan", 3, 2)
    time.Sleep(500 * time.Millisecond)
    
    println("🔹 Think State Animation:")
    agent.AnimateState(os.Stdout, "think", 4, 1)
    time.Sleep(500 * time.Millisecond)
    
    println("🔹 Execute State Animation:")
    agent.AnimateState(os.Stdout, "execute", 6, 3)
}
