package main

import "fmt"

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I am full! I have eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I am still hungry! I have eaten %d\n", eaten)
	return feedMe(portion, eaten)
}

func main() {
	fmt.Println(feedMe(1, 0))
}
