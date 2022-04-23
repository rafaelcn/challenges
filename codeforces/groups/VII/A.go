package main

import (
	"bufio"
	"fmt"
	"os"
)

// A. Dominós

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var m, n int

	I("%d %d\n", &m, &n)

	// the size of the domino is two by one so we will have at most ans dominoes
	// on the board.
	ans := m * n / 2

	O("%d\n", ans)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
