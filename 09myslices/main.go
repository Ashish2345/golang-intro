package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Hello to the slices section of the Go series")

	var myFruits = []string{"apple", "banana"}
	fmt.Println("Fruits Lists :", myFruits)
	fmt.Printf("The type is %T \n", myFruits) // Slices

	//Append
	myFruits = append(myFruits, "Mango", "Peach")
	fmt.Println("New Fruits Lists :", myFruits)

	//Slice
	myFruits = append(myFruits[:2])
	fmt.Println("New Fruits Lists :", myFruits)

	//Another way
	highScore := make([]int, 4)
	highScore[0] = 20
	highScore[1] = 60
	highScore[2] = 40
	highScore[3] = 50

	//We cannot add new element like
	// highScore[4] = 90
	highScore = append(highScore, 200, 5)
	fmt.Println("Hish Score Lists :", highScore)

	//Its a slice that usase lists underthe hood

	//To Sort
	sort.Ints(highScore)
	fmt.Println("Hish Score Lists :", highScore)

	// var a []int // Nil Slice
	// var a = []int{} //Empty slciea

	// a := make([]int, 5, 10) //Length 5 and capacity of 10
	// cap(a), len(a)

	//Merge 2 slice

	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// a = append(a, b...)
	// print(a) A will have both values

	// Also
	// when i do
	// a := []int{1, 2, 3, 4}
	// b := a[:3]

	// Eventhough i have copied the array till 3 still b will have the capacity of 4

	// Its called the ripple effect
	// also if i append
	// b = append(b, 15)
	// on the forth array it will add that calue the neww array will be for b = [1,2,3,15] and for a = [1,2,3,15]
	// but if only the b is in bound to a array

	// if I change th index of 2 from b
	// b[1] = 6

	// It will change the value of a and b as well as they share same array

	//Some Emaple to know it clearly

	// You can edit this code!
	// Click here and start typing.

	// a := []int{1, 2, 3, 4, 5, 6}

	// b := a[:3]

	// b = append(b, 88)
	// fmt.Println(a)
	// fmt.Println(b)

	// [1 2 3 88 5 6]
	// [1 2 3 88]

	//Another Case

	// a := []int{1, 2, 3, 4, 5, 6}

	// b := a[3:]
	// b = append(b, 66)

	//As now b doesnot share the same array as a and now they are independent

	// fmt.Println(a)
	// fmt.Println(b)

	// [1 2 3 4 5 6]
	// [4 5 6 66]

	//Another

	// b := a[0:6:6] //third six is the capacity of the slice

	//Another
	// a := []int{1, 2, 3, 4, 5, 6}
	// b := make([]int, 6) //6 is the length
	// copy(b, a) // To copy a slice to b // Also both need to be slice

}
