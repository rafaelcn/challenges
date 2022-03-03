package main

import (
	"fmt"
	"sort"
)

func main() {

	var t int
	fmt.Scanf("%d\n", &t)

	for t > 0 {
		var n int
		fmt.Scanf("%d\n", &n)

		var s []string
		var si string

		for n > 0 {
			fmt.Scanf("%s\n", &si)
			s = append(s, si)
			n--
		}

		if len(s) > 2 {
			// TODO: look for palindrome substrings
		} else {
			if palindrome(s[0] + s[1]) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}

		t--
	}
}

func palindrome(s string) bool {

	c := true

	for i, j := 0, len(s)-1; i < len(s); i++ {
		if s[i] != s[j] {
			c = false
			break
		}

		j--
	}

	return c
}
