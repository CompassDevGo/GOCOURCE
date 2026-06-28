package main

import "fmt"

func main() {
	fmt.Println("Игра: сбор сокровищ в подземелье")
	fmt.Println()
	chests := []string{"Пусто", "Золото", "Ловушка", "Зелье", "Пусто", "Алмаз", "Ловушка"}

	health := 3
	treasures := 0
	targetTreasures := 3
	for health > 0 && treasures < targetTreasures {
		fmt.Printf("Здоровье: %d | Сокровища: %d/%d\n", health, treasures, targetTreasures)
		fmt.Println("Открываем сундуки...")
		for index, item := range chests {
			fmt.Printf("Сундук %d: ", index+1)
			if item == "Пусто" {
				fmt.Println("Пусто, идем дальше")
				continue
			}
			if item == "Ловушка" {
				fmt.Println("ЛОВУШКА! Потеряно 1 здоровье")
				health--
				if health <= 0 {
					fmt.Println("Здоровье кончилось!")
					break
				}
				continue
			}
			fmt.Printf("Найдено: %s!\n", item)
			treasures++
			if treasures >= targetTreasures {
				fmt.Println("Собрано достаточно сокровищ!")
				break
			}
		}
		fmt.Println("--- ИТОГ ИГРЫ ---")
		if treasures >= targetTreasures {
			fmt.Println("Победа! Вы собрали все сокровища!")
		} else {
			fmt.Println("Поражение! Здоровье кончилось...")
		}
	}
}
