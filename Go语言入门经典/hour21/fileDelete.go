package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("file01.copy.txt")
	if err != nil {
		log.Fatal(err)
	}
}
