package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var LCM = 1

type Test struct {
	DivisibleValue int
	IfTrue         int
	IfFalse        int
}

func (t Test) doTest(n int) int {

	if n >= t.DivisibleValue && n%t.DivisibleValue == 0 {
		return t.IfTrue
	}
	return t.IfFalse
}

func NewTest(s1, s2, s3 string) Test {
	divString := strings.Split(s1, " by ")[1]
	div, _ := strconv.Atoi(divString)

	ifTrueString := strings.Split(s2, " monkey ")[1]
	ifTrue, _ := strconv.Atoi(ifTrueString)

	ifFalseString := strings.Split(s3, " monkey ")[1]
	ifFalse, _ := strconv.Atoi(ifFalseString)

	return Test{
		DivisibleValue: div,
		IfTrue:         ifTrue,
		IfFalse:        ifFalse,
	}
}

type Operation struct {
	Operator string
	Value    string
}

func NewOperation(s string) Operation {
	arr := strings.Split(strings.Split(s, " = ")[1], " ")
	return Operation{Operator: arr[1], Value: arr[2]}
}

func (o Operation) doOperation(old int) int {
	old = old % LCM // part 2
	var value int
	if o.Value == "old" {
		value = old
	} else {
		v, _ := strconv.Atoi(o.Value)
		value = v
	}

	if o.Operator == "*" {
		return old * value
	}

	if o.Operator == "+" {
		return old + value
	}
	return 0
}

type Monkey struct {
	InspectedCount int
	Items          []int
	Operation      Operation
	Test           Test
}

func (m *Monkey) throw() (int, int) {
	m.InspectedCount++
	item := m.Items[0]
	m.Items = m.Items[1:]

	newItem := m.Operation.doOperation(item) / 1 // Remove division by 3 in part 2
	monkeyToThrowTo := m.Test.doTest(newItem)

	return monkeyToThrowTo, newItem

}

func NewMonkey(s string) Monkey {
	var items []int
	lines := strings.Split(s, "\n")
	for _, level := range strings.Split(strings.Split(lines[1], ": ")[1], ", ") {
		levelInt, _ := strconv.Atoi(level)
		items = append(items, levelInt)
	}
	operation := NewOperation(lines[2])
	test := NewTest(lines[3], lines[4], lines[5])

	return Monkey{Items: items, Operation: operation, Test: test}
}

func main() {
	var monkeys []Monkey

	fileName := "input.txt"
	content, _ := os.ReadFile(fileName)
	lines := strings.Split(string(content), "\n\n")
	for _, line := range lines {
		monkeys = append(monkeys, NewMonkey(line))
	}

	for _, monkey := range monkeys {
		LCM *= monkey.Test.DivisibleValue
	}

	rounds := 10000

	for n := 0; n < rounds; n++ {
		for i := range monkeys {
			for range monkeys[i].Items {
				monkeyToThrowTo, newItem := monkeys[i].throw()
				monkeys[monkeyToThrowTo].Items = append(monkeys[monkeyToThrowTo].Items, newItem)
			}
		}
	}
	var bucket []int
	for _, monkey := range monkeys {
		bucket = append(bucket, monkey.InspectedCount)
	}
	sort.Ints(bucket)
	fmt.Println(bucket[len(bucket)-1])
	fmt.Println(bucket[len(bucket)-2])
	fmt.Println(bucket[len(bucket)-1] * bucket[len(bucket)-2])
}
