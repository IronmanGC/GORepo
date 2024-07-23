// Write go program to print 1-1000 in concurrent fashion

package main

import (
	"fmt"
	"sync"
)

func printSeries(startIndex int, endIndex int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := startIndex; i <= endIndex; i++ {
		fmt.Println("Index :", i)
	}
}

func main() {
	fmt.Println("Start Execution")
	numGoroutines := 2 // Number of goroutines to use
	var wg sync.WaitGroup

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go printSeries(1, 500, &wg)
		go printSeries(501, 1000, &wg)
	}

	wg.Wait()
	fmt.Println("End Execution")
}
