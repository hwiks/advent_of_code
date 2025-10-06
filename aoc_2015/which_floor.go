package aoc_2015

import (
	"fmt"
	"os"
)

func Run_which_floor() {

	var instructions string

	content, err := os.ReadFile("floors_instructions.txt")
	check_error(err)

	instructions = string(content)

	var floor int
	floor = 0

	var basement_command int
	basement_command = 0

	for i := range len(instructions) {
		if instructions[i] == '(' {
			floor = floor + 1
		} else if instructions[i] == ')' {
			floor = floor - 1
		}

		if (floor == -1) && (basement_command == 0) {
			basement_command = i + 1
			fmt.Println(basement_command)
		}
	}

	fmt.Println(floor)
}
