package main

import "fmt"

func main() {
	fmt.Printf("Welcome to the array section \n")

	var fruitList [4]string

	fruitList[0] = "Apple"

	fmt.Println("Arry of fruit is: ", fruitList, len(fruitList))

	var vegLists = [3]string{"Potato", "Beans", "Mushoroom"}
	fmt.Println("Arry of vegy is: ", vegLists, len(vegLists), vegLists[2])

	var bikeLists = [...]string{"honda", "cbz", "pulsor", "nice"} //Implicit length declaration
	fmt.Println("Arry of vegy is: ", bikeLists, len(bikeLists), bikeLists[2])

	twodarray := [2][2]string{{"one", "two"}, {"three", "three"}}
	fmt.Println("Arry of vegy is: ", twodarray)

}
