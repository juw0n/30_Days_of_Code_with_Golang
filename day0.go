package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //create new reader, assuming bufio imported
	input, _ := reader.ReadString('\n')
	fmt.Println("Hellow, World.")
	fmt.Println(input)

	//fmt.Println("Hellow, World.\n", input)
}
