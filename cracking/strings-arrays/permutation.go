package main

import "fmt"

func main() {
	var s string
	var p string

	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &p)

	fmt.Println(permutation(s, p))
}

func permutation(s, p string) bool {
	if len(s) != len(p) {
		return false
	}

	c := [2]int{}

	for i := 0; i < len(s); i++ {
		c[0] |= (1 << (s[i] - 'a'))
		c[1] |= (1 << (p[i] - 'a'))
	}

	return c[0] == c[1]
}
