package main

import "fmt"

// B - GCD Arrays

func main() {

	var t int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		var l, r, k int
		fmt.Scanf("%d %d %d\n", &l, &r, &k)

		var a []int

		if l == r {
			if l > 1 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			for i := l; i <= r; i++ {
				a = append(a, i)

				if i%20 == 0 {
					a = do(a)
				}
			}

			for k > 0 {
				if len(a) >= 2 {
					i := a[0]
					j := a[1]
					a = a[2:]

					r = i * j
					a = append(a, r)
				}

				k--
			}

			if k == 0 {
				r = a[0]
				if len(a) > 1 {
					for i := 1; i < len(a); i++ {
						r = gcd(r, a[i])
					}
				}
				if r > 1 {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
			}
		}

		t--
	}
}

func gcd(a, b int) int {

	if b == 0 {
		return a
	}

	for a%b > 0 {
		r := a % b
		a = b
		b = r
	}

	return b
}

func do(a []int) []int {
	if len(a) >= 2 {
		i := a[0]
		j := a[1]
		a = a[2:]

		r := i * j
		a = append(a, r)
	}

	return a
}
