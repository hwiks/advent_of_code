package aoc_2015

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extract_lines(input_file_name string) []string {

	var list []string

	list_file, err := os.Open(input_file_name)
	check_error(err)
	defer list_file.Close()

	scanner := bufio.NewScanner(list_file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	list_file.Close()

	return list
}

func is_nice(child string) bool {

	return has_3vowels(child) && has_repeating(child) && !has_forbidden(child)

}

func is_nice_updated(child string) bool {

	return has_pair_twice(child) && has_sandwich(child)

}

func has_sandwich(child string) bool {
	for i := range len(child) - 2 {

		if child[i] == child[i+2] {
			return true
		}
	}
	return false
}

func has_pair_twice(child string) bool {

	for i := range child[:len(child)-1] {

		var pair string = string(child[i]) + string(child[i+1])
		var smaller_child = strings.Replace(child, pair, "", -1)

		if len(smaller_child) <= len(child)-4 {
			return true
		}

	}

	return false

}

func has_forbidden(child string) bool {
	// forbidden sequences are: ab, cd, pq, or xy

	return strings.Contains(child, "ab") ||
		strings.Contains(child, "cd") ||
		strings.Contains(child, "pq") ||
		strings.Contains(child, "xy")

}

func has_repeating(child string) bool {
	for i := range len(child) - 1 {

		if child[i] == child[i+1] {
			return true
		}
	}
	return false
}

func has_3vowels(child string) bool {

	var vowels []byte

	for i := range len(child) {
		var letter byte = child[i]

		if letter == 'a' ||
			letter == 'e' ||
			letter == 'i' ||
			letter == 'o' ||
			letter == 'u' {
			vowels = append(vowels, letter)
		}

	}
	if len(vowels) >= 3 {
		return true
	} else {
		return false
	}
}

func Run_santas_list(updated bool) {

	var santas_list []string = extract_lines("aoc_2015/santas_list.txt")

	var nice_children []string

	for _, child := range santas_list[:4] {

		if updated {
			fmt.Println("len of child", len(child))
			if is_nice_updated(child) {

				nice_children = append(nice_children, child)
			}
		} else {

			if is_nice(child) {
				nice_children = append(nice_children, child)
			}
		}
	}

	fmt.Println("How many were nice?")
	fmt.Println(len(nice_children))

}
