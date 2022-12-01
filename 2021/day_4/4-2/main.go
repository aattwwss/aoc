package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(strings.Fields("1 2 3   4 5"))
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var cmdList []string
	var boards [][][]string
	var board [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if len(cmdList) == 0 {
			cmdList = strings.Split(line, ",")
			continue
		}
		if line == "" {
			if len(board) == 0 {
				continue
			}
			boards = append(boards, board)
			board = [][]string{}
			continue
		}
		board = append(board, strings.Fields(line))

	}

	// fmt.Println(arr[11])
	// fmt.Println(cmdList)
	// testBoard := [][]string{{"1", "2", "3", "4", "5"}, {"6", "7", "8", "9", "10"}, {"11", "12", "13", "14", "15"}, {"16", "17", "18", "19", "20"}, {"21", "22", "23", "24", "25"}}
	// setPlayingNumberToBoard(&testBoard, "25")
	// setPlayingNumberToBoard(&testBoard, "2")
	// fmt.Println(testBoard)
	// fmt.Println(isBoardWon(testBoard))
	boardWon := []int{}
exit:
	for _, num := range cmdList {
		for i, board := range boards {
			board := board
			setPlayingNumberToBoard(&board, num)
			if isBoardWon(board) && !slices.Contains(boardWon, i) {
				boardWon = append(boardWon, i)
			}
			if len(boardWon) == len(boards) {

				numValue, _ := strconv.Atoi(num)
				sumInBoard := calculateSumOfWinningBoard(board)
				fmt.Println(sumInBoard * numValue)
				fmt.Println(boardWon)
				break exit
			}
		}

	}
}
func calculateSumOfWinningBoard(board [][]string) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(board[i][j])
			if board[i][j] != "-1" {
				value, _ := strconv.Atoi(board[i][j])
				sum += value
			}
		}
	}
	return sum
}
func setPlayingNumberToBoard(board *[][]string, num string) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (*board)[i][j] == num {
				(*board)[i][j] = "-1"
			}
		}
	}

}

func isBoardWon(board [][]string) bool {

	yPoint := 0
	xPoint := 0
	for i := 0; i < 5; i++ {
		// fmt.Println(board)
		for j := 0; j < 5; j++ {
			//	fmt.Println(board[i][j])
			if board[i][j] == "-1" {
				xPoint++
			}
			if board[j][i] == "-1" {
				yPoint++
			}
		}
		if yPoint == 5 || xPoint == 5 {
			return true
		}
		yPoint = 0
		xPoint = 0

	}
	return false
}
