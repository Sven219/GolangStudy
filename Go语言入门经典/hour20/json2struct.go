package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Persion struct {
	Name    string   `json:"aaa"`
	Age     int      `json:"age`
	Hobbies []string `json:"habboes`
}

func main() {
	jsonStringData := `{"name":"George","Age":40,"Hobbies":["Cycling","Cheese","Techno"]}`
	jsonByteData := []byte(jsonStringData)
	p := Persion{}
	err := json.Unmarshal(jsonByteData, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}
