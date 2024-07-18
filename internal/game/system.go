package game

import (
    "math/rand"
    "time"
)

type System struct {
    Player *Player
}

func NewSystem() *System {
    return &System{
        Player: NewPlayer("System"),
    }
}

func (s *System) Bowl() rune {
    rand.Seed(time.Now().UnixNano())
    return rune('A' + rand.Intn(26))
}