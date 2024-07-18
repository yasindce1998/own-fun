package main

import (
    "fmt"
    "os"

    "github.com/yasindce1998/english-cricket/internal/game"
    "github.com/yasindce1998/english-cricket/internal/ui"
)

func main() {
    fmt.Println("Welcome to English Cricket!")
    fmt.Println("Choose your game mode:")
    fmt.Println("1. Single Team (You vs. System)")
    fmt.Println("2. Two Team (You vs. Human)")

    var mode int
    fmt.Scan(&mode)

    if mode != 1 && mode != 2 {
        fmt.Println("Invalid choice. Exiting.")
        os.Exit(1)
    }

    fmt.Println("Select number of overs:")
    fmt.Println("1. One")
    fmt.Println("2. Two")
    fmt.Println("3. Three")

    var overs int
    fmt.Scan(&overs)

    if overs < 1 || overs > 3 {
        fmt.Println("Invalid choice. Exiting.")
        os.Exit(1)
    }

    fmt.Println("Select difficulty level:")
    fmt.Println("1. Easy")
    fmt.Println("2. Medium")
    fmt.Println("3. Hard")

    var difficultyChoice int
    fmt.Scan(&difficultyChoice)

    var difficulty game.DifficultyLevel
    switch difficultyChoice {
    case 1:
        difficulty = game.Easy
    case 2:
        difficulty = game.Medium
    case 3:
        difficulty = game.Hard
    default:
        fmt.Println("Invalid choice. Exiting.")
        os.Exit(1)
    }

    g := game.NewGame(game.GameMode(mode-1), overs, difficulty)
    g.Start() // This will handle the coin toss
    ui.RunGame(g)
}