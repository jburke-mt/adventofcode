package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func get_called_numbers(scanner *bufio.Scanner) []int {
	var result []int
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	calledNumbers := scanner.Text()
	for _, val := range strings.Split(calledNumbers, ",") {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}
	return result
}

func get_winning_board(calledNumbers []int, boards []*BingoBoard) (*BingoBoard, int) {
	for _, num := range calledNumbers {
		for _, board := range boards {
			board.MarkSquare(num)
			winning_row, winning_col := board.GetWinningRowOrCol()
			if winning_row != nil || winning_col != nil {
				return board, num
			}
		}
	}
	return nil, -1
}

func main() {
	const filename = "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	calledNumbers := get_called_numbers(scanner)
	fmt.Printf("Called numbers: %v\n", calledNumbers)

	var boards []*BingoBoard
	for scanner.Scan() {
		boards = append(boards, ReadBingoBoard(scanner))
	}
	scanner_err := scanner.Err()
	if scanner_err != nil && scanner_err != io.EOF {
		panic(scanner_err)
	}

	winning_board, winning_number := get_winning_board(calledNumbers, boards)
	sum := 0
	for _, unmarked := range winning_board.GetMatchingSquares(false) {
		sum += unmarked.number
	}
	fmt.Println("Winning board:")
	winning_board.OutputBoard()
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Winning number: %d\n", winning_number)
	fmt.Printf("Product: %d\n", sum*winning_number)
}