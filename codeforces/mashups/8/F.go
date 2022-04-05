package main

import (
	"fmt"
)

// F. Quirino, o mago

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

	var n, m, k int
	fmt.Scanf("%d %d %d\n", &n, &m, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	d := int(1e4)

	for i := 0; i < n; i++ {
		if a[i] > 0 && a[i] <= k {
			if abs(m-(i+1))*10 < d {
				d = abs(m-(i+1)) * 10
			}
		}
	}

	fmt.Println(d)
}
