package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://codeforces.com/problemset/problem/935/A
func main() {
	var err any
	var numberStr string
	var numberOfEmployee int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numberStr = scanner.Text()
	numberOfEmployee, err = strconv.Atoi(numberStr)
	if err != nil {
		numberOfEmployee = 0
	}
	var result int
	for i := 1; i <= (numberOfEmployee / 2); i++ {
		var temp = numberOfEmployee - i
		if temp%i == 0 {
			result++
		}
	}
	fmt.Println(result)
}
