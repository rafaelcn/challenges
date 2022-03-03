package main

import (
	"fmt"
)

// C. And Matching

type Pair struct {
	First, Second int
}

type Pairs []Pair

func main() {

	var t, n, k int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		fmt.Scanf("%d %d\n", &n, &k)

		var a []int
		var pairs Pairs

		if n > 0 {
			if k == 0 {
				pairs = append(pairs, Pair{0, 3}, Pair{1, 2})
				goto end
			} else if n/2 < k {
				goto end
			}

			for i := 0; i < n; i++ {
				a = append(a, i)
			}

			for i := 0; i < n; i++ {
				for j := i+1; j < n; j++ {

					c := a[i] != a[j]

					if len(pairs) < n/2 && !pairs.contains(a[i], a[j]) && c {
						possible := a[i] & a[j]

						// possible has to contribute to minimize k (?)
						if k > 0 {
							if possible == 0 {
								continue
							}
							k -= possible
							pairs = append(pairs, Pair{a[i], a[j]})
							i++
						} else if possible == 0 {
							pairs = append(pairs, Pair{a[i], a[j]})
							i++
						}
					}
				}
			}
		}

	end:
		if len(pairs) == 0 {
			fmt.Println("-1")
		} else {
			for _, p := range pairs {
				fmt.Println(p.First, p.Second)
			}
		}

		t--
	}
}


func (p Pairs) contains(a, b int) bool {

	m := make(map[int]int)

	for _, pair := range p {
		m[pair.First] += 1
		m[pair.Second] += 1

		c1 := pair.First == a || pair.Second == a
		c2 := pair.First == b || pair.Second == b


		if (m[pair.First] > 1 || m[pair.Second] > 1) || c1 || c2 {
			return true
		}
	}


	return false
}