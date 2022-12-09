package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
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

func checkUp(grid [][]int, r int, c int) bool {
	for i := r - 1; i >= 0; i-- {
		if grid[i][c] >= grid[r][c] {
			return false
		}
	}
	return true
}

func checkDown(grid [][]int, r int, c int) bool {
	for i := r + 1; i < len(grid); i++ {
		if grid[i][c] >= grid[r][c] {
			return false
		}
	}
	return true
}

func checkRight(grid [][]int, r int, c int) bool {
	for i := c + 1; i < len(grid[r]); i++ {
		if grid[r][i] >= grid[r][c] {
			return false
		}
	}
	return true
}

func checkLeft(grid [][]int, r int, c int) bool {
	for i := c - 1; i >= 0; i-- {
		if grid[r][i] >= grid[r][c] {
			return false
		}
	}
	return true
}

func isVisible(grid [][]int, r int, c int) bool {
	results := []bool{}
	results = append(results, checkLeft(grid, r, c), checkRight(grid, r, c), checkUp(grid, r, c), checkDown(grid, r, c))
	return slices.Contains(results, true)
}

func part1() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	for scanner.Scan() {
		s := scanner.Text()
		row := []int{}
		for i := 0; i < len(s); i++ {
			v, _ := strconv.Atoi(string(s[i]))
			row = append(row, v)
		}
		grid = append(grid, row)
	}
	count := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[r])-1; c++ {
			if isVisible(grid, r, c) {
				count++
			}
		}
	}
	fmt.Println(len(grid[0]))
	count += (len(grid)*4 - 4)
	fmt.Println(count)
}

func scoreDown(grid [][]int, r int, c int) int {
	score := 0
	for i := r + 1; i < len(grid); i++ {
		score++
		if grid[i][c] >= grid[r][c] {
			break
		}
	}
	return score
}
func scoreUp(grid [][]int, r int, c int) int {
	score := 0
	for i := r - 1; i >= 0; i-- {
		score++
		if grid[i][c] >= grid[r][c] {
			break
		}
	}
	return score
}
func scoreRight(grid [][]int, r int, c int) int {
	score := 0
	for i := c + 1; i < len(grid[r]); i++ {
		score++
		if grid[r][i] >= grid[r][c] {
			break
		}
	}
	return score
}
func scoreLeft(grid [][]int, r int, c int) int {
	score := 0
	for i := c - 1; i >= 0; i-- {
		score++
		if grid[r][i] >= grid[r][c] {
			break
		}
	}
	return score
}
func calculate(grid [][]int, r int, c int) int {
	return scoreLeft(grid, r, c) * scoreRight(grid, r, c) * scoreUp(grid, r, c) * scoreDown(grid, r, c)
}

func part2() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	for scanner.Scan() {
		s := scanner.Text()
		row := []int{}
		for i := 0; i < len(s); i++ {
			v, _ := strconv.Atoi(string(s[i]))
			row = append(row, v)
		}
		grid = append(grid, row)
	}
	scores := []int{}
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[r])-1; c++ {
			scores = append(scores, calculate(grid, r, c))
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)-1])
}
