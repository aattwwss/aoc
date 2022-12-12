package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}
}
