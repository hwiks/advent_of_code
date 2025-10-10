package aoc_2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func get_edges(instruction string) []int {
	re := regexp.MustCompile("[0-9]+")
	var edges_string []string = re.FindAllString(instruction, -1)

	edges := make([]int, len(edges_string))

	for i, s := range edges_string {
		edges[i], _ = strconv.Atoi(s)
	}
	return edges
}

func turn_on(lights [1000][1000]int, edges []int) [1000][1000]int {
	var first = edges[:2]
	var last = edges[2:]

	for i := first[0]; i <= last[0]; i++ {
		for j := first[1]; j <= last[1]; j++ {
			lights[i][j] = lights[i][j] + 1
		}

	}

	return lights

}

func turn_off(lights [1000][1000]int, edges []int) [1000][1000]int {
	var first = edges[:2]
	var last = edges[2:]

	for i := first[0]; i <= last[0]; i++ {
		for j := first[1]; j <= last[1]; j++ {
			if lights[i][j] != 0 {
				lights[i][j] = lights[i][j] - 1
			}
		}

	}

	return lights
}

func toggle(lights [1000][1000]int, edges []int) [1000][1000]int {
	var first = edges[:2]
	var last = edges[2:]

	for i := first[0]; i <= last[0]; i++ {
		for j := first[1]; j <= last[1]; j++ {
			lights[i][j] = lights[i][j] + 2
		}

	}

	return lights
}

func Run_lights() {

	var lights_instructions []string = extract_lines("aoc_2015/lights.txt")

	var lights [1000][1000]int

	for i := range lights_instructions {
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

	for i := range 1000 {
		for j := range 1000 {
			light_turned_on = light_turned_on + lights[i][j]
		}
	}

	fmt.Println(light_turned_on)

}
