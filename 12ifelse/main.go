package main

import "fmt"

func main() {

	fmt.Printf(("Condition statement on this go series!\n"))
	var result string
	loginCount := 23

	if loginCount < 10 {
		result = "Regular Users"
	} else if loginCount > 10 && loginCount < 18 { //&& for and operator and || for or operator and != operator is not equalsto
		result = "Result is in between 10 and 30 characters"

	} else {
		result = "Something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even Numbers")
	} else {
		fmt.Println("ODD Numbers")

	}

	if num := 20; num < 10 { //I can put the variable on the go here also
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is more than 10")

	}

}
