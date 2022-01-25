package main

import (
	"fmt"
	"math"
)

// B. Corretor Automático

func main() {

	var n int
	var grade float64

	fmt.Scanf("%d\n", &n)

	// save state of questions
	questions := n

	for n > 0 {
		var answer, correctAnswer string

		fmt.Scanf("%s %s\n", &answer, &correctAnswer)

		if answer == correctAnswer {
			grade += 1
		} else if grader(answer, correctAnswer) {
			grade += 1
		}

		n--
	}

	result := (grade / float64(questions))*10

	fmt.Printf("Nota: %.2f", result)
}

func grader(y, z string) bool {

	mapper := map[string]float64{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}

	mappedY := mapper[y]
	mappedZ := mapper[z]

	return math.Abs(mappedY-mappedZ)/math.Max(1, mappedZ) <= math.Pow(10, -2)
}
