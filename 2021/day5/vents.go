package main

import (
	"fmt"
	"math"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func NewLine(x1, y1, x2, y2 int) *Line {
	return &Line{x1: x1, y1: y1, x2: x2, y2: y2}
}

func ReadLine(lineStr string) *Line {
	// (x1, y1) -> (x2, y2)
	lineParts := strings.Fields(lineStr)
	var x1, y1, x2, y2 int
	fmt.Sscanf(lineParts[0], "%d,%d", &x1, &y1)
	fmt.Sscanf(lineParts[2], "%d,%d", &x2, &y2)
	return NewLine(x1, y1, x2, y2)
}

func (line *Line) OutputLine() {
	fmt.Printf("(%d, %d) -> (%d, %d)\n", line.x1, line.y1, line.x2, line.y2)
}

type Grid struct {
	points map[int]map[int]int // map location to number of vents
}

func NewGrid() *Grid {
	return &Grid{points: make(map[int]map[int]int)}
}

func (grid *Grid) AddVent(x int, y int) {
	if _, ok := grid.points[x]; !ok {
		grid.points[x] = make(map[int]int)
	}
	if _, ok := grid.points[x][y]; !ok {
		grid.points[x][y] = 0
	}
	grid.points[x][y]++
}

func (grid *Grid) FillLine(line *Line, onlyHorizontalOrVertical bool) {
	x := line.x1
	y := line.y1

	xStep := line.x2 - line.x1
	yStep := line.y2 - line.y1

	if xStep != 0 && yStep != 0 && onlyHorizontalOrVertical {
		return
	}
	if xStep != 0 {
		xStep /= int(math.Abs(float64(xStep)))
	}
	if yStep != 0 {
		yStep /= int(math.Abs(float64(yStep)))
	}
	for x != line.x2 || y != line.y2 {
		grid.AddVent(x, y)
		x += xStep
		y += yStep
	}
	grid.AddVent(line.x2, line.y2)
}

func (grid *Grid) CountVentsAboveThreshold(threshold int) int {
	count := 0
	for x := range grid.points {
		for y := range grid.points[x] {
			if grid.points[x][y] >= threshold {
				count++
			}
		}
	}
	return count
}

func (grid *Grid) OutputGrid() {
	for x := range grid.points {
		for y := range grid.points[x] {
			fmt.Printf("(%d, %d): %d\n", x, y, grid.points[x][y])
		}
	}
}