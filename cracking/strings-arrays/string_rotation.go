package main

import "fmt"

func main() {
	words := [2]string{}

	fmt.Scanf("%s", &words[0])
	fmt.Scanf("%s", &words[1])

	fmt.Println(rotate(words))
	fmt.Println(rotate2(words))
}

// finds out if word[1] is a rotated version of word[0]
func rotate(words [2]string) bool {
	s := words[0]
	p := words[1]

	// O(2n) == O(n)
	for i := range p {
		if s[0] == p[i] {
			j := 0
			for k := i; k < len(p)+i; k++ {
				if s[j] != p[k%len(p)] {
					return false
				}
				j++
			}
		}
	}

	return true
}

func rotate2(words [2]string) bool {
	s := words[0]
	p := words[1] + words[1]

	for i := range p {
		if s[0] == p[i] {
			for j := 0; j < len(s); j++ {
				if s[j] != p[(i+j)%len(p)] {
					return false
				}
			}
		}
	}

	return true
}
