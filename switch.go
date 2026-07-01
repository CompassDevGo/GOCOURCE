package main

import "fmt"

//Функция обработки события по приоритету
func processEventPriority(priority int) {
	fmt.Println("\n=== Обработка по приоритету")
	switch {
	case priority >= 3:
		fmt.Println("[!!!] Отправка экстренного уведомления")
		fallthrough
	case priority >= 2:
		fmt.Println("[!!] Запись в журнал важных событий")
		fallthrough
	case priority >= 1:
		fmt.Println("[!] Обновление статистики")
	default:
		fmt.Println("[] Событие проигнорировано")
	}

}

//Функция определения по категории типа данных
func categorizeData(data interface{}) {
	fmt.Println("\n=== Определение категории данных ===")

	switch typed := data.(type) {
	case int:
		switch {
		case typed < 0:
			fmt.Printf("Отрицательное значение: %d\n", typed)
		case typed >= 0 && typed <= 100:
			fmt.Printf("Процент выполнения: %d%%\n", typed)
		default:
			fmt.Printf("Счетчик %d единиц\n", typed)
		}
	case string:
		switch typed {
		case "start", "begin", "init":
			fmt.Println("Команда инициализации системы")
		case "stop", "end", "finish":
			fmt.Println("Команда завершения работы")
		default:
			fmt.Printf("Текстовое сообщение: \"%s\"\n", typed)
		}
	case float64:
		fmt.Printf("Коэффициент: %.3f\n", typed)
	case bool:
		if typed {
			fmt.Println("Флаг активирован: ДА")
		} else {
			fmt.Println("Флаг активирован: НЕТ")
		}
	default:
		fmt.Printf("Неподдерживаемый тип данных: %T\n", data)
	}
}

// Главная функция демонстрации
func main() {

	fmt.Println("=== Система обработки игровых событий ===")
	// Тест обработки приоритетов
	processEventPriority(3)
	processEventPriority(1)
	processEventPriority(0)

	// Тест категории данных
	categorizeData(75)
	categorizeData(-5)
	categorizeData(500)
	categorizeData("start")
	categorizeData("custom_message")
	categorizeData(true)
	categorizeData(false)

}
