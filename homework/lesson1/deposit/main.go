package main

import (
	"fmt"
)

func main() {

	// Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет

	var deposit, interest, capitalDeposit float64

	fmt.Println("Введите сумму вклада")
	fmt.Scanln(&deposit)
	fmt.Println("Введите процент банковского вклада")
	fmt.Scanln(&interest)

	if (deposit > 0) && (interest > 0) && (interest < 100) {
		for i := 1; i <= 5; i++ {
			capitalDeposit = (capitalDeposit + deposit*interest/100)
			fmt.Printf("Год %d = ", i)
			fmt.Println(fmt.Sprintf("%.2f", deposit+capitalDeposit))
		}
	} else {
		fmt.Println("Введенные значения не валидны")
	}

}
