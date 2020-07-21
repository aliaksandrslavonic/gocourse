package main

import "fmt"

func main() {
	number := 98
	fibonacci(number)
}

func fibonacci(number int) {
	fmt.Println("Ряд Фибоначчи:")
	var f0, f1, fSum uint64 = 0, 1, 0
	fmt.Printf("f1 - %d\nf2 - %d\n", f0, f1)
	for i := 1; i <= number; i++ {
		fSum = f0 + f1
		f0 = f1
		f1 = fSum
		fmt.Printf("f%d - %d\n", i+2, fSum)
	}
}
