package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math/big"
)

func factorial(n int64, acc *big.Int) *big.Int {
    if n <= 1 {
        return acc
    }
    
    return factorial(n-1, acc.Mul(acc, big.NewInt(n)))
}

func extraLongFactorials(n int32) {
    fmt.Println(factorial(int64(n), big.NewInt(1)))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    extraLongFactorials(n)
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

