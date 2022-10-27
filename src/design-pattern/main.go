package main

import (
	"design-pattern/singleton"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", singleton.GetInstance())
		}()
	}
	s1 := singleton.GetInstance()
	s1.AddOne()
	fmt.Println(s1.AddOne())
	s2 := singleton.GetInstance()
	s2.AddOne()
	fmt.Printf("s1 %p\n", s1)
	fmt.Printf("s2 %p\n", s2)

	time.Sleep(time.Second * 10)
}
