package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boringQuitChannel(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		quit <- true
	}()
	return c
}

func main() {
	quit := make(chan bool)
	c := boringQuitChannel("Joe", quit)

	timeout := time.After(12 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			fmt.Println("They said quit")
			return
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
