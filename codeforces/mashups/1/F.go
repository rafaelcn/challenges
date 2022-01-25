package main

import (
	"fmt"
)

// F. Número Estranho

func main() {

	var input string
	var output string
	var strangers []string = []string{"144", "14", "1"}

	fmt.Scanf("%s\n", &input)

	for i := 0; i <= len(input) || output == input; {

		for j, strange := range strangers {
			if i+len(strange) <= len(input) {
				if input[i:i+len(strange)] == strange {
					output += strange
					i += len(strange)

					if i >= len(input) {
						goto end
					}

					break
				}
			}

			if j == len(strangers)-1 {
				goto end
			}
		}
	}
	
end:
	if output == input {
		fmt.Print("YES")
		return
	}

	fmt.Print("NO")
}
