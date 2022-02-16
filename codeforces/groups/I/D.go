package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// D. Deivis e a igualdade

func main() {

	var n int

	scanf("%d\n", &n)

	if n == 1 {
		scanf("%s\n")
		println(0)
	} else {
		total := 0
		a := make([]int, n)

		for i := 0; i < n; i++ {
			var ai int
			scanf("%d", &ai)
			a[i] = ai
		}

		sort.Ints(a)

		for i := 0; i < n-1; i++ {
			total += a[n-1] - a[i]
		}

		println(total)
	}

	writer.Flush()
}

// Use buffering when reading/writing data as the default fmt.Scanx/fmt.Printx
// functions doesn't.

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}

func println(f interface{}) {
	fmt.Fprintln(writer, f)
}

// Data structures

type Pair struct {
	First, Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Util functions

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
