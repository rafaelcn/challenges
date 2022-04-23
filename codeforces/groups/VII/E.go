package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// E. Ajude Faustito!

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

type Pair struct {
	DragonStrength uint64
	DragonBonus    uint64
}

func main() {

	var n int
	var s uint64

	I("%d %d\n", &s, &n)

	p := make([]Pair, n)

	for i := 0; i < n; i++ {
		I("%d %d\n", &p[i].DragonStrength, &p[i].DragonBonus)
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].DragonStrength < p[j].DragonStrength
	})

	pass := true

	for i := 0; i < n; i++ {
		if s > p[i].DragonStrength {
			s += p[i].DragonBonus
		} else {
			pass = false
			break
		}
	}

	if pass {
		O("YES\n")
	} else {
		O("NO\n")
	}

	w.Flush()
}

func I(format string, a ...interface{}) {
	fmt.Fscanf(r, format, a...)
}

func O(format string, a ...interface{}) {
	fmt.Fprintf(w, format, a...)
}
