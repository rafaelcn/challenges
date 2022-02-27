package main

import (
	"bufio"
	"fmt"
	"os"
)

// A. Equação com XOR

func main() {

	var a, b uint32

	scanf("%d %d\n", &a, &b)

	if a == b {
		println(0)
	} else {
		println(a^b)
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

func pow(b, n int) int {

	if n == 0 {
		return 1
	}

	a := 1

	for n >= 1 {
		a = b * a
		n--
	}

	return a
}

func flip(b int) int {

	if b == 0 {
		return 1
	}

	return 0
}

func bin(dec int) []int {

	bin := []int{}

	for dec > 0 {
		if dec%2 != 0 {
			bin = append(bin, 1)
		} else {
			bin = append(bin, 0)
		}

		dec = dec / 2
	}

	return bin
}

func dec(b []int) int {

	dec := 0

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	for i, v := range b {
		dec += pow(2, i) * v
	}

	return dec
}
