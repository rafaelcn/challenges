package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// B. Deivis e o menor retângulo

func main() {

	var n, m int

	scanf("%d %d\n", &n, &m)

	// create a hash table of rows and columns on which the key (the index) maps
	// to a value "1" in the matrix.
	rows := make(map[int]int)
	columns := make(map[int]int)

	for i := 0; i < n; i++ {
		var in string
		scanf("%s\n", &in)

		if strings.Contains(in, "1") {
			rows[i]++

			for j, r := range in {
				s := string(r)
				if s == "1" {
					columns[j]++
				}
			}
		}
	}

	// get indexes and sort them
	rowIndices := []int{}
	columnIndices := []int{}

	for k := range columns {
		columnIndices = append(columnIndices, k)
	}
	for k := range rows {
		rowIndices = append(rowIndices, k)
	}
	sort.Ints(columnIndices)
	sort.Ints(rowIndices)

	if len(columnIndices) > 0 {
		// fill column gaps
		for i := columnIndices[0]; i <= columnIndices[len(columnIndices)-1]; i++ {
			if _, ok := columns[i]; !ok {
				columns[i]++
			}
		}
	}
	if len(rowIndices) > 0 {
		// fill row gaps
		for i := rowIndices[0]; i <= rowIndices[len(rowIndices)-1]; i++ {
			if _, ok := rows[i]; !ok {
				rows[i]++
			}
		}
	}

	printf("%dx%d\n", len(columns), len(rows))

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