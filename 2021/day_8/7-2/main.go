package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var arr []int
	for scanner.Scan() {
		fishesTimerString := strings.Split(scanner.Text(), ",")
		for _, fishString := range fishesTimerString {
			num, _ := strconv.Atoi(fishString)
			arr = append(arr, num)
		}

	}
	avg := (float64(sumArray(arr)) / float64(len(arr)) / 1.0)
	fmt.Println(avg)
	fmt.Println(sumArrayFromPoint(arr, int(avg)))
}

func sumArrayFromPoint(arr []int, num int) int {
	var sum int
	for i := range arr {
		temp := math.Abs(float64(arr[i] - num))
		sumToPoint := sumToPoint(int(temp))
		sum += int(sumToPoint)

	}
	return sum
}

func sumArray(arr []int) int {
	var sum int
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func sumToPoint(p int) int {
	var sum int
	for p != 0 {
		sum += p
		p--
	}
	return sum
}
