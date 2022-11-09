package main

import (
	"fmt"
	"sync"
	"time"
)

/*
View basic_6. It's good example for workerpool
*/

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int, error chan<- error) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)

		// Make errors

		if j%2 == 0 {
			results <- j * 2
		} else {
			error <- fmt.Errorf("error on job %v", j)
		}

		wg.Done()

	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	errors := make(chan error, 100)

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		go worker(w, &wg, jobs, results, errors)
	}

	for j := 1; j <= 100; j++ {
		jobs <- j
		wg.Add(1)
	}
	close(jobs)

	wg.Wait()

	for a := 1; a <= 100; a++ {
		select {
		case err := <-errors:
			fmt.Println("finish with error: ", err.Error())
		default:
			<-results
		}
	}
}
