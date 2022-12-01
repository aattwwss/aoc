package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var x, y int
	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Split(s, " ")
		cmd := arr[0]
		value, _ := strconv.Atoi(arr[1])
		switch cmd {
		case "up":
			y -= value
		case "down":
			y += value
		case "forward":
			x += value
		default:
			continue
		}

	}
	fmt.Println(x * y)
}
