package main

import (
	"bufio"
	"fmt"
	"os"
)

// C. Deivis e as velas

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n, k, c int
	I("%d %d\n", &n, &k)

	for n >= k {
		c += k
		n -= k
		n++
	}

	O("%d\n", c+n)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
