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

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	coordsCount := map[Coordinate]int{}
	for scanner.Scan() {

		coords := parseTextIntoCoords(scanner.Text())
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

func parseTextIntoCoords(s string) []Coordinate {
	var res []Coordinate
	if s == "" {
		return res
	}
	startEndStringArr := strings.Split(s, " -> ")
	start := startEndStringArr[0]
	end := startEndStringArr[1]
	startCoords := parseCoordsTextIntoCoords(start)
	endCoords := parseCoordsTextIntoCoords(end)

	if startCoords.X != endCoords.X && endCoords.Y != startCoords.Y {
		return res
	}
	if startCoords == endCoords {
		return []Coordinate{startCoords}
	}
	diffCoords := Coordinate{
		X: startCoords.X - endCoords.X,
		Y: startCoords.Y - endCoords.Y,
	}
	if diffCoords.X == 0 {
		if startCoords.Y > endCoords.Y {
			for i := endCoords.Y; i <= startCoords.Y; i++ {
				res = append(res, Coordinate{startCoords.X, i})
			}
		} else {
			for i := startCoords.Y; i <= endCoords.Y; i++ {
				res = append(res, Coordinate{startCoords.X, i})
			}
		}
	} else {
		if startCoords.X > endCoords.X {
			for i := endCoords.X; i <= startCoords.X; i++ {
				res = append(res, Coordinate{i, startCoords.Y})
			}
		} else {
			for i := startCoords.X; i <= endCoords.X; i++ {
				res = append(res, Coordinate{i, startCoords.Y})
			}
		}
	}
	return res
}

func parseCoordsTextIntoCoords(s string) Coordinate {
	XY := strings.Split(s, ",")
	x, _ := strconv.Atoi(XY[0])
	y, _ := strconv.Atoi(XY[1])
	return Coordinate{x, y}
}
