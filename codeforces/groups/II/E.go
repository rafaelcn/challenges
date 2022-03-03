package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {

	var n int

	scanf("%d\n", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i
	}

	// to minimize the XOR operation we have to pair numbers with different
	// number of bits used to represent them.

	s := a[1 : len(a)-1]
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))

	answer := append([]int{}, s...)

	answer = append(answer, a[0])
	answer = append(answer, a[len(a)-1])


	for _, v := range answer {
		printf("%d ", v)
	}
	println()
}

func msbi(n int) int {
	// returns the index of the most significant bit in a number
	msb := -1

	for n > 0 {
		n /= 2
		msb++
	}

	return msb
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
