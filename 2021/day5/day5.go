package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const filename = "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	grid := NewGrid()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []*Line
	for scanner.Scan() {
		line := ReadLine(scanner.Text())
		lines = append(lines, line)
	}

	for _, line := range lines {
		grid.FillLine(line, false)
	}
	grid.OutputGrid()

	countAbove2 := grid.CountVentsAboveThreshold(2)
	fmt.Printf("Count of vents above 2: %d\n", countAbove2)
}