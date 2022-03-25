package main

import (
	"fmt"
)

// F. O Smartphone e a Fila

func main() {

	var s1, s2 string

	fmt.Scanf("%s\n%s\n", &s1, &s2)

	v1 := 0
	v2 := 0

	unknown := 0

	for _, c := range s1 {
		if c == '+' {
			v1++
		} else {
			v1--
		}
	}
	for _, c := range s2 {
		if c == '+' {
			v2++
		} else if c == '-' {
			v2--
		} else {
			unknown++
		}
	}

	var probability float64

	if unknown == 0 {
		if v1 == v2 {
			probability = 1.0
		} else {
			probability = 0.0
		}
	} else {
		states := 1 << unknown
		correct := 0.0
		possibilities := []Tuple{}

		q := Queue{}
		q.enqueue(Tuple{Possibility: "+", Value: 1})
		q.enqueue(Tuple{Possibility: "-", Value: -1})

		for q.size() > 0 {
			t := q.front()
			q.pop()

			if len(t.Possibility) < unknown {
				q.enqueue(Tuple{t.Possibility + "+", t.Value+1})
				q.enqueue(Tuple{t.Possibility + "-", t.Value-1})
			} else {
				possibilities = append(possibilities, t)
			}
		}

		for _, possibility := range possibilities {
			if possibility.Value+v2 == v1 {
				correct++
			}
		}

		//fmt.Println(possibilities, states, correct, unknown)

		probability = correct / float64(states)
	}

	fmt.Printf("%.12f\n", probability)
}

type Tuple struct {
	Possibility string
	Value int
}

type Queue struct {
	v    []Tuple
	Size uint
}

func (q *Queue) enqueue(c Tuple) {

	q.v = append(q.v, c)
	q.Size++
}

func (q *Queue) front() Tuple {
	return q.v[0]
}

func (q *Queue) pop() {

	if len(q.v) > 0 {
		q.v = q.v[1:]
		q.Size--
	}
}

func (q *Queue) size() uint {
	return q.Size
}
