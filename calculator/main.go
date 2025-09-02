package main

import (
	"fmt"
)

func main() {
	var a int32
	var b int32

	fmt.Println("Please input your first number ")
	fmt.Scanln(&a)
	fmt.Println("Please input your second number ")
	fmt.Scanln(&b)
	fmt.Println("Here is the summ of those two numbers")
	fmt.Println(a+b)
}
