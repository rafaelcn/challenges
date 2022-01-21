package main

import "fmt"

// A. Promoção ou Engação?

func main() {

	var n, r, m, p float64

	fmt.Scanf("%f %f\n", &n, &r)
	fmt.Scanf("%f %f\n", &m, &p)

	if n > 1 {
		r = r/n
	}

	// Se a mesma quantidade de itens comprados a preço regular for menor do que
	// o preço em promoção então isso é uma enganação.

	if r*m > p {
		fmt.Printf("Promocao")
	} else {
		fmt.Printf("Enganacao")
	}
}
