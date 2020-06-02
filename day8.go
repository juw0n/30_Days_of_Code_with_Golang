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

	phoneBook := make(map[string]string, n)

	for i := 0; i < n; i++ {
		newName := readLine(reader)
		split := strings.Split(newName, " ")
		phoneBook[split[0]] = split[1]
	}

	for {
		existingName := readLine(reader)

		if existingName == "" {
			break
		}

		if phone, ok := phoneBook[existingName]; ok {
			fmt.Printf("%s=%s\n", existingName, phone)
		} else {
			fmt.Printf("Not found\n")
		}
	}

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
		panic("invalid input entered")
	}
}
