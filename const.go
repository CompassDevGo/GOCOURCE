package main

import "fmt"

const (
	PI          = 3.14159
	earthRadius = 6371.0 // in kilometers
)
const (
	earthCircumference = 2 * PI * earthRadius
	earthArea          = 4 * PI * earthRadius * earthRadius
)

func main() {
	fmt.Println("Калькулятор круга")

	fmt.Printf("Число Пи = %.5f\n", PI)
	fmt.Printf("Радиус Земли = %.1f км\n", earthRadius)
	fmt.Printf("Длина окружности Земли = %.1f км\n", earthCircumference)
	fmt.Printf("Площадь Земли = %.1f км²\n", earthArea)

	customRadius := 10.0
	customArea := PI * customRadius * customRadius

	fmt.Printf("Радиус = %.1f км\n", customRadius)
	fmt.Printf("Площадь = %.1f км²\n", customArea)

}
