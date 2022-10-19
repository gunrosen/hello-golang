package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://codeforces.com/problemset/problem/1248/A
func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}
func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readArrInt64(in *bufio.Reader) []int64 {
	numbs := readLineNumbs(in)
	arr := make([]int64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseInt(n, 10, 64)
		arr[i] = val
	}
	return arr
}

func extractInfo(in *bufio.Reader) (int64, int64) {
	odd := int64(0)
	even := int64(0)
	var size int64
	fmt.Fscan(in, &size)
	for i := int64(0); i < size; i++ {
		var temp int64
		fmt.Fscan(in, &temp)
		if temp%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return odd, even
}

func main() {
	var numberTestcase int64
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &numberTestcase)

	for i := int64(0); i < numberTestcase; i++ {
		odd1, even1 := extractInfo(in)
		odd2, even2 := extractInfo(in)
		fmt.Println(odd1*odd2 + even1*even2)
	}
}
