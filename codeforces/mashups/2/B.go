package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var x, y, z int

	x0, y0, z0 := 0, 0, 0

	scanf("%d\n", &n)

	for n > 0 {
		scanf("%d %d %d\n", &x, &y, &z)
		x0 += x
		y0 += y
		z0 += z
		n--
	}

	if x0 == 0 && y0 == 0 && z0 == 0 {
		printf("YES")
	} else {
		printf("NO")
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
