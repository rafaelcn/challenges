package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// C. Premium ou Popular?

func main() {

	var p, n int

	scanf("%d %d\n", &p, &n)

	i := 0
	stack := []int{}

	for i < n {
		var op, number int
		scanf("%d %d\n", &op, &number)

		if op == 1 {
			// number is price
			stack = append(stack, number)
		} else {
			// number is the amount of revogations
			for j := 0; j < number; j++ {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			}
		}

		i++
	}

	for _, price := range stack {
		p += price
	}

	println(p)

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
