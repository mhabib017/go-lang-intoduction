package main

import "fmt"

func sumDiffMultiply(a,b int) (int ,int,int){
	return a+b, a-b, a*b
}

func main(){
	a , b, c := sumDiffMultiply(5,6)
	fmt.Println(a,b,c)
}