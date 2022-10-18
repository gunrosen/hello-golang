package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://codeforces.com/problemset/problem/509/A

func valueOf(x int, y int) int {
	if x == 1 || y == 1 {
		return 1
	} else {
		return valueOf(x-1, y) + valueOf(x, y-1)
	}
}
func main() {
	var err any
	var numberStr string
	var number int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numberStr = scanner.Text()
	number, err = strconv.Atoi(numberStr)
	if err != nil {
		number = 0
	}
	fmt.Println(valueOf(number, number))
}
