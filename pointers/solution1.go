package main

import "fmt"

func main() {

	var x int = 10

	var ptr *int = &x

	fmt.Println("The value of x =", x)
	fmt.Println("The address of x =", ptr)
	fmt.Println("The value of ptr =", *ptr)
}
