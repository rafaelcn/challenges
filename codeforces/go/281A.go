package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	reader *bufio.Reader = bufio.NewReader(os.Stdin)
)

func solve() {

}

func main() {
	var word string

	fmt.Fscanf(reader, "%s\n", &word)

	b := word[0]

	if b >= 97 && b <= 122 {
		b -= 32
	}

	fmt.Fprintf(writer, "%s\n", string(b)+word[1:])

	writer.Flush()
}
