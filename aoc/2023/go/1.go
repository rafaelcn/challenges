package main

import (
	"fmt"
	"io"
	"strconv"
	"unicode"
)

func main() {
	sum := 0

	for {
		line := ""

		n, err := fmt.Scanf("%s", &line)
		if err == io.EOF || n == 0 {
			break
		}

		first := -1
		last := -1

		// naive solution
		for _, character := range line {
			if unicode.IsDigit(character) {
				if first == -1 {
					first, _ = strconv.Atoi(string(character))
				} else {
					last, _ = strconv.Atoi(string(character))
				}

			}
		}

		if last == -1 {
			last = first
		}

		sum += (first * 10) + last
	}

	fmt.Println(sum)
}
