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
		fmt.Print("Угадайте число (0..9): ")
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Ошибка ввода.")
			continue
		}

		attemptsCount++

		if guess < number {
			fmt.Println("Число больше")
		} else if guess > number {
			fmt.Println("Число меньше")
		}
	}
	if guess == number {
		fmt.Printf("Вы победили, угадав с %d попытки\n", attemptsCount)
	} else {
		fmt.Printf("Вы проиграли, было загадано число %d\n", number)
	}
}

func main() {
	GuessTheNumber()
}
