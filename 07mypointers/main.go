//Go program memory spaces is divided into 2 main areas. Stack and heap.
// Stack is used for execution of locaal variables and function calls
// Heap is used for dynamic memory locations

//If pointer is declared and not used it's value is nil

//imp//
// if we try to deference or access the pointed by the nill pointerit will be panic

package main

import "fmt"

func main() {

	fmt.Printf("Some exampels on pointers\n")

	var ptr *int
	//if i try to print pointer tha *ptr i will get the runtime error or panic,
	mainPointer := 66

	ptr = &mainPointer
	fmt.Println("Value of pointer is", *ptr)

	// myNumber := 23
	// var ptr = &myNumber //& it gives us an address

	// fmt.Println("Memeory address of pointer ptr is ", ptr)
	// fmt.Println("Value of pointer ptr is ", *ptr) //* it gives us an vales //it is called dereferencing

	// *ptr = *ptr + 3
	// fmt.Println("New Value of pointer ptr is ", *ptr, myNumber)

}
