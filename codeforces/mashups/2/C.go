package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// C. Deivis e a maior diferença

func main() {

	var n int

	scanf("%d\n", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		var ai int
		scanf("%d", &ai)
		a[i] = ai
	}

	scanf("\n")

	if n == 2 {
		printf("%d", int(math.Abs(float64(a[1] - a[0]))))
	} else {

		max := 0

		for i, j := 0, 1; j < n; i++ {
			c := int(math.Abs(float64(a[j] - a[i])))
			if c > max {
				max = c
			}

			j++
		}

		printf("%d", max)
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

// Data structures

type Pair struct {
	First, Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
