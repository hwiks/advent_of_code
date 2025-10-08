package aoc_2015

import (
	"fmt"
	"os"
)

func get_next_house(current_house map[string]int, instruction string) map[string]int {

	var new_house = make(map[string]int)

	if instruction == "<" {
		new_house["x"] = current_house["x"] - 1
		new_house["y"] = current_house["y"]
	} else if instruction == ">" {
		new_house["x"] = current_house["x"] + 1
		new_house["y"] = current_house["y"]
	} else if instruction == "^" {
		new_house["y"] = current_house["y"] + 1
		new_house["x"] = current_house["x"]
	} else if instruction == "v" {
		new_house["y"] = current_house["y"] - 1
		new_house["x"] = current_house["x"]
	}

	return new_house

}

func get_all_houses(directions string) []map[string]int {
	var starting_point = make(map[string]int)
	starting_point["x"] = 0
	starting_point["y"] = 0

	var all_houses []map[string]int
	all_houses = append(all_houses, starting_point)
	for i := range len(directions) {
		var next_house map[string]int = get_next_house(all_houses[i], string(directions[i]))
		all_houses = append(all_houses, next_house)
	}

	return all_houses
}

func get_unique_houses(all_houses []map[string]int) []map[string]int {
	var unique_houses []map[string]int

	for i := range len(all_houses) {
		var already_there bool = false

		for j := range len(unique_houses) {
			var same_x bool = all_houses[i]["x"] == unique_houses[j]["x"]
			var same_y bool = all_houses[i]["y"] == unique_houses[j]["y"]

			if same_x && same_y {
				already_there = true
			}
		}

		if !already_there {
			unique_houses = append(unique_houses, all_houses[i])
		}

	}
	return unique_houses
}

func Run_which_houses(robot_santa bool) {

	fmt.Println("How many houses receive at least one present?")

	directions_file, err := os.ReadFile("aoc_2015/directions.txt")
	check_error(err)

	var directions string = string(directions_file)

	if robot_santa {
		fmt.Println("With the help of robot santa: ")

		var directions_santa string
		var directions_robot_santa string
		var santa_turn bool = true

		for i := range directions {
			if santa_turn {
				directions_santa = directions_santa + string(directions[i])
				santa_turn = false
			} else {
				directions_robot_santa = directions_robot_santa + string(directions[i])
				santa_turn = true
			}
		}

		var all_houses_santa = get_all_houses(directions_santa)
		var all_houses_robot_santa = get_all_houses(directions_robot_santa)

		var all_houses []map[string]int = append(all_houses_santa, all_houses_robot_santa...)

		var unique_houses = get_unique_houses(all_houses)
		fmt.Println(len(unique_houses))

	} else {

		var all_houses = get_all_houses(directions)
		var unique_houses = get_unique_houses(all_houses)
		fmt.Println(len(unique_houses))
	}

}
