package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Persion struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	hobbies := []string{"Cyclong", "Cheese", "Techno"}
	p := Persion{
		Name:    "George",
		Age:     40,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)
}
