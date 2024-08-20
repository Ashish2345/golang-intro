package main

import "fmt"

func main() {
	fmt.Println("Welcome to the maps section of this go series!")

	// languages := make(map[string]string) //Creating a empty map
	// // anotheType := map[string]int{ //String as the key and int as the value
	// // 	"Foo":1,
	// // }

	// languages["PY"] = "PYTHON"
	// languages["JS"] = "Java Script"
	// languages["GO"] = "Golang"
	// fmt.Println("List of programming languages in key value pairs", languages)
	// fmt.Println("Golang value is", languages["GO"])

	// //To delete
	// delete(languages, "JS")
	// fmt.Println("List of programming languages in key value pairs", languages)

	// //Loop thorugh the key val pair

	// for key, value := range languages {
	// 	fmt.Printf("Key is %s and value is %s \n", key, value)

	// nameAge := make(map[string]int)
	// nameAge["foo"] = 1
	// val, ok := nameAge["foo"] //if there is not key in the map it will return 0 if we specified val is int if str empty //So we user ,ok syntak
	// fmt.Println(val, ok)

	//Another
	// if we want the slice datatype in map
	// sliceDAta := make(map[string][]int)

}
