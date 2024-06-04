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
	var toss string
	fmt.Println("Choose heads or tails:\n 1.heads\n 2.tails")
	fmt.Scanln(&toss)
	if Toss() == toss {
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

func ChooseBatOrBowl() {
	
}