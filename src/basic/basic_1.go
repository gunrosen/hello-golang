package main

//About interface

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.write([]byte("Hello World"))

	myIncrement := IntCounter(0)
	var inc Increment = &myIncrement
	fmt.Println(inc.increment())
	fmt.Println(inc.increment())
	fmt.Println(inc.increment())
}

// By struct
type Writer interface {
	write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// By pinter
type Increment interface {
	increment() int
}
type IntCounter int

func (ic *IntCounter) increment() int {
	*ic++
	return int(*ic)
}
