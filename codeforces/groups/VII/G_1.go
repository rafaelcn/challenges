package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Slice(cc, func(i, j int) bool {
		return cc[i].ClassEnd < cc[j].ClassEnd
	})
	sort.Slice(pc, func(i, j int) bool {
		return pc[i].ClassBegin < pc[j].ClassBegin
	})

	ccf := cc[0]
	pcl := pc[m-1]

	D(cc, pc, pcl.ClassBegin, ccf.ClassEnd)

	sort.Slice(pc, func(i, j int) bool {
		return pc[i].ClassEnd < pc[j].ClassEnd
	})
	sort.Slice(cc, func(i, j int) bool {
		return cc[i].ClassBegin < cc[j].ClassBegin
	})

	pcf := pc[0]
	ccl := cc[n-1]

	D(pc, cc, ccl.ClassBegin, pcf.ClassEnd)

	ans := max(pcl.ClassBegin - ccf.ClassEnd, ccl.ClassBegin - pcf.ClassEnd)

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

func I(format string, args ...interface{}) {
	fmt.Fscanf(r, format, args...)
}

func O(format string, args ...interface{}) {
	fmt.Fprintf(w, format, args...)
}

// for debug purposes
const hostname = "elementary"

func D(args ...interface{}) {
	if h, err := os.Hostname(); h == hostname && err == nil {
		fmt.Fprintln(os.Stderr, args...)
	}
}