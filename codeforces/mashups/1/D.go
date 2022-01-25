package main

import "fmt"

// D. Wallace e as Subsequências

func main() {

	var t int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		b := make([]int, 7)

		fmt.Scanf("%d %d %d %d %d %d %d\n", &b[0], &b[1], &b[2], &b[3], &b[4],
			&b[5], &b[6])

		var a1, a2, a3 int

		a3 = b[6] - b[5]
		a2 = b[6] - b[4]

		if b[0]+b[1]+b[2] == b[6] {
			a1 = b[2]
		} else {
			a1 = b[6] - b[2]
		}

		fmt.Println(a3, a2, a1)

		t--
	}
}
