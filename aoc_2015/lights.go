package aoc_2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func turn_on(lights [1000000][3]int, edges []int) [1000000][3]int {
	var first = edges[:2]
	var last = edges[2:]

	for i := range lights {
		if light_in_range(lights[i], first, last) {
			lights[i][2] = 1
		}
	}
	return lights

}

func light_in_range(light [3]int, first []int, last []int) bool {

	var x int = light[0]
	var y int = light[1]

	return (x >= first[0]) && (x <= last[0]) && (y >= first[1]) && (y <= last[1])
}

func get_edges(instruction string) []int {
	re := regexp.MustCompile("[0-9]+")
	var edges_string []string = re.FindAllString(instruction, -1)

	edges := make([]int, len(edges_string))

	for i, s := range edges_string {
		edges[i], _ = strconv.Atoi(s)
	}
	return edges
}

func turn_off(lights [1000000][3]int, edges []int) [1000000][3]int {
	var first = edges[:2]
	var last = edges[2:]

	for i := range lights {
		if light_in_range(lights[i], first, last) {
			lights[i][2] = 0
		}
	}
	return lights
}

func toggle(lights [1000000][3]int, edges []int) [1000000][3]int {
	var first = edges[:2]
	var last = edges[2:]

	for i := range lights {
		if light_in_range(lights[i], first, last) {

			fmt.Println("light ", lights[i])
			if lights[i][2] == 0 {
				lights[i][2] = 1
			} else if lights[i][2] == 1 {
				lights[i][2] = 0
			}
		}
	}
	return lights
}

func Run_lights() {

	var lights_instructions []string = extract_lines("aoc_2015/lights.txt")

	var lights [1000][1000]int

	for i := range lights_instructions[:1] {
		fmt.Println(lights_instructions[i])
		var edges []int = get_edges(lights_instructions[i])

		if strings.Contains(lights_instructions[i], "toggle") {
			lights = toggle(lights, edges)
		} else if strings.Contains(lights_instructions[i], "turn on") {
			lights = turn_on(lights, edges)
		} else if strings.Contains(lights_instructions[i], "turn off") {
			lights = turn_off(lights, edges)
		}

	}

	var light_turned_on int = 0

	for i := range lights {
		if lights[i][2] == 1 {
			light_turned_on = light_turned_on + 1
		}
	}

	fmt.Println(light_turned_on)
	fmt.Println(all_lights)

}
