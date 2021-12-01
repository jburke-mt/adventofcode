package main

import (
	"fmt"
	"os"

	"jburke.com/advent_util"
)

func main() {
	const filename = "input.txt"

	inputSlice, err := advent_util.ReadInts(filename)
	
	if (err != nil) {
		fmt.Println("Error reading input: ", err)
		os.Exit(1)
	}
	count := 0

	for index := 1; index < len(inputSlice); index++ {
		if inputSlice[index] > inputSlice[index-1] {
			count++
		}
	}

	fmt.Printf("Count: %d", count)
}