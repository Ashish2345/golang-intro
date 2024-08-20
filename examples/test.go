// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	someFunc := func(argsss ...string) string {
		fmt.Println(argsss)
		return "ok"
	}
	fmt.Println(someFunc("nice", "beautiful"))
}
