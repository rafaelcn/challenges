package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0

	rgb := map[string]int{"red": 12, "green": 13, "blue": 14}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.Trim(line, "\n")
		if len(line) == 0 {
			break
		}

		marked := true

		sep := strings.Index(line, ":")
		sets := strings.Split(line[sep+1:], ";")

		for _, set := range sets {
			colors := strings.Split(string(set), ",")

			//fmt.Println(colors)

			for _, color := range colors {
				information := strings.Split(strings.TrimLeft(color, " "), " ")

				n := information[0]
				c := information[1]

				//fmt.Printf("   | %s | number: %s color: %s\n", information, n, c)

				if max, ok := rgb[strings.Trim(c, " ")]; ok {
					colorNumber, _ := strconv.Atoi(n)

					//fmt.Printf("   | -> %d %d", colorNumber, max)

					if colorNumber > max {
						marked = false
					}

					//fmt.Println()
				}
			}
		}

		if marked {
			id, _ := strconv.Atoi(strings.Trim(line[4:sep], " "))
			sum += id
		}

	}

	fmt.Println(sum)
}
