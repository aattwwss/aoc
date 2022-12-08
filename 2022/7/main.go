package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AFile struct {
	Name string
	Size int
}

func newAFile(name string, size int) AFile {
	return AFile{name, size}
}

type Dir struct {
	Name  string
	Files *[]AFile
	Dirs  *[]Dir
	Prev  *Dir
}

func newDir(name string, prev *Dir) Dir {
	return Dir{name, &[]AFile{}, &[]Dir{}, prev}
}

func (d Dir) getTotalSize(threshold int) int {
	var size int
	for _, file := range *d.Files {
		if size >= threshold {
			size += file.Size
		}
	}
	for _, dir := range *d.Dirs {
		size += dir.getTotalSize(threshold)
	}
	return size
}

func handleCD(prev *Dir, curr *Dir, s string) {
	arr := strings.Split(s, " ")
	if arr[2] == ".." {
		temp := curr.Prev
		curr.Prev = prev
		prev = temp
		return
	} else {
		var temp *Dir
		prev = curr
		for _, dir := range *curr.Dirs {
			if dir.Name == arr[2] {
				temp = &dir
			}
		}
		curr = temp
		*curr.Dirs = append(*curr.Dirs, newDir(arr[2], curr))
	}

}

func handleFileOrDir(curr *Dir, first string, name string) {
	if first == "dir" {
		fmt.Println("handling dir")
		*curr.Dirs = append(*curr.Dirs, newDir(name, curr))
	} else {
		fmt.Println("handling files")
		size, _ := strconv.Atoi(first)
		*curr.Files = append(*curr.Files, newAFile(name, size))
	}
}

func main() {
	part1()
	// part2()
}

func part1() {
	fileName := "test.txt"
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	isRoot := true
	var curr Dir
	var prev Dir
	root := &curr
	doProcessFiles := false
	for scanner.Scan() {
		s := scanner.Text()
		if isRoot {
			fmt.Println("init root as current <--" + s)
			curr = newDir(s, &prev)
			isRoot = false
			continue
		}

		arr := strings.Split(s, " ")
		if doProcessFiles && arr[0] == "$" {
			fmt.Println("End of processing file <--" + s)
			doProcessFiles = false
		} else if doProcessFiles {
			fmt.Println("handleFile <--" + s)
			handleFileOrDir(&curr, arr[0], arr[1])
		}
		if arr[0] == "$" {
			if arr[1] == "cd" {
				fmt.Println("handle cd <--" + s)
				handleCD(&prev, &curr, s)
			} else if arr[1] == "ls" {
				fmt.Println("handle ls <--" + s)
				doProcessFiles = true
			}
		}
		// fmt.Println(s)
	}
	mar, _ := json.Marshal(curr)
	fmt.Println(mar)
	fmt.Println(root.getTotalSize(100000))

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
