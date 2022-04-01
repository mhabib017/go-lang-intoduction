package main

import (
	"fmt"
)
func main(){
	i:=4
	switch i {
	case 1:
		fmt.Println("Case one")
		break
	case 2:
		fmt.Println("Case two")
		break
	default:
		fmt.Println("Default")
	}

	switch {
	case i<3:
		fmt.Println("value less than three")
		break
	case i>3:
		fmt.Println("value greater than three")
		break
	default:
		fmt.Println("Value equal to three")
	}
}