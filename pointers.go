package main

import "fmt"

func changeVal(a int) {
	a = 0
}

func changeValPointer(a *int) {
	*a = 0
}

func main() {
	a := 10
	changeVal(a)
	fmt.Println("A =>", a)
	a = 10
	changeValPointer(&a)
	fmt.Println("A =>", a)
}