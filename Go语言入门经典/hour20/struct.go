package main

import "fmt"

type Persion struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	hobbies := []string{"Cycling", "Cheese", "Techno"}
	p := Persion{
		Name:    "George",
		Age:     40,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)
}
