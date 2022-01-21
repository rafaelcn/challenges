package main

import (
	"fmt"
)

// G. Transformação com Primo

func main() {

	var x, y, t int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		fmt.Scanf("%d %d\n", &x, &y)

		found := false
		computed := x - y

		if computed <= 1 {
			fmt.Print("NO")
			return
		}

		for i := 2; i < x/2; i++ {
			if isPrime(i) {
				if computed%i == 0 {
					found = true
					break
				}
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}

}

func isPrime(n int) bool {

	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
