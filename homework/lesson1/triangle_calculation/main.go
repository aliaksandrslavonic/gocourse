package main

import (
	"fmt"
	"math"
)

func main() {

	// Даны катеты прямоугольного треугольника.
	// Найти его площадь, периметр и гипотенузу.
	// Используйте тип данных float64 и функции из пакета math.

	var firstLeg, secondLeg float64 = 3, 5

	fmt.Println("Введите длинну катета a, по умолчанию a = ", firstLeg)
	fmt.Scanln(&firstLeg)

	fmt.Println("Введите длинну катета b, по умолчанию b = ", secondLeg)
	fmt.Scanln(&secondLeg)

	if (firstLeg > 0) && (secondLeg > 0) {
		area := (firstLeg * secondLeg) / 2
		fmt.Println("Площадь прямоугольного треугольника = ", fmt.Sprintf("%.2f", area))

		hypotenuse := math.Sqrt(firstLeg*firstLeg + secondLeg*secondLeg)
		fmt.Println("Гипотенуза = ", fmt.Sprintf("%.2f", hypotenuse))

		perimeter := firstLeg + secondLeg + hypotenuse
		fmt.Println("Переиметр = ", fmt.Sprintf("%.2f", perimeter))
	} else {
		fmt.Println("Введенные значения не могут быть отрицательными либо равны нулю")
	}

}
