package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// B. Reflorestamento

func main() {

	var n, m, q int

	scanf("%d %d %d\n", &n, &m, &q)

	a := make([]int, N+2)

	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}
	scanf("\n")

	d := make([]int, N+2)

	for q > 0 {
		var l, r, x, m int

		scanf("%d %d %d %d\n", &l, &r, &x, &m)
		l -= 1
		r -= 1

		if m <= 5 {
			d[l] += x
			d[r+1] -= x
		}

		q--
	}

	var v int = 0

	for i := 0; i <= n; i++ {
		v += d[i]
		// the problem wants to know how many trees, in total, were planted on
		// each a_i section.
		a[i] += v
	}

	for i := 0; i < n; i++ {
		fmt.Print(a[i], " ")
	}
}

// Constants

const (
	N  = 1e5
	B  = 1e9
	BB = 1e18
)

// Data structures

type Pair struct {
	First, Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Util functions

func tos(n int) string {
	return strconv.Itoa(n)
}

func toi(s string) int {
	// ignore error
	n, _ := strconv.Atoi(s)

	return n
}

func pow(b, n uint) uint {

	if n == 0 {
		return 1
	}

	var a uint = 1

	for n >= 1 {
		a = b * a
		n--
	}

	return a
}

func flip(b uint) uint {

	if b == 0 {
		return 1
	}

	return 0
}

func bin(dec uint) []uint {

	bin := []uint{}

	for dec > 0 {
		if dec%2 != 0 {
			bin = append(bin, 1)
		} else {
			bin = append(bin, 0)
		}

		dec = dec / 2
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return bin
}

func dec(b []uint) uint {

	var dec uint = 0

	for i, v := range b {
		dec += (2 << i) * v
	}

	return dec
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
