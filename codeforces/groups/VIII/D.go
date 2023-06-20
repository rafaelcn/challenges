package main

import (
	"bufio"
	"fmt"
	"os"
)

// D. Escolhendo os pontos

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var i, n, d uint64

	I("%d %d\n", &n, &d)

	a := make([]uint64, n)

	for i = 0; i < n; i++ {
		I("%d", &a[i])
	}
	I("\n")

	if n > 2 {
		triplets := 0

		for i = 0; i < n; i++ {
			// calculate the upper bound (up to which points can be chosen)

		}

		O("%d\n", triplets)
	} else {
		O("0\n")
	}

	w.Flush()
}

func upper(a []int, b int) int {

	l, r := 0, len(a)-1

	search := 0

	for l <= r {
		m := (l + r) / 2

		if a[m] < b {
			l = m + 1
			search = a[m]
		} else {
			r = m - 1
		}
	}

	return search
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
