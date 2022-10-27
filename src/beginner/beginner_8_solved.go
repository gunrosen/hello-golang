package main

// https://codeforces.com/problemset/problem/996/B

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var in *bufio.Scanner
var out *bufio.Writer

func main() {
	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)
	defer out.Flush()

	in.Scan();
	n, _ := strconv.ParseInt(in.Text(), 10, 64);

	var res int64;
	var temp int64 = math.MaxInt64;

	for i:=int64(1); i<=n; i++ {
		in.Scan();
		a, _ := strconv.ParseInt(in.Text(), 10, 64);
		if (temp > (a-i+n)/n) {
			temp = (a-i+n)/n;
			res = i;
		}
	}

	out.WriteString(strconv.FormatInt(res, 10));
}
