package main

import (
	"bufio"
	"fmt"
	"os"
)

// B. Soma Maxima

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var i, n, p uint64

	I("%d %d\n", &n, &p)

	a := make([]uint64, n)

	for i = 0; i < n; i++ {
		I("%d", &a[i])
	}
	I("\n")

	var sum, interval, pointer uint64 = 0, 0, 0

	for i = 0; i < n; i++ {
		for pointer < n && sum+a[pointer] <= p {
			sum += a[pointer]
			pointer++
		}
		interval = max(interval, pointer-i)
		sum -= a[i]
	}

	O("%d\n", interval)

	w.Flush()
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}

	return b
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
