package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/706/B

func main() {
	var numberOfDay, numberOfShop int
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &numberOfShop)
	prices := make([]int, numberOfShop)
	for i := int(0); i < numberOfShop; i++ {
		fmt.Fscan(in, &prices[i])
	}
	sort.Ints(prices)
	fmt.Fscan(in, &numberOfDay)
	for i := int(0); i < numberOfDay; i++ {
		var money int
		fmt.Fscan(in, &money)
		left := 0
		right := len(prices)
		for left < right {
			mid := (left + right) / 2
			if prices[mid] <= money {
				left = mid + 1
			} else {
				right = mid
			}
		}
		fmt.Println(left)
	}
}
