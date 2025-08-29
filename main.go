package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GuessTheNumber() {
	rand.Seed(time.Now().UnixNano())
	count := 0
	guess := -1
	number := rand.Intn(10)

	for count < 6 && guess != number {
		fmt.Print("Guess the number (0..9): ")
		fmt.Scanln(&guess)

		if guess != number {
			if guess < number {
				fmt.Println("Your number is less")
			} else {
				fmt.Println("Your number is bigger")
			}
			count++
		}
		if guess == number {
			fmt.Println("You Won!")
		} else {
			fmt.Println("Your lose")
		}
	}
}

func main() {
	GuessTheNumber()
}
