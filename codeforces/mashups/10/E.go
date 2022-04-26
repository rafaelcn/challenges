package main

import (
	"bufio"
	"fmt"
	"os"
)

// E. Números Realmente Grandes

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func sum(n uint64) uint64 {
	var sum uint64 = 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {

	var n, y, big, l, r, m uint64
	I("%d %d\n", &n, &y)

	l, r = 1, n

	for l <= r {
		m = (l + r) / 2

		D(l, r, m)

		if m-sum(m) < y {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	D(l, r, m)

	big = n - ((l + r) / 2)

	O("%d\n", big)

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
