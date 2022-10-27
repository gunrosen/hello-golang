package main

import "fmt"

// Luu vao heap
// thay vi stack

// This command will show how compiler acts with variable
// go build --gcflags="-m -l"
//./advance_1.go:16:2: moved to heap: x

// Go tu dong escape from stack to heap if :
// + size of object not clarify in compile time
// + object maybe be referenced by others after return from function
func main() {
	n := answer()
	fmt.Printf("%p\n", n)
	println(*n / 2) // x co the bi gc
}

func answer() *int {
	x := 42
	fmt.Printf("%p\n", &x)
	return &x
}
