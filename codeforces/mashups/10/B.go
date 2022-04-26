package main

import (
	"bufio"
	"fmt"
	"os"
)

// B. Enunciados Grandes

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func solve() {

	var s string

	I("%s\n", &s)

	if len(s) <= 10 {
		O("%s\n", s)
		return
	}

	f := string(s[0])
	l := string(s[len(s)-1])

	O("%s%d%s\n", f, len(s)-2, l)
}

func main() {

	var t int

	I("%d\n", &t)

	for t > 0 {
		solve()
		t--
	}

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
