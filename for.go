package main

import "fmt"
func main (){
	i:=1
	
	for i<=3 {
		fmt.Println("value of i", i)
		i++
	}
	fmt.Println("value of i after loop",i)

	for j:=1; j<5; j++ {
		fmt.Println("value of j",j)
	}

	for k:=1; k<5 ;k++ {
		if(k ==3 ){
			continue
		}
		fmt.Println("value of k ", k)
	}
}