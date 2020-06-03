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
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int64(nTemp)

	baseTwo := strconv.FormatInt(n, 2)

	var totalCount, count int
	/*for i := 0; i < len(baseTwo); i++ {
		if string(baseTwo[i]) == "1" {
			count++
			if count > totalFig {
				totalFig = count
			}
		} else {
			count = 0
		}
	}
	fmt.Println(totalFig)*/

	for i := 0; i < len(baseTwo); i++ {
		if string(baseTwo[i]) == "1" {
			count++

			if count > totalCount {
				totalCount = count
			}
		} else {
			count = 0
		}
	}
	fmt.Println(totalCount)
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
		panic("Invalid input entered")
	}
}
