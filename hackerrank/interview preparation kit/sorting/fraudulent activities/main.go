package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Int32Slice []int32

func (a Int32Slice) Len() int           { return len(a) }
func (a Int32Slice) Less(i, j int) bool { return a[i] < a[j] }
func (a Int32Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func activityNotifications(expenditure []int32, d int32) int32 {

	var i int32
	var size int32
	var median int32
	var notices int32

	size = int32(len(expenditure))

	for i = 0; i+d <= size-1; i++ {
		segment := make(Int32Slice, d)
		copy(segment, expenditure[i:i+d])

		sort.Sort(segment)

		// calculate median of the segment
		if len(segment)%2 == 0 {
			m := len(segment) / 2
			s := segment[m : m+2]

			median = s[0] + s[1]/2
		} else {
			m := math.Floor(float64(len(segment)) / 2)
			median = segment[int32(m)]
		}

		if expenditure[i+d] >= 2*median {
			notices++
		}
	}

	return notices
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
