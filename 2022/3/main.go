package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

const abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	// part1()
	part2()
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func getPriority(letter string) int {
	return strings.Index(abc, letter) + 1
}

func mapStringIntoMap(s string) map[byte]bool {
	m := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = true
	}
	return m
}
func part1() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		s := scanner.Text()
		first := s[:len(s)/2]
		firstMap := mapStringIntoMap(first)
		second := s[len(s)/2:]
		for i := 0; i < len(second); i++ {
			if firstMap[second[i]] {
				fmt.Println(string(second[i]))
				score += getPriority(string(second[i]))
				firstMap[second[i]] = false
			}
		}

	}
	fmt.Println(score)
}

func getCommonLetterPoints(arr []string) int {
	m := map[rune]int{}
	for _, x := range abc {
		m[x] = 0
	}
	for _, s := range arr {
		sMap := map[rune]bool{}
		for _, b := range s {
			sMap[b] = true
		}
		for k := range sMap {
			if _, ok := m[k]; ok {
				m[k] = m[k] + 1
			} else {
				m[k] = 1
			}
		}
	}
	score := 0
	for k, v := range m {
		if v == 3 {
			score += getPriority(string(k))
		}
	}
	return score
}

func part2() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := 1
	var arr []string
	score := 0
	for scanner.Scan() {
		s := scanner.Text()
		arr = append(arr, s)
		if line%3 == 0 {
			score += getCommonLetterPoints(arr)
			arr = []string{}
		}
		line++
	}
	fmt.Println(score)
}
