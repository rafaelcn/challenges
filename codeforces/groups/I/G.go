package main

import (
	"bufio"
	"fmt"
	"os"
)

// G. Algumas somas

func sum(n int) int {
	// sum digits from the number n
	s := 0

	for n > 0 {
		a := n % 10
		s += a
		n = n / 10
	}

	return s
}

func main() {

	var n, a, b int

	scanf("%d %d %d\n", &n, &a, &b)

	r := 0

	for i := 0; i <= n; i++ {
		s := sum(i)

		if s >= a && s <= b {
			r += i
		}
	}

	println(r)

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
