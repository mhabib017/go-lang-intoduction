package main

import "fmt"

type rect struct {
	with int
	height int
}

func (r *rect) area () int {
	return r.height * r.with
}

func (r *rect ) peri()int {
	return  2*r.height + 2* r.with
}

func main(){
	r:= rect{with: 4, height: 5}
	fmt.Println(r)
	fmt.Println("Area", r.area())
	fmt.Println("Peri", r.peri())
	 
	r2 := &r
	fmt.Println("Area", r2.area())
}