package main

import (
	"bufio"
	"fmt"
	"os"
)

// B. Mexe com XOR

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func solve() {

	var a, b int
	I("%d %d\n", &a, &b)

	xor := 0

	switch a % 4 { // 1 <= a <= 3
	case 1:
		xor = a - 1
	case 2:
		xor = 1
	case 3:
		xor = a
	}

	if xor == b {
		O("%d\n", a)
	} else if b^xor != a {
		O("%d\n", a+1)
	} else {
		O("%d\n", a+2)
	}
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

// for debug purposes
var (
	localHostname   = "elementary"
	juryHostname, _ = os.Hostname()
)

func D(args ...interface{}) {
	if localHostname == juryHostname {
		fmt.Fprintln(os.Stderr, args...)
	}
}
