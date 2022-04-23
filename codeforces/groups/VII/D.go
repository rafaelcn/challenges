package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// D. Black Fraude

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

type Tuple struct {
	PriceBlackFriday int
	PriceNormal      int
	PriceDifference  int
}

func main() {

	var n, k, t int

	I("%d %d\n", &n, &k)

	p := make([]Tuple, n)

	for i := 0; i < n; i++ {
		I("%d", &t)
		p[i] = Tuple{t, 0, 0}
	}
	I("\n")

	for i := 0; i < n; i++ {
		I("%d", &t)

		p[i].PriceNormal = t
		p[i].PriceDifference = t - p[i].PriceBlackFriday
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].PriceDifference > p[j].PriceDifference
	})

	total := 0

	for i := 0; i < k; i++ {
		total += p[i].PriceBlackFriday
	}

	for i := k; i < n; i++ {
		total += min(p[i].PriceBlackFriday, p[i].PriceNormal)
	}

	O("%d\n", total)

	w.Flush()
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
