package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

// D. Melhor presente de aniversário (?)

func solve() {

	var n, i uint64
	scanf("%d\n", &n)

	a := make([]uint64, n)

	for i = 0; i < n; i++ {
		scanf("%d", &a[i])
	}

	scanf("\n")

	m := make(map[int]uint64)

	for _, ai := range a {
		// add up numbers that are represented by the same number of bits
		bs := bits.Len64(ai)

		m[bs]++
	}

	var pairs uint64 = 0

	for _, v := range m {
		if v >= 2 {
			pairs += uint64((v * (v - 1)) / 2)
		}
	}

	println(pairs)
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

// Constants

const (
	B  = 1e9
	BB = 1e18
)

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

func tos(n int) string {
	return strconv.Itoa(n)
}

func toi(s string) int {
	// ignore error
	n, _ := strconv.Atoi(s)

	return n
}

func bin(dec uint64) []uint64 {

	bin := []uint64{}

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
