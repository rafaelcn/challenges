package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// E. Treinamento

func main() {

	n, i, days := 0, 0, 0

	scanf("%d\n", &n)

	a := make([]int, n)

	for i < n {
		scanf("%d", &a[i])
		i++
	}

	scanf("\n")

	sort.Ints(a)

	// 1 1 3 4 n
	//         |
	//
	// j > a[0] nop
	// j = 1
	// j > a[2] nop
	// j = 2
	// j > a[3] nop
	// j = 3

	for i, j := 0, 1; i < len(a); i++ {
		//fmt.Printf("v: %v i: %d j: %d d: %d -> ", a, i, j, days)

		if j > a[i] {
			continue
		} else {
			j++
			days++
		}

		//fmt.Printf("v: %v i: %d j: %d d: %d\n", a, i, j, days)
	}

	// O(n*log(n) + n)
	println(days)

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
