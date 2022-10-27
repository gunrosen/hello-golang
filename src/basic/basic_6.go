package main

import (
	"fmt"
	"time"
)

// Cho n
// Xuat ra man hinh tat ca con so thu tu 1 -> n
// using worker pool

const numberOfWorker = 1

var workers = [numberOfWorker][]int{}

func main() {
	start := time.Now()
	number := 50

	jobs := make(chan int, number)
	results := make(chan int, number)

	// Create worker pool
	for i := 0; i < numberOfWorker; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < number; j++ {
		fmt.Println(<-results)
	}
	for j := 0; j < numberOfWorker; j++ {
		fmt.Printf("worker-%d %v \n", j, workers[j])
	}
	elapsed := time.Since(start)
	fmt.Printf("Task took %s", elapsed)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for n := range jobs {
		workers[id] = append(workers[id], n)
		results <- Fibonacci(n)
	}
}
