package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// C. Inversion Graph

func construct(a []int) map[int][]int {

	// construct connected components
	k := 0
	g := make(map[int][]int)

	if len(a) == 1 {
		g[0] = []int{a[0]}
	}

	// Implement a DSU

	/*for i := 0; i < len(a); i++ {
		b := make([]int, 0)

		for j := i + 1; j <= len(a)-1; j++ {
			if a[i] > a[j] {
				// not sure it will work for every case
				if len(g) > 0 && search(g[k-1], a[j]) {
					continue
				} else if len(g) > 0 && search(g[k-1], a[i]) {
					continue
				}

				if !search(b, a[j]) {
					b = append(b, a[j])
				}

				if !search(b, a[i]) {
					b = append(b, a[i])
				}
			}
		}

		if len(b) > 0 {
			g[k] = b
		}

		k++
	}

	fmt.Printf("%v\n", g)*/

	return g
}

func search(a []int, b int) bool {

	sort.Ints(a)
	i := sort.Search(len(a), func(i int) bool { return a[i] >= b })

	if i < len(a) && a[i] == b {
		return true
	}

	return false
}

func solve() {

	var n int

	scanf("%d\n", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}

	scanf("\n")

	println(len(construct(a)))
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
