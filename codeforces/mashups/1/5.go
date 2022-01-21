package main

import "fmt"

// E. As Meias de Ana Paula

func main() {

	var n, m, days int

	fmt.Scanf("%d %d\n", &n, &m)

	for n > 0 {
		days++

		if days % m == 0 {
			n++
		}

		n--
	}

	fmt.Print(days)
}
