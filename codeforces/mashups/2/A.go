package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var s string

	scanf("%d\n", &n)
	scanf("%s\n", &s)

	taken := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			taken++
		}
	}

	fmt.Printf("%d", taken)

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
