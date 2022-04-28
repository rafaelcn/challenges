package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// E. Alfredo e o Teclado Problemático

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func solve() {
	var s string
	I("%s\n", &s)

	solution := make(map[string]int)

	for i := 0; i < len(s); i++ {
		j := i
		// increase the pointer of repeated letters.
		for j+1 < len(s) && s[j+1] == s[i] {
			j++
		}
		// if there's a repeated number of letters, such letters will always be
		// in the form 2*n, where n is the number of repeated letters.
		if (j-i)%2 == 0 {
			solution[string(s[i])]++
		}
		i = j
	}

	keys := make([]string, 0, len(solution))

	for k := range solution {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		O("%s", k)
	}

	O("\n")
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
