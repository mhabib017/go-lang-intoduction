package main

import "fmt"

func add(a int, b int)int {
	return a+b
}

func subtract(a , b int)int {
	return a-b
}

func main(){
	fmt.Println("Add 2+4 =>", add(2,4))
	fmt.Println("subtract 12+4 =>", subtract(12,4))
}