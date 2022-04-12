package main

import "fmt"

// C. Treino do BamBam

type Pair struct {
	Value int
	Index int
}

func main() {

	var n int
	max := Pair{Value: 0, Index: 0}

	fmt.Scanf("%d", &n)

	a := make([]int, 3)

	for i := 0; i < n; i++ {
		var ai int
		fmt.Scanf("%d", &ai)

		a[i%3] += ai

		if a[i%3] > max.Value {
			max.Index = i%3
			max.Value = a[i%3]
		}
	}

	//fmt.Println(a[0], a[1], a[2], max.Index, max.Value)

	switch max.Index {
	case 0:
		fmt.Println("chest")
	case 1:
		fmt.Println("biceps")
	case 2:
		fmt.Println("back")
	}

}
