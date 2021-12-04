package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type BingoSquare struct {
	row int
	col int
	marked bool
	number int
}

type BingoBoard struct {
	squaresByNumber map[int]*BingoSquare // map square value to square location and marked status
	squaresByLocation map[int]*BingoSquare // map square location to square value and marked status
	hasWon bool
}

func ReadBingoBoard(scanner *bufio.Scanner) *BingoBoard {
	board := &BingoBoard{
		squaresByNumber: make(map[int]*BingoSquare),
		squaresByLocation: make(map[int]*BingoSquare),
	}
	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			board.squaresByNumber[num] = &BingoSquare{
				row: i,
				col: j,
				marked: false,
				number: num,
			}
			board.squaresByLocation[i*5+j] = board.squaresByNumber[num]
		}
		
	}
	return board
}

func (board *BingoBoard) MarkSquare(num int) {
	if square, ok := board.squaresByNumber[num]; ok {
		square.marked = true
	}
}

func (board *BingoBoard) GetMatchingSquares(isMarked bool) []*BingoSquare {
	var matchingSquares []*BingoSquare
	for _, square := range board.squaresByNumber {
		if square.marked == isMarked {
			matchingSquares = append(matchingSquares, square)
		}
	}
	return matchingSquares
}

func (board *BingoBoard) GetRow(row int) []*BingoSquare {
	var rowSquares []*BingoSquare
	for i := 0; i < 5; i++ {
		rowSquares = append(rowSquares, board.squaresByLocation[row*5+i])
	}
	return rowSquares
}

func (board *BingoBoard) GetCol(col int) []*BingoSquare {
	var colSquares []*BingoSquare
	for i := 0; i < 5; i++ {
		colSquares = append(colSquares, board.squaresByLocation[i*5+col])
	}
	return colSquares
}

func (board *BingoBoard) GetWinningRowOrCol() ([]*BingoSquare, []*BingoSquare) {
	markedRowCounts := make([]int, 5)
	markedColCounts := make([]int, 5)

	for _, square := range board.GetMatchingSquares(true) {
		markedRowCounts[square.row]++
		if markedRowCounts[square.row] == 5 {
			return board.GetRow(square.row), nil
		}

		markedColCounts[square.col]++
		if markedColCounts[square.col] == 5 {
			return nil, board.GetCol(square.col)
		}
	}

	return nil, nil
}

func (board *BingoBoard) OutputBoard() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			square := board.squaresByLocation[i*5+j]
			if square.marked {
				fmt.Printf("(%d) ", square.number)
			} else {
				fmt.Printf("%d ", square.number)
			}
		}
		fmt.Println()
	}
}

func (board *BingoBoard) GetScore(winningNumber int) int {
	sum := 0
	for _, square := range board.GetMatchingSquares(false) {
		sum += square.number
	}
	score := sum * winningNumber
	fmt.Printf("Winning number: %d\n", winningNumber)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Score: %d\n", score)
	return score
}