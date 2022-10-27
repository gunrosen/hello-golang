package main

import (
	"fmt"
	"sync"
)

// Mutex: sync
var wg = sync.WaitGroup{}
var counter = 0
var mutex = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)

		mutex.RLock() // read lock
		go sayHello()

		mutex.Lock() // write lock
		go increment()

	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	mutex.RUnlock() // read unlock
	wg.Done()
}

func increment() {
	counter++
	mutex.Unlock() // write unlock
	wg.Done()
}
