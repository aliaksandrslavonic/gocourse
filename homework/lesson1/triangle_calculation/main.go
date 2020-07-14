package main

import (
	"fmt"
	"math"
)

func main() {

	// Даны катеты прямоугольного треугольника.
	// Найти его площадь, периметр и гипотенузу.
	// Используйте тип данных float64 и функции из пакета math.

	var a, b, c, S, P float64 = 3, 5, 0, 0, 0

	fmt.Println("Введите длинну катета a, по умолчанию a = ", a)
	fmt.Scanln(&a)

	fmt.Println("Введите длинну катета b, по умолчанию b = ", b)
	fmt.Scanln(&b)

	if (a > 0) && (b > 0) && (a != 0) && (b != 0) {
		S = (a * b) / 2
		fmt.Println("Площадь прямоугольного треугольника (S) = ", fmt.Sprintf("%.2f", S))

		c = math.Sqrt(a*a + b*b)
		fmt.Println("Гипотенуза = ", fmt.Sprintf("%.2f", c))

		P = a + b + c
		fmt.Println("Переиметр (P) = ", fmt.Sprintf("%.2f", P))
	} else {
		fmt.Println("Введенные значения не могут быть отрицательными либо равны нулю")
	}

}
