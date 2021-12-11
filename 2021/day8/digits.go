package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//  aaa
// b   c
// b   c
//  ddd
// e   f
// e   f
//  ggg

func KnownDigitsFromInput(input string) (map[int]string, []string) {
	digits := make(map[int]string)
	unknown := make([]string, 0)
	for _, pattern := range strings.Fields(input) {
		patternLength := len(pattern)
		switch patternLength {
		case 2:
			digits[1] = pattern
		case 3:
			digits[7] = pattern
		case 4:
			digits[4] = pattern
		case 7:
			digits[8] = pattern
		default:
			unknown = append(unknown, pattern)
		}
	}
	return digits, unknown
}

func all_characters_in_input(input string, characters string) bool {
	for _, char := range characters {
		if !strings.ContainsRune(input, char) {
			return false
		}
	}
	return true
}

func UnknownDigitsFromInput(known_digits map[int]string, unknowns []string) (map[int]string, []string) {
	result := make(map[string]int)
	remaining_unknowns := make([]string, 0)
	for k, v := range known_digits {
		result[v] = k
	}
	for _, unknown := range unknowns {
		if _, ok := result[unknown]; !ok {
			unknown_length := len(unknown)
			if unknown_length == 5 {
				if (all_characters_in_input(unknown, known_digits[7])) {
					result[unknown] = 3
					known_digits[3] = unknown
				} else {
					remaining_unknowns = append(remaining_unknowns, unknown)
				}
			} else if unknown_length == 6 {
				if (!all_characters_in_input(unknown, known_digits[7])) {
					result[unknown] = 6
					known_digits[6] = unknown
				} else {
					if (all_characters_in_input(unknown, known_digits[4])) {
						result[unknown] = 9
						known_digits[9] = unknown
					} else {
						result[unknown] = 0
						known_digits[0] = unknown
					}
				}
			}
		}
	}
	return_value := make(map[int]string)
	for k, v := range result {
		return_value[v] = k
	}
	return return_value, remaining_unknowns
}

// remaining unkwowns are 2 and 5
func RemainingUnknowns(known_digits map[int]string, unknowns []string) map[string]int {
	result := make(map[string]int)
	for k, v := range known_digits {
		result[v] = k
	}
	for _, unknown := range unknowns {
		if _, ok := result[unknown]; !ok {
			if (all_characters_in_input(known_digits[6], unknown)) {
				result[unknown] = 5
				known_digits[5] = unknown
			} else {
				result[unknown] = 2
				known_digits[2] = unknown
			}
		}
	}
	return result
}

func SortStringChars(input string) string {
	string_slice := []rune(input)
	sort.Slice(string_slice, func (i, j int) bool { return string_slice[i] < string_slice[j]})
	return string(string_slice)
}

func SortKeys(digits map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range digits {
		result[SortStringChars(k)] = v
	}
	return result
}

func OutputToNumber(digits map[string]int, output string) int {
	value := make([]string, 4)
	for i, digit := range strings.Fields(output) {
		value[i] = fmt.Sprintf("%d", digits[SortStringChars(digit)])
	}
	result, err := strconv.Atoi(strings.Join(value, ""))
	if err != nil {
		panic(err)
	}
	return result
}