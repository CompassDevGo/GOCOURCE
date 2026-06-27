/*
package main

import "fmt"

	func main() {
		for playerLevel := 1; playerLevel <= 5; playerLevel++ {
			fmt.Printf("Достигнут уровень: %d\n", playerLevel)
		}
	}
*/
/*package main

import "fmt"

func main() {
	enemyHealth := 100
	playerDamage := 25
	attackCount := 0
	for {
		attackCount++
		enemyHealth -= playerDamage

		fmt.Printf("Атака %d: осталось %d здоровья\n", attackCount, enemyHealth)

		if enemyHealth <= 0 {
			fmt.Println("Враг повержен!")
			break
		}
	}
}
*/

package main

import "fmt"

func main() {
	spellNames := []string{
		"Огненный шар",
		"Ледяная стрела",
		"Молния",
		"Исцеление",
	}
	fmt.Println("Доступные заклинания:")
	for _, spellNames := range spellNames {
		fmt.Printf(" %s\n", spellNames)
	}
}
