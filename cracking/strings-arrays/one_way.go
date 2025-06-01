package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. if they are the same size (only edits are allowed)
	// 2. if they have different size they should differ only by one character
	words := [2]string{}

	fmt.Scanf("%s", &words[0])
	fmt.Scanf("%s", &words[1])

	fmt.Println(permute(words))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func permute(words [2]string) bool {
	difference := abs(len(words[0]) - len(words[1]))

	if difference > 1 {
		return false
	} else if difference == 0 {
		// check for how many letters each word differ, if it's greater than 1
		// then it's not possible to make the words equal
		if (differ(words)) <= 1 {
			return true
		}
	}

	return false
}

// compute by how much characters the words differ considering the input are only
// ASCI encoded strings (one byte per character)
func differ(words [2]string) int {
	c := [2]uint{}

	for i := 0; i < len(words[0]); i++ {
		c[0] |= (1 << (words[0][i] - 'a'))
		c[1] |= (1 << (words[1][i] - 'a'))
	}

	b := 0
	x := c[0] ^ c[1]

	//fmt.Println(decode(x), decode(c[0]), decode(c[1]))

	// count how many bits are one from the result of the XOR between the two words
	for i := 0; i < 32; i++ {
		if (x>>i)&1 == 1 {
			b++
		}
	}

	// the difference will always differ by two bits if there's only one character
	// different from one word to another so I have to compensate for that difference
	// subtracting one unit
	return b - 1
}

func decode(x uint) string {
	s := strings.Builder{}

	for i := 0; i < 32; i++ {
		s.Write([]byte{byte((x>>i)&1) + 48})
	}

	return s.String()
}
