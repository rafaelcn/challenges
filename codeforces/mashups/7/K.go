package main

import (
	"fmt"
)

func main() {

	var ba, bb, ca, cb, c int64

	fmt.Scanf("%d %d %d %d\n", &ba, &bb, &ca, &cb)

	for ba-ca >= bb+cb {
		ba -= ca
		bb += cb
		c++
	}

	fmt.Println(c)
}
