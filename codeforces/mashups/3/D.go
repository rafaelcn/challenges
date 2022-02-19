package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// D. Sequência Leia e Fale

var (
	table = []string{"1", "11", "21", "1211", "111221"}
)

func generate(n int) string {

	if n <= len(table) {
		return table[n-1]
	}

	// generate numbers of the table
	for i := len(table)-1; i < n-1; i++ {
		last := string(table[i][0])
		number := ""
		counter := 0

		for _, c := range table[i] {
			s := string(c)

			if last == s {
				counter++
			} else {
				number += strconv.Itoa(counter) + last
				counter = 1
			}

			last = s
		}

		number += strconv.Itoa(counter) + last

		table = append(table, number)
	}

	return table[n-1]
}

func main() {

	var n int

	scanf("%d\n", &n)

	println(generate(n))

	writer.Flush()
}

// Use buffering when reading/writing data as the default fmt.Scanx/fmt.Printx
// functions doesn't.

var (
	writer = bufio.NewWriter(os.Stdout)
	reader = bufio.NewReader(os.Stdin)
)

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}

func println(f ...interface{}) {
	fmt.Fprintln(writer, f...)
}

// Data structures

type Pair struct {
	First, Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }