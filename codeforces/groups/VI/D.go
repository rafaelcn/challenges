package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// D. Valor exato

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n, money int
	I("%d %d\n", &n, &money)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		I("%d", &a[i])
	}

	sort.Ints(a)

	// find x and y using binary search such that:
	// x+y = money
	// x <= y
	// minimize x-y

	ansX, ansY := 0, math.MaxInt32

	left := 0
	right := n - 1

	for left < right {
		x := a[left]
		y := a[right]

		if x+y > money {
			right--
		} else if x+y == money {
			ansX, ansY = x, y
			left++
			right--
		} else {
			left++
		}
	}

	O("%d %d\n", ansX, ansY)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
