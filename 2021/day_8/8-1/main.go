package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")[1]
		displays := strings.Split(line, " ")

		for i := range displays {
			length := len(displays[i])
			if length == 2 || length == 4 || length == 3 || length == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}
