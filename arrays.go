package main

import "fmt"
func main(){
	var a [5]int
	a[3] = 100
	fmt.Println("Array a:", a)
	b:= []int {1,3,4,5}
	fmt.Println("Array b", b, len(b))

	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
	fmt.Println("2d: ", twoD)
	twoD2 :=[2][3]int{{1,2,3},{4,5,6}}
   
	fmt.Println("2d2: ", twoD2)
}