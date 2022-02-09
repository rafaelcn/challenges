package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {

	var x int

	scanf("%d\n", &x)

	if x == 2 {
		printf("%d %d\n", 1, 1)
		return
	}

	// a*b+gcd^2(a,b)/gcd(a, b) = x
	// if a*b > x can we break it (?)


	/*for a := 1; a < x; a++ {
		for b := 1; b < x; b++ {
			m := a*b
			d := gcd(a, b)
			e := d*d

			computed := (m + e)/d

			if m >= x || computed > x {
				break
			} else if computed == x {
				printf("%d %d\n", a, b)
				return
			}
		}
	}*/

	printf("%d %d\n", 1, x-1)
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

// Data structures

type Pair struct {
	First, Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Utils

func gcd(a, b int) int {

	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}

func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
