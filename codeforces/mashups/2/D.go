package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// D. Resultado da corrida

func main() {

	var n int
	var t Pairs

	scanf("%d\n", &n)

	i := 1

	for i <= n {
		var ti int

		scanf("%d", &ti)

		p := Pair{PositionInserted: i, TimeTaken: ti}
		t = append(t, p)

		i++
	}

	sort.Sort(t)

	for i := 0; i < len(t); i++ {
		t[i].Rank = i + 1
		if i > 0 && t[i].TimeTaken == t[i-1].TimeTaken {
			t[i].Rank = t[i-1].Rank
		}
	}

	// [{2 2 1} {4 2 1} {1 5 3} {3 10 4}]
	//                     ^
	//     ^
	//                              ^
	//             ^

	r := make(map[int]int, len(t))

	for i := 0; i < len(t); i++ {
		r[t[i].PositionInserted] = t[i].Rank
	}

	for i := 1; i <= len(r); i++ {
		printf("%d ", r[i])
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
	PositionInserted, TimeTaken, Rank int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].TimeTaken < p[j].TimeTaken }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
