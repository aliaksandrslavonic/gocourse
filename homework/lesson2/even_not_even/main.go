package main

import "fmt"

func main() {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scanln(&number)
	fmt.Println(isEven(number))
	fmt.Println(division_by_three(number))
}

func isEven(num int) bool {
	return num%2 == 0
}

func division_by_three(number int) (answer string) {
	if (number % 3) == 0 {
		answer = "Делится без остатка на 3"
	} else {
		answer = "Не делится без остатка на 3"
	}
	return answer
}
