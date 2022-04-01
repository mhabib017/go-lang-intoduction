package main

import "fmt"
func main(){
	nums := []int {1,24,34}
	for _, num := range nums {
		fmt.Println("Num", num)
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s =>%s\n",k,v)
	}

	for k := range kvs {
        fmt.Println("key:", k)
    }
	for i, c := range "go" {
        fmt.Println(i, c)
    }
}