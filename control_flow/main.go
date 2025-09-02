package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("Please input your number")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Printf("%d is even number", number)
	} else {
		fmt.Printf("%d is odd number", number)
	}
}
