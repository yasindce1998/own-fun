package game

import (
    "fmt"
    "math/rand"
    "time"
)

type GameMode int
type DifficultyLevel int

const (
    SingleTeam GameMode = iota
    TwoTeam
)

const (
    Easy DifficultyLevel = iota
    Medium
    Hard
)

type Game struct {
    Mode           GameMode
    Difficulty     DifficultyLevel
    Overs          int
    CurrentBat     *Player
    CurrentBowl    *Player
    Player1        *Player
    Player2        *Player
    System         *System
    TotalBalls     int
    CurrentBall    int
    IsOver         bool
    HighestWord    string
    TotalWordCount int
}

func NewGame(mode GameMode, overs int, difficulty DifficultyLevel) *Game {
    game := &Game{
        Mode:        mode,
        Difficulty:  difficulty,
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
        g.Player1 = NewPlayer("You")
        g.Player2 = g.System.Player
    } else {
        g.Player1 = NewPlayer("Player 1")
        g.Player2 = NewPlayer("Player 2")
    }

    g.CoinToss()
    
    // Initialize CurrentBat and CurrentBowl based on coin toss result
    if g.CurrentBat == nil {
        g.CurrentBat = g.Player2
        g.CurrentBowl = g.Player1
    }
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
        g.CurrentBat = g.Player1
        g.CurrentBowl = g.Player2
    }
}

func (g *Game) ChooseBatOrBowl() {
    fmt.Println("Choose bat or bowl")
    fmt.Println("1. Bat")
    fmt.Println("2. Bowl")
    var choice int
    fmt.Scan(&choice)

    if choice == 1 {
        g.CurrentBat = g.Player1
        g.CurrentBowl = g.Player2
    } else {
        g.CurrentBat = g.Player2
        g.CurrentBowl = g.Player1
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
    g.TotalWordCount++
    
    if len(word) > len(g.HighestWord) {
        g.HighestWord = word
    }
    
    return score, nil
}

func (g *Game) SwitchInnings() {
    g.CurrentBat, g.CurrentBowl = g.CurrentBowl, g.CurrentBat
    g.CurrentBall = 1
    g.IsOver = false
}

func (g *Game) GetAverageWordLength() float64 {
    if g.TotalWordCount == 0 {
        return 0
    }
    return float64(g.Player1.Score+g.Player2.Score) / float64(g.TotalWordCount)
}