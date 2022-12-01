package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	counter := 0
	prev := 0
	for scanner.Scan() {
		intVar, _ := strconv.Atoi(scanner.Text())
		if prev == 0 {
			prev = intVar
		} else {
			if intVar > prev {
				counter += 1
			}
			prev = intVar
		}
	}
	fmt.Println(counter)
}
