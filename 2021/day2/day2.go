package main

import (
	"fmt"

	"jburke.com/advent_util"
)

func movement_part1(inputCommands []string) complex128 {
	commands := map[string]complex128 {
		"forward": complex(1, 0),
		"down": complex(0, 1),
		"up": complex(0, -1),
	};
	// Part 1
	location := complex(0, 0)
	for _, line := range inputCommands {
		var direction string
		var distance float64

		_, err := fmt.Sscanf(line, "%s %f", &direction, &distance)

		if err != nil {
			panic(err)
		}
		location += commands[direction] * complex(distance, 0)
	}
	return location
}

func output_location(location complex128) {
	fmt.Printf("Location: %v\n", location)
	fmt.Printf("Product: %f\n", real(location)*imag(location))
}

// Part 2
// still use the same commands, real part increases horizontal by X units and depth by aim * X units, imaginary part increases aim by X units
func movement_part2(inputCommands []string) complex128 {
	commands := map[string]complex128 {
		"forward": complex(1, 0),
		"down": complex(0, 1),
		"up": complex(0, -1),
	};
	location := complex(0, 0)
	aim := 0.0
	for _, line := range inputCommands {
		var direction string
		var distance float64

		_, err := fmt.Sscanf(line, "%s %f", &direction, &distance)

		if err != nil {
			panic(err)
		}
		aim += imag(commands[direction]) * distance // if command is forward this will just be an increase of 0
		movementDelta := real(commands[direction]) * distance // if command is down or up this will just be an increase of 0
		location += complex(movementDelta, aim * movementDelta)
	}
	return location
}

func main() {
	const filename = "input.txt"

	inputSlice, err := advent_util.ReadLines(filename)

	if err != nil {
		panic(err)
	}

	location := movement_part1(inputSlice)
	output_location(location)

	location = movement_part2(inputSlice)
	output_location(location)	
}