package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// D. Jogo Familiar

func main() {

	var n, m int

	scanf("%d %d\n", &n, &m)

	a := make([]int, n)
	// fill familiars
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}

	for i := 0; i < m; i++ {
		var ai int
		scanf("%d", &ai)

		if ai > 0 {
			// rotate elements (I could walk as well, something as a[i%len(a)])
			ai = ai % len(a)

			left := a[:ai]
			rigth := a[ai:]

			a = append([]int{}, rigth...)
			a = append(a, left...)
		} else if ai == -1 {
			// remove element in front of the queue
			a = a[1:]
		}
	}

	if len(a) > 0 {
		println(a[0])
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

// Util functions

func toStr(n int) string {
	return strconv.Itoa(n)
}

func toInt(s string) int {

	// ignore error
	n, _ := strconv.Atoi(s)

	return n
}
