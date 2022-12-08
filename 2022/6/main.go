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
	part2()
}

func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func hasRepeated(ss string) bool {
	m := map[rune]bool{}
	for _, s := range ss {
		if _, ok := m[s]; ok {
			return true
		}
		m[s] = true
	}
	return false
}

func part1() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	l, r := 0, 4
	for scanner.Scan() {
		s := scanner.Text()
		for r < len(s) {
			ss := s[l:r]
			if !hasRepeated(ss) {
				fmt.Println(ss)
				fmt.Println(r)
				break
			}
			l++
			r++
		}
		// fmt.Println(s)
	}

}

func part2() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	l, r := 0, 14
	for scanner.Scan() {
		s := scanner.Text()
		for r < len(s) {
			ss := s[l:r]
			if !hasRepeated(ss) {
				fmt.Println(ss)
				fmt.Println(r)
				break
			}
			l++
			r++
		}
		// fmt.Println(s)
	}
}
