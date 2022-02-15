package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// I. Almoço em Manhattan

func main() {

	var n, xc, yc int

	scanf("%d %d %d\n", &n, &xc, &yc)

	// list of costs
	i := 1
	c := make(map[int]int)

	for i <= n {
		var xi, yi, ci int
		scanf("%d %d %d\n", &xi, &yi, &ci)

		md := abs(xi-xc) + abs(yi-yc)

		if _, ok := c[md*2+ci]; !ok {
			c[md*2+ci] = i
		}

		i++
	}

	costs := []int{}
	for k := range c {
		costs = append(costs, k)
	}
	sort.Ints(costs)

	fmt.Printf("%d %d", costs[0], c[costs[0]])

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

func println(f ...interface{}) {
	fmt.Fprintln(writer, f...)
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

func abs(a int) int {
	if a < 0 {
		a = -a
	}

	return a
}
