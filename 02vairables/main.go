package main

import "fmt"

const LoginToken string = "Somerandomtoken"

func main() {

	var username string = "Aashish"

	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var score uint64 = 255
	fmt.Println(score)
	fmt.Printf("Variable is of type %T \n", score)

	var price float32 = 196.225555566666666
	fmt.Println(price)
	fmt.Printf("Variable is of type %T \n", price)

	//Default Values
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T \n", anotherVariable)

	//Implecit  type Here we didnot mention the type but lexer does the thing
	var anotherVariable1 = "Aashish RM"
	fmt.Println(anotherVariable1)
	fmt.Printf("Variable is of type %T \n", anotherVariable1)

	//No Var type
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type %T \n", numberOfUsers)

	//No Var type
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)

}
