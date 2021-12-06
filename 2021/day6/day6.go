package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_initial_counts(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var counts []int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	for _, c := range strings.Split(line, ",") {
		count, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		counts = append(counts, count)
	}

	return counts
}

// broke
func update_counts(counts []int) []int {
	var new_counts []int
	new_fish := make([]int, 0)
	for _, count := range counts {
		if count == 0 {
			new_counts = append(new_counts, 6)
			new_fish = append(new_fish, 8)
		} else {
			new_counts = append(new_counts, count - 1)
		}
	}
	new_counts = append(new_counts, new_fish...)
	return new_counts
}

// woke
func map_initial_counts(counts []int, max_days int) map[int]int {
	m := make(map[int]int)
	for i := 0; i <= max_days; i++ {
		m[i] = 0
	}
	for _, count := range counts {
		m[count]++
	}
	return m
}

func update_counts_map(counts map[int]int) map[int]int {
	new_counts := make(map[int]int)
	for count, count_count := range counts {
		if count == 0 {
			new_counts[6] += count_count
			new_counts[8] += count_count
		} else {
			new_counts[count - 1] += count_count
		}
	}
	return new_counts
}

func total_count(counts map[int]int) int {
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

func main() {
	const filename = "input.txt"
	counts := read_initial_counts(filename)

	const max_days = 8
	counts_map := map_initial_counts(counts, max_days)
	fmt.Printf("Initial counts: %v\n", counts_map)

	const days = 256
	for i := 0; i < days; i++ {
		counts_map = update_counts_map(counts_map)
		fmt.Printf("Day %d: %d\n", i + 1, total_count(counts_map))
	}
	fmt.Printf("Count: %d\n", total_count(counts_map))
}