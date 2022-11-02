package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://gist.github.com/YumaInaura/8d52e73dac7dc361745bf568c3c4ba37

/*
	Channel makes go-routines can communicate with each other.
	Through channel, go routines send/receive messages
*/
/*
	Blocking happens when
	Goroutine try to receive message but channel is empty. And other goroutine still running
*/
/*
	Deadlock happens when a group of goroutines is waiting for each other and none of them is able to proceed
*/
// now, the boring function additional parametter
// `c chan string` is a channel

// <-chan string means receives-only channel of string.
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c
}

func main() {
	chan1 := boring("Chan1")
	chan2 := boring("Chan2")
	for i := 0; i < 11; i++ {
		fmt.Println(<-chan1)
		fmt.Println(<-chan2)
	}
	fmt.Println("You're both boring. I'm leaving")
}
