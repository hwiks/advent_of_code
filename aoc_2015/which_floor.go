package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var instructions string

	content, err := os.ReadFile("floors_instructions.txt")

	if err != nil {
		log.Fatal(err)
	}

	instructions = string(content)

	var floor int
	floor = 0

	var basement_command int
	basement_command = 0

	for i := 0; i < len(instructions); i++ {
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
