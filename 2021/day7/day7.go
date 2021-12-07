package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"

	"jburke.com/advent_util"
)

// part 1
func sumDistances(positions []int, value int) int {
	sum := 0
	for _, v := range positions {
		sum += int(math.Abs(float64(v - value)))
	}
	return sum
}

func meanPosition(positions []int) int {
	sum := 0
	for _, v := range positions {
		sum += v
	}
	return int(math.Floor(float64(sum) / float64(len(positions))))
}

// part 2
func sumToN(n int) int {
	return n * (n + 1) / 2
}

func sumIncreasingDistances(positions []int, value int) int {
	sum := 0
	for _, v := range positions {
		sum += sumToN(int(math.Abs(float64(v - value))))
	}
	return sum
}

func main() {
	const filename = "input.txt"
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	positions, err := advent_util.StrToInts(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of positions: %d\n", len(positions))

	sort.Ints(positions)

	fmt.Printf("Sorted positions: %v\n", positions)

	median := positions[len(positions)/2]

	fmt.Printf("Median: %d\n", median)

	fmt.Printf("Sum of distances: %d\n", sumDistances(positions, median))

	mean := meanPosition(positions)
	fmt.Printf("Mean: %d\n", mean)

	fmt.Printf("Sum of distances: %d\n", sumIncreasingDistances(positions, mean))
}