package main

import (
	"fmt"
	"io"
	"strconv"
	"unicode"
)

func main() {
	sum := 0

	// o t t f f s s e n
	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for {
		line := ""
		fixed := "" // the fixed line

		_, err := fmt.Scanf("%s", &line)
		if err == io.EOF {
			break
		}

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				fixed += string(line[i])
				continue
			}

			// for each letter, try all numbers (which might not be that bad at)
			matched := false

			for k, v := range numbers {
				if i+len(k) > len(line) {
					continue
				}

				//fmt.Printf("i: %d, l: %d, i+len(k): %d | ", i, len(line), i+len(k))
				//fmt.Println(k, line[i:i+len(k)])
				if line[i:i+len(k)] == k {
					fixed += v
					i += len(k) - 1
					matched = true
					break
				}

			}

			if !matched {
				fixed += string(line[i])
			}

		}

		// build the result
		first := -1
		last := -1

		fmt.Println(fixed)

		for _, character := range fixed {
			if unicode.IsDigit(character) {
				if first == -1 { // first hasn't been set
					first, _ = strconv.Atoi(string(character))
				} else {
					last, _ = strconv.Atoi(string(character))
				}
			}
		}

		if last == -1 { // if the last number hasn't been set yet
			last = first
		}

		sum += (first * 10) + last

	}

	fmt.Println(sum)
}
