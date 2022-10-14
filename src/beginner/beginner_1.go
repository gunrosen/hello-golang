package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://codeforces.com/problemset/problem/1097/A
func main() {
	var tempString string
	var onTable string
	var inHand []string
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	onTable = scanner.Text()

	scanner.Scan()
	tempString = scanner.Text()
	inHand = strings.Fields(tempString)

	for i := 0; i < len(inHand); i++ {
		if strings.Contains(onTable, inHand[i][0:1]) || strings.Contains(onTable, inHand[i][1:2]) {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
