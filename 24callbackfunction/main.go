package main

import "fmt"

func someCallback(connt string, callback func(string)) string {
	callback(connt)
	return "some"
}

func main() {
	someCallback("Hello", func(nm string) {
		fmt.Println("Niveone")
	})
}
