package main

import (
	"fmt"
	"math/rand"
	"time"
)
var r *rand.Rand

func init() {
	// Initialize a new random source with the current time as the seed
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	if Toss() == "heads" {
		fmt.Println("You Won the Toss")
	} else {
		fmt.Println("You lose the Toss")
	}
	fmt.Println("Welcome to English cricket")
}

func Toss() (result string) {
	words := []string{"heads", "tails"}
	index := r.Intn(len(words))
	return words[index]
}