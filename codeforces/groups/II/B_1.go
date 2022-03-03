package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
)

// B. XOR Máximo

func main() {

	var n uint64

	scanf("%d\n", &n)

	// alternate solution, doesn't work yet...
	println(n, ^n&mask(bits.Len64(n)))

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

// Utils functions

func mask(b int) uint64 {

	if (b <= 0) {
		return 0
	}

	// b == 4, n = 15
	// b == 5, n = 31
	// b == 6, n = 63

	n := math.Pow(2, float64(b)) - 1

	return uint64(n)
}
