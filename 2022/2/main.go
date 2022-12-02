package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

func main() {
	part1()
	// part2()
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func calculateScore(opponent string, me string) int {
	if (opponent == "A" && me == "X") || (opponent == "B" && me == "Y") || (opponent == "C" && me == "Z") {
		return 3
	}

	if (opponent == "A" && me == "Y") || (opponent == "B" && me == "Z") || (opponent == "C" && me == "X") {
		return 6
	}
	return 0
}

func handScore(me string) int {
	switch me {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		return 0

	}
}

func part1() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	score := 0
	counter := 0
	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Split(s, "")
		points := calculateScore(arr[0], arr[2])
		handScore := handScore(arr[2])
		score += (points + handScore)
		fmt.Println(points + handScore)
		counter++
		if counter == 111111110 {
			break
		}
	}
	fmt.Println(score)
}

func part2() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}

}
