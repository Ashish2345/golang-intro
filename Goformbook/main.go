package main

import "fmt"

func main() {

	var num1, num2 int
	var operator string

	fmt.Println("Enter 2 number for calculating the numbers")
	fmt.Scanln(&num1, &num2)

	fmt.Println("Enter operator (+,-,*,/)")
	fmt.Scanln(&operator)

	calculation := func(operator string, num1, num2 int) int {
		switch operator {

		case "+":
			return num1 + num2
		case "-":
			return num1 - num2

		case "*":
			return num1 * num2

		case "/":
			if num2 == 0 {
				panic("invalid operator")
			}
			return num1 / num2

		default:
			panic("Unsupported Opearator")
		}
	}
	result := calculation(operator, num1, num2)
	fmt.Println(result)

}
