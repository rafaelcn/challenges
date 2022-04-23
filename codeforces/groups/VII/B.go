package main

import (
	"bufio"
	"fmt"
	"os"
)

// B. Moeda CIC

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var t, n uint64

	I("%d\n", &t)

	for t > 0 {
		I("%d\n", &n)

		c1 := n / 3
		c2 := c1

		if n%3 == 1 {
			c1++
		} else if n%3 == 2 {
			c2++
		}

		O("%d %d\n", c1, c2)

		t--
	}

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
