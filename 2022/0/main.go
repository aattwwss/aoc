package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

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

func part1() {
	fileName := "test.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}

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
