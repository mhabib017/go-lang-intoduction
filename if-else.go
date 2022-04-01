package main

import "fmt" 
func main(){
	i:=4
	if i <3 {
		fmt.Println("value is under three")
	}
	if(i>3) {
		fmt.Println("Value is above three")
	}

	if j:=40; j<5 {
		fmt.Println("number is below five")
	} else if k:=10; k<j {
		fmt.Println("k is less than j")
	}else {
		fmt.Println("number is above five")
	}

}