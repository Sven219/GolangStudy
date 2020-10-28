package main

import (
	"io/ioutil"
	"log"
)

func main() {
	b := make([]byte, 0)
	err := ioutil.WriteFile("file02.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
