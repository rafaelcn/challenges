package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// C. Deivis e as matrículas

func main() {

	var n int

	scanf("%d\n", &n)

	if n == 1 {
		scanf("%s\n")
		println(0)
	} else {
		i := 0
		triples := make(Triples, 0)

		for i < n {
			var name string
			scanf("%s\n", &name)

			t := Triple{Name: name, PositionInserted: i}
			triples = append(triples, t)

			i++
		}

		sort.Sort(triples)

		for i := 0; i < len(triples); i++ {
			triples[i].RegistrationNumber = i
		}

		output := make(map[int]int, len(triples))

		for i := 0; i < len(triples); i++ {
			output[triples[i].PositionInserted] = triples[i].RegistrationNumber
		}

		for i := 0; i < len(output); i++ {
			printf("%d ", output[i])
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

func println(f interface{}) {
	fmt.Fprintln(writer, f)
}

// Data structures

type Triple struct {
	Name               string
	PositionInserted   int
	RegistrationNumber int
}

type Triples []Triple

func (p Triples) Len() int           { return len(p) }
func (p Triples) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p Triples) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Util functions

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
