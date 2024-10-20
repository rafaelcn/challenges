package main

import "fmt"

func main() {
	var solutions int
	var results []int

	fmt.Scanf("%d\n", &solutions)

	for solutions > 0 {
		var n, m, p int
		fmt.Scanf("%d %d %d\n", &n, &m, &p)

		results = append(results, n, m, p)

		solutions--
	}

	computed := 0
	applicants := 0

	for i := 0; i+3 <= len(results); i += 3 {

		// 0 1 2 | 3 4 5 | 6 7 8
		// 1 1 0 | 0 1 0 | 1 0 1
		// 0+3   | 3+6   | 6+3

		segment := results[i : i+3]

		for _, v := range segment {
			if v > 0 {
				computed++
			}
		}

		if computed >= 2 {
			applicants++
		}

		computed = 0
	}

	fmt.Print(applicants)

}
