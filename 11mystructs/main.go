package main

import "fmt"

func main() {

	fmt.Println("Struct section of go series")
	// No super class and no inheritance in golang
	//Implesit
	aashish := User{"Aashish", 16, "ashihs@gmaail.com", true}
	//Explicit
	aashish2 := User{
		Name: "aashish",
		Age:  22,
	}

	//AS we see on explicit we may or maynot add other fields
	//Also in future if we want to add other fields we have to use append
	aashish2.Status = true
	fmt.Println(aashish, aashish2)
	fmt.Println("Updating the values of struct")
	aashish2.Name = "rajesh"
	fmt.Println(aashish, aashish2)
	fmt.Printf("%+v\n", aashish) // We print the value of struct

	// fmt.Printf("My name is %v and i am %v age and my email is %v and i am %v", aashish.Name, aashish.Age, aashish.Email, aashish.Status)

	//Anonymous slice // Also we cannot reuse it
	guardian := struct {
		firstName string
	}{"Raj kumar"}
	fmt.Println(guardian)

	//1st curly braces is to define the structure and another is for value
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
	// someSlice []string
}
