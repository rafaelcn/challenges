package main

import "fmt"

func solve() {

	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)

	if a == 0 && b > 0 {
		fmt.Println(1)
		return
	} else if b == 0 && a > 0 {
		fmt.Println(a + 1)
		return
	}

	payA := a
	payB := b * 2

	total := payA + payB

	fmt.Println(total + 1)
}

func main() {

	var t int
	fmt.Scanf("%d\n", &t)

	for t > 0 {
		solve()
		t--
	}
}
