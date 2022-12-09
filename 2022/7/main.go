package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	// part2()
}

type Dir struct {
	Name  string
	Prev  *Dir
	Dirs  []*Dir
	Files []int
}

func (d *Dir) GetDirSize() int {
	size := 0
	for _, file := range d.Files {
		size += file
	}
	for i := range d.Dirs {
		dirSize := d.Dirs[i].GetDirSize()
		size += dirSize
	}
	return size
}

func (d *Dir) ToJson() string {
	b, err := json.Marshal(d)
	if err != nil {
		fmt.Print(d.Name)
		fmt.Println(err)
	}
	return string(b)
}

func (d *Dir) GetDirSizeMap() map[string]int {
	m := map[string]int{d.Name: d.GetDirSize()}
	q := []*Dir{d}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, dir := range curr.Dirs {
			if _, ok := m[dir.Name]; !ok {
				m[dir.Name] = dir.GetDirSize()
				q = append(q, dir)
			}
		}
	}
	return m
}

func part1() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var root Dir
	curr := &root
	isRoot := true
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		fmt.Printf("curr: %v\n", curr.ToJson())
		if isRoot {
			root = Dir{Name: "/", Files: []int{}, Dirs: []*Dir{}}
			isRoot = false
			continue
		}
		arr := strings.Split(s, " ")
		if s == "$ cd .." {
			curr = curr.Prev
			continue
		}
		if strings.HasPrefix(s, "$ cd") {
			fmt.Println(curr.Dirs)
			for i := range curr.Dirs {
				if strings.HasSuffix(curr.Dirs[i].Name, arr[2]+"/") {
					temp := curr
					curr = curr.Dirs[i]
					curr.Prev = temp
					break
				}
			}
			continue
		}
		if s == "$ ls" {
			continue
		}
		if strings.HasPrefix(s, "dir") {
			next := Dir{Name: curr.Name + arr[1] + "/", Files: []int{}, Dirs: []*Dir{}}
			curr.Dirs = append(curr.Dirs, &next)
			continue
		}
		// is file
		size, _ := strconv.Atoi(arr[0])
		curr.Files = append(curr.Files, size)
	}
	m := root.GetDirSizeMap()
	total := 0
	toDelete := []int{}
	requiredSize := 30000000 - (70000000 - m["/"])
	for k, v := range m {
		if v <= 100000 {
			fmt.Printf("key: %s value: %v\n", k, v)
			total += v
		}
		if v >= requiredSize {
			toDelete = append(toDelete, v)
		}
	}
	fmt.Print("Total: ")
	fmt.Println(total)
	// fmt.Println(m)
	// part 2 here
	sort.Ints(toDelete)
	fmt.Println(toDelete[0])

}

func part2() {
	fileName := "input.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}
}
