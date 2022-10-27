package main

// Goroutine

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	// Wait number of goroutine
	wg.Add(2)

	go func() {
		count("sheep")
		// Set goroutine done
		wg.Done()
	}()
	go func() {
		count("fish")
		// Set goroutine done
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Done!")
}

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}
