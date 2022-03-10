package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// D. Pedras preciosas

func main() {

	var n, q int

	scanf("%d %d\n", &n, &q)

	a := make([]int, 1e5+2)
	p := make([]int64, 1e5+2)
	f := make([]int64, 1e5+2) // prefix sum of sorted a

	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}
	scanf("\n")

	s := make([]int, n)
	copy(s, a[:n])
	sort.Ints(s)

	p[0] = int64(a[0])
	for i := 1; i < n; i++ {
		p[i] = int64(a[i]) + p[i-1]
	}

	f[0] = int64(s[0])
	for i := 1; i < n; i++ {
		f[i] = int64(s[i]) + f[i-1]
	}

	//println("", s, f[:n], "")
	//writer.Flush()

	for q > 0 {
		var l, r, t int
		scanf("%d %d %d\n", &t, &l, &r)

		l -= 1
		r -= 1

		if t == 1 {
			if l == 0 {
				println(p[r])
			} else {
				println(p[r] - p[l-1])
			}
		} else if t == 2 {
			if l == 0 {
				println(f[r])
			} else {
				println(f[r] - f[l-1])
			}
		}

		q--
	}

	writer.Flush()
}

// Constants

const (
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
