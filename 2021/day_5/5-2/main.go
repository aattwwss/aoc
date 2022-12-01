package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func NewCoordFromString(s string) Coordinate {
	xyString := strings.Split(s, ",")
	x, _ := strconv.Atoi(xyString[0])
	y, _ := strconv.Atoi(xyString[1])
	return Coordinate{X: x, Y: y}
}

func (c Coordinate) Difference(c2 Coordinate) (int, int) {
	return c.X - c2.X, c.Y - c2.Y
}

type Line struct {
	Start Coordinate
	End   Coordinate
}

func NewLineFromString(s string) Line {
	startEndString := strings.Split(s, " -> ")
	start := NewCoordFromString(startEndString[0])
	end := NewCoordFromString(startEndString[1])
	return Line{Start: start, End: end}
}

func (l Line) GetCoordinatesWithin() []Coordinate {
	var res []Coordinate
	xInc := Sgn(l.End.X - l.Start.X)
	yInc := Sgn(l.End.Y - l.Start.Y)
	for i, j := l.Start.X, l.Start.Y; i != l.End.X || j != l.End.Y; i, j = i+xInc, j+yInc {
		res = append(res, Coordinate{X: i, Y: j})
	}
	res = append(res, l.End)
	return res
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	coordsCount := map[Coordinate]int{}
	for scanner.Scan() {
		line := NewLineFromString(scanner.Text())
		coords := line.GetCoordinatesWithin()
		for _, coord := range coords {
			_, ok := coordsCount[coord]
			if ok {
				coordsCount[coord]++
			} else {
				coordsCount[coord] = 1
			}
		}
	}
	count := 0
	for _, v := range coordsCount {
		if v > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func Sgn(a int) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
