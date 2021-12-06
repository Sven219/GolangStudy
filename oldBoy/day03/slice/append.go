package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := a1[:]

	s1 = append(s1[0:1], s1[2:]...)
	fmt.Println(s1, a1)
	s1 = append(s1, 1222)
	fmt.Println(s1, a1)

}
