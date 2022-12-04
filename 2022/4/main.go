package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func main() {
	// part1()
	part2()
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func genStringFromInputRange(s string) []int {
	arr := strings.Split(s, "-")
	left, _ := strconv.Atoi(arr[0])
	right, _ := strconv.Atoi(arr[1])
	res := []int{}
	for i := left; i <= right; i++ {
		res = append(res, i)
	}
	return res
}

func isIntArrayInAnother(a1, a2 []int) bool {
	return a1[0] <= a2[0] && a1[len(a1)-1] >= a2[len(a2)-1]
}

func isIntArraySomeOverlapped(a1, a2 []int) bool {

	m := map[int]bool{}
	for i := range a1 {
		m[a1[i]] = true
	}
	for i := range a2 {
		if m[a2[i]] {
			return true
		}
	}
	return false
}

func isOverlapped(s1, s2 []int) bool {

	return isIntArrayInAnother(s1, s2) || isIntArrayInAnother(s2, s1)

}

func hasSomeOverlap(s1, s2 []int) bool {

	return isIntArraySomeOverlapped(s1, s2) || isIntArraySomeOverlapped(s2, s1)

}

func part1() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Split(s, ",")
		first := genStringFromInputRange(arr[0])
		second := genStringFromInputRange(arr[1])
		if isOverlapped(first, second) {
			score++
		}
	}
	fmt.Println(score)

}

func part2() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Split(s, ",")
		first := genStringFromInputRange(arr[0])
		second := genStringFromInputRange(arr[1])
		if hasSomeOverlap(first, second) {
			score++
		}
	}
	fmt.Println(score)
}
