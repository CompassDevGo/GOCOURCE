package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)
	secretNumber := randomGenerator.Intn(12) + 1

	minRange := 1
	maxRange := 12

	fmt.Println("Добро пожаловать в игру 'Угадай число'!")
	fmt.Println("Я загадал число от 1 до 12.")
	fmt.Println("У вас есть 3 попытки, чтобы его угадать!")

	maxAttempts := 3

	var playerGuess int

	for attemp := 1; attemp <= maxAttempts; attemp++ {
		fmt.Printf("\nПопытка %d из %d: ", attemp, maxAttempts)
		fmt.Scan(&playerGuess)

		if playerGuess == secretNumber {
			possibleNumbers := []int{}

			for num := minRange; num <= maxRange; num++ {
				if num != playerGuess {
					possibleNumbers = append(possibleNumbers, num)
				}
			}
			if len(possibleNumbers) > 0 {
				randomIndex := randomGenerator.Intn(len(possibleNumbers))
				secretNumber = possibleNumbers[randomIndex]
			}
		}

		if playerGuess == secretNumber {
			fmt.Println("Поздравляю! Вы угадали число!")
			fmt.Printf("Загаданное число было: %d\n", secretNumber)
			return
		} else if playerGuess < secretNumber {
			fmt.Println("Загаданное число больше!")
			minRange = playerGuess + 1

		} else {
			fmt.Println("Загаданное число меньше!")
			maxRange = playerGuess - 1
		}

	}
	fmt.Println("Попытки закончились")
	fmt.Printf("Загаданное число было %d\n", secretNumber)

}
