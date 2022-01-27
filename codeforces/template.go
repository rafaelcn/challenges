package main

import "fmt"

func solve() {

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
