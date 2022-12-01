package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	sort.Ints(arr)
	fmt.Println(arr)
	pos := arr[len(arr)/2]
	fmt.Println(sumArrayFromPoint(arr, pos))
}

func sumArrayFromPoint(arr []int, num int) int {
	var sum int
	for i := range arr {
		temp := math.Abs(float64(arr[i] - num))
		sum += int(temp)

	}
	return sum
}
