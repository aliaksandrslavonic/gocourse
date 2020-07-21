package main

import "fmt"

func main() {
	number := 0
	fmt.Println("Введите число n до которого будет вестись поиск простых чисел: ")
	fmt.Scanln(&number)
	var isPrimeArray []int

	for i := 2; i <= number; i++ {
		if isprime(i) {
			isPrimeArray = append(isPrimeArray, i)
		}
	}

	for _, number := range isPrimeArray {
		fmt.Println(number)
	}
}

func isprime(n int) bool {
	if n == 1 {
		return false
	}

	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}
