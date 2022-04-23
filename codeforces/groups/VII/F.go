package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// F. Me Paga uma Coxinha

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n int

	I("%d\n", &n)

	a := make([]int, n)
	b := make([]int, n)

	var card1, card2 string

	I("%s\n", &card1)
	I("%s", &card2)

	// convert to vectors
	for i := 0; i < n; i++ {
		a[i] = int(card1[i] - '0')
		b[i] = int(card2[i] - '0')
	}

	sort.Ints(a)
	sort.Ints(b)

	var minimum, maximum, skip1, skip2 int

	// try to do a complete search over the two vectors

	for i := 0; i < n; i++ {
		if (skip1 >= n) && (skip2 >= n) {
			break
		}

		// look over card 1
		for (skip1 < n) && (b[skip1] < a[i]) {
			skip1++
		}

		// look over card 2
		for (skip2 < n) && (b[skip2] <= a[i]) {
			skip2++
		}

		if skip1 < n {
			skip1++
			minimum++
		}
		if skip2 < n {
			skip2++
			maximum++
		}
	}

	O("%d\n%d\n", n-minimum, maximum)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
