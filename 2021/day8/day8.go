package main

import (
	"fmt"
	"strings"

	"jburke.com/advent_util"
)

func part1(lines []string) {
	counts := make(map[int]int)

	// unique segments:
	// 1: 2 segments
	// 4: 4 segments
	// 7: 3 segments
	// 8: 7 segments
	for _, line := range lines {
		output := strings.Split(line, "|")[1]
		for _, digit := range strings.Fields(output) {
			digit_length := len(digit)
			if digit_length == 2 || digit_length == 3 || digit_length == 7 || digit_length == 4 {
				if _, ok := counts[digit_length]; !ok {
					counts[digit_length] = 1
				} else {
					counts[digit_length]++
				}
			}
		}
	}

	fmt.Printf("Final counts; %v\n", counts)
	sumCounts := advent_util.SumCounts(counts)
	fmt.Printf("Sum of counts: %d\n", sumCounts)
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		line_split := strings.Split(line, "|")
		signal_patterns := line_split[0]
		digit_mapping, unknowns := KnownDigitsFromInput(signal_patterns)
		fmt.Printf("digit_mapping: %v\n", digit_mapping)
		fmt.Printf("unknowns: %v\n", unknowns)
		mid_mapping, remaining := UnknownDigitsFromInput(digit_mapping, unknowns)
		fmt.Printf("final_mapping: %v\n", mid_mapping)
		fmt.Printf("remaining: %v\n", remaining)
		final_mapping := RemainingUnknowns(mid_mapping, remaining)
		fmt.Printf("final_mapping: %v\n", final_mapping)

		sorted_mapping := SortKeys(final_mapping)
		fmt.Printf("sorted_mapping: %v\n", sorted_mapping)

		output := line_split[1]
		fmt.Printf("output: %v\n", strings.Fields(output))
		num := OutputToNumber(sorted_mapping, line_split[1])
		fmt.Printf("num: %d\n", num)
		sum += num
	}
	return sum
}

func main() {
	const filename = "input.txt"

	lines, err := advent_util.ReadLines(filename)

	if err != nil {
		panic(err)
	}

	part1(lines)
	
	// part 2
	sum := part2(lines)
	fmt.Printf("sum: %d\n", sum)
}
