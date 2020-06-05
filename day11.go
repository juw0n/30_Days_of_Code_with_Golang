package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	var maxValue float64

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr)-2; j++ {
			topRow := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			middleRow := arr[i+1][j+1]
			bottomRow := arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			hourGlassSum := topRow + middleRow + bottomRow

			maxValue = math.Max(maxValue, float64(hourGlassSum))
		}
	}
	fmt.Println(maxValue)
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
