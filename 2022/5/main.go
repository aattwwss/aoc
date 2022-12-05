package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
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

type column []string

func (c *column) push(s string) {
	// *c = append(*c, s)
	*c = append(column{s}, *c...)
}

func (c *column) addToLast(s string) {
	*c = append(*c, s)
}

func (c *column) pop() string {
	s := (*c)[0]
	*c = (*c)[1:]
	return s
}

type columns struct {
	stacks map[string]*column
	size   int
}

func newColumns(n int) *columns {
	m := map[string]*column{}
	for i := 0; i < n; i++ {
		c := column{}
		m[fmt.Sprintf("%v", 1+i)] = &c
	}
	c := columns{m, n}
	return &c

}

func (c *columns) doMove(move move) {
	for i := 0; i < move.num; i++ {
		s := c.stacks[move.from].pop()
		c.stacks[move.to].push(s)
	}
}

func (c *columns) getTopLayer() string {
	s := ""
	for i := 0; i < c.size; i++ {
		index := fmt.Sprintf("%v", i+1)
		c := c.stacks[index]
		t := c.pop()
		if t != "" {
			s += t
		}
	}
	return s
}

type move struct {
	from string
	to   string
	num  int
}

func part1() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cratesLines := []string{}
	movesLines := []string{}
	toggle := false
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			toggle = true
		}
		if !toggle && s != "" {
			cratesLines = append(cratesLines, s)
		} else if toggle && s != "" {
			movesLines = append(movesLines, s)
		}
	}

	lables := strings.Fields(cratesLines[len(cratesLines)-1])

	n, _ := strconv.Atoi(lables[len(lables)-1])
	cratesLines = cratesLines[:len(cratesLines)-1]
	columns := newColumns(n)
	for _, line := range cratesLines {
		fmt.Println(line)
		for i := range line {
			if string(line[i]) == "[" {
				fmt.Printf("i: %v, label: %v, value: %s\n", i, i/4+1, string(line[i+1]))
				columns.stacks[fmt.Sprintf("%v", i/4+1)].addToLast(string(line[i+1]))
			}
		}
	}

	mar, _ := json.Marshal(columns.stacks)
	fmt.Println(string(mar))
	moves := []move{}
	for _, line := range movesLines {
		arr := strings.Split(line, " ")
		num, _ := strconv.Atoi(arr[1])
		move := move{
			from: arr[3],
			to:   arr[5],
			num:  num,
		}
		moves = append(moves, move)
	}

	for _, move := range moves {
		columns.doMove(move)
	}
	s := columns.getTopLayer()
	fmt.Println(s)
}
func part2() {
	// fileName := "input.txt"
	// file, _ := os.Open(fileName)
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	s := scanner.Text()
	// }
}
