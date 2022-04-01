package main

import "fmt"

func fact(a int) int {
	if a == 0 {
		return 1
	}
	return a * fact(a-1)
}

func main() {

	fmt.Println("Factorial of 5", fact(5))

	var fib func(a int) int
	fib = func(a int) int {
		if a<2 {
			return a
		}
		return fib(a-1)+fib(a-2)
	}

	fmt.Println("Fib of 10", fib(10))
}