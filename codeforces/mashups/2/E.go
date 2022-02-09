package main

import (
	"bufio"
	"fmt"
	"os"
)

// E. Tiago no Mundo dos Números

func solve() {

	var x int

	scanf("%d\n", &x)

	if x == 2 {
		printf("%d %d\n", 1, 1)
		return
	}

	printf("%d %d\n", 1, x-1)
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

// Data structures

type Pair struct {
	First, Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }