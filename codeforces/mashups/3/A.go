package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n, m int

	scanf("%d %d\n", &n, &m)

	profit := make(map[int]int)

	for i := 0; i < m; i++ {
		var ai int
		scanf("%d ", &ai)

		profit[ai] += ai
	}

	scanf("\n")

	type_ := 0
	max := 0

	for k, v := range profit {
		if v > max {
			max = v
			type_ = k
		} else if v == max && type_ < k {
			type_ = k
		}
	}

	println(max, type_)

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
