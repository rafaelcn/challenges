package main

import (
	"bufio"
	"fmt"
	"os"
)

// A. Talheres

func main() {

	var n, g, f, c, people int

	scanf("%d %d %d %d\n", &n, &g, &f, &c)

	people += c
	people += min(g, f)

	if people > n {
		people = n
	}

	println(people)

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
