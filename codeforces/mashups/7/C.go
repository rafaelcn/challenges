package main

import "fmt"

func main() {

	var t, n, m int64

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		fmt.Scanf("%d %d\n", &n, &m)

		mul := n * m
		if mul&1 == 0 {
			fmt.Println(mul / 2)
		} else {
			fmt.Println((mul + 1) / 2)
		}

		t--
	}
}
