package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 1e9 + 3
)

func main() {

	var n, a, b, people, maxi, i int64 = 0, 0, 0, 0, 0, 0

	// doesn't work. Too much data allocated...
	d := make([]int64, N)
	p := make([]int64, N)

	scanf("%d\n", &n)

	for n > 0 {
		scanf("%d %d\n", &a, &b)

		d[a] += 1
		d[b] -= 1 // don't count the last year (the year the person died)

		if b > maxi {
			maxi = b
		}

		n--
	}

	var acc, year int64

	for i = 0; i < maxi; i++ {
		acc += d[i]
		p[i] = acc

		if people < p[i] {
			people = p[i]
			year = i
		}
	}

	println(year, people)
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
