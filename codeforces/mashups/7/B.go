package main

import "fmt"

func main() {
	var n, x, y, z int64

	fmt.Scanf("%d\n", &n)

	for n > 0 {
		var a, b, c int64
		fmt.Scanf("%d %d %d\n", &a, &b, &c)

		x += a
		y += b
		z += c
		n--
	}

	if x == y && x == z {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
