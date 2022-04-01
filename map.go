package main

import "fmt"
func main(){
	m:= make(map[string]int)
	fmt.Println("map : ", m)
	
	mp:= map[string]int{"one":1, "two":2}
	fmt.Println("map : ", mp)

	delete(mp, "two")
	fmt.Println("map :", mp)

	mpn:= map[int]string{1: "one"}
	fmt.Println("map : ", mpn)


}