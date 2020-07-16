package main

import (
	"fmt"
)

func main() {

	// Написать программу для конвертации рублей в доллары.
	// Программа запрашивает сумму в рублях и выводит сумму в долларах.
	// Курс доллара задайте константой.

	const usd = 70.74
	var rub float32

	fmt.Println("Введите сумму в rub, которую вы хотите конвертировать: ")
	fmt.Scanln(&rub)
	if rub > 0 {
		fmt.Println("Курс USD = ", usd)
		convertationResult := rub / usd
		fmt.Println("Итого $ ", fmt.Sprintf("%.2f", convertationResult))
	} else {
		fmt.Println("Значение не может быть отрецательным либо ноль")
	}

}
