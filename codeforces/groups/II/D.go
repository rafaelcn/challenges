package main

import (
	"bufio"
	"fmt"
	"os"
)

// D. Melhor presente de aniversário (?)

func solve() {

	var n int
	scanf("%d\n", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}

	scanf("\n")

	pairs := []Pair{}

	// Non optimal solution, up to O(n^2*t) = 10^10*10 = 1e11 (see D_1.go)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i]&a[j] >= a[i]^a[j] {
				pairs = append(pairs, Pair{First: a[i], Second: a[j]})
			}
		}
	}

	println(len(pairs))
}

func main() {

	var t int

	scanf("%d\n", &t)

	for t > 0 {
		solve()
		t--
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
