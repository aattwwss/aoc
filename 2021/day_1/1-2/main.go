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
	var nums []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	prev := 0
	counter := 0
	for i := 0; i < len(nums)-2; i++ {
		currSum := 0
		// probably can dont need to computer this everytime, can reuse the prev to calculate currSum, but im lazy
		for j := 0; j < 3; j++ {
			currSum += nums[i+j]
		}
		if prev != 0 && currSum > prev {
			counter += 1
		}
		prev = currSum

	}
	fmt.Println(counter)
}
