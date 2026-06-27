package main

import "fmt"

func main() {
	//Данные нашего игрока
	playerAge := 16
	hasAccount := true
	isOnline := false

	fmt.Println("=== Проверка игрока ===")
	fmt.Printf(" Возраст: %d, Аккаунт: %t, Онлайн: %t\n\n", playerAge, hasAccount, isOnline)
	//1. Оператор НЕ (!) - переворачивает значение на противоположное
	isOffline := !isOnline
	fmt.Println("1.Оператор НЕ (!)")
	fmt.Printf(" Игрок оффлайн: %t\n", isOffline)
	fmt.Println(" ОбъяснениеЖ !false = true (перевернули значение)\n")
	//2. Оператор ИЛИ (||) - возвращает true, если хотя бы одно из условий истинно
	canPlayFree := playerAge >= 18 || hasAccount
	fmt.Println("2.Оператор ИЛИ (||)")
	fmt.Printf(" Может играть бесплатно: %t\n", canPlayFree)
	fmt.Println(" Объяснение: возраст 16 (нет) ИЛИ есть аккаут (да) = true")
	fmt.Println("Достаточно одного 'да'для результата true\n")

	//3. Оператор И (&&) - возвращает true, если оба условия истинны
	canStartGame := hasAccount && !isOnline // есть аккаунт и не онлайн
	fmt.Println("4. Комбинация операторов")
	fmt.Printf(" Можно начать игру: %t\n", canStartGame)
	fmt.Println(" Объяснение: есть аккаунт (да) И не онлайн (да) = true")
	fmt.Println(" Сначала ! делает true из false, затем && проверяет оба условия\n")
	//4. Комбинация операторов
	canStartGame = hasAccount && isOnline // есть аккаунт и онлайн
	fmt.Println("4. Комбинация операторов")
	fmt.Printf(" Можно начать игру: %t\n", canStartGame)
	fmt.Println(" Объяснение: есть аккаунт (да) И онлайн (нет) = false")
	fmt.Println(" Сначала ! делает true из false, затем && проверяет оба условия\n")

	//Итоговое решение
	fmt.Println("=== Итоговое решение ===")
	if canPlayFree && !isOnline {
		fmt.Println("Игрок может играть, но сначала нужно войти в аккаунт")
	} else if !canPlayFree {
		fmt.Println("Нужен аккаунт илиивозраст должен быть 18+")
	} else {
		fmt.Println("Игрок может играть")
	}
}
