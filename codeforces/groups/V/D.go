package main

import (
	"fmt"
	"math"
)

// D. Jogo Viciante

func main() {

	var l, r, x, y, k int

	fmt.Scanf("%d %d %d %d %d\n", &l, &r, &x, &y, &k)

	// a in [l,r]
	// b in [x,y]
	// k = a/b -> a = kb or b = a/k

	for i := l; i <= r; i++ {
		n := float64(i) / float64(k)
		//fmt.Println(n)

		if math.Trunc(n) == n {
			ni := int(n)
			if ni >= x && ni <= y {
				fmt.Println("YES")
				return
			}
		}
	}

	fmt.Println("NO")
}
