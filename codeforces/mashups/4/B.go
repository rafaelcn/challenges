package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// B. Sombra

func main() {

	var n, k int

	scanf("%d %d\n", &n, &k)

	a := make(map[int]int, n)

	for i := 0; i < n; i++ {
		var ai int
		scanf("%d", &ai)

		a[ai] += 1
	}

	j := 0
	uncovered := make(map[int][]int)

	for i := 1; i <= k; i++ {
		if _, ok := a[i]; !ok {
			uncovered[j] = append(uncovered[j], i)
		} else {
			j++
		}
	}

	//println(a)
	//println(uncovered)
	println(len(uncovered))

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
