package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	//"reflect"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// See how many strings to parse
	howMany, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkErr(err)

	// Collect as many strings as indicated
	var input []string
	for i := 0; i < int(howMany); i++ {
		newLine := readLine(reader)

		input = append(input, newLine)
	}

	// Loop the inputted strings
	for i := 0; i < len(input); i++ {
		// fmt.Println(input[i])
		var odd string
		var even string
		// concact the odd and even letters of each string
		for d := range input[i] {
			if d%2 == 0 {
				// even position
				// input[current string][slice of char position]
				even = even + string(input[i][d:d+1])
			} else {
				// odd position
				odd = odd + string(input[i][d:d+1])
			}
		}
		fmt.Printf("%s %s\n", even, odd)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
