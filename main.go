package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//create a game that wil pick a random number between 1 and 100
//and then prompt the user to guess the number
//if the user guesses wrong, tell them if it's higher or lower
//input comes from stdin and output goes to stdout
//input validation, only allow user to guess a number, and only allow guesses between 1 and 100
//also, only allow 10 guesses

type game struct {
	target int
	guess  int
}

func main() {
	g := newGame()
	g.play()
}

func newGame() *game {
	return &game{
		target: rand.Intn(100) + 1,
	}
}

func (g *game) play() {
	for i := 0; i < 10; i++ {
		g.guess = g.prompt()
		if g.guess == g.target {
			fmt.Println("You win!")
			return
		} else if g.guess > g.target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
	}
	fmt.Println("You lose!")
}

func (g *game) prompt() int {
	for {
		fmt.Print("Guess a number: ")
		var input string
		fmt.Scanln(&input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		if guess < 1 || guess > 100 {
			fmt.Println("Invalid input")
			continue
		}
		return guess
	}
}
