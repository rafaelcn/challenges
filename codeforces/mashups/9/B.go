package main

import (
	"fmt"
	"sort"
)

// B. Números menores ou iguais

func main() {

	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	sort.Ints(a)

	for i := 0; i < m; i++ {
		var bi int
		fmt.Scanf("%d", &bi)

		l := 0
		r := n - 1

		ans := -1

		for l <= r {
			mid := (l + r) / 2

			if bi >= a[mid] {
				l = mid + 1
				ans = mid
			} else {
				r = mid - 1
			}
		}

		fmt.Printf("%d ", ans+1)
	}
}
