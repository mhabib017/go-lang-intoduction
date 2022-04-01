package main

import "fmt"

func add(a string,nums ...int)(int){
	fmt.Println("Type", a)
	sum:=0
	for _, num:= range nums {
		sum+=num
	}
	return sum
}


func main(){
	fmt.Println("Sum => ", add("Numbers",1,2,3,4,5,6))
	ar := []int{1,2,3,4,5,6,7}
	fmt.Println("Sum => ", add("Array",ar...))
}