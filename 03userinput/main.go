package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user programs"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of the restaurant")

	//WE dont have try cache in go so we treat the error as a variable in Go
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating", input)
	fmt.Printf("Type of the rating is %T \n", input)

	input2, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating", input2)
	fmt.Printf("Type of the rating is %T \n", input2)

}
