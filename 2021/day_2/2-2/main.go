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
	var aim, x, y int
	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Split(s, " ")
		cmd := arr[0]
		value, _ := strconv.Atoi(arr[1])
		switch cmd {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			x += value
			y += (value * aim)
		default:
			continue
		}

	}
	fmt.Println(x * y)
}
