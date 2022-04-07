package main

import (
	"fmt"
	"sort"
)

// D. Olímpiada

type Tuple struct {
	Number    int
	Index     int
}

type Tuples []Tuple

func main() {

	var n, ni, shots int
	fmt.Scanf("%d\n", &n)

	tuples := make(Tuples, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ni)

		tuples[i].Index = i
		tuples[i].Number = ni
	}
	//fmt.Scanf("\n")

	// sort first by the number
	sort.Slice(tuples, func(i, j int) bool { return tuples[i].Number > tuples[j].Number })

	for i, v := range tuples {
		shots += i*v.Number + 1
	}

	fmt.Println(shots)
	for _, t := range tuples {
		fmt.Printf("%d ", t.Index+1)
	}
}
