package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A. Madoka and Math Dad

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func solve() {

	var n int
	scanf("%d\n", &n)

	if n == 1 {
		println("1")
	} else if n == 2 {
		println("2")
	} else {
		// combination of ones an twos
		// naive solution (start both strings an choose one)
		var out1, out2 string
		var sum1, sum2 int

		toggle := true

		for sum1 <= n || sum2 <= n {
			if toggle {
				sum1 += 1
				sum2 += 2
				out1 += "1"
				out2 += "2"
			} else {
				sum1 += 2
				sum2 += 1
				out1 += "2"
				out2 += "1"
			}

			if sum1 == n || sum2 == n {
				break
			}

			toggle = !toggle
		}

		if sum1 == sum2 {
			println(out2)
		} else {
			if sum1 == n {
				println(out1)
			} else if sum2 == n {
				println(out2)
			}
		}
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
	N  = 1e5
	M  = 1e6
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

type PriorityQueue struct {
	v    []int
	Size int
}

func (pq *PriorityQueue) build(a []int) {

	pq.v = a
	pq.Size += len(a)

	pq.heap()
}

func (pq *PriorityQueue) heap() {

	// builds a max heap

	n := pq.Size
	r := (n / 2) - 1

	for i := r; i >= 0; i-- {
		pq.heapify(i)
	}

}

func (pq *PriorityQueue) heapify(index int) {

	root := index
	lc := 2*index + 1
	rc := 2*index + 2

	if lc < pq.Size && pq.v[lc] > pq.v[root] {
		root = lc
	}
	if rc < pq.Size && pq.v[rc] > pq.v[root] {
		root = rc
	}

	if root != index {
		pq.v[index], pq.v[root] = pq.v[root], pq.v[index]
		pq.heapify(root)
	}
}

func (pq *PriorityQueue) add(n int) {

	pq.v = append(pq.v, n)
	pq.Size++

	pq.heapify(0)
}

func (pq *PriorityQueue) front() int {

	if pq.Size > 0 {
		// store the first element
		n := pq.v[0]
		// remove element
		pq.v = pq.v[1:]
		pq.Size--

		return n
	}

	return -1
}

func (pq *PriorityQueue) empty() bool {
	return pq.Size == 0
}

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
