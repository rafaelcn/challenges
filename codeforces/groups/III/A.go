package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// A. Transferindo Prisioneiros

func main() {

	var n, t, c int

	scanf("%d %d %d\n", &n, &t, &c)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		scanf("%d", &a[i])
	}
	scanf("\n")

	ans := 0

	// non optimal solution (n^2*logn)
	for i := 0; i <= n-1 && i+c <= n; i++ {					// O(n)
		s := []int{}

		s = append(s, a[i:i+c]...)
		sort.Ints(s)										// O(n*logn)

		if len(s) == c {
			r := sort.Search(len(s), func(i int) bool {		// O(logn)
				return s[i] > t
			})

			if r >= len(s) {
				ans++
			}
		}
	}

	println(ans)

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
