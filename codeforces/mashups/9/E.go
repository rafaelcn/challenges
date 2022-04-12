package main

import "fmt"

// E. O Palíndromo do Alfabeto

func main() {
	var s string
	fmt.Scanf("%s", &s)

	changes := 0

	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {

		//fmt.Println(string(s[i]), string(s[j]), changes)

		if s[i] != s[j] {
			changes++
		}
	}

	if changes == 1 || (changes == 0 && len(s)&1 == 1) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
