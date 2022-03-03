package main

import "fmt"

// D. Make Them Equal

func solve() {

	var n, k int

	fmt.Scanf("%d %d\n", &n, &k)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)

	result := 0

	for i := 0; i < n; i++ {
		var bi int
		fmt.Scanf("%d", &bi)

		a[i] = 1
		b[i] = bi
	}

	fmt.Scanf("\n")

	for i := 0; i < n; i++ {
		var ci int
		fmt.Scanf("%d", &ci)

		c[i] = ci
	}

	fmt.Scanf("\n")

	for i := 0; i < n; i++ {
		for j := 1; j < 1e3; j++ {
			if b[i] == 1 {
				result += c[i]
				break
			} else if a[i]+a[i]/j == b[i] {
				result += c[i]
				break
			}
		}
	}

	fmt.Println(result)
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
