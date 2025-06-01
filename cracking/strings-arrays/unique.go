package main

import "fmt"

func main() {
	var input string

	fmt.Scanf("%s", &input)
	fmt.Println(unique(input))
}

func unique(s string) bool {
	m := 0
	c := 0

	for i := range s {
		m = 1 << (s[i] - 'a')

		if m^c < c {
			return false
		}

		c |= m
	}

	return true
}

// if the c variable (which is the cache) XOR with the
// current bit mask is smaller than c itself we know we
// have seen this character before. that's why it works.
