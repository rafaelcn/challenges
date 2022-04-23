package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// C. Torneio de Xadrez

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n int

	I("%d\n", &n)

	a := make([]int, 2*n)

	for i := 0; i < 2*n; i++ {
		I("%d", &a[i])
	}
	I("\n")

	sort.Ints(a)

	// if the array is sorted, then the nth player should have a different
	// rating of the nth-1 element to win. That's because the middle of one
	// array of 2*n positions will always contains (if there are any) repeated
	// ratings...
	if a[n] == a[n-1] {
		O("NO\n")
	} else {
		O("YES\n")
	}

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
