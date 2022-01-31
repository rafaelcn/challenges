package main

import (
	"fmt"
)

// B. Minority

func solve() {

	var s string

	fmt.Scanf("%s\n", &s)

	if len(s) > 2 {
		num0 := 0
		num1 := 0

		max := 0

		for _, c := range s {
			ch := string(c)

			if ch == "0" {
				num0++
			} else if ch == "1" {
				num1++
			}

			if num1 > num0 {
				max = num0
			} else if num1 < num0 {
				max = num1
			}
		}

		fmt.Println(max)
	} else {
		fmt.Println("0")
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
