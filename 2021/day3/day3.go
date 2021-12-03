package main

import (
	"fmt"

	"jburke.com/advent_util"
)

const num_bits = uint(12)

func get_binary_lines(filename string) ([]uint16, error) {
	inputSlice, err := advent_util.ReadLines(filename)

	if err != nil {
		return nil, err
	}

	var result []uint16
	for _, line := range inputSlice {
		var value uint16

		_, err := fmt.Sscanf(line, "%b", &value)
		if err != nil {
			panic(err)
		}

		result = append(result, value)
	}

	return result, nil
}

func get_most_common_bit(input []uint16, bit_index uint) uint16 {
	var zero_count uint
	var one_count uint

	for _, value := range input {
		flag := uint16(1 << bit_index)
		if value & flag != 0 {
			one_count++
		} else {
			zero_count++
		}
	}

	if one_count > zero_count {
		return 1
	} else {
		return 0	
	}
}

func get_gamma_value(input [] uint16) uint {
	var result uint

	for i := uint(0); i < num_bits; i++ {
		bit := get_most_common_bit(input, i)
		result |= uint(bit << i)
	}

	return result
}

func get_mask() uint {
	result := uint(0)
	for i := uint(0); i < num_bits; i++ {
		result |= 1 << i
	}
	return result
}

func main() {
	input, err := get_binary_lines("input.txt")
	if err != nil {
		panic(err)
	}

	gamma := get_gamma_value(input)
	epsilon := ^gamma & get_mask()
	fmt.Printf("%d %b\n", gamma, gamma)
	fmt.Printf("%d %b\n", epsilon, epsilon)
	fmt.Printf("Product: %d\n", gamma * epsilon)
}