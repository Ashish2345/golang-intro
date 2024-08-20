package main

import "fmt"

func main() {

	fmt.Println("Starting the functio nsection of this go series...")
	greeter()
	result := adder(1, 2)
	fmt.Println("The addition resukt is", result)

	result2, msg := multiAdder(1, 2, 3, 4, 5, 6)
	fmt.Println("The addition resukt is", result2, msg)

}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

//Veriadic parameter
func multiAdder(manyNums ...int) (int, string) {
	total := 0

	for _, val := range manyNums {
		total += val
	}
	return total, "Your number is this"
}

func greeter() {
	fmt.Println("Namaste this is the greeter")
}
