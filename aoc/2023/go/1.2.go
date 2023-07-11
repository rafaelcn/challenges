package main

import (
	"fmt"
	"io"
	"strconv"
)

// merge sort
//
// [2,3,1,3,5,9,7]
//
// [2,3,1] + [3,5,9,7]
// 
// 2 3 3 1 5 9 7
//
// [2,3,3] + [1,5,9,7]
//
// [1,2,3,3,5,9,7]
//
// [1,2,3] + [3,5,    9,7]
//           [3,5] + [9,7]
//           [3,5] + [7,9]
//           [3,5,    7,9]


func sort(a []int) []int {
		l := len(a)
		h := l/2

		if h == 1 {
				return merge([]int{a[0]}, []int{a[1]})
		}

		pa := a[0:h]
		pb := a[h:l]


		return merge(sort(pa), sort(pb))
}

func merge(a, b []int) []int {
		s := []int{}

		la := len(a)
		lb := len(b)

		l := 0
		r := 0

		// two pointers solution
		for ; l < la && r < lb; {
				if a[l] < b[r] {
						s = append(s, a[l])
						l++
				} else {
						s = append(s, b[r])
						r++
				}
		}

		for l < la {
				s = append(s, a[l])
				l++
		}

		for r < lb {
				s = append(s, b[r])
				r++
		}

		//fmt.Printf("after merge a: %v b: %v -> s: %v\n", a, b, s)

		return s
}

func main() {
		l := ""
		i := 0

		elves := []int{0}

		for {
				_, err := fmt.Scanf("%s", &l)
				if (err == io.EOF) {
						break
				}

				if l == "" {
						i++
						elves = append(elves, 0)
						continue
				}

				n, _ := strconv.Atoi(l)

				elves[i] += n
				
				l = ""
		}

		sorted := sort(elves)
		s := len(sorted)

		sum := sorted[s-1]+sorted[s-2]+sorted[s-3]

		fmt.Println(sum)
}
