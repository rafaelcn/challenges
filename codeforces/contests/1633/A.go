package main

import (
	"fmt"
	"math"
)

// A. Div 7

func solve() {

	var n int

	fmt.Scanf("%d\n", &n)

	if n%7 == 0 {
		fmt.Println(n)
		return
	}

	digits := math.Floor(math.Log10(float64(n))) + 1

	if digits == 3 {
		for i := 100; n+i <= 999 && n < 900; i += 100 {
			if (n+i)%7 == 0 {
				fmt.Println(n + i)
				return
			}
		}

		for i := 10; i <= 100; i += 10 {
			if (n+i)%7 == 0 {
				fmt.Println(n + i)
				return
			}
		}
	}

	for i := 1; n+i <= 999; i++ {
		if (n+i)%7 == 0 {
			fmt.Println(n + i)
			break
		}
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
