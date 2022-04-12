package main

import (
	"fmt"
	"sort"
)

// D. Jogo de cartas pra dois

func main() {

	var n int
	fmt.Scanf("%d", &n)

	a := make([]int64, n)
	b := make([]int64, 2)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	for i := 0; i < n; i++ {
		b[i%2] += a[i]
	}

	fmt.Println(b[0]-b[1])
}
