package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type Knot struct {
	Curr    Coordinate
	Visited []Coordinate
}

func (k *Knot) countUniqueVisited() int {
	m := map[Coordinate]bool{}
	for i := range k.Visited {
		m[k.Visited[i]] = true
	}
	return len(m)
}

func (k *Knot) move(s string) {
	if s == "L" {
		k.Curr.X--
	} else if s == "R" {
		k.Curr.X++
	} else if s == "D" {
		k.Curr.Y--
	} else if s == "U" {
		k.Curr.Y++
	}
}

func (k *Knot) follow(front Knot) {
	// for each step the head takes, we calculate the tail next step
	diff := Coordinate{X: front.Curr.X - k.Curr.X, Y: front.Curr.Y - k.Curr.Y}
	if abs(diff.X) <= 1 && abs(diff.Y) <= 1 {
		return
	}
	if diff.X == 0 {
		k.Curr.Y += sign(diff.Y)
		return
	}
	if diff.Y == 0 {
		k.Curr.X += sign(diff.X)
		return
	}
	if diff.Y > 0 {
		k.Curr.Y++
	} else {
		k.Curr.Y--
	}
	if diff.X > 0 {
		k.Curr.X++
	} else {
		k.Curr.X--
	}
}

type Knots []Knot

func main() {
	n := 10
	knots := Knots{}
	for i := 0; i < n; i++ {
		knots = append(knots, Knot{Curr: Coordinate{0, 0}, Visited: []Coordinate{{0, 0}}})
	}

	file, _ := os.Open("C:\\Users\\Alvin\\Personal Projects\\aoc\\2022\\9\\input.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		s := fileScanner.Text()
		fmt.Println(s)
		arr := strings.Split(s, " ")
		direction := arr[0]
		steps, _ := strconv.Atoi(arr[1])
		for n := 0; n < steps; n++ {
			knots[0].move(direction)
			for i := 1; i < len(knots); i++ {
				knots[i].follow(knots[i-1])
				knots[i].Visited = append(knots[i].Visited, knots[i].Curr)
			}
		}
	}
	fmt.Println(knots[1].countUniqueVisited())
	fmt.Println(knots[9].countUniqueVisited())

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}
