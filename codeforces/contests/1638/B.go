package main

import (
	"bufio"
	"fmt"
	"os"
)

// B. Odd Swap Sort

func solve() {

	var n int

	scanf("%d\n", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanf("%d", &a[i])

		// try to sort
		if i > 0 && a[i] < a[i-1] && (a[i]+a[i-1])&1 == 1 {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}

	scanf("\n")

	sorted := true

	for i := 1; i <= n-1; i++ {
		if a[i] < a[i-1] {
			sorted = false
			break
		}
	}

	if sorted {
		println("Yes")
	} else {
		println("No")
	}
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
