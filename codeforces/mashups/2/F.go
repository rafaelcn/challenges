package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type ll int64
type lla []ll

func main() {

	var n, i int64

	scanf("%d\n", &n)

	var x, y ll

	x, y = 0, 0
	a := make(lla, n)

	for i < n {
		scanf("%d", &a[i])
		i++
	}

	sort.Sort(a)
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

//

func (a lla) Len() int           { return len(a) }
func (a lla) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a lla) Less(i, j int) bool { return a[i] < a[j] }
