package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://codeforces.com/problemset/problem/1064/A

func main() {
	var _ any
	var tempString string
	var tempStringArr []string
	var arr = []int{}
	var step int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	tempString = scanner.Text()
	tempStringArr = strings.Fields(tempString)

	for _, i := range tempStringArr {
		temp, _ := strconv.Atoi(i)
		arr = append(arr, temp)
	}
	sort.Ints(arr)
	if arr[0]+arr[1] > arr[2] {
		step = 0
	} else {
		step = arr[2] - (arr[0] + arr[1]) + 1
	}
	fmt.Println(step)
}
