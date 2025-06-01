package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Scanf("%s", &input)
	fmt.Println(rle(input))
}

func rle(input string) string {
	encoded := strings.Builder{}
	quantity := 0
	size := len(input)

	for i := 0; i < size; i++ {
		quantity++

		if i+1 >= size || input[i] != input[i+1] {
			encoded.WriteByte(input[i])
			encoded.WriteString(strconv.Itoa(quantity))
			quantity = 0
		}
	}

	if encoded.Len() >= len(input) {
		return input
	}

	return encoded.String()
}
