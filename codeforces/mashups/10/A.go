package main

import (
	"bufio"
	"fmt"
	"os"
)

// A. De graça, até injeção na testa!

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var s string
	var t, vaccines, number, groups int64

	I("%d %d\n", &t, &vaccines)

	for t > 0 {
		I("%s %d\n", &s, &number)

		switch s {
		case "-":
			//D(number, vaccines)
			if vaccines-number < 0 {
				groups++
			} else {
				vaccines -= number
			}
		case "+":
			vaccines += number
		}
		t--
	}

	O("%d %d\n", vaccines, groups)

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}

// for debug purposes
const hostname = "elementary"

func D(args ...interface{}) {
	if h, err := os.Hostname(); h == hostname && err == nil {
		fmt.Fprintln(os.Stderr, args...)
	}
}