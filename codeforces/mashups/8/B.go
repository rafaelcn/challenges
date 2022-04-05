package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// B. Comprando strings

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func solve() {

	var n, c0, c1, h int
	var s string

	fmt.Scanf("%d %d %d %d\n", &n, &c0, &c1, &h)
	fmt.Scanf("%s\n", &s)

	n0 := strings.Count(s, "0")
	n1 := strings.Count(s, "1")

	if max(c0, c1) <= h+min(c0, c1) {
		println(n1*c1 + n0*c0)
	} else {
		if c1 > c0 {
			println(n*c0 + (h * n1))
		} else {
			println(n*c1 + (h * n0))
		}
	}
}

func main() {

	var t int
	fmt.Scanf("%d\n", &t)

	for t > 0 {
		solve()
		t--
	}

	writer.Flush()
}

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}

func println(f ...interface{}) {
	fmt.Fprintln(writer, f...)
}
