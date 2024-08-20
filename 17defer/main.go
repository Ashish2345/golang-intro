package main

import "fmt"

func main() {
	// Usages last in first out
	defer fmt.Println("World")
	defer fmt.Println("World2")
	defer fmt.Println("World3")
	fmt.Println("Hello")

}
