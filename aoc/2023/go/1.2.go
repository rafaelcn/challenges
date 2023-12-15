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
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for {
		line := ""

		n, err := fmt.Scanf("%s", &line)
		if err == io.EOF || n == 0 {
			break
		}

		j := 0
		fl := [2]int{}

		for i := 0; i < len(line); i++ {
			// check before because of the continue condition below
			if j == 0 && fl[j] != 0 {
				j++
			}

			if unicode.IsDigit(rune(line[i])) {
				fl[j], _ = strconv.Atoi(string(line[i]))
				continue
			}

			for k, v := range numbers {
				if i+len(k) > len(line) {
					continue
				}

				//fmt.Printf("i: %d, l: %d, i+len(k): %d | ", i, len(line), i+len(k))
				//fmt.Println(k, line[i:i+len(k)])
				if line[i:i+len(k)] == k {
					i += len(k) - 2 // numbers that share a letter will be considered!
					fl[j] = v
					break
				}
			}
		}

		if fl[1] == 0 {
			fl[1] = fl[0]
		}

		sum += fl[0]*10 + fl[1]
	}

	fmt.Println(sum)
}
