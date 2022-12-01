package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	max := 0
	sum := 0
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {

			if sum > max {
				max = sum
			}
			sum = 0

		}
		num, _ := strconv.Atoi(s)
		sum += num
	}

	fmt.Println(max)
}

func part2() {
	sum := 0
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var arr []int
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			arr = append(arr, sum)
			sum = 0
		}
		num, _ := strconv.Atoi(s)
		sum += num
	}

	length := len(arr)
	sort.Ints(arr)
	ans := arr[length-1] + arr[length-2] + arr[length-3]
	fmt.Println(ans)
}
