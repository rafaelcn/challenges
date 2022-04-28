package main

import (
	"bufio"
	"fmt"
	"os"
)

// C.

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {

	var n, ans int

	I("%d\n", &n)

	var s string
	I("%s\n", &s)

	for i := 0; i < n-1; {
		if s[i] != s[i+1] {
			ans++
			i += 2
		} else {
			i++
		}
	}

	O("%d\n", len(s)-ans)

	w.Flush()

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
