package main

import "fmt"

var maxLevel = 50

func createHero() {
	var mana int
	var name string = "Legolas"
	var exp = 2500

	level := 12
	gold := 750
	isOnline := true

	mana = 200

	exp = exp + 500
	gold -= 100
	level ++

	fmt.Println("=== Hero Profile ===")
	fmt.Printf("Name %s\n", name)
	fmt.Printf("Level: %d / %d\n", level, maxLevel)
	fmt.Printf("Experiance: %d\n", exp)
	fmt.Printf("Gold: %d\n", gold)
	fmt.Printf("Mana: %d\n", mana)
	fmt.Printf("Online: %t\n", isOnline)
}

func main () {
	createHero()
}