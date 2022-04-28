package main

import (
	"bufio"
	"fmt"
	"os"
)

// F. Matemágica

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n, m int
	I("%d %d\n", &n, &m)

	// a^2 + b = n
	// a + b^2 = m
	solutions := 0

	for i := 0; i*i <= 1000; i++ {
		for j := 0; j*j <= 1000; j++ {
			if i*i+j == n && i+j*j == m {
				solutions++
			}
		}
	}

	O("%d\n", solutions)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}

// for debug purposes
var (
	localHostname   = "elementary"
	juryHostname, _ = os.Hostname()
)

func D(args ...interface{}) {
	if localHostname == juryHostname {
		fmt.Fprintln(os.Stderr, args...)
	}
}
