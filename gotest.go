package main

import (
	"fmt"
	"sync"
	"time"
)

func heavyComputation(id int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()
	// Simulate heavy computation by calculating sum of squares for a large range
	total := 0
	for i := 0; i < 100000; i++ {
		total += i * i
	}
	resultChan <- total
}

func main() {
	start := time.Now() // Start measuring time

	var wg sync.WaitGroup
	resultChan := make(chan int, 500)
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go heavyComputation(i, &wg, resultChan)
	}
	wg.Wait()
	close(resultChan)

	total := 0
	for result := range resultChan {
		total += result
	}

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Println("Total time:", elapsed)
	fmt.Println("Total result:", total)
}
