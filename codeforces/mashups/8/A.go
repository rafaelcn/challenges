package main

import (
	"fmt"
	"strings"
)

// A. Enzo jogando Zoma

func main() {

	var n int
	var s string

	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &s)

	l := strings.Count(s, "L")
	r := strings.Count(s, "R")

	fmt.Println(l + r + 1)
}
