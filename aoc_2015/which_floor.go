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

	for i := 0; i < len(instructions); i++ {
		if string(instructions[i]) == "(" {
			floor = floor + 1
		} else if string(instructions[i]) == ")" {
			floor = floor - 1
		}
	}

	fmt.Println(floor)
}
