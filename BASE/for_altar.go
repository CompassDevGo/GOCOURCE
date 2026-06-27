package main

import "fmt"

func main() {
	altarLevel := 3

	fmt.Println("Строим магический алтарь")
	fmt.Println()

	for currentLevel := 1; currentLevel <= altarLevel; currentLevel++ {
		for spaceCount := 1; spaceCount <= altarLevel-currentLevel; spaceCount++ {
			fmt.Print(" ")
		}
		for stoneCount := 1; stoneCount <= 2*currentLevel-1; stoneCount++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for currentLevel := altarLevel - 1; currentLevel >= 1; currentLevel-- {
		for spaceCount := 1; spaceCount <= altarLevel-currentLevel; spaceCount++ {
			fmt.Print(" ")
		}
		for stoneCount := 1; stoneCount <= 2*currentLevel-1; stoneCount++ {
			fmt.Print("*")
		}

		fmt.Println()

	}

	fmt.Println()
	fmt.Println("Магический алтарь готов!")
}
