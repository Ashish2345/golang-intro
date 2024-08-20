package main

import "fmt"

func main() {
	fmt.Println("Closures in Golang")

	//Anonymous function
	func(message string) string {
		fmt.Println("Hello in to the anonymous function")
		return message
	}("message")

	//
}
