package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Lco Web Request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // REsponsibility

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response: ", string(dataBytes))
}
