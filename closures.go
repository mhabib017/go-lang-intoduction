package main

import "fmt"

func intSeq() func() int {
	i:=0
	return func() int {
        i++
        return i
    }
}

func main(){
	nextNumber := intSeq()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	newNumber := intSeq()

	fmt.Println(newNumber())
	fmt.Println(newNumber())
	fmt.Println(newNumber())
}