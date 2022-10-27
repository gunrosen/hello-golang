package main

import (
	"fmt"
)

func main() {
	randomNumbers := []int{}
	for i := 1; i < 100; i++ {
		randomNumbers = append(randomNumbers, i)
	}
	//pipeline
	inputChan := generatePipeline(randomNumbers)

	// Fanout
	// 1 in - chia ra cho n out
	c1 := fanout("c1", inputChan)
	c2 := fanout("c2", inputChan)
	c3 := fanout("c3", inputChan)
	c4 := fanout("c4", inputChan)

	// Fanin
	c := fanin(c1, c2, c3, c4)

	sum := 0
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}

	fmt.Printf("Total sum of square: %d", sum)
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanout(name string, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("%d process in channel %s\n", n, name)
			out <- n * n
		}
		close(out)
	}()
	return out
}
func fanin(inputs ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range inputs {
			for n := range c {
				in <- n
			}
		}
	}()
	return in
}
