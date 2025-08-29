package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessTheNumber() {
	rand.Seed(time.Now().UnixNano())
	attemptsLimit := 5
	attemptsCount := 0
	guess := -1
	number := rand.Intn(10) + 1

	for attemptsCount < attemptsLimit && guess != number {
		fmt.Print("Guess the number (0..9): ")
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Input error.")
			continue
		}

		attemptsCount++

		if guess < number {
			fmt.Println("The number is bigger")
		} else if guess > number {
			fmt.Println("The number is less")
		}
	}
	if guess == number {
		fmt.Printf("You won by guessing on %d attempt\n", attemptsCount)
	} else {
		fmt.Printf("You lost, the number %d\n was guessed.", number)
	}
}

func main() {
	GuessTheNumber()
}
