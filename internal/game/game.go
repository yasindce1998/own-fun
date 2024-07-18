package game

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

type GameMode int

const (
    SingleTeam GameMode = iota
    TwoTeam
)

type Game struct {
    Mode        GameMode
    Overs       int
    CurrentBat  *Player
    CurrentBowl *Player
    System      *System
    TotalBalls  int
    CurrentBall int
    IsOver      bool
}

func NewGame(mode GameMode, overs int) *Game {
    game := &Game{
        Mode:        mode,
        Overs:       overs,
        System:      NewSystem(),
        TotalBalls:  overs * 6,
        CurrentBall: 1,
        IsOver:      false,
    }
    return game
}

func (g *Game) Start() {
    fmt.Println("Welcome to English Cricket!")
    
    if g.Mode == SingleTeam {
        g.CurrentBat = NewPlayer("You")
        g.CurrentBowl = g.System.Player
    } else {
        // Implement two-player mode later
        fmt.Println("Two-player mode not implemented yet.")
        return
    }

    g.CoinToss()
}

func (g *Game) CoinToss() {
    fmt.Println("Choose heads or tails for coin toss")
    fmt.Println("1. Heads")
    fmt.Println("2. Tails")
    var choice int
    fmt.Scan(&choice)

    rand.Seed(time.Now().UnixNano())
    toss := rand.Intn(2) + 1

    if (choice == 1 && toss == 1) || (choice == 2 && toss == 2) {
        fmt.Println("You won the toss!")
        g.ChooseBatOrBowl()
    } else {
        fmt.Println("You lost the toss. System chooses to bowl.")
    }
}

func (g *Game) ChooseBatOrBowl() {
    fmt.Println("Choose bat or bowl")
    fmt.Println("1. Bat")
    fmt.Println("2. Bowl")
    var choice int
    fmt.Scan(&choice)

    if choice == 2 {
        g.CurrentBat, g.CurrentBowl = g.CurrentBowl, g.CurrentBat
    }
}

func (g *Game) NextBall() rune {
    if g.CurrentBall > g.TotalBalls {
        g.IsOver = true
        return 0
    }
    letter := g.CurrentBowl.Bowl()
    g.CurrentBall++
    return letter
}

func (g *Game) PlayBall(word string) (int, error) {
    if g.IsOver {
        return 0, fmt.Errorf("game is over")
    }
    
    score := len(word)
    g.CurrentBat.Score += score
    return score, nil
}