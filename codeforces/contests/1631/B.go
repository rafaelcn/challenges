package main

import (
	"fmt"
)

// B. 

func main() {

	var t, n int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		fmt.Scanf("%d\n", &n)

		if n == 1 {
			fmt.Println("0")
			continue
		}

		var a []int

		for i := 0; i < n; i++ {
			var ai int
			fmt.Scanf("%d", &ai)

			a = append(a, ai)
		}

		fmt.Scanf("\n")

		fmt.Println("")

	}
}
