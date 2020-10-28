package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "hello world!"
	err := ioutil.WriteFile("file03.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
