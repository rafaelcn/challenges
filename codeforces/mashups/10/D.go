package main

import (
	"bufio"
	"fmt"
	"os"
)

// D. Duplicaixa

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n int
	I("%d\n", &n)

	pens := 1

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n--
			pens++
		}
	}

	O("%d\n", pens)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}

// for debug purposes
const hostname = "elementary"

func D(args ...interface{}) {
	if h, err := os.Hostname(); h == hostname && err == nil {
		fmt.Fprintln(os.Stderr, args...)
	}
}
