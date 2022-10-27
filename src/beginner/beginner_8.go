package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/996/B

func main() {
	var numberOfEntrance int64
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &numberOfEntrance)
	queue := make([]int64, numberOfEntrance)
	for i := int64(0); i < numberOfEntrance; i++ {
		var temp int64
		fmt.Fscan(in, &temp)
		if temp == 0 {
			queue[i] = i
		} else if temp < numberOfEntrance {
			if temp - i > 0{
				queue[i] =  i + numberOfEntrance
			} else {
				queue[i] = i
			}
		} else {
			queue[i] = temp + temp % numberOfEntrance
		}
	}
	if numberOfEntrance == 1 {
		fmt.Println(1)
		return
	}
	minIndex := int64(0)
	currentMin := queue[0]

	for i := int64(1); i < numberOfEntrance; i++ {
		if queue[i] < currentMin {
			minIndex = i
			currentMin = queue[i]
		}
	}
	fmt.Println(minIndex +1 )
}
