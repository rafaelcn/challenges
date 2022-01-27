package main

import (
	"fmt"
	"sort"
)

func main() {

	var t, n int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		fmt.Scanf("%d\n", &n)

		var a, b []int

		for i := 0; i < n; i++ {
			var ai int
			fmt.Scanf("%d", &ai)

			a = append(a, ai)
		}

		fmt.Scanf("\n")

		for i := 0; i < n; i++ {
			var bi int
			fmt.Scanf("%d", &bi)

			b = append(b, bi)
		}

		fmt.Scanf("\n")

		for i := 0; i < n; i++ {
			if b[i] > a[i] {
				tmp := b[i]
				b[i] = a[i]
				a[i] = tmp
			}
		}

		sort.Ints(a)
		sort.Ints(b)

		fmt.Println(b[n-1] * a[n-1])

		t--
	}
}
