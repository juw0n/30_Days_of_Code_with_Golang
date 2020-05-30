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
	reader := bufio.NewReader(os.Stdin)

	n, err := strconv.Atoi(readLine(reader))
	checkError(err)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = readLine(reader)
	}

	for _, character := range s {
		var evenChar, oddChar string
		for k := 0; k < len(character); k++ {
			if k%2 == 0 {
				evenChar += string(character[k])
			} else {
				oddChar += string(character[k])
			}
		}
		fmt.Printf("%s %s\n", evenChar, oddChar)

	}

}

func checkError(err error) {
	if err != nil {
		panic("invalid input entered")
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
