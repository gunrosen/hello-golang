package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://codeforces.com/problemset/problem/1167/A

func main() {
	var tempString string
	var numberTestcase int

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	tempString = scanner.Text()

	numberTestcase, _ = strconv.Atoi(tempString)

	for i := 0; i < numberTestcase; i++ {
		scanner.Scan()
		_ = scanner.Text()
		scanner.Scan()
		tempString = scanner.Text()
		if len(tempString) < 11 {
			fmt.Println("NO")
			continue
		}
		var canBePhoneNumber = false
		for j := 0; j <= len(tempString)-11; j++ {
			if tempString[j] == '8' {
				canBePhoneNumber = true
				break
			}
		}
		if canBePhoneNumber {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
