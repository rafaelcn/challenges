package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// G. Virtuaula

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

type Class struct {
	ClassBegin int
	ClassEnd   int
}

func main() {

	var n, m int

	I("%d\n", &n)

	// chess classes
	cc := make([]Class, n)

	for i := 0; i < n; i++ {
		I("%d %d\n", &cc[i].ClassBegin, &cc[i].ClassEnd)
	}

	I("%d\n", &m)

	// programming classes
	pc := make([]Class, m)

	for i := 0; i < m; i++ {
		I("%d %d\n", &pc[i].ClassBegin, &pc[i].ClassEnd)
	}

	maxL1, minR1 := -math.MaxInt32, math.MaxInt32
	maxL2, minR2 := -math.MaxInt32, math.MaxInt32

	for i := 0; i < n; i++ {
		maxL1 = max(maxL1, cc[i].ClassBegin)
		minR1 = min(minR1, cc[i].ClassEnd)
	}

	for i := 0; i < m; i++ {
		maxL2 = max(maxL2, pc[i].ClassBegin)
		minR2 = min(minR2, pc[i].ClassEnd)
	}

	//fmt.Println(maxL1, minR1, maxL2, minR2)

	ans := max(maxL2-minR1, maxL1-minR2)

	O("%d\n", max(ans, 0))

	w.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func I(format string, args ...interface{}) {
	fmt.Fscanf(r, format, args...)
}

func O(format string, args ...interface{}) {
	fmt.Fprintf(w, format, args...)
}
