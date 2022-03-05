package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// B. Quality vs Quantity

func solve() {

	var n int

	scanf("%d\n", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}
	scanf("\n")

	sort.Ints(a)

	var bsum, rsum, bcount, rcount int64 = int64(a[0]), int64(a[n-1]), 1, 1

	// limits the sum on the for loop
	j := n - 1

	// if it has found a way to paint the numbers
	found := false

	for i := 1; i <= j; i++ {
		bsum += int64(a[i])
		if bsum >= rsum && j-1 > i {
			//println(bsum, ">=", rsum, i, j)
			j -= 1
			rsum += int64(a[j])
			rcount++
		}
		bcount++

		if bsum < rsum && bcount > rcount {
			found = true
			break
		}
	}

	//println(bsum, rsum, bcount, rcount)

	if found {
		println("YES")
	} else {
		println("NO")
	}
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
