package advent_util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts(filename string) ([]int, error) {
	file, err := os.Open(filename)

	var result []int
	if err != nil {
		return result, err	
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, i)
	}
	return result, scanner.Err()
}