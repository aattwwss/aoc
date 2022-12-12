package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	CyclesNeeded int
	Value        int
}

func NewInstruction(s string) Instruction {
	arr := strings.Split(s, " ")
	if arr[0] == "addx" {
		value, _ := strconv.Atoi(arr[1])
		return Instruction{2, value}
	}
	return Instruction{1, 0}
}

type CPU struct {
	Tick     int
	Register int
	Score    int
	Values   []int
}

//func (c *CPU) execute(instruction Instruction) {
//	for i := 0; i < instruction.CyclesNeeded; i++ {
//		c.Tick++
//		if c.Tick%20 == 0 {
//			c.Score += c.Tick * c.Register
//		}
//	}
//	c.Register += instruction.Value
//}

func (c *CPU) execute(instruction Instruction) {
	for i := 0; i < instruction.CyclesNeeded; i++ {
		c.Values = append(c.Values, c.Register)
		c.Tick++
	}
	c.Register += instruction.Value
}

func main() {

	file, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	cpu := CPU{0, 1, 0, []int{}}
	for fileScanner.Scan() {
		s := fileScanner.Text()
		instruction := NewInstruction(s)
		cpu.execute(instruction)
	}
	var score int
	for i := 19; i < len(cpu.Values); i += 40 {
		score += cpu.Values[i] * (i + 1)
	}

	screen := [6][40]string{}

	for i := range cpu.Values {
		row := i / 40
		column := i % 40
		spritePosition := cpu.Values[i]
		if column == spritePosition || column+1 == spritePosition || column-1 == spritePosition {
			screen[row][column] = "#"
		} else {
			screen[row][column] = "."
		}
	}

	for i := range screen {
		fmt.Println(strings.Join(screen[i][:], ""))
	}

	fmt.Println(score)
}
