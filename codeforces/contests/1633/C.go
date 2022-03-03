package main

import "fmt"

// C. Kill the Monster

func solve() {
	var hc, hm, dc, dm, k, w, a int

	fmt.Scanf("%d %d\n", &hc, &hm)
	fmt.Scanf("%d %d\n", &dc, &dm)
	fmt.Scanf("%d %d %d\n", &k, &w, &a)

	thc := hc
	thm := hm

	turns := 0
	base := false

	// base case
	for i := 0; i < 100; i++ {
		turns++

		thm = hm - dc

		if thm <= 0 && hc > 0 {
			base = true
			break
		}

		thc = hc - dm

		if thc <= 0 {
			break
		}
	}

	if base {
		fmt.Println("YES")
	} else {
		// upgrades
	}
}

func main() {

	var t int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		solve()
		t--
	}
}

// Data structures

type Pair struct {
	First, Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
