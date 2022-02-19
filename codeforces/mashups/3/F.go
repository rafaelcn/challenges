package main

import (
	"bufio"
	"fmt"
	"os"
)

// F. Buracos dos dígitos

func count(n int) int {

	a := []int{}

	for n > 0 {
		t := n % 10

		a = append(a, t)

		n = n / 10
	}

	holes := 0

	if len(a) == 0 {
		// n == 0
		holes++
	} else {
		for _, v := range a {
			if v == 0 || v == 6 || v == 9 {
				holes++
			} else if v == 8 {
				holes += 2
			}
		}
	}

	return holes
}

func main() {

	var a, b int

	scanf("%d %d\n", &a, &b)

	maxHoles := a

	for i := a; i <= b; i++ {
		if count(i) > count(maxHoles) {
			maxHoles = i
		}
	}

	println(maxHoles)

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
