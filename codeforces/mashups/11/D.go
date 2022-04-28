package main

import (
	"bufio"
	"fmt"
	"os"
)

// D. Ano da Sorte

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n int
	I("%d\n", &n)

	ans, digits := 1, split(n)

	if n >= 10 {
		ans = (digits[len(digits)-1]+1)*pow(10, len(digits)-1) - n
	}

	O("%d\n", ans)

	w.Flush()
}

func split(a int) []int {

	digits := []int{}

	for a > 0 {
		d := a % 10
		digits = append(digits, d)
		a = a / 10
	}

	return digits
}

func pow(b, e int) int {
	if e == 0 {
		return 1
	}

	return b * pow(b, e-1)
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
