package main

import (
	"bufio"
	"fmt"
	"os"
)

// A. Substrings distintas

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n int
	var s string

	I("%d\n", &n)
	I("%s\n", &s)

	t := 0
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		c := s[i]

		if _, ok := m[c]; ok {
			if i > 0 && s[i-1] == c {
				m[c] = m[c] + 1
				t++
			}
		} else {
			m[c] = 1
			t++
		}
	}

	for k, v := range m {
		fmt.Println(string(k), v)
	}

	O("%d\n", t)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
