package main

import (
	"fmt"
	"sort"
)

// A - Download More RAM

type Pair struct {
	Key   int
	Value int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	var t, n, k int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		fmt.Scanf("%d %d\n", &n, &k)

		var a, b []int

		for i := 0; i < n; i++ {
			var a_i int
			fmt.Scanf("%d", &a_i)
			a = append(a, a_i)
		}

		fmt.Scanf("%d\n", nil)

		for i := 0; i < n; i++ {
			var b_i int
			fmt.Scanf("%d", &b_i)
			b = append(b, b_i)
		}

		fmt.Scanf("%d\n", nil)

		pairs := make(Pairs, len(a))

		for i := 0; i < len(a); i++ {
			p := Pair{
				Key:   a[i],
				Value: b[i],
			}
			pairs = append(pairs, p)
		}

		sort.Sort(pairs)

		for _, pair := range pairs {
			if k >= pair.Key {
				k += pair.Value
			}
		}

		fmt.Println(k)

		t--
	}
}
