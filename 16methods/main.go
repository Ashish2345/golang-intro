package main

import "fmt"

func main() {

	fmt.Println("Method section of go series")
	// No super class and no inheritance in golang

	aashish := User{"Aashish", 16, "ashihs@gmaail.com", true}
	fmt.Println(aashish)
	fmt.Printf("My name is %v and i am %v age and my email is %v and i am %v\n", aashish.Name, aashish.Age, aashish.Email, aashish.Status)
	aashish.GetStatus()
	aashish.NewEmail()
	//Whenever we pass the objects or structs it passes the copy
	//As we can see that the email value doesnot change it send the copy. So we use the concept of pointer
	fmt.Printf("My name is %v and i am %v age and my email is %v and i am %v\n", aashish.Name, aashish.Age, aashish.Email, aashish.Status)

}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is User active", u.Status)
}

//one thing
func (u User) NewEmail() {
	u.Email = "test@dev.go"
	fmt.Println("New email is", u.Email)

}
