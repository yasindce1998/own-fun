package game

import (
    "fmt"
)

type Player struct {
    Name  string
    Score int
}

func NewPlayer(name string) *Player {
    return &Player{
        Name:  name,
        Score: 0,
    }
}

func (p *Player) Bowl() rune {
    var letter string
    fmt.Printf("%s, enter a letter to bowl: ", p.Name)
    fmt.Scan(&letter)
    return rune(letter[0])
}