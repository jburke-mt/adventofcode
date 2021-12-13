package main

import (
	"fmt"
	"sort"
	"strings"

	"jburke.com/advent_util"
)

var score_illegal_mapping map[string]int
var matching_chars map[string]string
var score_completion_mapping map[string]int

func is_opening(char string) bool {
	return char == "(" || char == "[" || char == "<" || char == "{"
}

func is_closing(char string) bool {
	return char == ")" || char == "]" || char == ">" || char == "}"
}

func matches(opener string, closer string) bool {
	return closer == matching_chars[opener]
}

func find_illegal_character(line string) string {
	stack := make([]string, 0)
	for _, char := range strings.Split(line, "") {
		if is_opening(char) {
			stack = append(stack, char)
		} else if is_closing(char) {
			n := len(stack) - 1
			opener := stack[n]
			stack = stack[:n]
			if !matches(opener, char) {
				return char
			}
		}
	}

	return ""
}


func init() {
	score_illegal_mapping = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	score_completion_mapping = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	matching_chars = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
}

func score_illegal_character(char string) int {
	return score_illegal_mapping[char]
}

func complete_line(line string) string {
	stack := make([]string, 0)
	for _, char := range strings.Split(line, "") {
		if is_opening(char) {
			stack = append(stack, char)
		} else if is_closing(char) {
			n := len(stack) - 1
			stack = stack[:n]
		}
	}
	completions := make([]string, 0)
	for len(stack) > 0 {
		n := len(stack) - 1
		completions = append(completions, matching_chars[stack[n]])
		stack = stack[:n]
	}

	return strings.Join(completions, "")
}

func score_completion(completion string) int {
	total := 0
	for _, char := range strings.Split(completion, "") {
		total *= 5
		total += score_completion_mapping[char]
	}
	return total
}

func main() {
	const filename = "input.txt"

	lines, err := advent_util.ReadLines(filename)

	if err != nil {
		panic(err)
	}

	sum_scores := 0

	completion_scores := make([]int, 0)
	for _, line := range lines {
		illegal_character := find_illegal_character(line)
		if illegal_character != "" {
			fmt.Println(illegal_character)
			score := score_illegal_character(illegal_character)
			fmt.Printf("Score is %d\n", score)
			sum_scores += score
		} else {
			completion := complete_line(line)
			fmt.Printf("Completion is %s\n", completion)
			score := score_completion(completion)
			fmt.Printf("Score is %d\n", score)
			completion_scores = append(completion_scores, score)
		}
	}

	fmt.Printf("Sum of illegal scores is %d\n", sum_scores)

	sort.Ints(completion_scores)
	middle_score := completion_scores[len(completion_scores)/2]
	fmt.Printf("Middle score is %d\n", middle_score)
}