package main

import "fmt"

func main() {
	m, n := 0, 0

	fmt.Scanf("%d %d", &m, &n)

	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			o := 0
			fmt.Scanf("%d", &o)

			matrix[i][j] = o
		}
	}

	zero(matrix)

	print(matrix)
}

func zero(matrix [][]int) {
	lines := []int{}
	columns := []int{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				lines = append(lines, i)
				columns = append(columns, j)
			}
		}
	}

	fmt.Println(lines, columns)

	for _, line := range lines {
		matrix[line] = make([]int, len(matrix))
	}

	for _, column := range columns {
		for i := 0; i < len(matrix); i++ {
			matrix[i][column] = 0
		}
	}
}

func print(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
