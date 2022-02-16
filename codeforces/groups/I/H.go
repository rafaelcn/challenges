package main

import (
	"bufio"
	"fmt"
	"os"
)

// H. Deivis e o ano bonito

func check(a int) bool {

	m := make(map[int]int)

	for a > 0 {
		n := a % 10

		m[n]++

		a = a / 10
	}

	if len(m) == 4 {
		return true
	}

	return false
}

func main() {

	var n int

	scanf("%d\n", &n)

	for i := n + 1; i < 9999; i++ {
		if check(i) {
			println(i)
			break
		}
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
