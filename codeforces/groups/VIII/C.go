package main

import (
	"bufio"
	"fmt"
	"os"
)

// C. Valéria e a Leitura

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n, t int

	I("%d %d\n", &n, &t)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		I("%d", &a[i])
	}
	I("\n")

	sum, interval, pointer := 0, 0, 0

	for i := 0; i < n; i++ {
		for pointer < n && sum+a[pointer] <= t {
			sum += a[pointer]
			pointer++
		}
		interval = max(interval, pointer-i)
		sum -= a[i]
	}

	O("%d\n", interval)

	w.Flush()
}

func max(a, b int) int {
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
