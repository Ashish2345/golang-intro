package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	// days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	//Type 1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])

	// }

	// Type 2
	// for i := range days {
	// 	fmt.Println(days[i])

	// }

	// Type 3

	// for _, values := range days {

	// 	fmt.Printf("The value is %v\n", values)
	// }

	//Type 4 which is like the while loop

	randNum := 1

	for randNum < 10 {
		if randNum == 8 {
			break
		}
		if randNum == 6 {
			randNum++
			continue
		}

		if randNum == 9 {
			goto loc
		}
		fmt.Printf("The value is %v\n", randNum)
		randNum++

	}

	fmt.Println("Fresh one!!")

loc:
	fmt.Println("Inthte goto syntax")
}
