package main

import (
	"fmt"
	"sort"
)

// C. Meximum Array

func main() {

	var t int

	fmt.Scanf("%d\n", &t)

	for t > 0 {

		var n int
		var a []int

		fmt.Scanf("%d", &n)

		for n > 0 {
			var na int
			fmt.Scanf("%d", &na)
			a = append(a, na)
			n--
		}

		fmt.Scanf("\n")

		sort.Ints(a)

		for i := 0; i < len(a); i++ {
			
		}


		t--
	}

}
