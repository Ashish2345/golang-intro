package main

import (
	"fmt"
	"os"
)

func Say(s string) string {
	return fmt.Sprintf("%s", s)
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println(Say(os.Args[1]))
	} else {
		fmt.Println("No Args provided")
	}

}
