package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	l := ""
	i := 0

	elves := []int{0}

	maxI := 0
	maxCalories := 0

	for {
		_, err := fmt.Scanf("%s", &l)
		if err == io.EOF {
			break
		}

		if l == "" {
			i++
			elves = append(elves, 0)
			continue
		}

		n, _ := strconv.Atoi(l)

		elves[i] += n

		if elves[i] > maxCalories {
			maxI = i
			maxCalories = elves[i]
		}

		l = ""
	}

	fmt.Println(elves[maxI])
}
