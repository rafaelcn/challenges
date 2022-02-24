package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A. Formigas

func main() {

	var s string

	scanf("%s\n", &s)

	total := 0
	stack := Pairs{}

	for _, r := range s {
		c := string(r)

		switch c {
		case ".":
			if len(stack) > 0 {
				stack[len(stack)-1].Second += 1
			}
		case "r":
			if len(stack) > 0 {
				total += stack[len(stack)-1].Second
				stack = stack[:len(stack)-1]
			}

			p := Pair{First: c, Second: 0}
			stack = append(stack, p)
		case "s":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	// if there are remaining frogs
	for _, p := range stack {
		total += p.Second
	}

	println(total)

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
	First  string
	Second int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i].First < p[j].First }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Util functions

func toStr(n int) string {
	return strconv.Itoa(n)
}

func toInt(s string) int {

	// ignore error
	n, _ := strconv.Atoi(s)

	return n
}
