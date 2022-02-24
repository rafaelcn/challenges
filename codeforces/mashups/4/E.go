package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// E. Tamanduás vs Cupins

func main() {

	var x, y, z, capacity, enemies int

	scanf("%d %d %d\n", &x, &y, &z)

	m := map[byte]int{
		'F': x,
		'A': y,
		'R': z,
	}

	scanf("%d %d\n", &capacity, &enemies)

	var s string

	scanf("%s\n", &s)

	ts := []int{capacity}

	for i := 0; i < len(s); i++ {
		c := m[s[i]]

		for j := 0; j < len(ts); j++ {
			if ts[j]-c >= 0 {
				ts[j] -= c
				c = 0
				break
			}
		}

		if c > 0 {
			ts = append(ts, capacity)
			ts[len(ts)-1] -= c
		}
	}

	println(len(ts))

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
