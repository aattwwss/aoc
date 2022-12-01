package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Aquarium struct {
	FishTimer [9]int
}

func NewAquarium() Aquarium {
	return Aquarium{FishTimer: [9]int{}}
}

func (a *Aquarium) NextDay() {
	newArr := [len(a.FishTimer)]int{}
	for i := range a.FishTimer {
		newArr[i] = a.FishTimer[(i+1)%len(a.FishTimer)]
	}
	newArr[6] += newArr[8]
	a.FishTimer = newArr
}

func (a Aquarium) sumAllFishes() int {
	count := 0
	for _, fishes := range a.FishTimer {
		count += fishes
	}
	return count
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	aquarium := NewAquarium()
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fishesTimerString := strings.Split(scanner.Text(), ",")
		for _, fishString := range fishesTimerString {
			fishTimer, _ := strconv.Atoi(fishString)
			aquarium.FishTimer[fishTimer]++
		}

	}

	for i := 0; i < 256; i++ {
		aquarium.NextDay()
	}
	fmt.Println(aquarium.sumAllFishes())
}
