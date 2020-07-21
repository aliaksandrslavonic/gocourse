package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scanln(&number)
	fmt.Println(even_number(number))
	fmt.Println(division_by_three(number))
}

func even_number(number int) (answer string) {
	if (number % 2) == 0 {
		answer = "Четное"
	} else {
		answer = "Не четное"
	}
	return answer
}

func division_by_three(number int) (answer string) {
	if (number % 3) == 0 {
		answer = "Делится без остатка на 3"
	} else {
		answer = "Не делится без остатка на 3"
	}
	return answer
}
