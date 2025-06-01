package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Println(encode("Mr John Smith     ", 13))
}

func encode(url string, length int) string {
	encoded := []byte{}

	fmt.Println("encoding", url)

	for i := range url {
		if length == 0 {
			break
		}

		length--
		if url[i] == ' ' {
			encoded = append(encoded, []byte("%20")...)
			continue
		}

		encoded = append(encoded, url[i])
	}

	return string(encoded)
}
