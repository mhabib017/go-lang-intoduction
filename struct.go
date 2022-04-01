package main

import "fmt"

type person struct{
	name string
	age int
}

func newPerson(name string, age int) *person{
	p:= person{name, age}
	return &p
}

 func main(){
	fmt.Println( person{"habib", 25})
	fmt.Println( person{name:"habib", age:25})
	fmt.Println( person{name:"habib"})
	fmt.Println( &person{name:"habib"})
	fmt.Println(newPerson("Habib", 25))
	p := person{name: "habib", age: 25}
    fmt.Println(p.name)
 }