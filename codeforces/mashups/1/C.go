package main

import "fmt"

// C. Strings Similares

func main() {

	var t int

	fmt.Scanf("%d\n", &t)

	for t > 0 {
		var n int
		var s string

		fmt.Scanf("%d\n", &n)
		fmt.Scanf("%s\n", &s)

		if n == 1 {
			fmt.Println(s)
		} else {
			var w string

			for i := 0; i < len(s); i += 2 {
				w += string(s[i])
			}

			fmt.Println(w)
		}

		t--
	}
}
