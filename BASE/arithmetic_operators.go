package main

import "fmt"

func main() {
	// Арифметические операторы
	//a := 10
	//b := 5

	//fmt.Println("=== Арифметические операторы ===")
	//fmt.Printf(" a: %d, b: %d\n", a, b)
	//fmt.Printf(" a + b: %d\n", a+b)
	//fmt.Printf(" a - b: %d\n", a-b)
	//fmt.Printf(" a * b: %d\n", a*b)
	//fmt.Printf(" a / b: %d\n", a/b)
	//fmt.Printf(" a %% b: %d\n", a%b)

	var playerScore, enemyScore int = 150, 75
	var battleResult int
	battleResult = playerScore + enemyScore
	fmt.Println(" Общие очки в битве:", battleResult)
	battleResult = playerScore - enemyScore
	fmt.Println(" Преимущество игрока:", battleResult)
	battleResult = playerScore * 2
	fmt.Println(" Удвоенные очки игрока:", battleResult)
	battleResult = playerScore / 2
	fmt.Println(" Половина очков игрока:", battleResult)
	battleResult = playerScore % enemyScore
	fmt.Println(" Остаток при делении очков:", battleResult)

	var gameLevel int = 5
	var experiencePoints int = 100
	var bonusMultiplier int = 3
	totalExperience := gameLevel + experiencePoints*bonusMultiplier
	fmt.Println(" Без скобок:", totalExperience)
	totalExperienceWithParentheses := (gameLevel + experiencePoints) * bonusMultiplier
	fmt.Println(" Со скобками:", totalExperienceWithParentheses)

	var totalCoinsInt int = 100
	var playersCountInt int = 3

	coinsPerPlayer := totalCoinsInt / playersCountInt
	fmt.Println(" Монет на игрока (целые числа):", coinsPerPlayer)
	var totalCoins float64 = 100.0
	var playersCount float64 = 3.0

	coinsPerPlayerFloat := totalCoins / playersCount
	fmt.Println(" Монет на игрока (точно):", coinsPerPlayerFloat)
	coinsPerPlayerMod := totalCoinsInt % playersCountInt
	fmt.Println(" Монет на игрока (после деления):", coinsPerPlayerMod)

}
