package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var n, i int

	scanf("%d\n", &n)

	x, y := 0, 0
	a := make([]int, n)

	for i < n {
		scanf("%d", &a[i])
		i++
	}

	sort.Ints(a)
	i = 0

	for i < n {
		if i < n/2 {
			y += a[i]
		} else {
			x += a[i]
		}
		i++
	}

	printf("%d\n", (x*x + y*y))

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
