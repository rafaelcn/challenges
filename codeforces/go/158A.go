package main

import "fmt"

func main() {

	n, k := 0, 0

	fmt.Scanf("%d %d\n", &n, &k)

	var grades []int

	for n > 0 {
		var grade int
		fmt.Scanf("%d", &grade)

		grades = append(grades, grade)

		n--
	}

	advancers := 0

	for _, grade := range grades {
		if grade >= grades[k-1] && grade > 0 {
			advancers++
		}
	}

	fmt.Print(advancers)
}
