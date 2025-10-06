package aoc_2015

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func get_sizes(line string) []int {
	var sizes []string = strings.Split(line, "x")

	var sizes_numerical []int
	for _, dim := range sizes {
		j, _ := strconv.Atoi(string(dim))
		sizes_numerical = append(sizes_numerical, j)
	}
	return sizes_numerical

}

func get_smallest_area(sizes []int) int {
	sort.Ints(sizes)
	var two_smallest []int = sizes[:len(sizes)-1]

	return (two_smallest[0] * two_smallest[1])
}

func Run_wrap_presents() {

	var sizes []string

	presents_sizes, err := os.Open("aoc_2015/presents.txt")
	check_error(err)
	defer presents_sizes.Close()

	scanner := bufio.NewScanner(presents_sizes)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		sizes = append(sizes, scanner.Text())
	}

	presents_sizes.Close()

	var all_area_needed int
	all_area_needed = 0

	for _, line := range sizes {

		var sizes []int = get_sizes(line)

		var smallest_area int = get_smallest_area(sizes)

		var area_needed int = 2*(sizes[0]*sizes[1]+sizes[0]*sizes[2]+sizes[1]*sizes[2]) + smallest_area
		all_area_needed = all_area_needed + area_needed
	}

	fmt.Println(all_area_needed)

}
