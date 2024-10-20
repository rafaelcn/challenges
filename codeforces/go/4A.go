package main

import (
	"fmt"
)

func main() {
	var w int

	fmt.Scanf("%d", &w)

	// e + e = e
	// o + o = e
	// e + o = o

	if w&1 == 1 {
		fmt.Print("NO")
	} else {
		if factors(w) {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	}
}

func factors(n int) bool {

	for i := 2; i < n; i += 2 {
		for j := 2; j+i <= n; j += 2 {
			if i+j == n {
				return true
			}
		}
	}

	return false
}
