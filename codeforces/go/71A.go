package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	var words []string

	fmt.Scanf("%d\n", &n)

	for n > 0 {

		var word string
		fmt.Scanf("%s\n", &word)

		words = append(words, word)

		n--
	}

	for _, word := range words {
		if len(word) <= 10 {
			fmt.Println(word)
		} else {
			size := strconv.Itoa(len(word) - 2)
			modified := word[0:1] + size + word[len(word)-1:]

			fmt.Println(modified)
		}
	}
}
