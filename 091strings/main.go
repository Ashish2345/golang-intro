package main

import "fmt"

func main() {
	//Strings are immutable
	// String, runes, and bytes are convertable between eachother
	fmt.Println("Strings, Runes and Bytes")

	someRandomStrings := "Hello Coders"
	fmt.Println(someRandomStrings)

	// When i try to see the string for the index i will get the bytes
	fmt.Println(someRandomStrings[0]) //72
	// So to get the reight string we use
	fmt.Println(string(someRandomStrings[0]))

	//String in go are encoded with utf-8 encoding

	// But if we use

	someRandom2 := "शुभ सन्ध्या" //As it includes unicode with size more than 1 bytes . these byte doesnot represent a completecorrector
	unicodeString := []rune(someRandom2)
	fmt.Println(string(unicodeString[0]))

}
