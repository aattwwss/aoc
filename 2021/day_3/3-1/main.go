package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var arr []string
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	bArr := [12][1000]string{}
	for i, s := range arr {
		chars := []rune(s)
		for j, c := range chars {
			bArr[j][i] = string(c)
		}
	}
	var gamam, eps string

	for _, s := range bArr {
		countMap := CountOccurence(s[:])
		if countMap["0"] > countMap["1"] {
			gamam += "0"
			eps += "1"
		} else {
			gamam += "1"
			eps += "0"
		}
	}
	fmt.Println(len(bArr))
	fmt.Println(len(bArr[0]))
	gammaInt, _ := strconv.ParseInt(gamam, 2, 64)
	epsInt, _ := strconv.ParseInt(eps, 2, 64)
	fmt.Println(gammaInt * epsInt)
}

func CountOccurence(apps []string) map[string]int {
	dict := make(map[string]int)
	for _, v := range apps {
		dict[v]++
	}
	return dict
}
