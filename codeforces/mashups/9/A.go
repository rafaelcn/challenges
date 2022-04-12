package main

import "fmt"

// A. Três

func main() {

	var n, a, b, c int
	fmt.Scanf("%d", &n)

	rest := n - 2
	if rest%3 == 0 {
		a = n-3
		b = 1
		c = 2
	} else {
		a = rest
		b = 1
		c = 1
	}

	fmt.Printf("%d %d %d\n", a, b, c)
}
