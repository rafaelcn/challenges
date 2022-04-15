package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n, m int
	I("%d %d\n", &n, &m)

	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		I("%d", &a[i])
	}
	I("\n")

	for i := 0; i < m; i++ {
		I("%d", &b[i])
	}
	I("\n")

	sort.Ints(a)

	//fmt.Printf("%v\n%v", a, b)

	for i, v := range b {
		j := sort.SearchInts(a, v)

		if j < len(a) && a[j] == v {
			O("%d", j+1)
		} else {
			O("%d", -1)
		}

		if i < len(b)-1 {
			O(" ")
		}
	}

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
