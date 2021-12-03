package main

import (
	"fmt"

	"jburke.com/advent_util"
)

const num_bits = uint(12)

// common
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

func get_mask() uint {
	result := uint(0)
	for i := uint(0); i < num_bits; i++ {
		result |= 1 << i
	}
	return result
}

func print_binary_slice(input []uint16) {
	for _, value := range input {
		fmt.Printf("%b\n", value)
	}
}

// part 1 functions
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

// part 2 functions
func get_number_split(input []uint16, bit_index uint) (zeros []uint16, ones []uint16) {
	for _, value := range input {
		flag := uint16(1 << (num_bits - bit_index - 1))
		if value & flag != 0 {
			ones = append(ones, value)
		} else {
			zeros = append(zeros, value)
		}
	}

	return
}

func get_ratings(input []uint16, bit_index uint) ([]uint16, []uint16) {
	ones, zeros := get_number_split(input, bit_index)
	num_ones := len(ones)
	num_zeros := len(zeros)
	if num_ones > num_zeros {
		return ones, zeros
	} else {
		return zeros, ones
	}
}

func get_life_support_ratings(input []uint16) (oxygen uint, co2 uint) {
	most_common, least_common := get_ratings(input, 0)

	for cur_index := uint(1); cur_index < num_bits; cur_index++ {
		most_common, _ = get_ratings(most_common, cur_index)
		fmt.Printf("Most common for index %d\n", cur_index)
		print_binary_slice(most_common)
		if len(most_common) == 1 {
			oxygen = uint(most_common[0])
		}
		_, least_common = get_ratings(least_common, cur_index)
		fmt.Printf("Least common for index %d\n", cur_index)
		print_binary_slice(least_common)
		if len(least_common) == 1 {
			co2 = uint(least_common[0])
		}
		if oxygen != 0 && co2 != 0 {
			return
		}
	}
	return
}

func main() {
	input, err := get_binary_lines("input.txt")
	if err != nil {
		panic(err)
	}

	// part 1
	gamma := get_gamma_value(input)
	epsilon := ^gamma & get_mask()
	fmt.Printf("%d %b\n", gamma, gamma)
	fmt.Printf("%d %b\n", epsilon, epsilon)
	fmt.Printf("Product: %d\n", gamma * epsilon)

	// part 2

	oxygen, co2 := get_life_support_ratings(input)
	fmt.Printf("Oxygen: %d %b\n", oxygen, oxygen)
	fmt.Printf("CO2: %d %b\n", co2, co2)
	fmt.Printf("Product: %d\n", oxygen * co2)
}