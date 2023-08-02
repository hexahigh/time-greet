package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	hour := currentTime.Hour()

	if hour < 12 && hour > 5 {
		fmt.Println("Good morning")
	} else if hour < 17 {
		fmt.Println("Good afternoon")
	} else if hour < 20 {
		fmt.Println("Good evening")
	} else if hour < 24 {
		fmt.Println("Good night")
	} else if hour < 5 {
		fmt.Println("You should go to bed")
	}
}
