package main

import "fmt"

func main() {
	s := make([]int, 4)
	fmt.Println("Slice", s)
	s[0] = 1
	s[2] = 13
	fmt.Println("Slice", s)
	s = append(s, 34)
	fmt.Println("Slice", s)

	l := s[2:4]
	fmt.Println("Slice", l)
	m := s[:4]
	fmt.Println("Slice", m)
	n := s[3:]
	fmt.Println("Slice", n)
	
	
}