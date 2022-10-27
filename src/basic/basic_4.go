package main

import (
	"fmt"
	"sync"
)

// Channel
// Cac tai nguyen su dung chung giua cac goroutine
// Nen co goroutine chuyen gui, goroutine chuyen nhan
// Phai co gui vao va lay ra

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan int, 50) // Chua duoc 50 con so int - kho chua duoc 50 so int

	wg.Add(2)
	//Chuyen nhan
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		//for i := range ch {
		//	fmt.Println(i)
		//}
		wg.Done()
	}(ch)

	//Chuyen gui
	go func(ch chan<- int) {
		ch <- 42
		ch <- 42
		close(ch) // close channel
		wg.Done()
	}(ch)

	wg.Wait()

}
